package mock

import "fmt"

type Retriever struct {
	Contents string
}

func (r *Retriever) Post(url string, form map[string]string) string {
	//TODO implement me
	r.Contents = form["contents"]
	return "ok"
}

func (r *Retriever) Get(url string) string {
	return r.Contents
}

func (r *Retriever) String() string {
	return fmt.Sprintf(
		"Retriever:{Contents=%s}", r.Contents)
}
