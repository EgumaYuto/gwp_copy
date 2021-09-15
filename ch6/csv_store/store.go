package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

type Post struct {
	Id      int
	Content string
	Author  string
}

func (p Post) ToString() string {
	return strconv.Itoa(p.Id) + " " + p.Content + " " + p.Author
}

func main() {
	csvFile, err := os.Create("posts.csv")
	if err != nil {
		panic(err)
	}
	defer csvFile.Close()

	allPosts := []Post{
		{Id: 1, Content: "Hello1", Author: "Author1"},
		{Id: 2, Content: "Hello2", Author: "Author2"},
		{Id: 3, Content: "Hello3", Author: "Author2"},
		{Id: 4, Content: "Hello4", Author: "Author1"},
	}

	writer := csv.NewWriter(csvFile)
	for _, post := range allPosts {
		line := []string{strconv.Itoa(post.Id), post.Content, post.Author}
		err := writer.Write(line)
		if err != nil {
			panic(err)
		}
	}
	writer.Flush()

	file, err := os.Open("posts.csv")
	if err != nil {
		panic(err)
	}
	defer csvFile.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1
	records, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	var posts []Post
	for _, item := range records {
		id, _ := strconv.ParseInt(item[0], 0, 0)
		posts = append(posts, Post{Id: int(id), Content: item[1], Author: item[2]})
	}
	fmt.Println(posts[0].Id)
	fmt.Println(posts[0].Content)
	fmt.Println(posts[0].Author)
}
