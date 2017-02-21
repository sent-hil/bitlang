# runeio

runeio is a library that provides functions to help work with runes from an
underlying io.Reader.

## Getting started

    // See https://github.com/sent-hil/bitlang/blob/master/runeio/runeio.go#L4
    // for `RuneReader` interface.
    //
    // `bufio.Reader`, `bytes.Reader` and `strings.Reader` all implement the
    // interface  and can be used here.
    buf := bufio.NewStringReader("Hello World")
    runeio.NewRuneio(buf)

## Godoc

    https://godoc.org/github.com/sent-hil/bitlang/runeio

## Install

    go get -u github.com/sent-hil/bitlang/runeio
