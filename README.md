[![](https://godoc.org/github.com/4ydx/cdp?status.svg)](http://godoc.org/github.com/4ydx/cdp)

# About cdp

Generates structs as well as a few important methods for interacting with the
chrome devtools protocol.  Uses the latest protocol definitions.

Originally from https://github.com/mafredri/cdp but I have revised it to match my own needs.  Thanks mafredri!

# Use

```console
cd cmd/cdpgen/
go build
./run.sh
```

## Installing

```console
go get -u github.com/4ydx/cdp
```

### Updating protocol definitions

```console
cd cmd/cdpgen/
$ ./update.sh
```
