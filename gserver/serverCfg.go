package gserver

// ServerConfig  server cfg
type ServerConfig struct {
	ServerName string
	ServerID   int32
	Version    string

	Daemon     bool
	RestartNum int

	OpenHTTP bool
	HTTPPort int32

	StatsView     bool
	StatsViewPort int32

	NetType     string
	Port        int32
	Packet      int32
	Readtimeout int32 //读超时时间

	MsgTime int32
	MsgNum  int32

	ProtoPath string
	GoOut     string

	MongoConnStr string
	Mongodb      string

	RedisConnStr string
	RedisDB      int

	CfgPath string
	CfgType string

	LogWrite bool
	Loglevel string
	LogPath  string
	LogName  string

	ListenRangeBegin int
	ListenRangeEnd   int
	EPMDPort         int
	NodeName         string
	Cookie           string
}

// ServerCfg  Program overall configuration
var ServerCfg = ServerConfig{
	ServerName: "server",
	ServerID:   1,
	Version:    "0.0.1",

	Daemon:     false,
	RestartNum: 2,

	// http
	OpenHTTP: true,
	HTTPPort: 8080,

	StatsView:     true,
	StatsViewPort: 8087,
	// #network : tcp/udp
	NetType:     "tcp",
	Port:        3344,
	Packet:      2,
	Readtimeout: 0,

	MsgTime: 300,
	MsgNum:  500,

	// #protobuf path
	ProtoPath: "./proto",
	GoOut:     "./proto",

	MongoConnStr: "mongodb://localhost:27017",
	Mongodb:      "mygame",

	RedisConnStr: "127.0.0.1:6379",
	RedisDB:      0,

	CfgPath: "./config",
	CfgType: "",

	Loglevel: "info",
	LogPath:  "./log",
	LogName:  "log",
	LogWrite: false,

	ListenRangeBegin: 15151,
	ListenRangeEnd:   25151,
	EPMDPort:         4369,
	NodeName:         "server001@127.0.0.1",
	Cookie:           "123",
}
