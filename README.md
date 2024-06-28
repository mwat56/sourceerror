# SourceError

[![golang](https://img.shields.io/badge/Language-Go-green.svg)](https://golang.org/)
[![GoDoc](https://godoc.org/github.com/mwat56/sourceerror?status.svg)](https://godoc.org/github.com/mwat56/sourceerror)
[![Go Report](https://goreportcard.com/badge/github.com/mwat56/sourceerror)](https://goreportcard.com/report/github.com/mwat56/sourceerror)
[![Issues](https://img.shields.io/github/issues/mwat56/sourceerror.svg)](https://github.com/mwat56/sourceerror/issues?q=is%3Aopen+is%3Aissue)
[![Size](https://img.shields.io/github/repo-size/mwat56/sourceerror.svg)](https://github.com/mwat56/sourceerror/)
[![Tag](https://img.shields.io/github/tag/mwat56/sourceerror.svg)](https://github.com/mwat56/sourceerror/tags)
[![View examples](https://img.shields.io/badge/learn%20by-examples-0077b3.svg)](https://github.com/mwat56/sourceerror/blob/main/_demo/demo.go)
[![License](https://img.shields.io/github/mwat56/sourceerror.svg)](https://github.com/mwat56/sourceerror/blob/main/LICENSE)

- [SourceError](#sourceerror)
	- [Purpose](#purpose)
	- [Installation](#installation)
	- [Usage](#usage)
	- [Libraries](#libraries)
	- [Licence](#licence)

----

## Purpose

This small module offers the `ErrCodeLocation` error type that wraps another error instance, along with the file name, line number, and function name where the error occurred.
It can be used by calling the provided constructor function:

	// here some error occurs:
	err := someFunction()
	if nil != err {
		err = SourceError(err, 2)
		// `err` now wraps the original 'err` and points 2 lines
		// up i.e. the line where the error appeared.

		return err
		// or perform some proper error handling here
	}

## Installation

You can use `Go` to install this package for you:

    go get -u github.com/mwat56/sourceerror

## Usage

    //TODO

## Libraries

The following external libraries were used building `sourceerror`:

* (none)

## Licence

        Copyright © 2024 M.Watermann, 10247 Berlin, Germany
                        All rights reserved
                    EMail : <support@mwat.de>

> This program is free software; you can redistribute it and/or modify it under the terms of the GNU General Public License as published by the Free Software Foundation; either version 3 of the License, or (at your option) any later version.
>
> This software is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.
>
> You should have received a copy of the GNU General Public License along with this program. If not, see the [GNU General Public License](http://www.gnu.org/licenses/gpl.html) for details.

----
[![GFDL](https://www.gnu.org/graphics/gfdl-logo-tiny.png)](http://www.gnu.org/copyleft/fdl.html)
