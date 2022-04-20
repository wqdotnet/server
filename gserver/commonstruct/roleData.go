package commonstruct

type RoleData struct {
	Acconut      *AccountInfoStruct //账号信息
	RoleBaseInfo *RoleBaseInfo      //角色基础数据
	//道具
	//宗门
	//好友
}

func GetRoleAllData(roleid int32) *RoleData {
	return &RoleData{
		Acconut:      &AccountInfoStruct{},
		RoleBaseInfo: &RoleBaseInfo{},
	}
}
