package units

import "context"

// Registration 测试单元注册信息
type Registration struct {
	Name        string
	ID          string // 类似于 HTML 标签中的 'id' 属性, 用于 '#ID' 选择器
	Class       string // 类似于 HTML 标签中的 'class' 属性, 用于 '.CLASS' 选择器
	Description string
	Enabled     bool
	Priority    int
	OnError     OnErrorMethod

	Do func(c context.Context) error
}

// Unit 测试单元注册接口
type Unit interface {
	ListRegistrations(list []*Registration) []*Registration
}
