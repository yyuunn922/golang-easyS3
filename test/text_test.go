package test

import (
	"strings"
	"testing"
)

func TestText(t *testing.T) {
	var testText = "./../test/test/asdfv.png"

	temp := strings.Split(testText, "/")
	temp = strings.Split(temp[len(temp)-1], ".")

	t.Log(temp[0])

}
