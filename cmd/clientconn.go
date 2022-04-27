/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"context"
	"encoding/binary"
	"fmt"
	"net"
	"server/network"
	"server/proto/account"
	"server/proto/protocol_base"
	"server/tools"
	"strconv"
	"sync"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"google.golang.org/protobuf/proto"
)

var wg sync.WaitGroup

// clientconnCmd represents the clientconn command
var clientconnCmd = &cobra.Command{
	Use:   "clientconn",
	Short: "模拟客户端连接",
	Long:  `模拟客户端连接  args: 连接数量`,
	Run: func(cmd *cobra.Command, args []string) {
		num := 2

		wg = sync.WaitGroup{}

		if len(args) == 1 {
			num, _ = strconv.Atoi(args[0])
			num++
		}
		wg.Add(num - 1)
		fmt.Println(num)
		for i := 1; i < num; i++ {
			go conn(i)
		}
		wg.Wait()
	},
}

func init() {
	rootCmd.AddCommand(clientconnCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// clientconnCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// clientconnCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func conn(key int) {
	defer wg.Done()

	tcpaddr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:3344")
	conn, err := net.DialTCP("tcp", nil, tcpaddr)
	if err != nil {
		fmt.Println("Dial failed:", err)
		return
	}
	defer conn.Close()

	sendctx, sendcancelFunc := context.WithCancel(context.Background())
	defer sendcancelFunc()
	sendchan := make(chan []byte, 1)

	accountname := fmt.Sprintf("%v_%v", "test", key)
	password := "123456"

	go func() {
		defer sendcancelFunc()
		for {
			_, buf, err := network.UnpackToBlockFromReader(conn, 2)
			if err != nil {
				fmt.Println("read failed:", err)
				return
			}
			//module := int32(binary.BigEndian.Uint16(buf[2:4]))
			method := int32(binary.BigEndian.Uint16(buf[4:6]))
			msgbuf := buf[6:]

			switch method {
			case int32(account.MSG_ACCOUNT_Login):
				msg := &account.S2C_Login{}
				proto.Unmarshal(msgbuf, msg)
				fmt.Printf("S2C_Login: [%v]\n", msg)

				if msg.Retcode == 100005 {
					//角色为空， 创建角色
					SendToClient(int32(account.MSG_ACCOUNT_Module),
						int32(account.MSG_ACCOUNT_CreateRole),
						&account.C2S_CreateRole{
							RoleName: accountname,
							Sex:      1,
							HeadID:   23,
						}, sendchan)
				} else if msg.Retcode == 100004 {
					//注册账号
					SendToClient(int32(account.MSG_ACCOUNT_Module),
						int32(account.MSG_ACCOUNT_Register),
						&account.C2S_Register{
							Account:   accountname,
							Password:  password,
							Source:    accountname,
							Equipment: "pc",
							CDK:       "",
						}, sendchan)
				} else if msg.Retcode == 100002 {
					fmt.Println("密码错误:", msg.Retcode)
					return
				} else if msg.Retcode == 0 {
					fmt.Println("登陆成功:", msg.Retcode)
					// time.Sleep(time.Second * 1)
					// SendToClient(int32(account.MSG_ACCOUNT_Module),
					// 	int32(account.MSG_ACCOUNT_Ping),
					// 	&account.C2S_Ping{}, sendchan)
				}
			// case int32(account.MSG_ACCOUNT_Ping):
			// 	msg := &account.S2C_Ping{}
			// 	proto.Unmarshal(msgbuf, msg)
			// 	fmt.Println("ping:", msg.Timestamp)
			// 	time.Sleep(time.Second * 1)
			// 	SendToClient(int32(account.MSG_ACCOUNT_Module),
			// 		int32(account.MSG_ACCOUNT_Ping),
			// 		&account.C2S_Ping{}, sendchan)
			case int32(protocol_base.MSG_BASE_NoticeMsg):
				fmt.Println("被挤下线")
				return
			case int32(account.MSG_ACCOUNT_Register):
				msg := &account.S2C_Register{}
				proto.Unmarshal(msgbuf, msg)
				fmt.Printf("S2C_Register: [%v]\n", msg)

				//成功注册创建角色
				SendToClient(int32(account.MSG_ACCOUNT_Module),
					int32(account.MSG_ACCOUNT_CreateRole),
					&account.C2S_CreateRole{
						RoleName: accountname,
						Sex:      1,
						HeadID:   23,
					}, sendchan)

			case int32(account.MSG_ACCOUNT_CreateRole):
				msg := &account.S2C_CreateRole{}
				proto.Unmarshal(msgbuf, msg)
				fmt.Printf("S2C_CreateRole: [%v]\n", msg)
				SendToClient(int32(account.MSG_ACCOUNT_Module),
					int32(protocol_base.MSG_BASE_HeartBeat),
					&protocol_base.C2S_HeartBeat{}, sendchan)
			}
		}
	}()

	msg := &account.C2S_Login{
		Account:  accountname,
		Password: password,
	}
	SendToClient(int32(account.MSG_ACCOUNT_Module), int32(account.MSG_ACCOUNT_Login), msg, sendchan)

	for {
		select {
		case buf := <-sendchan:
			le := tools.IntToBytes(int32(len(buf)), 2)
			conn.Write(tools.BytesCombine(le, buf))
		case <-sendctx.Done():
			return
		}
	}

}

// //SendToClient 发送消息至客户端
func SendToClient(module int32, method int32, pb proto.Message, sendchan chan []byte) {
	//logrus.Debugf("client send msg [%v] [%v] [%v]", module, method, pb)
	data, err := proto.Marshal(pb)
	if err != nil {
		logrus.Errorf("proto encode error[%v] [%v][%v] [%v]", err.Error(), module, method, pb)
		return
	}
	mldulebuf := tools.IntToBytes(module, 2)
	methodbuf := tools.IntToBytes(method, 2)
	sendchan <- tools.BytesCombine(mldulebuf, methodbuf, data)
}
