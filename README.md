# rustemmer
Golang implementation Porter Stemming for Russian language.

[![Build Status](https://travis-ci.org/liderman/rustemmer.svg?branch=master)](https://travis-ci.org/liderman/rustemmer)

Installation
-----------
	go get github.com/liderman/rustemmer

Usage
-----------
Getting base word:
```go
    wordBase := rustemmer.GetWordBase("вазы")
    // wordBase = "ваз"
```

Normalization of the text:
```go
    text := "г. Москва, ул. Полярная, д. 31А, стр. 1"
    fmt.Print(
        rustemmer.NormalizeText(text),
    )
    // Displays:
    // г Москв ул Полярн д 31А стр 1
```

Requirements
-----------

* Need at least `go1.2` or newer.

Documentation
-----------

You can read package documentation [here](http:godoc.org/github.com/liderman/rustemmer).
