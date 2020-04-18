# Go Parse It 


Go-Parse-It is a simple html parser written in Go (Golang). It features a declarative way to parse html pages.


## Contents

- [Installation](#installation)
- [Quick start](#quick-start)
- [Testing](#testing)


## Installation

To install Gin package, you need to install Go and set your Go workspace first.

1. The first need [Go](https://golang.org/) installed, then you can use the below Go command to install Gin.

```sh
$ go get -u github.com/lucky-libora/gordom
```

2. Import it in your code:

```go
import "github.com/lucky-libora/gordom"
```

## Quick start

```go
package main

import (
    "fmt"
    "github.com/lucky-libora/gordom"
)

// gordom maps dom data to struct fields'
// css query passed by $ tag 
type GitHubProjectFile struct {
    // by default field value is set from a text of html node
    Name string
    // also it can be set from an attribute value by adding tag value:"[attribute_name]"
    Link string `value:"[href]"`
}

type GitHubProjectInfo struct {
    // get node by the attribute itemprop='about' and set node's text to the field
    Description string `$:"[itemprop='about']"`
    // int, uint, float types are automatically converted
    Commits     int8   `$:".commits .num"`
}

type GitHubProject struct {
    // structures are also supported. You can select a parent node for inner structure if it necessary
    Info GitHubProjectInfo `$:".repository-content "`
    // you can iterate even arrays
    Files []GitHubProjectFile `$:".js-navigation-item > .content a"`
    // WARNING: pointers can't be used as field types
}

func main() {
    project := &GitHubProject{}
    // pass pointer to your structure. Gordom will fill all fields
    err := gordom.ParseFromUrl("https://github.com/lucky-libora/gordom", project)
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println(project)	
}
```


## Testing

Run tests

```sh
./make.sh test
```
