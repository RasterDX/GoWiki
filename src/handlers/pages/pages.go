package pages

import (
	pg "../../utils/pages"
	tmpl "../../utils/templating"
	"fmt"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)
	http.ServeFile(w, r, r.URL.Path)
}

func ViewHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := pg.LoadPage(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	tmpl.RenderTemplate(w, "view", p)
}

func EditHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := pg.LoadPage(title)
	if err != nil {
		p = &pg.Page{Title: title}
	}
	tmpl.RenderTemplate(w, "edit", p)
}

func SaveHandler(w http.ResponseWriter, r *http.Request, title string) {
	body := r.FormValue("body")
	author := r.FormValue("author")
	p := &pg.Page{Title: title, Authors : []string{author}, Body: []byte(body)}
	err := p.Save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}