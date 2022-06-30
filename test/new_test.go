package test

import (
	"testing"

	"github.com/joho/godotenv"
	easyS3 "github.com/madoleeee/golang-easyS3"
)

func TestNew(t *testing.T) {
	// env를 로드합니다
	err := godotenv.Load(".env")
	if err != nil {
		panic(err.Error())
	}

	i, err := easyS3.New()
	if err != nil {
		panic(err.Error())
	}

	t.Log(i)

}
