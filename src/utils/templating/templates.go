package templating

import (
	".."
	"html/template"
)

var Templates = template.Must(template.ParseFiles(utils.GetTemplatePath() + "edit.html", utils.GetTemplatePath() +  "view.html", utils.GetTemplatePath() + "head.tmpl", utils.GetTemplatePath() + "header.tmpl"))
