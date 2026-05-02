package units

// Registration 测试单元注册信息
type Registration struct {
	Name        string
	Alias       string
	Description string
	Enabled     bool
	Priority    int
	OnError     OnErrorMethod

	Do func() error
}

// Unit 测试单元注册接口
type Unit interface {
	ListRegistrations(list []*Registration) []*Registration
}
