package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"os"
)

const destURL = "http://10.168.1.166:8888/set_fpga_bitstream"

func main() {
	var bufReader bytes.Buffer
	mpWriter := multipart.NewWriter(&bufReader)
	h := make(textproto.MIMEHeader)
	h.Set("Content-Type", "application/json")
	jsonWriter, err := mpWriter.CreatePart(h)
	if err != nil {
		fmt.Printf("createPart failed!")
		return
	}
	jsonString := "{\"partial_flag\":true,\"partion_number\":0}"
	_, err = jsonWriter.Write([]byte(jsonString))
	if err != nil {
		fmt.Printf("createPart failed!")
		return
	}

	fw, err := mpWriter.CreateFormFile("upload_file", os.Args[1])
	if err != nil {
		fmt.Printf("CreateFormFile failed!")
		return
	}
	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Printf("CreateFormFile failed!")
		return
	}
	defer f.Close()

	_, err = io.Copy(fw, f)
	if err != nil {
		fmt.Printf("io.Copy failed!")
		return
	}
	_ = mpWriter.Close()
	//log.Printf(bufReader.String())
	_, err = http.Post(destURL, mpWriter.FormDataContentType(), &bufReader)

	if err != nil {
		log.Printf("http.Post failed!")
		return
	}

}
