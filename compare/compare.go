package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

var dirOne map[string]int
var dirTwo map[string]int

func main() {
	if len(os.Args) != 3 {
		fmt.Println("please input FIRST_DIRECTORY and SECOND_DIRECTORY")
		return
	}
	srcDir := os.Args[1]
	dstDir := os.Args[2]

	dirOne = make(map[string]int)
	dirTwo = make(map[string]int)

	fmt.Println("First Directory")
	dirOne = lookDir(srcDir)
	fmt.Println("Second Directory")
	dirTwo = lookDir(dstDir)

	cmprDir(dirOne, dirTwo)
}

func cmprDir(src, dst map[string]int) {
	for dir, size := range src {
		val, ok := dst[dir]
		if !ok {
			fmt.Println(dir, "NEW")
		} else if size != val {
			fmt.Println(dir, "MODIFIED")
		}
	}
	for dir := range dst {
		_, ok := src[dir]
		if !ok {
			fmt.Println(dir, "DELETED")
		}
	}
}

func lookDir(dirAddress string) (m map[string]int) {
	m = make(map[string]int)
	err := filepath.Walk(dirAddress,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			// fmt.Println(path, info.Size())
			m[path] = int(info.Size())
			return nil
		})
	if err != nil {
		log.Println(err)
	}
	return m
}
