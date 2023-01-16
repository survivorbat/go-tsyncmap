# 🗺️️ Generic sync.Map

[![Go package](https://github.com/survivorbat/go-tsyncmap/actions/workflows/test.yaml/badge.svg)](https://github.com/survivorbat/go-tsyncmap/actions/workflows/test.yaml)
![GitHub](https://img.shields.io/github/license/survivorbat/go-tsyncmap)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/survivorbat/go-tsyncmap)

A wrap around sync.Map that uses generics.

```go
package main

import "github.com/survivorbat/go-tsyncmap"

func main() {
	myMap := tsyncmap.Map[string, string]{}
	
	myMap.Store("a", "b")
	
	myMap.Load("a") // Returns "b"
}
```

## ⬇️ Installation

`go get github.com/survivorbat/go-tsyncmap`

## 📋 Usage

```go
package main

import "github.com/survivorbat/go-tsyncmap"

func main() {
	myMap := tsyncmap.Map[string, string]{}

	myMap.Store("a", "b")

	myMap.Load("a") // Returns "b"
}
```

## 🔭 Plans

Not much.