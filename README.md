# csv2md

Stupidly simple tool to convert csv-file to [markdown](https://spec-md.com/) table.

Outputs result in stdout.

## Usage

```shell
csv2md example.csv > example.md # save result to new file
csv2md example.csv | less       # view result using pager

...anything is possible with redirection and piping
```

> **IMPORTANT:**
> * input file must be valid csv
> * whitespaces allowed only between double-quotes

Examples can be found here: https://people.sc.fsu.edu/~jburkardt/data/csv/csv.html

## Compilation

1) [Install go](https://go.dev/learn/).
2) Clone this repo.
3) Run `make help` to get help or `go run . <csv_path>` to build and run temporary binary.

## License

[MIT](LICENSE)
