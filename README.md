# csv2md

Stupidly simple tool to convert csv-file to [markdown](https://spec-md.com/) table.

Outputs result in stdout.

## Usage

```shell
csv2md (-h|-help|--help|help)    # to get this help
csv2md example.csv               # convert data from file and write result in stdout
csv2md < example.csv             # convert data from stdin and write result in stdout
cat example.csv | csv2md         # convert data from stdin and write result in stdout
csv2md example.csv > example.md  # convert data from file and write result in new file
csv2md example.csv | less        # convert data from file and write result in stdout using pager
csv2md                           # paste or type data to stdin by hands
                                 # press Ctrl+D to view result in stdout
csv2md > example.md              # paste or type data to stdin by hands
                                 # press Ctrl+D to write result in new file

...anything is possible with redirection and piping
```

> **IMPORTANT:**
> * input must be valid csv
> * whitespaces allowed only between double-quotes

Examples can be found here: https://people.sc.fsu.edu/~jburkardt/data/csv/csv.html

## Compilation

1) [Install go](https://go.dev/learn/).
2) Download this repo via zip or `git clone`.
3) Run `make help` to get help or `go run . <csv_path>` to build and run temporary binary.

## License

[MIT](LICENSE)
