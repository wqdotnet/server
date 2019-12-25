package cmd

import (
	"bytes"
	"fmt"
	"os/exec"
	"runtime"

	"github.com/spf13/cobra"
)

// protobufCmd build protobuf
var protobufCmd = &cobra.Command{
	Use:   "protobuf",
	Short: "Short",
	Long:  `long`,
	Run: func(cmd *cobra.Command, args []string) {
		//runtime.GOARCH 返回当前的系统架构；runtime.GOOS 返回当前的操作系统。
		//protoc --proto_path=d:/proto  --go_out=d:/proto  msg.pro
		//https://github.com/golang/go/wiki/SliceTricks#push-frontunshift 切片示例
		out, errout, err := Shellout(args...)
		if err != nil {
			fmt.Printf("error: %v\n", err)
		}
		if out != "" {
			fmt.Println("--- stdout ---")
			fmt.Println(out)
		}
		if errout != "" {
			fmt.Println("--- stderr ---")
			fmt.Println(errout)
		}
	},
}

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

func init() {
	rootCmd.AddCommand(protobufCmd)

}
