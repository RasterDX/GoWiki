package utils

import "html/template"

var Templates = template.Must(template.ParseFiles("edit.html", "view.html"))
