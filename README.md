
# Gcore-utils

A hobby project reimplementing GNU core utilities in Go, focusing on cross-platform compatibility.

## Overview

Gcore-utils aims to provide Go-based alternatives to common GNU core utilities, with an emphasis on:
- Cross-platform compatibility (Windows, Linux, macOS)
- Modern Go implementation
- Simple and maintainable codebase

## Installation

```bash
# Clone the repository
git clone https://github.com/Dlcuy22/Gcore-utils.git

# Navigate to project directory
cd Gcore-utils

# Enter your desired utilities to build
# example 
cd new

# Build the project
go build -o new.exe 
```

## Build All Utilities

```bash
make all
# binaries output to build/bin
```

## Utilities Status

| Utility | Status | Description | Cross-Platform |
|---------|--------|-------------|----------------|
| gcat    | ✅ Working | Concatenate and print files | Yes |
| new     | ✅ Working | Create new empty files (similar to touch) | Yes |
| cp      | ✅ Working | Copy files and directories | Yes |
| *More utilities planned* | 🚧 In Progress | | |

## Usage Examples

### gcat
```bash
# Print contents of a file
./gcat file.txt

```

### new (touch)
```bash
# Create a new empty file
./new file.txt

```

### cp
```bash
# Copy one file
./cp source.txt destination.txt

# Copy multiple files into a directory
./cp file1.txt file2.txt target_dir/

# Copy directories recursively
./cp -r source_dir target_dir
```


## Contributing

Contributions are welcome! Feel free to:
1. Fork the repository
2. Create a new branch for your feature
3. Submit a pull request

## Development Roadmap

- [x] gcat implementation
- [x] new (touch) implementation
- [x] cp utility
- [ ] Additional core utilities
- [ ] Unit tests
- [ ] CI/CD pipeline

## License

Copyright (C) 2025 Dlcuy22

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
 any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <https://www.gnu.org/licenses/>.

## Project Status

This project is under active development. New utilities will be added over time.


