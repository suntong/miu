////////////////////////////////////////////////////////////////////////////
// Program: miu
// Purpose: Monitor Internet Usage
// Authors: Tong Sun (c) 2017, All rights reserved
////////////////////////////////////////////////////////////////////////////

package main

import "os/exec"

func Exec(cmdStr string) {
	cmd := exec.Command("bash", "-c", cmdStr)
	err := cmd.Start()
	abortOn("Execute send", err)
	cmd.Wait()
}
