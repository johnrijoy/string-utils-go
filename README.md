# String utils Cli

 A Go Cli which can manipulate string inside files

## Commands

- Find Files
  
    `string-utils findFiles filePath [-f filePath]`

- Find text

    `string-utils regex filePath [-s] [-f filePath]`

- Replace text
  
    `string-utils replaceAll regex replaceText filePath [-s|t] [-f filePath]`

## Sample Commands

- Find files from glob pattern
    
    `go run . findFiles ./*.txt`

- Find files from input Path file

    `go run . findFiles -f "./sampleInput.txt"`

- Find 3 letter words
    
    `go run . findAll "\b\w{3}\b" ./sample.txt`

- Find 3 letter words in all files
    
    `go run . findAll "\b\w{3}\b" ./**`

- Find 3 letter words with input Path file
    
    `go run . findAll -f "\b\w{3}\b" "./sampleInput.txt"`

- Replace how with why
   
    `go run . replaceAll "how" ./sample.txt "why"`
    
    `go run . replaceAll "why" ./sample.txt "how"`

- find using capture groups
    
    `go run . findAll -s "\b(\w)(\w)" ./sample.txt`

- replace using capture groups
    
    `go run . replaceAll "\b(\w)(\w)" "$2$1" ./sample.txt`

- replace using capture groups, input path file and template file
    
    `go run . replaceAll -sft "\b(\w)(\w)" ./sampleTemplate.txt ./sampleInput.txt`


## Building Project

`go build -o bin/string-utils.exe main.go`