package main

type Config struct {
	Endpoint string
}

// 创建服务名到服务配置的映射
var CfgMap  map[string]*Config

// 先声明
func init() {
	// 后创建
	CtfMap := make(map[string]*Config, 4)
	CtfMap["hello"] =&Config{
		Endpoint: "http://localhost",
	}
}

 // 如何解决配置写死问题
 // 引入配置模块
 // 用服务名作为服务的唯一性id
 // 不同服务有不同的配置
 // 建立一个 服务名 ==> 服务配置的映射 就叫服务配置

 // map 1声明 2 创建 3赋值 4遍历

 // 一个东西 三步走 创建 设置 取值
 // 创建 var 类型变量名 类型


 // endpoint 抽象出上层ctfmap ctfmap肯定是包含endpoint 所以将ctfmap作为结构体 包含 属性endpoint
 // 也可以理解为对象加属性