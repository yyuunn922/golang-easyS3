package test

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/joho/godotenv"
	easys3 "github.com/madoleeee/golang-easyS3"
)

func TestSave(t *testing.T) {

	err := godotenv.Load(".env")

	temp, err := easys3.New()
	if err != nil {
		panic(err.Error())
	}

	files, err := ioutil.ReadDir("../tempFile")
	if err != nil {
		panic(err.Error())
	}

	var output []string

	for _, file := range files {
		aa, err := os.Open("../tempFile/" + file.Name())
		if err != nil {
			panic(err.Error())
		}
		defer aa.Close()

		result, err := temp.Save(aa, "/temp", "")
		if err != nil {
			panic(err.Error())
		}
		output = append(output, result)
	}

	t.Log(output)
}
