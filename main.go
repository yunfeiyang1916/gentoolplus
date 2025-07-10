package main

import (
	"github.com/yunfeiyang1916/gentoolplus/process"

	_ "github.com/yunfeiyang1916/gentoolplus/initialize"
)

func main() {
	// 生成所有model和query
	process.ProcessAllTables()
	// 处理表关联关系
	process.ProcessTableRelations()
}
