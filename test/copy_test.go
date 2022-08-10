package test

import (
	"fmt"
	"os"
	"testing"

	"github.com/joho/godotenv"
	easys3 "github.com/madoleeee/golang-easyS3"
)

func TestCopy(t *testing.T) {
	godotenv.Load(".env")

	s3, err := easys3.New()
	if err != nil {
		panic(err.Error())
	}

	file, err := os.Open("../tempFile/temp1.png")
	defer file.Close()
	if err != nil {
		panic(err.Error())
	}

	uploadFile, err := s3.Save(file, "/copyTestFolder", "")
	if err != nil {
		panic(err.Error())
	}

	temp, err := s3.Copy(uploadFile.FileUrl, "/copyTemp/testaa")
	if err != nil {
		panic(err.Error())
	}

	fmt.Print(temp)

}
