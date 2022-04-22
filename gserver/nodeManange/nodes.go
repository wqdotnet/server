package nodeManange

import (
	"fmt"
	"server/gserver/commonstruct"
	"sync"

	"github.com/ergo-services/ergo/gen"
	"github.com/ergo-services/ergo/node"
)

//ergo.Node 节点管理
var (
	nodesMap sync.Map
	//remoteMap sync.Map  //远程连接节点
)

type GenNodeName string

const (
	GateNode   GenNodeName = "gate"
	ServerNode GenNodeName = "server"
	DBNode     GenNodeName = "db"
)

//系统初始启动的几个 genserver
type GenServerName string

const (
	DBGenServer   GenServerName = "dbServer"
	GameGenServer GenServerName = "gameServer"
	CMDGenServer  GenServerName = "cmdServer"
)

func Start(command chan string) {
	for _, v := range commonstruct.ServerCfg.StartList {
		switch v {
		case "gateway":
			gateNode, _, gerr := StartGateSupNode(getNodeName(GateNode))
			if gerr != nil {
				panic(gerr)
			}
			nodesMap.Store(gateNode.Name(), gateNode)
		case "server":
			serverNode, _, serr := StartGameServerSupNode(getNodeName(ServerNode), command)
			if serr != nil {
				panic(serr)
			}
			nodesMap.Store(serverNode.Name(), serverNode)
		case "db":
			dbNode, _, derr := StartDataBaseSupSupNode(getNodeName(DBNode))
			if derr != nil {
				panic(derr)
			}
			//dbNode.Spawn("", gen.ProcessOptions{}, nil)
			nodesMap.Store(dbNode.Name(), dbNode)
		}
	}
}

func getNodeName(node GenNodeName) string {
	switch node {
	case GateNode:
		return fmt.Sprintf("gatewayNode_%v@127.0.0.1", commonstruct.ServerCfg.ServerID)
	case ServerNode:
		return fmt.Sprintf("serverNode_%v@127.0.0.1", commonstruct.ServerCfg.ServerID)
	case DBNode:
		return fmt.Sprintf("dbNode_%v@127.0.0.1", commonstruct.ServerCfg.ServerID)
	}
	return ""
}

func GetNode(nodename GenNodeName) node.Node {
	if v, ok := nodesMap.Load(getNodeName(nodename)); ok {
		return v.(node.Node)
	}
	return nil
}

func GetGenServer(genserver GenServerName) gen.Process {
	dbnode := GetNode(DBNode)
	return dbnode.ProcessByName(string(genserver))
}

func GetNodes() map[string]node.Node {
	nodemap := map[string]node.Node{}
	nodesMap.Range(func(key, value interface{}) bool {
		nodemap[key.(string)] = value.(node.Node)
		return true
	})
	return nodemap
}
