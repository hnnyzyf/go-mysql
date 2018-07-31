package hack

import (
	"reflect"
	"unsafe"
)

//String convert byte slice to string with no copy
//It may be unsafe because s is not immutable!
func String(b []byte) (s string) {
	pb := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	ps := (*reflect.StringHeader)(unsafe.Pointer(&s))
	ps.Data = pb.Data
	ps.Len = pb.Len
	return
}

//Slice convert string to slice with no copy
//Never try to change b!
func Slice(s string) (b []byte) {
	pb := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	ps := (*reflect.StringHeader)(unsafe.Pointer(&s))
	pb.Data = ps.Data
	pb.Len = ps.Len
	pb.Cap = ps.Len
	return
}
