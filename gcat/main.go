package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	_ "runtime"
)

func cat(filenames []string) error {
	for _, filename := range filenames {
		var file *os.File
		var err error

		if filename == "-" { // Handle stdin
			file = os.Stdin
		} else {
			file, err = os.Open(filename)
			if err != nil {
				return fmt.Errorf("error opening file %s: %v", filename, err)
			}
			defer file.Close() // Close file when done
		}

		// Optimized buffer size for reading (adjust if needed)
		bufferSize := 4 * 1024 * 1024 // 4MB
		buffer := make([]byte, bufferSize)

		for {
			nr, err := file.Read(buffer)
			if nr > 0 {
				nw, err := os.Stdout.Write(buffer[:nr])
				if err != nil {
					return fmt.Errorf("error writing to stdout: %v", err)
				}
				if nw != nr {
					return fmt.Errorf("error writing to stdout: wrote %d bytes, expected %d", nw, nr)
				}

			}

			if err == io.EOF {
				break
			} else if err != nil {
				return fmt.Errorf("error reading from file %s: %v", filename, err)
			}
		}
	}
	return nil
}

func main() {
	flag.Parse()
	filenames := flag.Args()

	if len(filenames) == 0 {
		filenames = []string{"-"} // Default to stdin if no arguments
	}

	err := cat(filenames)
	if err != nil {
		fmt.Fprintf(os.Stderr, "cat: %v\n", err)
		os.Exit(1)
	}
}

// Build instructions for cross-platform:
// For Linux:
// GOOS=linux GOARCH=amd64 go build -o cat_linux

// For macOS:
// GOOS=darwin GOARCH=amd64 go build -o cat_darwin

// For Windows:
// GOOS=windows GOARCH=amd64 go build -o cat_windows.exe

// For all platforms at once (using `for` loop in bash):
// platforms=("linux/amd64" "darwin/amd64" "windows/amd64")
// for platform in "${platforms[@]}"; do
//   GOOS=${platform%%/*} GOARCH=${platform/*/} go build -o cat_${GOOS}_${GOARCH}
// done
