package floats_test

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"testing"
)

func TestDecodeFloat(t *testing.T) {
	// 01000010110101111001100110011010
	b := []byte{0b01000010, 0b11010111, 0b10011001, 0b10011010}
	fmt.Println(b)

	var f float32
	buf := bytes.NewBuffer(b)
	err := binary.Read(buf, binary.LittleEndian, &f)
	fmt.Println(err)
	fmt.Println(f)
}
