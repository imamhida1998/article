package request

type GetListArticle struct {
	Query  string `form:"query" json:"query"`
	Author string `form:"author" json:"author"`
}
