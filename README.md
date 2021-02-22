# go-cipher-text

## How to build
- `go build encrypt.go util.go`
- `go build decrypt.go util.go`

## How to run
`k` is key, `t` is text
- `./encrypt -t 'Hello, World!' -k 3` prints `Khoor, Zruog!`
- `./decrypt -t 'Khoor, Zruog!' -k 3` prints `Hello, World!`
