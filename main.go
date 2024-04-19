package main

import (
	"bufio"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func normalizeUserInput(input string) string {
	return strings.ToLower(strings.TrimSpace(input))
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Old name:")
	oldName, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal("Wrong input")
	}
	oldName = normalizeUserInput(oldName)

	fmt.Println("New name:")
	newName, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	newName = normalizeUserInput(newName)

	subDirToSkip := ".git"

	err = filepath.Walk(".", func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
			return err
		}
		if info.IsDir() && info.Name() == subDirToSkip {
			fmt.Printf("skipping a dir without errors: %+v \n", info.Name())
			return filepath.SkipDir
		}

		if info.Name() == oldName {
			arrPath := strings.Split(path, "/")
			arrPath[len(arrPath) - 1] = newName
			fmt.Println(oldName, newName)

			newPath := strings.Join(arrPath, "/")
			fmt.Println(newPath)
			err = os.Rename(path, newPath)
			if err != nil {
				log.Fatal(err)
			}
		}
		fmt.Printf("visited file or dir: %q\n", path)
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
}
