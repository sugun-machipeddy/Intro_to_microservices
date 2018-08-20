package main

import (
	//"bytes"
	"unsafe"
	"fmt"

)


func ExampleBuffer_Grow() {
	//var b bytes.Buffer
	//b.Grow(10)
	//b.Write([]byte("Hello World"))
	//fmt.Printf("length:%d ", b.Len())
	//fmt.Printf("length:%d ", b.Cap())
	//fmt.Printf("length:%d ", unsafe.Sizeof(b))
	//s := ""
	//var c bytes.Buffer
	//c.Write([]byte("1.823597ms"))
	//fmt.Printf("length:%d ", c.Len())
	//fmt.Printf("length:%d ", c.Cap())
	//fmt.Printf("length:%d ", unsafe.Sizeof(s))
	//fmt.Printf("length:%d ", unsafe.Sizeof(c))

	//b.Grow(64)
	//bb := b.Bytes()
	//b.Write([]byte("64 bytes or fewer"))
	//fmt.Printf("%q", bb[:b.Len()])
	// Output: "64 bytes or fewer"
	var i int
	var u uint
	var up uintptr
	var s string


fmt.Printf("i Type:%T Size:%d\n", i, unsafe.Sizeof(i))
fmt.Printf("u Type:%T Size:%d\n", u, unsafe.Sizeof(u))
fmt.Printf("up Type:%T Size:%d\n", up, unsafe.Sizeof(up))
fmt.Printf("up Type:%T Size:%d\n", s, unsafe.Sizeof(s))
}


func main(){
	ExampleBuffer_Grow()

}
