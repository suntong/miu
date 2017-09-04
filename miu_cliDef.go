////////////////////////////////////////////////////////////////////////////
// Program: miu
// Purpose: Monitor Internet Usage
// Authors: Tong Sun (c) 2017, All rights reserved
////////////////////////////////////////////////////////////////////////////

package main

import (
	"github.com/mkideal/cli"
	clix "github.com/mkideal/cli/ext"
)

////////////////////////////////////////////////////////////////////////////
// Constant and data type/structure definitions

//==========================================================================
// miu

type rootT struct {
	cli.Helper
	Filei     *clix.Reader `cli:"*i,url" usage:"usage URL from ISP"`
	CSS       string       `cli:"*c,css" usage:"css selection for usage text from usage rul"`
	SendTo    string       `cli:"t,to" usage:"email address to send to"`
	Template  string       `cli:"p,template" usage:"mail sending command template string"`
	NoHTML    bool         `cli:"n,no-html" usage:"do not wrap up the output with html tags"`
	DaysShift int          `cli:"d,days" usage:"days to shift from billing date"`
	Verbose   cli.Counter  `cli:"v,verbose" usage:"Verbose mode (Multiple -v increase the verbosity)"`
}

var root = &cli.Command{
	Name:   "miu",
	Desc:   "Monitor Internet Usage\nVersion " + version + " built on " + date,
	Text:   "Tool to monitor Internet usage from ISP",
	Global: true,
	Argv:   func() interface{} { return new(rootT) },
	Fn:     miu,

	NumArg: cli.AtLeast(1),
}

// Template for main starts here
////////////////////////////////////////////////////////////////////////////
// Global variables definitions

//  var (
//          progname  = "miu"
//          version   = "0.1.0"
//          date = "2017-09-04"
//  )

//  var rootArgv *rootT

////////////////////////////////////////////////////////////////////////////
// Function definitions

// Function main
//  func main() {
//  	cli.SetUsageStyle(cli.ManualStyle) // up-down, for left-right, use NormalStyle
//  	//NOTE: You can set any writer implements io.Writer
//  	// default writer is os.Stdout
//  	if err := cli.Root(root,
//  		cli.Tree(defaultDef)).Run(os.Args[1:]); err != nil {
//  		fmt.Fprintln(os.Stderr, err)
//  	}
//  	fmt.Println("")
//  }

// Template for main dispatcher starts here
//==========================================================================
// Main dispatcher

//  func miu(ctx *cli.Context) error {
//  	ctx.JSON(ctx.RootArgv())
//  	ctx.JSON(ctx.Argv())
//  	fmt.Println()

//  	return nil
//  }

// Template for CLI handling starts here

////////////////////////////////////////////////////////////////////////////
// default

//  func defaultCLI(ctx *cli.Context) error {
//  	rootArgv = ctx.RootArgv().(*rootT)
//  	argv := ctx.Argv().(*defaultT)
//  	fmt.Printf("[default]:\n  %+v\n  %+v\n  %v\n", rootArgv, argv, ctx.Args())
//  	return nil
//  }

type defaultT struct {
}

var defaultDef = &cli.Command{
	Name: "default",
	Desc: "Default ISP",
	Text: "Usage:\n  miu default -i http://url/ -c css.sel",
	Argv: func() interface{} { return new(defaultT) },
	Fn:   defaultCLI,

	NumOption: cli.AtLeast(1),
}
