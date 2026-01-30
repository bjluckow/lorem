# lorem

A tiny Go CLI that generates a stream of lorem ipsum text.

## Install

### Go users

```
go install github.com/bjluckow/lorem@latest
```

Make sure `$GOBIN` or `$HOME/go/bin` is in your PATH.

### Build from source

```
git clone https://github.com/bjluckow/lorem
cd lorem
go build -o lorem .
```

---

## Usage

```
lorem -w WORDS -n LINES [-d DELIMITER]
```

### Flags

```
-n int
    number of lines (required, must be > 0, default 10)

-w int
    words per line (default 10)

-d string
    line delimiter (default "\n")
```

---

## Examples

```
lorem -w 4 -n 2
```

```
lorem ipsum dolor sit
amet consectetur adipiscing elit
```

Custom delimiter:

```
lorem -w 3 -n 2 -d "|"
```

```
lorem ipsum dolor|sit amet consectetur|
```

Pipe into other tools:

```
lorem -w 5 -n 100 | head
```
