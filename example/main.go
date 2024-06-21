// Binary goproxy HTTP(S)代理
//
//go:generate statik -src=web/public -dest=internal -f
package main

import (
	"github.com/ouqiang/goproxy/example/cmd"
	"github.com/ouqiang/goproxy/example/version"
)

var (
	// AppVersion 应用版本
	AppVersion string
	// BuildDate 构建日期
	BuildDate string
	// GitCommit 最后提交的git commit
	GitCommit string
)

func main() {
	version.Init(AppVersion, BuildDate, GitCommit)
	cmd.Execute()
}
