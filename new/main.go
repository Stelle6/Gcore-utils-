// Copyright (C) 2025 @Dlcuy22
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

func touch(filenames []string, timestamp time.Time, noCreate bool) error {
	for _, filename := range filenames {
		if noCreate {
			if _, err := os.Stat(filename); os.IsNotExist(err) {
				continue // Skip if file doesn't exist and noCreate is true
			} else if err != nil {
				return fmt.Errorf("error stating file %s: %v", filename, err)
			}
		}

		if timestamp.IsZero() {
			// Use current time if no timestamp is specified
			timestamp = time.Now()
		}

		err := os.Chtimes(filename, timestamp, timestamp)
		if err != nil {
			if os.IsNotExist(err) && !noCreate {
        // Create the file if it doesn't exist and noCreate is false
				file, err := os.Create(filename)
				if err != nil {
					return fmt.Errorf("error creating file %s: %v", filename, err)
				}
				file.Close() // Close the newly created file
				// Now try Chtimes again, it should succeed
				err = os.Chtimes(filename, timestamp, timestamp)
				if err != nil {
					return fmt.Errorf("error touching file %s after creation: %v", filename, err)
				}

			} else {
				return fmt.Errorf("error touching file %s: %v", filename, err)
			}
		}
	}
	return nil
}

func main() {
	timestampStr := flag.String("t", "", "Use the specified time instead of the current time")
	noCreate := flag.Bool("c", false, "Do not create any files")
	flag.Parse()
	filenames := flag.Args()

	if len(filenames) == 0 {
		fmt.Fprintln(os.Stderr, "new: missing file operand")
		os.Exit(1)
	}

	var timestamp time.Time
	if *timestampStr != "" {
		var err error
		timestamp, err = time.Parse("200601021504.05", *timestampStr) // Example format, customize as needed
		if err != nil {
			fmt.Fprintf(os.Stderr, "new: invalid date time format: %v\n", err)
			os.Exit(1)
		}
	}

	err := touch(filenames, timestamp, *noCreate)
	if err != nil {
		fmt.Fprintf(os.Stderr, "new: %v\n", err)
		os.Exit(1)
	}

}

// Build instructions for cross-platform:
// (Similar to the cat example, replace 'cat' with 'new')
// platforms=("linux/amd64" "darwin/amd64" "windows/amd64")
// for platform in "${platforms[@]}"; do
//   GOOS=${platform%%/*} GOARCH=${platform/*/} go build -o new_${GOOS}_${GOARCH}
// done
