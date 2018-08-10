package binlog


type reader struct{
	name string
	r io.Reader
}

func NewReader()