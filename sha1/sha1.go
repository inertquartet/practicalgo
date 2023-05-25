package main

import (
	"compress/gzip"
	"crypto/sha1"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	sig, err := sha1Sum("http.log.gz")
	if err != nil {
		log.Fatalf("sha1Sum failed: %s", err)
	}
	fmt.Println(sig)
}

// $ cat http.log.gz | gunzip | sha1sum
func sha1Sum(fileName string) (string, error) {
	// idiom: acquire a resoruce, check for acquisition error, defer its release
	file, err := os.Open(fileName)
	if err != nil {
		return "", nil
	}
	defer file.Close() // defers asre called in LIFO order

	r, err := gzip.NewReader(file)
	if err != nil {
		return "", err
	}

	// io.CopyN(os.Stdout, r, 1024)
	w := sha1.New()

	if _, err := io.Copy(w, r); err != nil {
		return "", err
	}

	sig := w.Sum(nil)

	return fmt.Sprintf("%x", sig), nil
}
