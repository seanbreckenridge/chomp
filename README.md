## chomp

```
usage: chomp [-h] [FLAGS]

Receives input from STDIN
Removes extra whitespace characters from the beginning/end of lines
Removes lines which have just whitespace (no content)

Flags:
  -max-capacity int
    	Maximum capacity for each line in kilobytes (default 64)
```

This is essentially a portable (and 3x faster):

```
sed -E -e 's/^\s*//; s/\s*$//; /^\s*$/d'
```

[I use this in scripts](https://gist.github.com/seanbreckenridge/02bf00bc50b3ad6a35088fb75e41e9e6) when trying to remove spaces from user input/command output. Its also often helpful when data wrangling, to be able to quickly ignore lines/spaces I'm not interested in.

### Install

Using `go install` to put it on your `$GOBIN`:

`go install github.com/seanbreckenridge/chomp@latest`

Manually:

```bash
git clone https://github.com/seanbreckenridge/chomp
cd ./chomp
go build .
# copy binary somewhere on your $PATH
sudo cp ./chomp /usr/local/bin
```

### Example

Typically this would be used by piping some command into it:

```
$ man -P cat rm | head
RM(1)               User Commands              RM(1)

NAME
       rm - remove files or directories

SYNOPSIS
       rm [OPTION]... [FILE]...

DESCRIPTION
       This manual page documents the GNU version of
```

```
$ man -P cat rm | head | chomp
RM(1)               User Commands              RM(1)
NAME
rm - remove files or directories
SYNOPSIS
rm [OPTION]... [FILE]...
DESCRIPTION
This manual page documents the GNU version of
```
