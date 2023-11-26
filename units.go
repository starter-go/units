package units

// Registration 测试单元注册信息
type Registration struct {
	Name     string
	Enabled  bool
	Priority int
	Test     func() error
}

// Units 测试单元注册接口
type Units interface {
	ListUnits(list []*Registration) []*Registration
}
