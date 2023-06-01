package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUnpack(t *testing.T) {
	for _, v := range testCase {
		assert.Equal(t, Unpack(v.Input), v.Output,
			fmt.Sprintf("Input Value: \"%s\"", v.Input))
	}
}

var testCase = []*struct {
	Input  string
	Output string
}{
	{"a4bc2d5e", "aaaabccddddde"},
	{"abcd", "abcd"},
	{"45", ""},
	{"", ""},
	{"qwe\\4\\5", "qwe45"},
	{"qwe\\45", "qwe44444"},
	{"qwe\\\\5", "qwe\\\\\\\\\\"},
}
