package main

import (
	"encoding/csv"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

const VERSION = "1.2.0"

func main() {
	log.SetFlags(0)
	filepath := flag.String("f", "", "File path")
	is_help := flag.Bool("help", false, "Get help")
	as_tsv := flag.Bool("t", false, "Parse as tsv")
	flag.Parse()

	if *is_help {
		usage(os.Stdout)
		os.Exit(0)
	}

	if *filepath == "" {
		data, err := readRawCsv(*as_tsv)
		if err != nil {
			log.Fatal(err)
		}
		print(data)
	} else {
		src, err := ExpandPath(*filepath)
		if err != nil {
			log.Fatal(err)
		}
		if _, err := os.Stat(src); err != nil {
			log.Fatal(err)
		}

		data, err := readCsvFile(src, *as_tsv)
		if err != nil {
			log.Fatal(err)
		}
		print(data)
	}
}

// print write converted data to stdout
func print(data [][]string) {
	if len(data) == 0 {
		usage(os.Stderr)
	}
	result := convert(data)
	for _, row := range result {
		fmt.Println(row)
	}
}

// ExpandPath return absolute path to file replacing ~ to user's home folder
func ExpandPath(path string) (string, error) {
	homepath, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	newpath, err := filepath.Abs(strings.Replace(path, "~", homepath, 1))
	return newpath, err
}

// readRawCsv read data from file
func readCsvFile(filePath string, as_tsv bool) ([][]string, error) {
    f, err := os.Open(filePath)
    if err != nil {
    	return nil, errors.New("Failed to open file '" + filePath + "': " + err.Error())
    }
    defer f.Close()

    csvReader := csv.NewReader(f)
	if as_tsv {
		csvReader.Comma = '\t'
	}

    records, err := csvReader.ReadAll()
    if err != nil {
    	return nil, errors.New("Failed to parse file '" + filePath + "': " + err.Error())
    }

    return records, nil
}

// readRawCsv read data from stdin
func readRawCsv(as_tsv bool) ([][]string, error) {
    csvReader := csv.NewReader(os.Stdin)
	if as_tsv {
		csvReader.Comma = '\t'
	}

    records, err := csvReader.ReadAll()
    if err != nil {
    	return nil, errors.New("Failed to parse input from stdin: " + err.Error())
    }

    return records, nil
}

// convert format data from file or stdin as markdown
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

// usage print help into writer
func usage(writer *os.File) {
	usage := []string{
		"csv2md v" + VERSION,
		"Anthony Axenov (c) 2022, MIT License",
		"https://github.com/anthonyaxenov/csv2md",
		"",
		"Usage:",
		"\tcsv2md [-help|--help] [-t] [-f <FILE>]",
		"",
		"Available arguments:",
		"\t-help|--help        - get this help",
		"\t-f=<FILE>|-f <FILE> - convert specified FILE",
		"\t-t                  - convert input as TSV",
		"",
		"FILE formats supported:",
		"\t- csv (default)",
		"\t- tsv (with -t flag)",
		"",
		"Path to FILE may be presented as:",
		"\t- absolute",
		"\t- relative to current working directory",
		"\t- relative to user home directory (~)",
	}
	for _, str := range usage {
		fmt.Fprintln(writer, str)
	}
}
