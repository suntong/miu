////////////////////////////////////////////////////////////////////////////
// Program: miu
// Purpose: Monitor Internet Usage
// Authors: Tong Sun (c) 2017, All rights reserved
////////////////////////////////////////////////////////////////////////////

package main

import (
	"fmt"
	"io"

	"github.com/andybalholm/cascadia"
	"golang.org/x/net/html"
)

// Cascadia filters the input buffer/stream `bi` with CSS selectors `css` and write to the output buffer/stream `bw`.
func Cascadia(bi io.Reader, bw io.Writer, css string) error {
	doc, err := html.Parse(bi)
	abortOn("Input", err)
	c, err := cascadia.Compile(css)
	abortOn("CSS Selector string "+css, err)

	if !Opts.NoHTML {
		fmt.Fprintf(bw, `<!DOCTYPE html>
<html>
  <head>
    <meta charset="utf-8">
    <base href="%s">
  </head>
<body>
`, Opts.Base)
	}

	// https://godoc.org/github.com/andybalholm/cascadia
	ns := c.MatchAll(doc)
	verbose(1, "%d elements for '%s':\n", len(ns), css)
	for _, n := range ns {
		html.Render(bw, n)
		fmt.Fprintf(bw, "\n")
	}
	if !Opts.NoHTML {
		fmt.Fprintln(bw, `</body>`)
	}
	return nil
}
