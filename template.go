////////////////////////////////////////////////////////////////////////////
// Program: miu
// Purpose: Monitor Internet Usage
// Authors: Tong Sun (c) 2017, All rights reserved
////////////////////////////////////////////////////////////////////////////

package main

import (
	"io"
	"text/template"
)

func processTemplate(tmplStr string, bw io.Writer, input interface{}) {
	tmpl, err := template.New("").Parse(tmplStr)
	abortOn("Parse template", err)
	// fmt.Printf("> %#v\n", input)
	err = tmpl.Execute(bw, input)
	abortOn("Executing template", err)
}
