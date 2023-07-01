# Gout Web Framework

## Features

* 路由支持GET、POST、DELETE、PUT功能
* 实现Context功能
* 嵌入log、cors、recovery等middleware 

## Getting Started

### Prerequisites
* **[Go](https://go.dev/)** (we test it with 1.19.2).

### Getting Gout
With [Go module](https://github.com/golang/go/wiki/Modules) support, simply add the following import
~~~
import "github.com/SkyDDDog/gout"
~~~
to your code, and then `go [build|run|test]` will automatically fetch the necessary dependencies.

Otherwise, run the following Go command to install the `gout` package:
```sh
$ go get -u github.com/SkyDDDog/gout
```
### Running Gin

First you need to import Gin package for using Gin, one simplest example likes the follow `example.go`:
