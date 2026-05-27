# extname

`extname` is a pathname utility in the style of `basename`, `dirname`, and `realpath`. It writes the extension of each pathname to standard output.

## Installation

```shell
go install github.com/CatsDeservePets/extname@latest
```

## Usage

```
usage: extname [-a] [-d] [-z] string [...]
  -a	print all extension segments (e.g. .tar.gz)
  -d	print the extension without a leading dot
  -z	end each output line with NUL, not newline
```

## Semantics

By default, `extname` matches [PurePath.suffix](https://docs.python.org/3/library/pathlib.html#pathlib.PurePath.suffix) in Python. With `-a`, it matches `''.join(PurePath(path).suffixes)`.

`extname` does not inspect the filesystem. It determines the extension from the final path element, so paths like `archive.tar.gz/` can still return an extension. Leading dots do not count as extensions by themselves, so *dotfiles* such as `.bashrc` produce an empty result.

### Example output

```shell
$ extname main.go
.go
$ extname archive.tar.gz
.gz
$ extname -a archive.tar.gz
.tar.gz
$ extname -a -d archive.tar.gz
tar.gz
$ extname /tmp/archive.tar.gz/
.gz
$ extname .bashrc

```
