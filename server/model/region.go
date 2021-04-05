package model

/**
	管理维度：
		- 对 terms 的弱分类管理
 */
type Region struct {
	Id string	//全局唯一
	Name string	//可读性

	Type string //类型，排他
	Tags []string //标签，非排他

	Status int
	Terms []Term
}