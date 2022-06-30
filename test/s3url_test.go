package test

import (
	"testing"

	"github.com/joho/godotenv"
	easys3 "github.com/madoleeee/golang-easyS3"
)

const (
	left  easys3.Direction = "TrimLeft"
	right easys3.Direction = "TrimRight"
)

func TestS3Url(t *testing.T) {

	// env를 로드합니다
	err := godotenv.Load(".env")
	if err != nil {
		panic(err.Error())
	}

	result, err := easys3.New()
	if err != nil {
		t.Fail()
	}

	t.Log(result.S3Url("1231223"))
}
