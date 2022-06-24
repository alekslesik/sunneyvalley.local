package page

// struct with common page data
type PageData struct {
	Title       string
	Description string
	Keywords    string
}

// show page Title
func (page *PageData) GetTitle() string {
	return "<title>" + page.Title + "</title>\n"
}

// show page Description
func (page *PageData) GetDescription() string {
	return "<meta name=\"description\" content=\"" + page.Description + "\" />"
}

// show page Keywords
func (page *PageData) GetKeywords() string {
	return "<meta name=\"keywords\" content=\"" + page.Keywords + "\" />"
}

// func (page *PageData) ShowMeta() string {
// 	meta := "<title>" + page.Title + "</title>\n
// 	<meta name=\"keywords\" content=\" + page.Keywords + " />\n
// 	<meta name="description" content=" + +" />\n"
// }
