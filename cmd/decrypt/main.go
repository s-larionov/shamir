package main

import (
	"encoding/base64"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"shamir/pkg/crypto"
	"shamir/pkg/shamir"
)

var (
	input  []byte
	output string
	parts  []string
)

func init() {
	stdin, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}
	input = stdin

	flag.StringVar(&output, "output", "", "filename for decoded content or empty for STDOUT")
	partsStr := flag.String("parts", "", "comma separated parts of the key")
	flag.Parse()
	parts = strings.Split(*partsStr, ",")
}

func main() {
	if len(parts) < 2 {
		printUsage("Count of parts of the key should of at least 2")
		return
	}

	keyParts, err := convertParts(parts)
	if err != nil {
		printUsage(fmt.Sprintf("Wrong parts format: %s", err.Error()))
		return
	}

	key, err := shamir.Combine(keyParts)
	if err != nil {
		printUsage(fmt.Sprintf("Unable to combine key: %s", err.Error()))
		return
	}

	decrypted, err := crypto.DecryptSimple(key, input)
	if err != nil {
		printUsage(fmt.Sprintf("Decryption error: %s", err.Error()))
		return
	}

	if output == "" {
		fmt.Printf("Decoded content:\n\n%s\n", decrypted)
	} else {
		if err := os.WriteFile(output, decrypted, 0600); err != nil {
			panic(err)
		}
		fmt.Printf("Decrypted file: %s\n", output)
	}
}

func convertParts(parts []string) ([][]byte, error) {
	converted := make([][]byte, 0, len(parts))

	for _, part := range parts {
		decodedPart, err := base64.StdEncoding.DecodeString(part)
		if err != nil {
			return nil, err
		}
		converted = append(converted, decodedPart)
	}

	return converted, nil
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
