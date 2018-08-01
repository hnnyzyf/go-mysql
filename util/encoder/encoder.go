package code

//we hope a Encoder could implement this interface 
type Encoder interface{
	ReadRune() ([]byte,error)
}

