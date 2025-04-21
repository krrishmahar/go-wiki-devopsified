package main

import (
	"errors"
	"html/template"
	"log"
	"net/http"
	"os"
	"regexp"
)

const (
	templateDir = "tmpl/"
	dataDir     = "data/"
)

var templates = template.Must(template.ParseFiles(templateDir+"edit.html", templateDir+"view.html", templateDir+"error.html"))
var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {
	filename := dataDir + p.Title + ".txt"
	return os.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := dataDir + title + ".txt"
	body, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r, m[2])
	}
}

func getTitle(w http.ResponseWriter, r *http.Request) (string, error) {
	m := validPath.FindStringSubmatch(r.URL.Path)
	if m == nil {
		http.NotFound(w, r)
		return "", errors.New("invalid Page Title")
	}
	return m[2], nil // The title is the second subexpression.
}

var linkRegexp = regexp.MustCompile(`\[(\w+)\]`)

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	// Apply link replacement as before
	htmlBody := linkRegexp.ReplaceAllFunc(p.Body, func(match []byte) []byte {
		pageName := match[1 : len(match)-1]
		link := []byte(`<a href="/view/` + string(pageName) + `">` + string(pageName) + `</a>`)
		return link
	})
	p.Body = htmlBody

	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		renderError(w, "Template error: "+err.Error())
		log.Printf("Template render error: %v", err)
	}
}

func renderError(w http.ResponseWriter, errMsg string) {
	err := templates.ExecuteTemplate(w, "error.html", errMsg)
	if err != nil {
		http.Error(w, "500 Internal Server Error: "+err.Error(), http.StatusInternalServerError)
	}
}

func viewHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}

	renderTemplate(w, "view", p)
}

func editHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	renderTemplate(w, "edit", p)
}

func saveHandler(w http.ResponseWriter, r *http.Request, title string) {
	if r.Method != http.MethodPost {
		renderError(w, "Direct access to /save/"+title+" is not allowed. Use the form to POST data.")
		return
	}
	body := r.FormValue("body")
	if len(body) == 0 {
		renderError(w, "Cannot save empty content for page: "+title)
		return
	}
	
	p := &Page{Title: title, Body: []byte(body)}
	err := p.save()
	if err != nil {
		http.Error(w, "500 Internal Server Error: "+err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/view/FrontPage", http.StatusFound)
	})

	log.Println(`Server is listening on port 8080`)
	http.HandleFunc("/view/", makeHandler(viewHandler))
	http.HandleFunc("/edit/", makeHandler(editHandler))
	http.HandleFunc("/save/", makeHandler(saveHandler))
	log.Fatal(http.ListenAndServe(":8080", nil))
	log.Println("Made with ❤️ by krrish")

}
