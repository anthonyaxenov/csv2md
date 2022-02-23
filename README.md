# csv2md

Stupidly simple tool to convert csv/tsv to [markdown](https://spec-md.com/) table.

Outputs result in stdout.

## Usage

```shell
csv2md [-help|--help] [-t] [-f <FILE>]
```

Available arguments:
* `-help` or `--help` - get help
* `-f=<FILE>` or `-f <FILE>` - convert specified `FILE`
* `-t` - convert input as tsv

Available `FILE` formats:
* csv (default);
* tsv (with `-t` argument).

Path to `FILE` may be presented as:
* absolute path;
* path relative to current working directory;
* path relative to user home directory (~).

Also if `PATH` contains whitespaces then you should double-quote it.

To save result as separate file you can use piping.

> **IMPORTANT:**
> 1. Input data must be valid csv/tsv
> 2. Whitespaces allowed only between double-quotes
> 3. Due to markdown spec first line of result table will always be presented as header.  
>    So if your raw data hasn't one you'll should add it before conversion or edit later in ready md.

### Examples

```
csv2md                                - paste or type csv to stdin and then
                                        press Ctrl+D to view result in stdout
csv2md -t > example.md                - paste or type tsv to stdin and then
                                        press Ctrl+D to write result in new file
csv2md -f example.csv                 - convert csv from file and view result in stdout
csv2md -t < example.tsv               - convert tsv from stdin and view result in stdout
csv2md -t < example.tsv > example.md  - convert tsv from stdin and write result in new file
cat example.csv | csv2md              - convert csv from stdin and view result in stdout
csv2md -t -f=example.tsv > example.md - convert tsv from file and write result in new file
csv2md -f example.csv | less          - convert csv from file and view result in stdout using pager

...anything is possible with redirection and piping, e.g. grep, sed, awk, etc.
```

You can generate some examples here: [csv](https://onlinerandomtools.com/generate-random-csv), [tsv](https://onlinerandomtools.com/generate-random-tsv)

## Compilation

1) [Install go](https://go.dev/learn/).
2) Download this repo via zip or `git clone`.
3) Run `make help` to get help about compilation or `go run . [ARGS...]` to build and run temporary binary.

## License

[MIT](LICENSE)
