package unitcore

import "github.com/starter-go/application"

// Context 是运行单元测试的上下文
type Context struct {
	RunList            []string
	RunAll             bool
	ApplicationContext application.Context
}
