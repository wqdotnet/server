package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"server/gserver"
	"server/tool"
	"strconv"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// protobufCmd build protobuf
var protobufCmd = &cobra.Command{
	Use:   "protobuf",
	Short: "protobuf [int] [obj] ",
	Long:  `protobuf [只生成 int 分钟内修改文件] [生成客户端pb]`,
	Run: func(cmd *cobra.Command, args []string) {
		var filetime int64 = 0

		if len(args) > 0 {
			i, err := strconv.ParseInt(args[0], 10, 64)
			if err == nil {
				filetime = i
			}
		}

		isclientpb := len(args) == 2

		pbpath := gserver.ServerCfg.ProtoPath
		outpath := gserver.ServerCfg.GoOut
		timeformat := tool.DateTimeFormat

		if !PathExists(pbpath) || !PathExists(outpath) {
			fmt.Println("文件夹不存在:", pbpath, outpath)
			return
		}

		//输出地址
		execstr := ""
		if isclientpb {
			execstr = "protoc -o %s/%s.pb  --proto_path=%s   --go_out=../ %s"
		} else {
			execstr = "protoc  --proto_path=%s   --go_out=../ %s"
		}

		files, _ := ioutil.ReadDir(pbpath)
		for _, onefile := range files {
			filename := onefile.Name()
			filebase := filename[0 : len(filename)-len(path.Ext(filename))]
			if !onefile.IsDir() && path.Ext(filename) == ".proto" {

				diff := getHourDiffer(onefile.ModTime().Format(timeformat), time.Now().Format(timeformat))
				if filetime == 0 || diff < 60*filetime {

					execstrpro := ""
					if isclientpb {
						execstrpro = fmt.Sprintf(execstr, outpath, filebase, pbpath, filename)
					} else {
						execstrpro = fmt.Sprintf(execstr, pbpath, filename)
					}

					_, errout, err := Shellout(execstrpro)
					if err != nil {
						log.Errorf("protoc [%s] ==>: %v errout:%v", filename, err, errout)
					} else {
						log.Infof("protoc [%s] ==> success", filename)
					}
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
	//return err == nil || !os.IsNotExist(err)
	//return !os.IsNotExist(err)
}

//获取相差时间
func getHourDiffer(startTime, endTime string) int64 {
	var hour int64
	t1, err := time.ParseInLocation(tool.DateTimeFormat, startTime, time.Local)
	t2, err2 := time.ParseInLocation(tool.DateTimeFormat, endTime, time.Local)

	if err == nil && err2 == nil && t1.Before(t2) {
		diff := t2.Unix() - t1.Unix() //
		hour = diff
		return hour
	}
	return hour
}
