package handler

import (
	"html/template"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/pradeep-veera89/campaign/internal/database"
	"github.com/pradeep-veera89/campaign/internal/forms"
	"github.com/pradeep-veera89/campaign/internal/models"
)

type Page struct {
	Title   string
	Name    string
	Data    map[string]interface{}
	Success string
	Error   string
	Form    *forms.Forms
}

//TemplateRedirect the page to another page.
func TemplateRedirect(w http.ResponseWriter, r *http.Request, p *Page) {

	var host = r.URL.Host
	var pageName = p.Name
	http.Redirect(w, r, host+pageName, http.StatusTemporaryRedirect)
}

// TemplateRender HTML page.
func TemplateRender(w http.ResponseWriter, r *http.Request, p *Page) error {

	var source = "templates/"
	t, err := template.ParseFiles(source + p.Name + ".html")
	if err != nil {
		return err
	}
	return t.Execute(w, p)
}

// SetCookie save the value as cookie
func SetCookie(w http.ResponseWriter, name, value string) {
	expire := time.Now().Add(20 * time.Minute) // Expires in 20 minutes
	cookie := http.Cookie{
		Name:     name,
		Value:    value,
		MaxAge:   86400,
		Secure:   true,
		HttpOnly: true,
		Expires:  expire,
	}
	http.SetCookie(w, &cookie)
}

// startHandler displays the start page
func (p *Page) StartHandler(w http.ResponseWriter, r *http.Request) {

	log.Println("StartPage")
	p.Title = "StartPage"
	p.Name = "index"
	if r.Method == "GET" {
		p.Form = forms.New(url.Values{})

		err := TemplateRender(w, r, p)
		if err != nil {
			log.Println("failed to execute page "+p.Title, err)
		}
	} else if r.Method == "POST" {

		lead := models.Lead{
			EMail:      r.FormValue("email"),
			FirstName:  r.FormValue("firstname"),
			LastName:   r.FormValue("lastname"),
			Salutation: r.FormValue("salutation"),
		}

		forms := forms.New(r.PostForm)

		forms.Required("email", "firstname", "lastname")
		forms.MinLength("email", 5)
		forms.MinLength("firstname", 3)
		forms.MinLength("lastname", 3)
		forms.RegisteredEmail("email", r)

		data := make(map[string]interface{})
		data["lead"] = lead

		if !forms.Valid() {
			p.Form = forms
			p.Data = data
			p.Error = "Failed to Register the Lead"
			p.Success = ""
			TemplateRender(w, r, p)
			return
		}

		conn := database.GetDB()
		conn.InsertLead(lead)

		p.Form = forms
		p.Data = data
		p.Success = "Lead Registrations is succesfull"
		p.Error = ""
		SetCookie(w, "email", lead.EMail)

		TemplateRender(w, r, p)
	}
}

// successHandler displays the successpage
func (p *Page) SuccessHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("SuccessPage")
	p.Title = "SuccessPage"
	p.Name = "success"
	err := TemplateRender(w, r, p)
	if err != nil {
		log.Println("failed to execute page " + p.Title)
	}
}

// errorhandler displays the errorpage.
func (p *Page) ErrorHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("ErrorPage")
	p.Title = "ErrorPage"
	p.Name = "error"
	err := TemplateRender(w, r, p)
	if err != nil {
		log.Println("failed to execute page " + p.Title)
	}
}
