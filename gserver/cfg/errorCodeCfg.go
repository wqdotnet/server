package cfg

//ErrorCodeCfg 错误码
type ErrorCodeCfg struct {
	Key  string `json:"key"`
	Name string `json:"Name"`
	Code int    `json:"Code"`
}

//GetErrorCodeNumber 错误提示码
func GetErrorCodeNumber(code string) string {
	for _, v := range GameCfg.ErrorCode.CfgList {
		if code == v.Key {
			return v.Name // strconv.Itoa(v.Code)
		}
	}
	return "0"
}
