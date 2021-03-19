package commonstruct

//AddExpItem 获取经验
type AddExpItem struct {
	Type     int32 //0 角色  1部队
	Key      int32
	AddExp   int64
	NewLevel int64
	NewExp   int64
	LostNum  int32
}
