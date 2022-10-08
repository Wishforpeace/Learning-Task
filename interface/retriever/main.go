package main

import (
	"fmt"
	"interface/mock"
	"interface/real"
)

const url = "http://www.imooc.com"

type Retriever interface {
	Get(url string) string
}

type Poster interface {
	Post(url string, form map[string]string) string
}

func download(r Retriever) string {
	return r.Get(url)
}

func post(poster Poster) {
	poster.Post(url,
		map[string]string{
			"name":   "ccmouse",
			"course": "golang",
		})
}

type RetrieverPoster interface {
	Retriever
	Poster
}

func serssion(s RetrieverPoster) string {
	s.Post(url, map[string]string{
		"contents": "another faked imooc com",
	})
	return s.Get(url)
}
func main() {
	var r Retriever
	retriever := mock.Retriever{"this is a fake imooc.com"}
	r = &retriever
	fmt.Printf("%T %v\n", r, r)
	//r = &real.Retriever{
	//	UserAgent: "Mozilla/5.0",
	//	TimeOut:   time.Minute,
	//}
	//fmt.Printf("%T %v\n", r, r)
	inspect(r)
	if mockRetriever, ok := r.(*mock.Retriever); ok {
		fmt.Println(mockRetriever.Contents)
	} else {
		fmt.Println("not a mock retriever")
	}
	fmt.Println("Try a session")
	fmt.Println(serssion(&retriever))
}
func inspect(r Retriever) {
	fmt.Println("Inspecting", r)
	fmt.Printf("> %T %v\n", r, r)
	fmt.Print("> Type switch:")
	fmt.Printf("%T %v\n", r, r)
	switch v := r.(type) {
	case *mock.Retriever:
		fmt.Println("Contents:", v.Contents)
	case *real.Retriever:
		fmt.Println("UserAgents:", v.UserAgent)
	}
}
