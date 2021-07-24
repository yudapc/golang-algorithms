package structsample

import (
	"fmt"
)

type Article struct {
	ID    string
	Title string
}

func SampleAppendSliceOfStruct() {
	fmt.Println("Hello, playground")
	var articles []Article
	articleOne := Article{
		ID:    "one",
		Title: "One Title",
	}
	articles = append(articles, articleOne)

	articleTwo := Article{
		ID:    "two",
		Title: "Two Title",
	}
	articles = append(articles, articleTwo)
	fmt.Println(articles)
}
