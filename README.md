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

type GitHubProjectFile struct {
    Link string `value:"[href]"`
    Name string
}

type GitHubProject struct {
    Files []GitHubProjectFile `$:".js-navigation-item > .content a"`
}

func main() {
    project := &GitHubProject{}
    err := gordom.ParseFromUrl("https://github.com/lucky-libora/go-parse-it", project)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(project)	
}
```


## Testing

Run tests

```sh
./make.sh test
```
