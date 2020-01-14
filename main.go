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

const destURL = "http://localhost:8888/set_fpga_bitstream"

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
	jsonString := "{\"partial_flag\":true,\"has_wrapper\":true}"
	_, err = jsonWriter.Write([]byte(jsonString))
	if err != nil {
		fmt.Printf("createPart failed!")
		return
	}

	fw, err := mpWriter.CreateFormFile("upload_file", "pl_part_pr_wrapper_inst0_pr_inst0_pr_inst_partial.bit")
	if err != nil {
		fmt.Printf("CreateFormFile failed!")
		return
	}
	f, err := os.Open("pl_part_pr_wrapper_inst0_pr_inst0_pr_inst_partial.bit")
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
