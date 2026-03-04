package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func copyFile(src, dst string) error {
	srcInfo, err := os.Stat(src)
	if err != nil {
		return fmt.Errorf("cannot stat source %s: %w", src, err)
	}

	if srcInfo.IsDir() {
		return fmt.Errorf("omitting directory %s", src)
	}

	in, err := os.Open(src)
	if err != nil {
		return fmt.Errorf("cannot open source %s: %w", src, err)
	}
	defer in.Close()

	out, err := os.OpenFile(dst, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, srcInfo.Mode())
	if err != nil {
		return fmt.Errorf("cannot open destination %s: %w", dst, err)
	}
	defer out.Close()

	if _, err := io.Copy(out, in); err != nil {
		return fmt.Errorf("failed copying %s to %s: %w", src, dst, err)
	}

	if err := out.Sync(); err != nil {
		return fmt.Errorf("failed syncing destination %s: %w", dst, err)
	}

	return nil
}

func copyDir(src, dst string) error {
	srcInfo, err := os.Stat(src)
	if err != nil {
		return fmt.Errorf("cannot stat source dir %s: %w", src, err)
	}

	if !srcInfo.IsDir() {
		return fmt.Errorf("source %s is not a directory", src)
	}

	if err := os.MkdirAll(dst, srcInfo.Mode()); err != nil {
		return fmt.Errorf("cannot create destination dir %s: %w", dst, err)
	}

	entries, err := os.ReadDir(src)
	if err != nil {
		return fmt.Errorf("cannot read source dir %s: %w", src, err)
	}

	for _, entry := range entries {
		srcPath := filepath.Join(src, entry.Name())
		dstPath := filepath.Join(dst, entry.Name())

		if entry.IsDir() {
			if err := copyDir(srcPath, dstPath); err != nil {
				return err
			}
			continue
		}

		if err := copyFile(srcPath, dstPath); err != nil {
			return err
		}
	}

	return nil
}

func main() {
	recursive := flag.Bool("r", false, "copy directories recursively")
	flag.BoolVar(recursive, "R", false, "copy directories recursively")
	flag.Parse()

	args := flag.Args()
	if len(args) < 2 {
		fmt.Fprintln(os.Stderr, "cp: missing file operand")
		os.Exit(1)
	}

	sources := args[:len(args)-1]
	destination := args[len(args)-1]

	dstInfo, dstErr := os.Stat(destination)
	dstExists := dstErr == nil
	dstIsDir := dstExists && dstInfo.IsDir()

	if len(sources) > 1 && !dstIsDir {
		fmt.Fprintf(os.Stderr, "cp: target '%s' is not a directory\n", destination)
		os.Exit(1)
	}

	for _, src := range sources {
		srcInfo, err := os.Stat(src)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cp: cannot stat '%s': %v\n", src, err)
			os.Exit(1)
		}

		target := destination
		if dstIsDir {
			target = filepath.Join(destination, filepath.Base(src))
		}

		if srcInfo.IsDir() {
			if !*recursive {
				fmt.Fprintf(os.Stderr, "cp: -r not specified; omitting directory '%s'\n", src)
				os.Exit(1)
			}
			if err := copyDir(src, target); err != nil {
				fmt.Fprintf(os.Stderr, "cp: %v\n", err)
				os.Exit(1)
			}
			continue
		}

		if err := copyFile(src, target); err != nil {
			fmt.Fprintf(os.Stderr, "cp: %v\n", err)
			os.Exit(1)
		}
	}
}
