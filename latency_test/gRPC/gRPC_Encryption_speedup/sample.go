package main

import (
	"bytes"
	//"unsafe"
	//"fmt"
	"os"
	"fmt"
	"bufio"
	"image"
	"image/png"


)


func ExampleBuffer_Grow() {
	//image to bytes

	image1 := "Plot3.png"
	file, err := os.Open(image1)

  	if err != nil {
          fmt.Println(err)
          os.Exit(1)
  	}
	
	defer file.Close()

	fileInfo, _ := file.Stat()
  	var size int64 = fileInfo.Size()
  	bytes1 := make([]byte, size)

	buffer := bufio.NewReader(file)
  	_, err = buffer.Read(bytes1)
	
	
	//fmt.Printf("%d", bytes)

	//bytes to image
	img, _, _ := image.Decode(bytes.NewReader(bytes1))
	
   	out, err := os.Create("./QRImg.png")

   	if err != nil {
             fmt.Println(err)
             os.Exit(1)
   	}

   	err = png.Encode(out, img)

   	if err != nil {
            fmt.Println(err)
            os.Exit(1)
   	}
}


func main(){
	ExampleBuffer_Grow()

}
