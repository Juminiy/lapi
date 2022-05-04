package utils

import (
	"io/ioutil"
	"mime/multipart"
)

func FileParsing(file *multipart.FileHeader) ([]byte,error){
	fileTmp,err := file.Open()
	if err != nil {
		return nil,err
	}
	defer fileTmp.Close()
	fileByte,err := ioutil.ReadAll(fileTmp)
	if err != nil {
		return nil,err
	}
	return fileByte,nil
}