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
	Short: "protobuf 协议生成",
	Long:  `long`,
	Run: func(cmd *cobra.Command, args []string) {
		var filetime int64 = 0

		if len(args) > 0 {
			i, err := strconv.ParseInt(args[0], 10, 64)
			if err == nil {
				filetime = i
			}
		}

		pbpath := gserver.ServerCfg.ProtoPath
		outpath := gserver.ServerCfg.GoOut
		timeformat := tool.DateTimeFormat

		if !PathExists(pbpath) || !PathExists(outpath) {
			fmt.Println("文件夹不存在:", pbpath, outpath)
			return
		}
		//execstr := "protoc  --proto_path=%s   --go_out=../ %s"
		execstr := "protoc -o %s/%s.pb  --proto_path=%s   --go_out=../ %s"

		files, _ := ioutil.ReadDir(pbpath)
		for _, onefile := range files {
			filename := onefile.Name()
			filebase := filename[0 : len(filename)-len(path.Ext(filename))]
			if !onefile.IsDir() && path.Ext(filename) == ".proto" {

				diff := getHourDiffer(onefile.ModTime().Format(timeformat), time.Now().Format(timeformat))
				if filetime == 0 || diff < 60*filetime {

					//execstrpro := fmt.Sprintf(execstr, pbpath, filename)
					execstrpro := fmt.Sprintf(execstr, outpath, filebase, pbpath, filename)

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
	// 或者
	//return err == nil || !os.IsNotExist(err)
	// 或者
	//return !os.IsNotExist(err)
}

//获取相差时间
func getHourDiffer(startTime, endTime string) int64 {
	var hour int64
	t1, err := time.ParseInLocation(tool.DateTimeFormat, startTime, time.Local)
	t2, err := time.ParseInLocation(tool.DateTimeFormat, endTime, time.Local)

	if err == nil && t1.Before(t2) {
		diff := t2.Unix() - t1.Unix() //
		hour = diff
		return hour
	}
	return hour
}
