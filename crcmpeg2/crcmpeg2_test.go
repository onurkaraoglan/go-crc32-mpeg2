package crcmpeg2

import (
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCRC(t *testing.T) {
	var expected uint32 = 0x9EA45415
	hexStr := "1a2b3c"
	data, _ := hex.DecodeString(hexStr)
	result := CalculateCrcMpeg2(data)
	assert.Equal(t, expected, result, "Calculation Error")

	expected = 0x56D1FD12
	src := []byte("Hello, I'm Onur!")
	encodedStr := hex.EncodeToString(src) // Result = 48656c6c6f2c2049276d204f6e757221
	data, _ = hex.DecodeString(encodedStr)
	result = CalculateCrcMpeg2(data)
	assert.Equal(t, expected, result, "Calculation Error")
}
