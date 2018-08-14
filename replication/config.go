package replication

//Config represent the init infoamtion for a Dumper
type Config struct {
	//connction information
	Host   string
	User   string
	Passwd string

	//binlog
	FileName string
	Position uint32

	//slave id information
	ServerId uint32

	//slave uuid information
	Uuid string

	//Client version because if slave version is lower than master
	//some functions may not work,for example checksum
	//so that we should tell master that we have the same version with you
	Version string

	//other function
	CheckSum bool
	Gtid     bool
}
