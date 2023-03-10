package lib

import (
	"bufio"
	"bytes"
	"io"
	"log"
	"os"
	"strings"
)

func SetStdin(fName string) *os.File {
	f, err := os.Open(fName)
	if err == nil {
		os.Stdin = f
		return f
	} else {
		panic("file not found: " + fName)
	}
}

func CaptureStdout(f func()) []byte {
	old := os.Stdout
	r, w, err := os.Pipe()
	if err != nil {
		log.Fatal(err)
	}
	os.Stdout = w
	outCh := make(chan []byte)
	go func() {
		var b bytes.Buffer
		if _, err := io.Copy(&b, r); err != nil {
			log.Println(err)
		}
		outCh <- b.Bytes()
	}()
	f()
	// back to normal state
	w.Close()
	out := <-outCh
	if r != nil {
		r.Close()
	}
	os.Stdout = old
	return out
}

func ReadLinesFromFile(fileName string) []string {
	f, err := os.Open(fileName)
	if err != nil {
		log.Fatal(fileName)
	}
	defer f.Close()
	sc := bufio.NewScanner(f)
	return ReadLinesFromScanner(sc)
}

func ReadLinesFromScanner(sc *bufio.Scanner) []string {
	result := make([]string, 0)
	for sc.Scan() {
		result = append(result, strings.TrimSpace(sc.Text()))
	}
	return result
}

func ReadTextFromScanner(sc *bufio.Scanner) string {
	sc.Scan()
	return strings.TrimSpace(sc.Text())
}
