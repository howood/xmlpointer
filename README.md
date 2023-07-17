[![GitHub release](http://img.shields.io/github/release/howood/xmlpointer.svg?style=flat-square)][release]
[![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white&style=flat-square)](https://pkg.go.dev/github.com/howood/xmlpointer)
[![Build Status](https://github.com/howood/xmlpointer/actions/workflows/test.yml/badge.svg?branch=master)][actions]
[![Test Coverage](https://api.codeclimate.com/v1/badges/00e0b66cf675d519a2a8/test_coverage)](https://codeclimate.com/github/howood/xmlpointer/test_coverage)
[![Go Report Card](https://goreportcard.com/badge/github.com/howood/xmlpointer)](https://goreportcard.com/report/github.com/howood/xmlpointer)
[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat-square)][license]

[release]: https://github.com/howood/xmlpointer/releases
[actions]: https://github.com/howood/cryptotools/actions
[license]: https://github.com/howood/xmlpointer/blob/master/LICENSE

# xmlpointer

XMLPointer provides XML Pointers and Decodeing to Map[string]interface{}.

# Install

```
$ go get -u github.com/howood/xmlpointer
```

# Usage

```
	// Create new
	xp, err := NewXMLPointer(xmlDataTest)
	if err != nil {
		...
	}

	// Pointing with key
	xmldata, err := xp.Query("Doc.Body.Item")
	if err != nil {
		...
	}

	// Map
	xmlMap := xp.Data


```