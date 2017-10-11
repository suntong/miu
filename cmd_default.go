////////////////////////////////////////////////////////////////////////////
// Program: miu
// Purpose: Monitor Internet Usage
// Authors: Tong Sun (c) 2017, All rights reserved
////////////////////////////////////////////////////////////////////////////

package main

import (
	"bytes"
	"io/ioutil"
	"os"
	"regexp"

	"github.com/mkideal/cli"
)

////////////////////////////////////////////////////////////////////////////
// default

func defaultCLI(ctx *cli.Context) error {
	rootArgv = ctx.RootArgv().(*rootT)
	// argv := ctx.Argv().(*defaultT)
	// fmt.Printf("[default]:\n  %+v\n  %+v\n  %v\n", rootArgv, argv, ctx.Args())
	Opts.Template, Opts.SendTo, Opts.Base, Opts.NoHTML, Opts.Verbose =
		rootArgv.Template, rootArgv.SendTo,
		rootArgv.Base, rootArgv.NoHTML, rootArgv.Verbose.Value()
	if Opts.Base == "" {
		Opts.Base = regexp.MustCompile(`(?i)https?://.*?/`).
			FindString(rootArgv.Filei.Name())
	}

	fileo, err := ioutil.TempFile(os.TempDir(), progname+".tmp.")
	abortOn("Create temp file", err)
	defer os.Remove(fileo.Name())
	Opts.TempFile = fileo.Name()

	Cascadia(rootArgv.Filei, fileo, rootArgv.CSS)

	// Only proceed when need to send out email
	if Opts.SendTo == "" || Opts.Template == "" {
		verbose(1, "Sending email step skipped")
		return nil
	}

	buf := bytes.NewBufferString("")
	processTemplate(Opts.Template, buf, Opts)
	// fmt.Printf("] %s\n", buf.String())
	Exec(buf.String())

	return nil
}
