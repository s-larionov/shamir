package main

import (
	"encoding/base64"
	"flag"
	"fmt"
	"io"
	"os"

	"shamir/pkg/crypto"
	"shamir/pkg/shamir"
)

const (
	defaultParts     = 5
	defaultThreshold = 3
	aesKeySize       = 32
)

var (
	input     []byte
	output    string
	parts     int
	threshold int
)

func init() {
	stdin, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}
	input = stdin

	flag.IntVar(&parts, "parts", defaultParts, "how many parts of the key should be created")
	flag.IntVar(&threshold, "threshold", defaultThreshold, "how many parts of the key are required for recombination")
	flag.StringVar(&output, "output", "out", "filename for encoded content")
	flag.Parse()
}

func main() {
	if threshold < 2 {
		printUsage("Threshold should be at least 2")
		return
	}

	if parts <= threshold {
		printUsage("Count of parts could not be less than threshold")
		return
	}

	if len(input) == 0 {
		printUsage("Encoded content should not be empty")
		return
	}

	key, err := crypto.RandBytes(aesKeySize)
	if err != nil {
		panic(err)
	}

	encrypted, err := crypto.EncryptSimple(key, input)
	if err != nil {
		panic(err)
	}

	keyParts, err := shamir.Split(key, parts, threshold)
	if err != nil {
		panic(err)
	}

	if err := os.WriteFile(output, encrypted, 0600); err != nil {
		panic(err)
	}

	fmt.Printf("Encrypted file: %s\n", output)
	fmt.Printf("Key parts (required %d of %d):\n", threshold, parts)
	for _, part := range keyParts {
		fmt.Printf("– %s\n", base64.StdEncoding.EncodeToString(part))
	}
}

func printUsage(err string) {
	if err != "" {
		fmt.Println(err)
		fmt.Println()
	}

	fmt.Printf("cat file | %s [args]\n\n", os.Args[0])

	flag.Usage()
	return
}
