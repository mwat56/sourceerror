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

This module offers the `ErrSource` error type that wraps another error instance, along with the file name, line number, and function name where the initial error occurred, along with original error's message text and a call stack.

The public fields should be considered R/O - there really isn't any reason to modify those fields apart from confusing yourself :-)

The fields are as follows:

	- `File`: The source file where the error was encountered.
	- `Function`: The function wherein the error was encountered
	- `Line`: The code line within the `File`.
	- `Stack`: The call stack to where the error was created.

The `ErrSource` type provides the methods `Error()` , `String()` and `Unwrap()` as required by the `error` interface.

The `ErrSource` can be very useful especially during development to help finding problems in the source code.
In case the error call-stacks are not needed just set the `NOSTACK` flag to `true` (which will save some time and memory).

Once the source code is free of avoidable errors, just set the `NODEBUG` flag to `true` – without any need to change the source code otherwise.

## Installation

You can use `Go` to install this package for you:

    go get -u github.com/mwat56/sourceerror@latest

## Usage

It can be used by calling the provided constructor function `New()`:

	import (
		se "github.com/mwat56/sourceerror"
	)

	// ...

	// IF the call-stacks are not needed:
	se.NOSTACK = true

	// uncomment the next line when your code is production ready:
	// se.NODEBUG = true

	// ...

	// here some error occurs:
	err := someFunction()
	if nil != err {
		err = se.New(err, 2)
		// `err` now wraps the original `err` and points
		// two lines up i.e. to the line where the error
		// was encountered.

		return err
		// ... or perform some proper error handling here
	}

	// ...

## Libraries

No external libraries were used building `sourceerror`.

## Licence

        Copyright © 2024, 2025  M.Watermann, 10247 Berlin, Germany
                        All rights reserved
                    EMail : <support@mwat.de>

> This program is free software; you can redistribute it and/or modify it under the terms of the GNU General Public License as published by the Free Software Foundation; either version 3 of the License, or (at your option) any later version.
>
> This software is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.
>
> You should have received a copy of the GNU General Public License along with this program. If not, see the [GNU General Public License](http://www.gnu.org/licenses/gpl.html) for details.

----
[![GFDL](https://www.gnu.org/graphics/gfdl-logo-tiny.png)](http://www.gnu.org/copyleft/fdl.html)
