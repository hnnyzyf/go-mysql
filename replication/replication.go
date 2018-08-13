package replication

import (
	"github.com/hnnyzyf/go-mysql/binlog"
	"github.com/hnnyzyf/go-mysql/client"
	"github.com/hnnyzyf/go-mysql/protocol"
	"github.com/hnnyzyf/go-mysql/util/binary"
	"github.com/hnnyzyf/go-mysql/util/hack"
	"github.com/juju/errors"

	"time"
)

var errRunnable = errors.Errorf("Replicatin:Dumper is still running,please stop first")

//Config represent the init infoamtion for a Dumper
type Config struct {
	//connction information
	Host   string
	User   string
	Passwd string

	//tradition dump information
	ServerId uint32

	//IsGtid
	FileName string
	Position uint32

	//whether use gtid
	IsGtid bool
}

//Dumper represent a traditional binlog Dumper
//Dumper use position and filename
type Dumper struct {
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

	//the error message
	err error
}

func NewDumper(name string) *Dumper {
	return &Dumper{
		name: name,
	}
}

//ChangeMaster set the basic information for Dumper
func (d *Dumper) ChangeMaster(cfg *Config) error {
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
func (d *Dumper) Start() (chan binlog.EventHandler, error) {
	//if Dumper is running,try to stop it first
	if d.isRunnable {
		return nil, errRunnable
	}

	//connect to server with not db
	if s, err := client.Connect(d.cfg.Host, d.cfg.User, d.cfg.Passwd, ""); err != nil {
		return nil, errors.Trace(err)
	} else {
		d.s = s
	}

	//resgister slave
	if err := d.sendComRegisterSlave(); err != nil {
		return nil, errors.Trace(err)
	}

	//read resopnse packet
	if err := d.readResponse(); err != nil {
		return nil, errors.Trace(err)
	}

	time.Sleep(20 * time.Second)
	//send binlog dump
	if err := d.sendComBinlogDump(); err != nil {
		return nil, errors.Trace(err)
	}

	//set runable
	d.isRunnable = true

	time.Sleep(20 * time.Second)
	//init output channel
	output := make(chan binlog.EventHandler)
	go func() {
		for {
			//read packet
			if b, err := d.s.ReadPacket(); err != nil {
				d.err = err
				break
			} else {
				switch {
				//handle binlog event
				case IsEventPacket(b):
					//read header
					if event, err := d.readEvent(b[2:]); err != nil {
						d.err = err
					} else {
						//asgin event
						output <- event
					}
				//handle error packet
				case client.IsErrPacket(b):
					d.err = client.ReadErrPacket(b)
					break
				case client.IsEofPacket(b):
					continue
				//other packet
				default:
					d.err = errors.Errorf("Replication:Unsupported packet")
					break
				}
			}
		}

		close(output)
	}()

	return output, nil
}

//sendComRegisterSlave resgiter slave inforamtion
func (d *Dumper) sendComRegisterSlave() error {
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

	//write Hostname length
	if err := payload.WriteInt1(uint8(len(d.name))); err != nil {
		return errors.Trace(err)
	}

	//write hostname
	if err := payload.WriteStringWithFixLen(hack.Slice(d.name)); err != nil {
		return errors.Trace(err)
	}

	//write user length
	if err := payload.WriteInt1(uint8(len(d.cfg.User))); err != nil {
		return errors.Trace(err)
	}

	//write user
	if err := payload.WriteStringWithFixLen(hack.Slice(d.cfg.User)); err != nil {
		return errors.Trace(err)
	}

	//write password length
	if err := payload.WriteInt1(uint8(len(d.cfg.Passwd))); err != nil {
		return errors.Trace(err)
	}

	//write password
	if err := payload.WriteStringWithFixLen(hack.Slice(d.cfg.Passwd)); err != nil {
		return errors.Trace(err)
	}

	//skip slave port,replication rank and master id
	if err := payload.Skip(10); err != nil {
		return errors.Trace(err)
	}

	//send dump command
	if err := d.s.WriteCommand(size, payload.Bytes()); err != nil {
		return errors.Trace(err)
	}

	return nil

}

//createDumpCommand create payload of COM_BINLOG_DUMPpayload
func (d *Dumper) sendComBinlogDump() error {
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
	if err := d.s.WriteCommand(size, payload.Bytes()); err != nil {
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

//Stop close the connection to master
func (d *Dumper) Stop() {
	d.isRunnable = false
	//close session
	d.s.Close()
}

//GtidDumper represent a gtid binlog Dumper
//GtidDumper use gtid instead of position and filename
type GtidDumper struct{}

func NewGtidDumper() *GtidDumper {
	return &GtidDumper{}
}
