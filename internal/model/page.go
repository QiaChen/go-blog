package model

// PageCreateUpdateBase 创建/修改内容基类
type PageCreateUpdateBase struct {
	Title       string
	UrlPath     string
	Description string
	Content     string
	Template    string
	Reply       int
}

type PageCreateInput struct {
	PageCreateUpdateBase
}
type PageUpdateInput struct {
	PageCreateUpdateBase
	Id uint
}
type PageCreateOutput struct {
	Id uint
}
type PagetGetListItem struct {
	PageCreateUpdateBase
	Id uint
}

type PagetGetListOutput struct {
	List []PagetGetListItem
}
