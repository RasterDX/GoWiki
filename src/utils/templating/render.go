package templating

import (
	pgs "../pages"
	"net/http"
)

func RenderTemplate(w http.ResponseWriter, tmpl string, p *pgs.Page) {
	err := Templates.ExecuteTemplate(w, tmpl + ".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}