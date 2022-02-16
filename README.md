# csv2md

Stupidly simple tool to convert csv-file to [markdown](https://spec-md.com/) table.

Outputs result in stdout.

Building:

```shell
make help
```

Usage:

```shell
csv2md example.csv > example.md # makes new file
csv2md example.csv | less       # view result using *pager*
...anything is possible with redirection and piping
```

> **IMPORTANT:** input must be valid csv and whitespaces are allowed only between double-quotes.

Examples can be found here: https://people.sc.fsu.edu/~jburkardt/data/csv/csv.html
