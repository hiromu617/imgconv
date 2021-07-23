package main

import (
	"os"
	"path/filepath"
	"fmt"
	"flag"
	"imgconv/imgconv"
)

func main() {
	var beforeExt = flag.String("before", "jpeg", "Extension before conversion")
	var afterExt = flag.String("after", "png", "Extension before conversion")
	flag.Parse()
	dirs := flag.Args()

	if *beforeExt == *afterExt {
		fmt.Printf("beforeとafterに同じ拡張子%sが指定されています", *beforeExt)
		return
	}

	for _, dir := range dirs {
		dirwalk(dir, *beforeExt, *afterExt)
	}
}

// ディレクトリの中を巡回する
func dirwalk(dir, beforeExt, afterExt string){
	if _, err := os.Stat(dir); err != nil {
    fmt.Printf("%v\n", err)
		return
	}
	
	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error{
		if filepath.Ext(path) == "." + beforeExt {
			imgconv.Imgconv(path, beforeExt, afterExt)
		}

		return nil
	})
}

