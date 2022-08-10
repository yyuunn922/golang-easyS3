package test

import (
	"testing"

	"github.com/joho/godotenv"
	easys3 "github.com/madoleeee/golang-easyS3"
)

func TestGetList(t *testing.T) {
	godotenv.Load(".env")

	s3, err := easys3.New()
	if err != nil {
		panic(err.Error())
	}

	output, err := s3.Load("temp/")
	if err != nil {
		t.Error(err.Error())
	}

	t.Log(output)

}
