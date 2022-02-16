package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

const VERSION = "1.0.0"

func main() {
	if len(os.Args) < 2 {
		usage()
		os.Exit(1)
	}

	src, err := ExpandPath(os.Args[1])
	if err != nil {
		panic(err)
	}
	if _, err := os.Stat(src); err != nil {
		panic(err)
	}

	result := convert(readCsv(src))
	for _, row := range result {
		fmt.Println(row)
	}
}

func usage() {
	fmt.Fprintln(os.Stderr, "csv2md v" + VERSION)
	fmt.Fprintln(os.Stderr, "Usage: csv2md data.csv > data.md")
}

func ExpandPath(path string) (string, error) {
	homepath, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	newpath, err := filepath.Abs(strings.Replace(path, "~", homepath, 1))
	return newpath, err
}

func readCsv(filePath string) [][]string {
    f, err := os.Open(filePath)
    if err != nil {
        log.Fatal("Unable to read input file " + filePath, err)
    }
    defer f.Close()

    csvReader := csv.NewReader(f)
    records, err := csvReader.ReadAll()
    if err != nil {
        log.Fatal("Unable to parse file as CSV for " + filePath, err)
    }

    return records
}

func convert(records [][]string) []string {
	var result []string
	for idx, row := range records {
		str := "| "
		for _, col := range row {
			str += col + " | "
		}
		result = append(result, str)
		if idx == 0 {
			str := "| "
			for range row {
				str += "--- | "
			}
			result = append(result, str)
		}
	}
	return result
}
