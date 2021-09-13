package gserver

import "server/gserver/commonstruct"

// ServerCfg  Program overall configuration
var ServerCfg = commonstruct.ServerConfig{
	ServerName: "server",
	ServerID:   1,
	Version:    "1.0.0",

	Daemon:     false,
	RestartNum: 2,

	// http
	OpenHTTP: false,
	HTTPPort: 8080,

	StatsView:     false,
	StatsViewPort: 8081,

	// #network : tcp/udp
	NetType:     "tcp",
	Port:        3344,
	Packet:      2,
	Readtimeout: 0,

	MsgTime: 300,
	MsgNum:  500,

	// #protobuf path
	ProtoPath: "./proto",
	GoOut:     "./msgproto",

	MongoConnStr: "mongodb://localhost:27017",
	Mongodb:      "mygame",

	RedisConnStr: "localhost:6379",
	RedisDB:      0,

	CfgPath:     "./config",
	CfgType:     "json",
	WatchConfig: false,

	Loglevel: "info",
	LogPath:  "./log",
	LogName:  "log",
	LogWrite: false,

	ListenRangeBegin: 15151,
	ListenRangeEnd:   25151,
	EPMDPort:         4369,
	Cookie:           "123",
}
