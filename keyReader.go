package main

import (
	"os"
	"golang.org/x/term"
)

func captureKeyPresses(keyPressChannel chan string) {
	fd := int(os.Stdin.Fd())
	oldState, err := term.MakeRaw(fd)
	if err != nil {
		// fallback: just exit goroutine on error
		return
	}
	defer term.Restore(fd, oldState)
	var quit = false

	buf := make([]byte, 3) // 1..3 bytes to capture common escape sequences
	for !quit {
		n, err := os.Stdin.Read(buf)
		if err != nil {
			return
		}
		if n > 0 {
			// send the raw bytes as a string; arrow keys / function keys are escape sequences
			keyPressChannel <- string(buf[:n])
		}
	}
}