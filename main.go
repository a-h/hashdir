package main

import (
	"crypto/sha256"
	"flag"
	"fmt"
	"hash"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"slices"
	"strings"
)

var help = `hashdir

Calculate the hash of all files in a directory tree. Empty directories are ignored.

Usage:

  hashdir
  hashdir -ignore-git=false
  hashdir ./dir`

var flagHelp = flag.Bool("help", false, "Set to view help")
var flagIgnoreGit = flag.Bool("ignore-git", true, "Set to false to include .git directories")
var flagIgnore = flag.String("ignore", "", "Set to a space separated list of directories to ignore.")

func main() {
	flag.Parse()
	if *flagHelp {
		fmt.Println(help)
		return
	}

	if len(flag.Args()) > 1 {
		fmt.Println("Unexpected args.")
		fmt.Println()
		fmt.Println(help)
		os.Exit(1)
	}
	dir := "."
	if len(flag.Args()) == 1 {
		dir = flag.Args()[0]
	}
	ignore := strings.Fields(*flagIgnore)
	if *flagIgnoreGit {
		ignore = append(ignore, ".git")
	}
	if err := walk(dir, ignore); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func walk(dir string, ignore []string) error {
	basePath, err := filepath.Abs(dir)
	if err != nil {
		return err
	}

	totalHash := sha256.New()

	var currentDir string
	var dirHash hash.Hash
	err = filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
		absolutePath, err := filepath.Abs(path)
		if err != nil {
			return err
		}

		// Open file.
		r, err := os.Open(absolutePath)
		if err != nil {
			return fmt.Errorf("failed to open file: %w", err)
		}
		// Check to see if it's a directory.
		fi, err := r.Stat()
		if err != nil {
			return err
		}
		if fi.IsDir() {
			relDir, err := filepath.Rel(basePath, absolutePath)
			if err != nil {
				return fmt.Errorf("failed to get relative path: %w", err)
			}
			if slices.Contains(ignore, relDir) {
				return filepath.SkipDir
			}
			if currentDir != "" {
				fmt.Printf("%x %s\n\n", dirHash.Sum(nil), currentDir)
			}
			currentDir = relDir
			dirHash = sha256.New()
			return nil
		}

		// Calculate the SHA256 hash value of the file.
		fileHash := sha256.New()
		w := io.MultiWriter(totalHash, dirHash, fileHash)
		_, err = io.Copy(w, r)
		if err != nil {
			return fmt.Errorf("failed to hash file: %w", err)
		}

		// Get the relative path.
		relativePath, err := filepath.Rel(basePath, absolutePath)
		if err != nil {
			return fmt.Errorf("failed to get relative path: %w", err)
		}

		// Print results.
		fmt.Printf("%x %s\n", fileHash.Sum(nil), relativePath)

		return nil
	})
	if currentDir != "" {
		fmt.Printf("%x %s\n\n", dirHash.Sum(nil), currentDir)
	}
	if err != nil {
		return err
	}

	// Print overall hash.
	fmt.Printf("%x .\n", totalHash.Sum(nil))
	return nil
}
