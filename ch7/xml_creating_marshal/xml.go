package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
)

type Post struct {
	XMLName xml.Name `xml:"post"`
	Id      string   `xml:"id,attr"`
	Content string   `xml:"content"`
	Author  Author   `xml:"author"`
}

type Author struct {
	Id   string `xml:"id,attr"`
	Name string `xml:",chardata"`
}

func main() {
	post := Post{
		Id:      "1",
		Content: "Hello World!",
		Author: Author{
			Id:   "2",
			Name: "Saw Sheong",
		},
	}
	output, err := xml.Marshal(&post)
	if err != nil {
		fmt.Println("Error marshaling to XML:", err)
		return
	}
	err = ioutil.WriteFile("post.xml", output, 0644)
	if err != nil {
		fmt.Println("Error writing XML to  file:", err)
		return
	}
}
