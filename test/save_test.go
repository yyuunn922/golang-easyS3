package test

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/joho/godotenv"
	easys3 "github.com/madoleeee/golang-easyS3"
)

func TestSave(t *testing.T) {
	// load env
	godotenv.Load(".env")

	s3, err := easys3.New()
	if err != nil {
		t.Log("error")
		t.Fail()
	}

	temp, err := ioutil.ReadDir("../tempFile")
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	}

	var output []easys3.UploadFileInfo
	for _, value := range temp {
		file, err := os.Open("../tempFile/" + value.Name())
		if err != nil {
			t.Log(err.Error())
			t.Fail()
		}

		temp, err := s3.Save(file, "/tempTest", "")
		output = append(output, temp)
	}
	t.Log(output)
}
