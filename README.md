# tGo

[![GoDoc](https://godoc.org/github.com/gotokatsuya/tgo?status.svg)](https://godoc.org/github.com/gotokatsuya/tgo)
[![Go Report Card](http://goreportcard.com/badge/gotokatsuya/tgo)](http://goreportcard.com/report/gotokatsuya/tgo)
[![Build Status](https://travis-ci.org/gotokatsuya/tgo.svg?branch=master)](https://travis-ci.org/gotokatsuya/tgo)

*This package is currently under heavy development and should be used with care.*

Trivago go extensions and utilities.
This is a library package containing tools that aid gotokatsuya with golang development across different projects.

This package and all subpackage match the golang standard library package names along with a "t" prefix.
I.e. type that would be placed in the "net" package can be found in the "tnet" package, etc..
This prefix was chosen to allow mixing standard libary and tgo without having to rename package imports all the time.
