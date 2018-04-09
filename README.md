# Astra
Fetch information about go source file easily.

## Install
``` bash
go get -u github.com/vetcher/go-astra
```

## Description

Package astra use `ast.File` from [standard pkg](http://golang.org/pkg/go/ast) to collect information about source file.
It can collect information about:
* Imports
    * Docs
    * Packages
    * Aliases
* Constants
    * Name
    * Docs
    * Types
* Variables
    * Name
    * Docs
    * Types
    * **Ignore variables which declared by function call!**
* Interfaces
    * Name
    * Docs
    * Functions
* Structures
    * Docs
    * Fields (with tags)
    * Methods
* Functions
    * Name
    * Docs
    * Arguments
    * Results
* Methods (functions with receivers)
    * Name
    * Docs
    * Receiver
    * Arguments
    * Results
    * Linked structure

## Usage example
``` go
package main

import (
	"encoding/json"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"

	"github.com/vetcher/go-astra"
)

func main() {
	currentDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	path := filepath.Join(currentDir, "./test/service.go")
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, path, nil, parser.ParseComments|parser.AllErrors)
	if err != nil {
		panic(fmt.Errorf("error when parse file: %v", err))
	}
	file, err := astra.ParseFile(f)
	if err != nil {
		fmt.Println(err)
	}
	t, err := json.Marshal(file)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(t))
}
```
