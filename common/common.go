package common

type Post struct {
	Id     uint64 `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
	Image1 []byte `json:"image1"`
}
