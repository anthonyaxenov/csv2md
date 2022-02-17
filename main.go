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

const VERSION = "1.1.0"

func main() {
	log.SetFlags(0)
	switch len(os.Args) {
		case 1: // first we read data from stdin and then convert it
			data, err := readRawCsv()
			if err != nil {
				log.Fatal(err)
			}
			print(data)

		case 2: // but if 2nd arg is present
			// probably user wants to get help
			help1 := flag.Bool("h", false, "Get help")
			help2 := flag.Bool("help", false, "Get help")
			flag.Parse()
			if os.Args[1] == "help" || *help1 || *help2 {
				usage(os.Stdout)
				os.Exit(0)
			}

			// ...or to convert data from file
			src, err := ExpandPath(os.Args[1])
			if err != nil {
				log.Fatal(err)
			}
			if _, err := os.Stat(src); err != nil {
				log.Fatal(err)
			}

			data, err := readCsvFile(src)
			if err != nil {
				log.Fatal(err)
			}
			print(data)

		// otherwise let's show usage help and exit (probably inaccessible code, but anyway)
		default:
			usage(os.Stdout)
			os.Exit(0)
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
func readCsvFile(filePath string) ([][]string, error) {
    f, err := os.Open(filePath)
    if err != nil {
    	return nil, errors.New("Failed to open file '" + filePath + "': " + err.Error())
    }
    defer f.Close()

    csvReader := csv.NewReader(f)
    records, err := csvReader.ReadAll()
    if err != nil {
    	return nil, errors.New("Failed to parse CSV from file '" + filePath + "': " + err.Error())
    }

    return records, nil
}

// readRawCsv read data from stdin
func readRawCsv() ([][]string, error) {
    csvReader := csv.NewReader(os.Stdin)
    records, err := csvReader.ReadAll()
    if err != nil {
    	return nil, errors.New("Failed to parse CSV from stdin: " + err.Error())
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
		"Anthony Axenov (c) 2022, MIT license",
		"https://github.com/anthonyaxenov/csv2md",
		"",
		"Usage:",
		"\tcsv2md (-h|-help|--help|help)    # to get this help",
		"\tcsv2md example.csv               # convert data from file and write result in stdout",
		"\tcsv2md < example.csv             # convert data from stdin and write result in stdout",
		"\tcat example.csv | csv2md         # convert data from stdin and write result in stdout",
		"\tcsv2md example.csv > example.md  # convert data from file and write result in new file",
		"\tcsv2md example.csv | less        # convert data from file and write result in stdout using pager",
		"\tcsv2md                           # paste or type data to stdin by hands",
		"\t                                 # press Ctrl+D to view result in stdout",
		"\tcsv2md > example.md              # paste or type data to stdin by hands",
		"\t                                 # press Ctrl+D to write result in new file",
	}
	for _, str := range usage {
		fmt.Fprintln(writer, str)
	}
}
