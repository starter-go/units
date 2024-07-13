package units

import (
	"testing"

	"github.com/starter-go/application"
)

// Config ...
type Config struct {

	// Args 启动的命令行参数
	Args []string

	// Cases : 一个名称列表（以','号分隔）; 或者，'all' 表示执行所以的测试用例
	Cases string

	// 测试的目标模块
	Module application.Module

	// Properties 附加的属性表
	Properties map[string]string

	// 测试上下文
	T *testing.T

	// 使用 panic 代替 error
	UsePanic bool
}

// Run ...
func Run(cfg *Config) error {

	preparePropsForRun(cfg)

	// runner
	runner := NewRunner()

	runner.Module(cfg.Module)
	runner.Testing(cfg.T)
	runner.EnablePanic(cfg.UsePanic)
	runner.SetProperties(cfg.Properties)
	// runner.ModuleT(nil)

	err := runner.Run(cfg.Args)
	if err != nil && cfg.T != nil {
		cfg.T.Error(err)
	}
	return err
}

func preparePropsForRun(cfg *Config) {

	// props
	// test.units.run-all = 0
	// test.units.run-list = a,b,c,d

	const (
		runAll  = "test.units.run-all"
		runList = "test.units.run-list"
	)

	cases := cfg.Cases

	props := cfg.Properties
	if props == nil {
		props = make(map[string]string)
	}

	if cases == "" {
		// nop
	} else if cases == "all" {
		props[runAll] = "1"
		props[runList] = ""
	} else {
		props[runAll] = "0"
		props[runList] = cases
	}

	cfg.Properties = props

}
