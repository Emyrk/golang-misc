package util

import (
	"bufio"
	"crypto/rand"
	"os"
)

func Pause() {
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}

// RandomByteSliceOfLen()
// Returns a random set of bytes of a given length
func RandomByteSliceOfLen(sliceLen int) []byte {
	if sliceLen <= 0 {
		return nil
	}
	answer := make([]byte, sliceLen)
	_, err := rand.Read(answer)
	if err != nil {
		return nil
	}
	return answer
}
