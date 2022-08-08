package test

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/joho/godotenv"
	easys3 "github.com/madoleeee/golang-easyS3"
)

func TestXxx(t *testing.T) {

	godotenv.Load(".env")

	e, err := easys3.New()
	e.Region = os.Getenv("ap-northeast-2")
	e.AccessKey = os.Getenv("AWS_ACCESSKEY")
	e.SecretKey = os.Getenv("AWS_SECRETKEY")
	e.BucketName = os.Getenv("AWS_BUCKET_NAME")

	if err != nil {
		t.Log(err.Error())
	}

	fileLists, err := ioutil.ReadDir("./../tempFile")
	if err != nil {
		t.Log(err.Error())
	}

	for _, value := range fileLists {

		tempFile, err := os.Open("./../tempFile/" + value.Name())
		if err != nil {
			panic(err.Error())
		}
		defer tempFile.Close()

		temp, err := e.Save(tempFile, "/newTemp", "")
		if err != nil {
			t.Log(err.Error())
		}
		t.Log(temp)

	}

}
