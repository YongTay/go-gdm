package model

type Dependence struct {
	Package string `xml:"package"`
	Version string `xml:"version"`
	Url     string `xml:"url"`
}

type Dependencies struct {
	Dependence []Dependence `xml:"dependence"`
}

type Project struct {
	Dependencies Dependencies `xml:"dependencies"`
	Workspace    string       `xml:"workspace"`
}
