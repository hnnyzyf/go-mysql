package replication

import (
	"fmt"
	"sync"

	"github.com/hnnyzyf/go-mysql/binlog"
	"github.com/hnnyzyf/go-mysql/client"
	"github.com/hnnyzyf/go-mysql/protocol"
	"github.com/hnnyzyf/go-mysql/util/binary"
	"github.com/hnnyzyf/go-mysql/util/hack"
	"github.com/juju/errors"
)

var errRunnable = errors.Errorf("Replicatin:Dumper is still running,please stop first")

//Dumper represent a traditional binlog Dumper
//Dumper use position and filename
type Dumper struct {
	//to keep goroutine safety
	mutex sync.Mutex

	//the name of Dumper
	name string

	//the Session we used to connect to server
	s *client.Session

	//the init configuration
	cfg *Config

	//current binlog name
	file string
	//current binlog pos
	pos uint32

	//state
	isRunnable bool

	//eventhandler channel
	output chan binlog.EventHandler
	//the error message
	err chan error
}

func NewDumper(name string) *Dumper {
	return &Dumper{
		name: name,
		err:  make(chan error),
	}
}

//ChangeMaster set the basic information for Dumper
func (d *Dumper) ChangeMaster(cfg *Config) error {
	d.mutex.Lock()
	defer d.mutex.Unlock()

	if d.isRunnable {
		return errRunnable
	} else {
		d.cfg = cfg
		d.file = cfg.FileName
		d.pos = cfg.Position
		return nil
	}
}

//Start begins to dump binlog from server
func (d *Dumper) Start() error {
	d.mutex.Lock()
	defer d.mutex.Unlock()

	//if Dumper is running,try to stop it first
	if d.isRunnable {
		return errRunnable
	} else {
		d.isRunnable = true
		d.output = make(chan binlog.EventHandler)
	}

	//connect to server with not db
	if err := d.connect(); err != nil {
		return errors.Trace(err)
	}

	//init the basice slave variables
	if err := d.init(); err != nil {
		return errors.Trace(err)
	}

	//resgister slave
	if err := d.RequestComRegisterSlave(); err != nil {
		return errors.Trace(err)
	}

	//read resopnse packet
	if err := d.readResponse(); err != nil {
		return errors.Trace(err)
	}

	//send binlog dump
	if err := d.RequestComBinlogDump(); err != nil {
		return errors.Trace(err)
	}

	//init output channel
	go func() {
		defer d.Stop()
		for {
			//read packet
			if b, err := d.s.ReadPacket(); err != nil {
				d.err <- errors.Trace(err)
				return
			} else {
				switch {
				//handle binlog event
				case IsEventPacket(b):
					//read header
					if event, err := d.readEvent(b[2:]); err != nil {
						d.err <- errors.Trace(err)
						return
					} else {
						//asgin event
						d.output <- event
					}
				//handle error packet
				case client.IsErrPacket(b):
					d.err <- errors.Trace(client.ReadErrPacket(b))
					return
				case client.IsEofPacket(b):
					continue
				//other packet
				default:
					d.err <- errors.Errorf("Replication:unsupported packet")
					return
				}
			}
		}
	}()

	return nil
}

//Stop close the connection to master
func (d *Dumper) Stop() {
	d.mutex.Lock()
	defer d.mutex.Unlock()

	if d.isRunnable {
		//modify state
		d.isRunnable = false
		//close session
		d.s.Close()
	}
}

//Next return the event we read or error we read from
func (d *Dumper) Next() (binlog.EventHandler, error) {
	if !d.isRunnable {
		return nil, errors.Errorf("Replication:dumper is not still running")
	}

	//get result
	select {
	case event := <-d.output:
		return event, nil
	case err := <-d.err:
		return nil, err
	}
}

//connect connect to server according to the configurations
func (d *Dumper) connect() error {
	//init configuration
	cfg := client.NewConfig()
	if len(d.cfg.Version) != 0 {
		cfg.SetConnectionAttrs("_client_version", d.cfg.Version)
	}

	//connect tot server
	if s, err := client.ConnectWithConfig(d.cfg.Host, d.cfg.User, d.cfg.Passwd, "", cfg); err != nil {
		return errors.Trace(err)
	} else {
		d.s = s
	}

	return nil
}

//connect connect to server according to the configurations
//look details in sql/rpl_slave.cc line 4233
func (d *Dumper) init() error {
	//check master version
	if err := d.checkVersion(); err != nil {
		return errors.Trace(err)
	}

	//check master id
	if err := d.checkServerId(); err != nil {
		return errors.Trace(err)
	}

	//set slave uuid and binlog checksum
	if err := d.checkUuid(); err != nil {
		return errors.Trace(err)
	}
}

func (d *Dumper) checkVersion() error {
	return nil
}

func (d *Dumper) checkServerId() error {
	return nil
}

func (d *Dumper) checkUuid() error {
	return nil
}

//RequestComRegisterSlave resgiter slave inforamtion
func (d *Dumper) RequestComRegisterSlave() error {
	//init size
	size := 1 + 4 + 1 + len(d.name) + 1 + len(d.cfg.User) + 1 + len(d.cfg.Passwd) + 2 + 4 + 4

	//create payload
	payload := binary.NewBuffer(make([]byte, size))

	//write command name
	if err := payload.WriteInt1(protocol.COM_REGISTER_SLAVE); err != nil {
		return errors.Trace(err)
	}

	//write server id
	if err := payload.WriteInt4(d.cfg.ServerId); err != nil {
		return errors.Trace(err)
	}

	//write Hostname
	if len(d.name) > 0xff {
		return errors.Errorf("Replication:failed to register slave,hostname is too long")
	} else {
		if err := payload.WriteLengthEncodedString(hack.Slice(d.name)); err != nil {
			return errors.Trace(err)
		}
	}

	//write user length
	if len(d.cfg.User) > 0xff {
		return errors.Errorf("Replication:failed to register slave,username is too long")
	} else {
		if err := payload.WriteLengthEncodedString(hack.Slice(d.cfg.User)); err != nil {
			return errors.Trace(err)
		}
	}

	//write password length
	if len(d.cfg.Passwd) > 0xff {
		return errors.Errorf("Replication:failed to register slave,password is too long")
	} else {
		if err := payload.WriteLengthEncodedString(hack.Slice(d.cfg.Passwd)); err != nil {
			return errors.Trace(err)
		}
	}

	//skip slave port,replication rank and master id
	if err := payload.Skip(10); err != nil {
		return errors.Trace(err)
	}

	//send dump command
	if err := d.s.RequestCommand(size, payload.Bytes()); err != nil {
		return errors.Trace(err)
	}

	return nil

}

//createDumpCommand create payload of COM_BINLOG_DUMPpayload
func (d *Dumper) RequestComBinlogDump() error {
	//init size
	size := 1 + 4 + 2 + 4 + len(d.file)

	//create payload
	payload := binary.NewBuffer(make([]byte, size))

	//wirte command name
	if err := payload.WriteInt1(protocol.COM_BINLOG_DUMP); err != nil {
		return errors.Trace(err)
	}

	//write binglog position
	if err := payload.WriteInt4(d.pos); err != nil {
		return errors.Trace(err)
	}

	//write flags,only one choice
	if err := payload.WriteInt2(BINLOG_DUMP_NEVER_STOP); err != nil {
		return errors.Trace(err)
	}

	//write server id
	if err := payload.WriteInt4(d.cfg.ServerId); err != nil {
		return errors.Trace(err)
	}

	//write binlog filename
	if err := payload.WriteStringWithFixLen(hack.Slice(d.file)); err != nil {
		return errors.Trace(err)
	}

	//send dump command
	if err := d.s.RequestCommand(size, payload.Bytes()); err != nil {
		return errors.Trace(err)
	}

	return nil

}

//readRespnse return the response information
func (d *Dumper) readResponse() error {
	//read packet from connection
	buffer, err := d.s.ReadPacket()
	if err != nil {
		return errors.Trace(err)
	}

	switch {
	case client.IsOkPacket(buffer):
		return nil
	case client.IsErrPacket(buffer):
		return client.ReadErrPacket(buffer)
	default:
		return errors.Errorf("Replication:Unsupported response packet")
	}
}

//readEvent read event from buffer
func (d *Dumper) readEvent(buffer []byte) (binlog.EventHandler, error) {
	if event, err := binlog.ReadEvent(buffer); err != nil {
		return nil, errors.Trace(err)
	} else {
		//record the  the
		d.pos += uint32(len(buffer))
		return event, nil
	}
}

//GtidDumper represent a gtid binlog Dumper
//GtidDumper use gtid instead of position and filename
type GtidDumper struct{}

func NewGtidDumper() *GtidDumper {
	return &GtidDumper{}
}
