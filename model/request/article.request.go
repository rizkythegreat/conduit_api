package request

type ArticleCreateRequest struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}
