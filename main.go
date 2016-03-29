package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/faiq/dopepope-populate/populate"
)

func main() {
	pwd, _ := os.Getwd()
	fileName := pwd + "/speeches/final.txt"
	fileName, _ = filepath.Abs(fileName)
	lines, err := populate.ReadLines(fileName)
	if err != nil {
		fmt.Println(err)
	}
	err = populate.CleanLinesAndSave(lines)
	if err != nil {
		fmt.Println(err)
	}
}
