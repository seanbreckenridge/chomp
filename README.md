### chomp

```
usage: chomp [-h]

Receives input from STDIN
Removes extra whitespace characters from the beginning/end of lines
Removes lines which have just whitespace (no content)
```

This is essentially a portable:

```
sed -E -e 's/^\s*//; s/\s*$//; /^\s*$/d'
```

I use this in scripts when trying to remove spaces from user input/command output.

---

To install:

Using `go get` to put it on your `$GOBIN`:

`go get github.com/seanbreckenridge/chomp`

Manually:

```bash
git clone https://github.com/seanbreckenridge/chomp
cd ./chomp
go build .
# copy binary somewhere on your $PATH
sudo cp ./chomp /usr/local/bin
```
