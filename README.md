# String utils Cli

 A Go Cli which can manipulate string inside files

## Sample Commands

- Find 3 letter words
    `go run . findAll "\b\w{3}\b" ./sample.txt`

- Find 3 letter words in all files
    `go run . findAll "\b\w{3}\b" ./**`

- Find files from glob pattern
    `go run . findFiles ./*.txt`

- Replace how with why
    `go run . replaceAll "how" ./sample.txt "why"`
    `go run . replaceAll "why" ./sample.txt "how"`

- find using capture groups
    `go run . findAll -s "\b(\w)(\w)" ./sample.txt`

- replace using capture groups
    `go run . replaceAll "\b(\w)(\w)" ./sample.txt "$2$1"`

## Building Project

go build -o bin/string-utils.exe main.go