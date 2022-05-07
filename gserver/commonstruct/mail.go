package commonstruct

type RoleMail struct {
	RoleID   int32
	MailList []*MailInfo
}

type MailInfo struct {
	UUID      string
	Title     string
	Body      string
	Annex     []string
	Logotype  bool  //是否已读
	Timestamp int64 //发邮件时间戳
}
