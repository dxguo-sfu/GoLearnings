package basics

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBasicsSwap(t *testing.T) {
	s1 := "hello"
	s2 := "world"
	s3, s4 := Swap(s1, s2)

	assert.True(t, s1 == s4 && s2 == s3, "swap did not work correctly")
}

/*
	go get github.com/stretchr/testify/assert
	go test basics -v
	go test basics -bench=.

	To find the test package, the files must all sit in the same folder (basics)
*/
