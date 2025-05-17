package main

import (
	"fmt"
	"io"
)

const greeting = "auto-tune pg vacuum"

func Run(out io.Writer) {
	fmt.Fprint(out, greeting)
}
