package main

import (
	"gentoolplus/process"

	_ "gentoolplus/initialize"
)

func main() {

	// 生成所有model和query
	process.ProcessAllTables()
	// 处理表关联关系
	process.ProcessTableRelations()
}
