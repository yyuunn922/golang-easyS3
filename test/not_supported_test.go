package test

import (
	"io/ioutil"
	"testing"

	"github.com/joho/godotenv"
	easys3 "github.com/madoleeee/golang-easyS3"
)

func TestNotSupportedFile(t *testing.T) {
	godotenv.Load(".env")

	s3, err := easys3.New()
	if err != nil {
		t.Error(err.Error())
	}

	files, err := ioutil.ReadDir("../tempFile")
	if err != nil {
		t.Error(err.Error())
	}

	for _, value := range files {
		_, err := s3.Save(value, "/tempTest", "")
		if err != nil {
			t.Log(err.Error())
		}
	}
}
