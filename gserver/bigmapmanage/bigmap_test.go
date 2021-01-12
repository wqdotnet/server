package bigmapmanage

import (
	"container/list"
	"fmt"
	"server/db"
	"server/gserver/cfg"
	"server/gserver/commonstruct"
	"server/gserver/process"
	"server/gserver/timedtasks"
	"server/msgproto/bigmap"
	"server/tool"
	"testing"
	"time"

	log "github.com/sirupsen/logrus"
)

func init() {
	//ctx, cancelFunc := context.WithCancel(context.Background())
	db.StartRedis("127.0.0.1:6379", 0)
	StartBigmapGoroutine()
	//启动定时器
	timedtasks.StartCronTasks()
	//大地图loop
	timedtasks.AddTasks("bigmaploop", "* * * * * ?", func() {
		SendMsgBigMap("BigMapLoop_OneSecond")
	})

	cfg.InitViperConfig("../../config", "json")

}

func TestList(t *testing.T) {
	// 创建一个 list
	l := list.New()
	//把4元素放在最后
	e4 := l.PushBack(4)
	//把1元素放在最前
	e1 := l.PushFront(1)
	//在e4元素前面插入3
	l.InsertBefore(3, e4)
	//在e1后面插入2
	e2 := l.InsertAfter(2, e1)
	// 遍历所有元素并打印其内容
	fmt.Println(" 元素 ")

	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value, " ")
	}
	//获取l 最前的元素
	et1 := l.Front()
	fmt.Println("list 最前的元素 Front  ", et1.Value)
	//获取l 最后的元素
	et2 := l.Back()
	fmt.Println("list 最后的元素  Back ", et2.Value)
	//获取l的长度
	fmt.Println("list 的长度为： Len ", l.Len())
	//向后移动
	l.MoveAfter(e1, e2)
	fmt.Println("把1元素移动到2元素的后面 向后移动后 MoveAfter :")
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value, " ")
	}
	//向前移动
	l.MoveBefore(e1, e2)
	fmt.Println("\n把1元素移动到2元素的前面 向前移动后 MoveBefore :")
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value, " ")
	}
	//移动到最后面
	l.MoveToBack(e1)
	fmt.Println("\n 1元素出现在最后面 MoveToBack ")
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value, " ")
	}
	//移动到最前面
	l.MoveToFront(e1)
	fmt.Println("\n 1元素出现在最前面 MoveToFront ")
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value, " ")
	}
	//删除元素
	fmt.Println("")
	l.Remove(e1)
	fmt.Println("\n e1元素移除后 Remove ")
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value, " ")
	}
	// init 可以用作 clear
	l.Init()
	fmt.Println("\n list init()后 ")
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value, " ")
	}
	fmt.Println("list 的长度Init  ", l.Len())
}

func TestMove(t *testing.T) {

	fc := func(troopsid int32, roleid int32, path []int32) {
		ch := make(chan commonstruct.ProcessMsg)
		process.Register(roleid, ch)
		SendTroopsMove(commonstruct.TroopsStruct{TroopsID: troopsid, State: 1, Roleid: roleid, AreasList: path})
		for troops := range ch {
			switch troops.MsgType {
			case "TroopsMove":
				datat := troops.Data.(commonstruct.TroopsStruct)
				log.Infof("[%v][%v]role receive:[%v]", time.Now().Format("15:04:05"), tool.GoID(), datat)

				// if datat.AreasIndex == 25 {
				// 	SendStopMove(datat.TroopsID)
				// }
			case "OverMove":
				datat := troops.Data.(commonstruct.TroopsStruct)
				log.Infof("[%v] OverMove2: [%v]  [%v]", time.Now().Format("15:04:05"), time.Now().Unix(), datat)
				return

			}
		}
	}

	// unix := time.Now().Unix()
	// arlist := []int32{21, 22, 23, 24, 25, 26, 27, 28, 29}
	go fc(102, 111, []int32{21, 22, 23, 24, 25, 26, 27, 28, 29})

	// log.Info("OverMove1:", int64(len(arlist)*3)+unix)

	time.Sleep(time.Second * 50)
}

func TestInit(t *testing.T) {
	saveAreasInfo(AreasInfo{AreasIndex: 1048})
	log.Info("1048:", getAreasInfo(1048))
	log.Info("748:", getAreasInfo(748))
}

func TestBigmapConfigInit(t *testing.T) {
	areaslist := &bigmap.S2C_AreasInfo{}

	AreasRange(func(value AreasInfo) bool {
		log.Info("value.Type:", value.Type, "    value.State:", value.State)
		//只发送被占领的，正在发生战斗的
		if value.Occupy > 0 || value.State > 0 {
			areaslist.AreasInfoList = append(areaslist.AreasInfoList,
				&bigmap.P_AreasInfo{AreasIndex: value.AreasIndex,
					Type:  value.Type,
					State: value.State})
		}
		return true
	})

	log.Info(len(areaslist.AreasInfoList))

}
