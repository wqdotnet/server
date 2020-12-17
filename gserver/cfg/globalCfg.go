package cfg

var (
	//GlobalCfg 全局配置
	GlobalCfg global
)

type global struct {
	//大地图配置
	MapInfo mapinfo
}

//错误提示
var (
	ERROR_AccountNull    = "未找到账号 or 密码错误"
	ERROR_AccountExists  = "账号已存在"
	ERROR_RoleNameExists = "角色名已存在"
)
