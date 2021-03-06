package pages

import (
	".."
	"io/ioutil"
)

type Page struct {
	Title string
	Authors []string
	Body []byte
	Date string
	LastEdit string
	Content content
}

func (p * Page) Save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(utils.GetPagesPath() + "/" + filename, p.Body, 0600)
}

func LoadPage(title string) (*Page, error) {
	filename := "res/pages/" + title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}