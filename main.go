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
	is_help := flag.Bool("help", false, "Get help")
	filepath := flag.String("f", "", "File path")
	header := flag.String("h", "", "Add main header (h1)")
	as_tsv := flag.Bool("t", false, "Parse input as tsv")
	aligned := flag.Bool("a", false, "Align columns width")
	flag.Parse()

	// show help
	if *is_help {
		usage(os.Stdout)
		os.Exit(0)
	}

	// if filepath is not provided then convert data from stdin
	if len(*filepath) == 0 {
		data, err := ReadStdin(*as_tsv)
		if err != nil {
			log.Fatal(err)
		}
		Print(Convert(*header, data, *aligned))
	} else { // otherwise convert data from file
		src, err := ExpandPath(*filepath)
		if err != nil {
			log.Fatal(err)
		}
		if _, err := os.Stat(src); err != nil {
			log.Fatal(err)
		}

		data, err := ReadFile(src, *as_tsv)
		if err != nil {
			log.Fatal(err)
		}
		Print(Convert(*header, data, *aligned))
	}
}

// Print writes converted data to stdout
func Print(data []string) {
	if len(data) == 0 {
		usage(os.Stderr)
	}

	for _, row := range data {
		fmt.Println(row)
	}
}

// ExpandPath returns absolute path to file replacing ~ to user's home folder
func ExpandPath(path string) (string, error) {
	homepath, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	newpath, err := filepath.Abs(strings.Replace(path, "~", homepath, 1))
	return newpath, err
}

// ReadFile reads data from file
func ReadFile(filePath string, as_tsv bool) ([][]string, error) {
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

// ReadStdin reads data from stdin
func ReadStdin(as_tsv bool) ([][]string, error) {
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

// Convert formats data from file or stdin as markdown
func Convert(header string, records [][]string, aligned bool) []string {
	var result []string

	// add h1 if passed
	header = strings.Trim(header, "\t\r\n ")
	if len(header) != 0 {
		result = append(result, "# " + header)
		result = append(result, "")
	}

	// if user wants aligned columns width then we
	// count max length of every value in every column
	widths := make(map[int]int)
	if aligned {
		for _, row := range records {
			for col_idx, col := range row {
				length := len(col)
				if len(widths) == 0 || widths[col_idx] < length {
					widths[col_idx] = length
				}
			}
		}
	}

	// build markdown table
	for row_idx, row := range records {

		// table content
		str := "| "
		for col_idx, col := range row {
			if aligned {
				str += fmt.Sprintf("%-*s | ", widths[col_idx], col)
			} else {
				str += col + " | "
			}
		}
		result = append(result, str)

		// content separator only after first row (header)
		if row_idx == 0 {
			str := "| "
			for col_idx := range row {
				if !aligned || widths[col_idx] < 3 {
					str += "--- | "
				} else {
					str += strings.Repeat("-", widths[col_idx]) + " | "
				}
			}
			result = append(result, str)
		}
	}
	return result
}

// usage Print help into writer
func usage(writer *os.File) {
	usage := []string{
		"csv2md v" + VERSION,
		"Anthony Axenov (c) 2022, MIT License",
		"https://github.com/anthonyaxenov/csv2md",
		"",
		"Usage:",
		"\tcsv2md [-help|--help] [-f=<FILE>] [-h=<HEADER>] [-t] [-a]",
		"",
		"Available arguments:",
		"\t-help|--help   - show this help",
		"\t-f=<FILE>      - Convert specified FILE",
		"\t-h=<HEADER>    - add main header (h1) to result",
		"\t-t             - Convert input as tsv",
		"\t-a             - align columns width",
		"",
		"FILE formats supported:",
		"\t- csv (default)",
		"\t- tsv (with -t flag)",
		"",
		"Path to FILE may be presented as:",
		"\t- absolute",
		"\t- relative to current working directory",
		"\t- relative to user home directory (~)",
		"",
		"Both HEADER and FILE path with whitespaces must be double-quoted.",
		"To save result as separate file you should use redirection operators (> or >>).",
	}
	for _, str := range usage {
		fmt.Fprintln(writer, str)
	}
}
