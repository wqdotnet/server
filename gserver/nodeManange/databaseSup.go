package nodemanange

import (
	genserver "server/gserver/genServer"

	"github.com/halturin/ergo"
)

type DataBaseSup struct {
	ergo.Supervisor
}

func (ds *DataBaseSup) Init(args ...interface{}) ergo.SupervisorSpec {
	return ergo.SupervisorSpec{
		Name: "DataBaseSup",
		Children: []ergo.SupervisorChildSpec{
			{
				Name:  "dbServer",
				Child: &genserver.DbGenServer{},
				//Restart: ergo.SupervisorChildRestartTemporary,
				Restart: ergo.SupervisorChildRestartTransient,
				//Restart: ergo.SupervisorChildRestartPermanent,
				Args: []interface{}{},

				// temporary:进程永远都不会被重启
				// transient: 只有进程异常终止的时候会被重启
				// permanent:遇到任何错误导致进程终止就会重启
			},
		},
		Strategy: ergo.SupervisorStrategy{
			Type: ergo.SupervisorStrategyOneForOne,
			//Type: ergo.SupervisorStrategyOneForAll,
			//Type: ergo.SupervisorStrategyRestForOne,

			//重启策略
			// one_for_one : 把子进程当成各自独立的,一个进程出现问题其它进程不会受到崩溃的进程的影响.该子进程死掉,只有这个进程会被重启
			// one_for_all : 如果子进程终止,所有其它子进程也都会被终止,然后所有进程都会被重启.
			// rest_for_one:如果一个子进程终止,在这个进程启动之后启动的进程都会被终止掉.然后终止掉的进程和连带关闭的进程都会被重启.
			// simple_one_for_one 是one_for_one的简化版 ,所有子进程都动态添加同一种进程的实例
			Intensity: 3, //次数
			Period:    5, //时间  1 -0 代表不重启
		},
	}
}

func StartDataBaseSupSupNode(nodeName string) (*ergo.Node, *ergo.Process, error) {
	opts := ergo.NodeOptions{
		ListenRangeBegin: uint16(serverCfg.ListenRangeBegin),
		ListenRangeEnd:   uint16(serverCfg.ListenRangeEnd),
		EPMDPort:         uint16(serverCfg.EPMDPort),
	}

	node := ergo.CreateNode(nodeName, serverCfg.Cookie, opts)
	// Spawn supervisor process
	process, err := node.Spawn("database_sup", ergo.ProcessOptions{}, &DataBaseSup{})
	return node, process, err
}
