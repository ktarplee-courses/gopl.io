// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 173.

// Bytecounter demonstrates an implementation of io.Writer that counts bytes.
package main

import (
	"fmt"
	"io"
	"os"
)

//!+bytecounter

type ByteCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p)) // convert int to ByteCounter
	return len(p), nil
}

//!-bytecounter

func main() {
	//!+main
	var c ByteCounter
	c.Write([]byte("hello"))
	fmt.Println(c) // "5", = len("hello")

	var _ io.Writer = &c

	c = 0 // reset the counter
	var name = "Dolly"
	fmt.Fprintf(io.MultiWriter(&c, os.Stdout), "hello, %s", name)
	fmt.Println(c) // "12", = len("hello, Dolly")
	//!-main

	/*
		echo "hello, Dolly" | tee file | wc -c
	*/

	// io.Pipe() is like the shell pipe operator
	// echo "Hello" | wc -c

	// combinedR := io.MultiReader(r1, r2, r3) is like `cat`

	// io.TeeReader is somewhat like `tee`

}
