package main

import (
	"bytes"
	"fmt"
	"syscall/js"

	"github.com/rwcarlsen/goexif/exif"
)

func processImage(this js.Value, p []js.Value) any {
	if len(p) != 1 {
		return "Invalid no of arguments passed"
	}

	imgBytes := make([]uint8, p[0].Get("byteLength").Int())
	js.CopyBytesToGo(imgBytes, p[0])

	r := bytes.NewReader(imgBytes)
	e, err := exif.Decode(r)
	if err != nil {
		fmt.Println("decode err: ", err.Error())
		return nil
	}

	imgJson, err := e.MarshalJSON()
	if err != nil {
		fmt.Println("unmarshall err: ", err.Error())
		return nil
	}

	return string(imgJson)
}

func main() {
	js.Global().Set("process", js.FuncOf(processImage))
	<-make(chan struct{})
}
