package main

// THE LS IS CURRENTLY NOT AVAILABLE
// IM STUCK MAN :(

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"path/filepath"
	"sort"
	"strconv"
	"syscall"
	"time"
)

func main() {
	var (
		longFormat    bool
		humanReadable bool
	)

	flag.BoolVar(&longFormat, "l", false, "Use a long listing format.")
	flag.BoolVar(&humanReadable, "h", false, "Print sizes in human readable format.")
	flag.Parse()

	args := flag.Args()
	paths := []string{"."} // Default to current directory if no arguments are given
	if len(args) > 0 {
		paths = args
	}

	for _, path := range paths {
		files, err := ioutil.ReadDir(path)
		if err != nil {
			log.Fatalf("Error reading directory %s: %v", path, err)
		}

		if len(paths) > 1 { // Print directory name if listing multiple directories
			fmt.Printf("%s:\n", path)
		}

		fileInfos := make([]fileInfo, 0, len(files)) // Store file info for sorting
		for _, file := range files {
			fileInfos = append(fileInfos, newFileInfo(path, file))
		}

		sort.Slice(fileInfos, func(i, j int) bool {
			return fileInfos[i].Name < fileInfos[j].Name
		})

		for _, fi := range fileInfos {
			if longFormat {
				fmt.Println(fi.getLongFormat(humanReadable))
			} else {
				fmt.Println(fi.Name)
			}
		}
		if len(paths) > 1 {
			fmt.Println() // Add extra newline between directories
		}

	}
}

type fileInfo struct {
	Path       string
	Name       string
	Size       int64
	Mode       os.FileMode
	ModTime    time.Time
	Uid        int
	Gid        int
	UserName   string
	GroupName  string
	IsDir      bool
	LinkTarget string // For symbolic links
}

func newFileInfo(path string, file os.FileInfo) fileInfo {
	fi := fileInfo{
		Path:    path,
		Name:    file.Name(),
		Size:    file.Size(),
		Mode:    file.Mode(),
		ModTime: file.ModTime(),
		IsDir:   file.IsDir(),
	}

	stat := file.Sys().(*syscall.Stat_t)
	fi.Uid = int(stat.Uid)
	fi.Gid = int(stat.Gid)

	if user, err := user.LookupId(strconv.Itoa(fi.Uid)); err == nil {
		fi.UserName = user.Name
	} else {
		fi.UserName = strconv.Itoa(fi.Uid) // Fallback to UID if lookup fails
	}

	if group, err := user.LookupGroupId(strconv.Itoa(fi.Gid)); err == nil {
		fi.GroupName = group.Name
	} else {
		fi.GroupName = strconv.Itoa(fi.Gid) // Fallback to GID
	}

	if file.Mode()&os.ModeSymlink != 0 {
		target, err := filepath.EvalSymlinks(filepath.Join(path, file.Name()))
		if err == nil {
			fi.LinkTarget = target
		}
	}

	return fi
}

func (fi fileInfo) getLongFormat(humanReadable bool) string {
	permissions := fi.Mode.String()

	var sizeStr string
	if humanReadable {
		sizeStr = humanReadableSize(fi.Size)
	} else {
		sizeStr = strconv.FormatInt(fi.Size, 10)
	}

	modTime := fi.ModTime.Format("Jan 02 15:04")
	name := fi.Name

	if fi.LinkTarget != "" {
		name += " -> " + fi.LinkTarget
	}

	return fmt.Sprintf("%s %d %s %s %s %s %s",
		permissions, 1, fi.UserName, fi.GroupName, sizeStr, modTime, name)
}

func humanReadableSize(size int64) string {
	const unit = 1024
	if size < unit {
		return fmt.Sprintf("%d B", size)
	}
	div := int64(unit)
	exp := 0
	for n := size / div; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %ciB", float64(size)/float64(div), "KMGTPE"[exp])
}
