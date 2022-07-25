package form


type FormData struct {
	formName    string
	formPhone   string
	formMessage string
}

// setters
func (fd *FormData) SetName(name string) {
	fd.formName = name
}

func (fd *FormData) SetPhone(phone string) {
	fd.formPhone = phone
}

func (fd *FormData) SetMessage(message string) {
	fd.formMessage = message
}

// getters
func (fd *FormData) GetName(name string) string {

	return fd.formName
}

func (fd *FormData) GetPhone(phone string) string {
	return fd.formPhone
}

func (fd *FormData) GetMessage(message string) string {
	return fd.formMessage
}
