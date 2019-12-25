package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"

	"github.com/spf13/cobra"
)

// protobufCmd build protobuf
var protobufCmd = &cobra.Command{
	Use:   "protobuf",
	Short: "Short",
	Long:  `long`,
	Run: func(cmd *cobra.Command, args []string) {

		//protoc --proto_path=d:/proto  --go_out=d:/proto  msg.pro
		//execstr := fmt.Sprintf("protoc --proto_path=%s  --go_out=%s %s", ServerCfg.ProtoPath, ServerCfg.GoOut)

		// fmt.Println(execstr)
		// dir, _ := os.Getwd()
		// // exPath := filepath.Dir(dir)

		// fmt.Println(dir)
		// pathstr, err := filepath.Abs(path.Join(dir, "../proto/"))
		// if err != nil {
		// 	fmt.Println("error:", err)
		// 	return
		// }

		if !PathExists(ServerCfg.ProtoPath) || !PathExists(ServerCfg.GoOut) {
			fmt.Println("文件夹不存在:", ServerCfg.ProtoPath, ServerCfg.GoOut)
			return
		}

		execstr := "protoc --proto_path=%s  --go_out=%s %s"

		files, _ := ioutil.ReadDir(ServerCfg.ProtoPath)
		for _, onefile := range files {
			if !onefile.IsDir() && path.Ext(onefile.Name()) == ".proto" {
				execstr = fmt.Sprintf(execstr, ServerCfg.ProtoPath, ServerCfg.GoOut, onefile.Name())
				_, errout, err := Shellout(execstr)
				if err != nil {
					fmt.Printf("protoc [%s] ==>: %v\n", onefile.Name(), errout)
				} else {
					fmt.Printf("protoc [%s] ==>success", onefile.Name())
				}

			}
		}

	},
}

func init() {
	rootCmd.AddCommand(protobufCmd)

}

// PathExists 判断文件夹是否存在
func PathExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
	// 或者
	//return err == nil || !os.IsNotExist(err)
	// 或者
	//return !os.IsNotExist(err)
}
