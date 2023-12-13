package lib

import (
	"bufio"
	"bytes"
	"io"
	"log"
	"os"
	"strconv"
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
	return ReadStringsFromScanner(sc)
}

func ReadLines() []string {
	return ReadStringsFromScanner(bufio.NewScanner(os.Stdin))
}

func ReadByteSlices() [][]byte {
	return ReadByteSlicesFromScanner(bufio.NewScanner(os.Stdin))
}

func ReadStringsFromScanner(sc *bufio.Scanner) []string {
	result := make([]string, 0)
	for sc.Scan() {
		result = append(result, strings.TrimSpace(sc.Text()))
	}
	return result
}

func ReadByteSlicesFromScanner(sc *bufio.Scanner) [][]byte {
	result := make([][]byte, 0)
	for sc.Scan() {
		line := make([]byte, len(sc.Bytes()))
		copy(line, sc.Bytes())
		result = append(result, line)
	}
	return result
}

func ReadTextFromScanner(sc *bufio.Scanner) string {
	sc.Scan()
	return strings.TrimSpace(sc.Text())
}


func ScanFloat(sc *bufio.Scanner) float64 {
	sc.Scan()
	result, _ := strconv.ParseFloat(sc.Text(), 64)
	return result
}
func ScanInt(sc *bufio.Scanner) int {
	sc.Scan()
	num, _ := strconv.Atoi(sc.Text())
	return num
}

func ScanText(sc *bufio.Scanner) string {
	sc.Scan()
	return strings.TrimSpace(sc.Text())
}
