[![Build Status](https://travis-ci.org/howood/xmlpointer.svg?branch=master)](https://travis-ci.org/howood/xmlpointer)
[![GitHub release](http://img.shields.io/github/release/howood/xmlpointer.svg?style=flat-square)][release]
[![godoc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](http://godoc.org/github.com/howood/xmlpointer)
[![Test Coverage](https://api.codeclimate.com/v1/badges/00e0b66cf675d519a2a8/test_coverage)](https://codeclimate.com/github/howood/xmlpointer/test_coverage)
[![Go Report Card](https://goreportcard.com/badge/github.com/howood/xmlpointer)](https://goreportcard.com/report/github.com/howood/xmlpointer)
[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat-square)][license]

[release]: https://github.com/howood/xmlpointer/releases
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