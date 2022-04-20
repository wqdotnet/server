package commonstruct

// ServerConfig  server cfg
type ServerConfig struct {
	ServerName string
	ServerID   int32
	Version    string

	Daemon     bool
	RestartNum int

	OpenHTTP bool
	HTTPPort int32
	OpenWS   bool

	OpenPyroscope bool
	PyroscopeHost string

	NetType       string
	Port          int32
	Packet        int32
	Readtimeout   int32 //读超时时间
	MaxConnectNum int32 //最大连接数

	MsgTime int32
	MsgNum  int32

	ProtoPath string
	GoOut     string

	MongoConnStr string
	Mongodb      string

	RedisConnStr string
	RedisDB      int

	CfgPath     string
	CfgType     string
	WatchConfig bool

	LogWrite bool
	Loglevel string
	LogPath  string
	LogName  string

	ListenBegin int
	ListenEnd   int
	Cookie      string

	StartList []string
}

// ServerCfg  Program overall configuration
var ServerCfg = ServerConfig{
	ServerName: "server",
	ServerID:   1,
	Version:    "1.0.0",

	Daemon:     false,
	RestartNum: 2,

	// http
	OpenHTTP: false,
	HTTPPort: 8080,
	OpenWS:   true,

	OpenPyroscope: false,
	PyroscopeHost: "http://localhost:4040",

	// #network : tcp/udp
	NetType:       "tcp",
	Port:          3344,
	Packet:        2,
	Readtimeout:   0,
	MaxConnectNum: 2000,

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

	ListenBegin: 15151,
	ListenEnd:   25151,
	Cookie:      "123",
	StartList:   []string{"db", "gateway", "server"},
}
