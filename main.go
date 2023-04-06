/*
In this program, we first define two slices: one for the filenames of the files we want to copy,
and one for the names of the folders we want to copy. We then use a sync.WaitGroup to ensure that
all goroutines finish before the program exits.

We use two for loops to iterate over both slices. For each file, we use a goroutine to copy the file
to a destination folder, and for each folder, we use a goroutine to copy the entire folder and its
contents to a destination folder.

Each goroutine is defined as an anonymous function that takes the name of the file or folder as an
argument, allowing us to use the go keyword to create a separate concurrent thread of execution for
each copy operation.

After all the goroutines have finished, we print a message to the console indicating that all files
and folders have been copied.

The copyFolder function is used to recursively copy a folder and its contents. We use the os.
Stat function to get information about the source folder, and the os.MkdirAll function to create
the destination folder with the same permissions as the source folder.

We use the os.ReadDir function to get a list of all the entries in the source folder, and then
iterate over them using a for loop. For each entry, we check whether it is a file or directory
using the entry.IsDir method. If it is a directory, we recursively call copyFolder with the source
and destination paths. If it is a file, we use the io.Copy function to copy the file to the destination path.
*/
package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	// Get source and destination directories from command line arguments
	if len(os.Args) < 3 {
		fmt.Println("Usage: conc_copy <source_dir> <dest_dir>")
		return
	}
	srcDir := os.Args[1]
	destDir := os.Args[2]

	// Get list of files and folders to copy
	files, err := getFilesInDir(srcDir)
	if err != nil {
		fmt.Println(err)
		return
	}
	folders, err := getFoldersInDir(srcDir)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Copy files concurrently
	for _, file := range files {
		wg.Add(1)
		go func(filename string) {
			defer wg.Done()

			src, err := os.Open(filepath.Join(srcDir, filename))
			if err != nil {
				fmt.Println(err)
				return
			}
			defer src.Close()

			dst, err := os.Create(filepath.Join(destDir, filename))
			if err != nil {
				fmt.Println(err)
				return
			}
			defer dst.Close()

			_, err = io.Copy(dst, src)
			if err != nil {
				fmt.Println(err)
				return
			}

			fmt.Println("File copied:", filename)
		}(file)
	}

	// Copy folders concurrently
	for _, folder := range folders {
		wg.Add(1)
		go func(foldername string) {
			defer wg.Done()

			err := copyFolder(filepath.Join(srcDir, foldername), filepath.Join(destDir, foldername))
			if err != nil {
				fmt.Println(err)
				return
			}

			fmt.Println("Folder copied:", foldername)
		}(folder)
	}

	// Wait for all goroutines to finish
	wg.Wait()
	fmt.Println("All files and folders copied.")
}

func copyFolder(src string, dest string) error {
	info, err := os.Stat(src)
	if err != nil {
		return err
	}

	err = os.MkdirAll(dest, info.Mode())
	if err != nil {
		return err
	}

	entries, err := os.ReadDir(src)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		srcPath := filepath.Join(src, entry.Name())
		destPath := filepath.Join(dest, entry.Name())

		if entry.IsDir() {
			err := copyFolder(srcPath, destPath)
			if err != nil {
				return err
			}
		} else {
			srcFile, err := os.Open(srcPath)
			if err != nil {
				return err
			}
			defer srcFile.Close()

			destFile, err := os.Create(destPath)
			if err != nil {
				return err
			}
			defer destFile.Close()

			_, err = io.Copy(destFile, srcFile)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func getFilesInDir(dir string) ([]string, error) {
	var files []string

	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	for _, entry := range entries {
		if !entry.IsDir() {
			files = append(files, entry.Name())
		}
	}

	return files, nil
}

func getFoldersInDir(dir string) ([]string, error) {
	var folders []string

	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	for _, entry := range entries {
		if entry.IsDir() {
			folders = append(folders, entry.Name())
		}
	}

	return folders, nil
}
