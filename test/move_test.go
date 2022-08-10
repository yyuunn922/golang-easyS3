package test

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
	easys3 "github.com/madoleeee/golang-easyS3"
)

func TestMove(t *testing.T) {
	godotenv.Load(".env")

	s3, err := easys3.New()
	if err != nil {
		panic(err.Error())
	}

	file, err := os.Open("../tempFile/temp1.png")
	if err != nil {
		panic(err.Error())
	}
	result, err := s3.Save(file, "/moveTest", "")
	if err != nil {
		panic(err.Error())
	}

	moveResult, err := s3.Move(result.FileUrl, "/moveTestTest/t123123123")
	if err != nil {
		panic(err.Error())
	}

	t.Log(moveResult)

}
