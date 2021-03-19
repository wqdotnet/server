package cmd

import (
	"bytes"
	"os/exec"
	"runtime"
)

//runtime.GOARCH 返回当前的系统架构；runtime.GOOS 返回当前的操作系统。
//https://github.com/golang/go/wiki/SliceTricks#push-frontunshift 切片示例

// Shellout system shell
func Shellout(command ...string) (string, string, error) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	var cmdtype, cmdarg string
	if runtime.GOOS == "windows" {
		cmdtype = "cmd"
		cmdarg = "/C"
	} else {
		cmdtype = "bash"
		cmdarg = "-c"
	}

	command = append([]string{cmdarg}, command...)

	cmd := exec.Command(cmdtype, command...)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	return stdout.String(), stderr.String(), err

}
