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

	CfgPath     string
	CfgType     string
	WatchConfig bool

	LogWrite bool
	Loglevel string
	LogPath  string
	LogName  string

	ListenRangeBegin int
	ListenRangeEnd   int
	EPMDPort         int
	Cookie           string

	Pid int
}
