package test

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
	easys3 "github.com/madoleeee/golang-easyS3"
)

func TestDelete(t *testing.T) {
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

	temp, err := s3.Save(file, "/deleteTest", "")
	if err != nil {
		panic(err.Error())
	}
	t.Logf("생성된 데이터 이름 : %v", temp.FileUrl)

	s3.Delete(temp.FileUrl)

	t.Log("삭제")

}
