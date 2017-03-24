package util

import (
	"bufio"
	"os"
)

func Pause() {
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}
