package gserver

import (
	"fmt"

	"github.com/halturin/ergo"
	"github.com/halturin/ergo/etf"
)

type serverSup struct {
	ergo.Supervisor
}

func (ds *serverSup) Init(args ...interface{}) ergo.SupervisorSpec {
	return ergo.SupervisorSpec{
		Name: "demoSupervisorSup",
		Children: []ergo.SupervisorChildSpec{
			{
				Name:    "demoServer01",
				Child:   &demoGenServ{},
				Restart: ergo.SupervisorChildRestartTemporary,
				// Restart: ergo.SupervisorChildRestartTransient,
				// Restart: ergo.SupervisorChildRestartPermanent,

				// temporary:进程永远都不会被重启
				// transient: 只有进程异常终止的时候会被重启
				// permanent:遇到任何错误导致进程终止就会重启
			},
			{
				Name:    "demoServer02",
				Child:   &demoGenServ{},
				Restart: ergo.SupervisorChildRestartPermanent,
				Args:    []interface{}{12345},
			},
			{
				Name:    "demoServer03",
				Child:   &demoGenServ{},
				Restart: ergo.SupervisorChildRestartPermanent,
				Args:    []interface{}{"abc", 67890},
			},
		},
		Strategy: ergo.SupervisorStrategy{
			Type: ergo.SupervisorStrategyOneForAll,
			// Type:      ergo.SupervisorStrategyRestForOne,
			// Type:      ergo.SupervisorStrategyOneForOne,

			//重启策略
			// one_for_one : 把子进程当成各自独立的,一个进程出现问题其它进程不会受到崩溃的进程的影响.该子进程死掉,只有这个进程会被重启
			// one_for_all : 如果子进程终止,所有其它子进程也都会被终止,然后所有进程都会被重启.
			// rest_for_one:如果一个子进程终止,在这个进程启动之后启动的进程都会被终止掉.然后终止掉的进程和连带关闭的进程都会被重启.
			// simple_one_for_one 是one_for_one的简化版 ,所有子进程都动态添加同一种进程的实例
			Intensity: 2, //次数
			Period:    5, //时间  1 -0 代表不重启
		},
	}
}

// GenServer implementation structure
type demoGenServ struct {
	ergo.GenServer
	process *ergo.Process
}

type state struct {
	i int
}

// Init initializes process state using arbitrary arguments
// Init(...) -> state
func (dgs *demoGenServ) Init(p *ergo.Process, args ...interface{}) interface{} {
	fmt.Printf("Init (%s): args %v \n", p.Name(), args)
	dgs.process = p
	return state{i: 12345}
}

// HandleCast serves incoming messages sending via gen_server:cast
// HandleCast -> ("noreply", state) - noreply
//		         ("stop", reason) - stop with reason
func (dgs *demoGenServ) HandleCast(message etf.Term, state interface{}) (string, interface{}) {
	fmt.Printf("HandleCast (%s): %#v\n", dgs.process.Name(), message)
	switch message {
	case etf.Atom("stop"):
		return "stop", "they said"
	}
	return "noreply", state
}

// HandleCall serves incoming messages sending via gen_server:call
// HandleCall -> ("reply", message, state) - reply
//				 ("noreply", _, state) - noreply
//		         ("stop", reason, _) - normal stop
func (dgs *demoGenServ) HandleCall(from etf.Tuple, message etf.Term, state interface{}) (string, etf.Term, interface{}) {
	fmt.Printf("HandleCall (%s): %#v, From: %#v\n", dgs.process.Name(), message, from)

	reply := etf.Term(etf.Tuple{etf.Atom("error"), etf.Atom("unknown_request")})

	switch message {
	case etf.Atom("hello"):
		reply = etf.Term(etf.Atom("hi"))
	}
	return "reply", reply, state
}

// HandleInfo serves all another incoming messages (Pid ! message)
// HandleInfo -> ("noreply", state) - noreply
//		         ("stop", reason) - normal stop
func (dgs *demoGenServ) HandleInfo(message etf.Term, state interface{}) (string, interface{}) {
	fmt.Printf("HandleInfo (%s): %#v\n", dgs.process.Name(), message)
	return "noreply", state
}

// Terminate called when process died
func (dgs *demoGenServ) Terminate(reason string, state interface{}) {
	fmt.Printf("Terminate (%s): %#v\n", dgs.process.Name(), reason)
}
