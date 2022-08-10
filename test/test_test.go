package test

import (
	"testing"
)

func TestTest(t *testing.T) {
	s := "/temp/"

	t.Log(s[len(s)-1:])
}
