package page

import "html/template"

// struct with common page data
type PageData struct {
	Title       string
	Description string
	Keywords    string
	LeftMenu    bool
}

// show page Title
func (page *PageData) GetTitle() template.HTML {
	return template.HTML("<title>" + page.Title + "</title>")
}

// show page Description
func (page *PageData) GetDescription() template.HTML {
	return template.HTML("<meta name=\"description\" content=\"" + page.Description + "\" />")
}

// show page Keywords
func (page *PageData) GetKeywords() template.HTML {
	return template.HTML("<meta name=\"keywords\" content=\"" + page.Keywords + "\" />")
}

// func (page *PageData) ShowMeta() string {
// 	meta := "<title>" + page.Title + "</title>\n
// 	<meta name=\"keywords\" content=\" + page.Keywords + " />\n
// 	<meta name="description" content=" + +" />\n"
// }
