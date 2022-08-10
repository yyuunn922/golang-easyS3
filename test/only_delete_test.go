package test

import (
	"testing"

	"github.com/joho/godotenv"
	easys3 "github.com/madoleeee/golang-easyS3"
)

func TestOnlyDelete(t *testing.T) {
	godotenv.Load(".env")

	s3, err := easys3.New()

	if err != nil {
		panic(err.Error())
	}

	result, err := s3.Delete("/moveTest/a75e5a82-6caf-47a4-ba42-cf3f32b113d5")
	if err != nil {
		panic(err.Error())
	}

	t.Log(result)
}
