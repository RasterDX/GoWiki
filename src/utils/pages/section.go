package pages

type section struct {
	Title string
	Intro string
	Subsections []subsection

}

type subsection struct {
	Title string
	Body string
}