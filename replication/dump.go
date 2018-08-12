package replication

import (
	"github.com/hnnyzyf/go-mysql/binlog"
	"github.com/hnnyzyf/go-mysql/client"
	"github.com/hnnyzyf/go-mysql/protocol"
	"github.com/hnnyzyf/go-mysql/util/hack"
	"github.com/juju/errors"
)

//Config represent the init infoamtion for a dumper
type Config struct {
	//connction information
	Host   string
	User   string
	Passwd string

	//tradition dump information
	ServerId uint32

	//IsGtid
	FileName string
	Position string

	//whether use gtid
	IsGtid bool
}

//Dumper represent a traditional binlog dumper
//Dumper use position and filename
type Dumper struct {
	//the name of dumper
	name string

	//the Session we used to connect to server
	s *client.Session

	//the init configuration
	cfg *Config

	//current binlog name
	file string
	//current binlog pos
	pos uint32

	//output channel return the event
	output chan binlog.EventHandler

	//the error message
	err error
}

func NewDumper(name string, cfg *Config) (*Dumper, error) {
	if len(cfg.User) > 0xFF {
		return nil, errors.Errorf("Replication:username is more than 255 bytes")
	}

	if len(cfg.Passwd) > 0xFF {
		return nil, errors.Errorf("Replication:password is more than 255 bytes")
	}

	if len(name) > 0xFF {
		return nil, errors.Errorf("Replication:password is more than 255 bytes")
	}

	return &Dumper{
		cfg:  cfg,
		file: cfg.FileName,
		pos:  cfg.Position,
	}, nil
}

//Start begins to dump binlog from server
func (d *Dumper) Start() (chan binlog.EventHandler, error) {
	//init output channel
	d.output = make(chan binlog.EventHandler)

	//connect to server with not db
	s, err := client.Connect(d.cfg.Host, d.cfg.User, d.cfg.Passwd, "")
	if err != nil {
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

	//send binlog dump
	if err := d.sendComBinlogDump(); err != nil {
		return nil, errors.Trace(err)
	}

	//begin to read
	go func() {
		for {
			//read packet
			if b, err := d.s.ReadPacket(); err != nil {
				d.err = errors.Trace(err)
				break
			} else {
				switch {
				//handle binlog event
				case IsEventPacket(b):
					//read header
					if event, err := binlog.ReadEvent(buffer[2:]); err != nil {
						d.err = errors.Trace(err)
					} else {
						//asgin event
						d.output <- event
					}
				//handle error packet
				case IsErrPacket(b):
					d.err = client.ReadErrPacket(b)
					break
				//other packeet
				default:
					d.err = errors.Errorf("Replication:Unsupport pakcet type")
					break
				}
			}
		}

		Close(d.output)
	}()

	return d.output, nil
}

func (d *Dumper) Stop() {

	//close session
	d.s.Close()

	//close output channel
	close(d.output)
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
	if err := payload.WriteInt1(uint8(len(c.cfg.Passwd))); err != nil {
		return errors.Trace(err)
	}

	//write password
	if err := payload.WriteStringWithFixLen(hack.Slice(c.cfg.Passwd)); err != nil {
		return errors.Trace(err)
	}

	//skip slave port,replication rank and master id
	if err := payload.Skip(10); err != nil {
		return errors.Trace(err)
	}

	//send dump command
	if err := d.s.WriteCommand(uint32(size), payload.ToReader()); err != nil {
		return errors.Trace(err)
	}

	return nil

}

//createDumpCommand create payload of COM_BINLOG_DUMPpayload
func (d *Dumper) sendComBinlogDump() error {
	//init size
	size := 1 + 4 + 2 + 4 + len(d.name) + 1

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

	//write flags
	if err := payload.WriteInt2(0x0000); err != nil {
		return errors.Trace(err)
	}

	//write server id
	if err := payload.WriteInt4(d.cfg.ServerId); err != nil {
		return errors.Trace(err)
	}

	//write binlog filename
	if err := payload.WriteStringWithNull(hack.Slice(d.name)); err != nil {
		return errors.Trace(err)
	}

	//send dump command
	if err := d.s.WriteCommand(uint32(size), payload.ToReader()); err != nil {
		return errors.Trace(err)
	}

	return nil

}

//readRespnse return the response information
func (d *dumper) readResponse() error {
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

//GtidDumper represent a gtid binlog dumper
//GtidDumper use gtid instead of position and filename
type GtidDumper struct{}

func NewGtidDumper() *GtidDumper {
	return &GtidDumper{}
}
