[![Go Report Card][goreportcard-badge]][goreportcard] [![GoDoc][godoc-badge]][godoc]

# About cdp

Generates structs as well as a few important methods for interacting with the
chrome devtools protocol.  Uses the latest protocol definitions.

Originally from https://github.com/mafredri/cdp but I have revised it to match my own needs.  Thanks mafredri!

# Use

cd cmd/cdpgen/
./run.sh

## Installing

```console
go get -u github.com/4ydx/cdp
```

### Updating protocol definitions

```console
$ ./update.sh
```
