package app

import (
	"github.com/gorilla/mux"
	"github.com/sayedppqq/go-microservices-project/movies/pkg/models"
	"net/http"
	"text/template"
)

type movieTemplateData struct {
	Movie  models.Movie
	Movies []models.Movie
}

func (app *application) moviesList(w http.ResponseWriter, r *http.Request) {

	// Get movies list from API
	var mtd movieTemplateData
	app.infoLog.Println("Calling movies API...")
	err := app.getAPIContent(*app.apis.movies, &mtd.Movies)
	if err != nil {
		app.errorLog.Println(err.Error())
		return
	}
	app.infoLog.Println(mtd.Movies)

	// Load template files
	files := []string{
		"./ui/html/movies/list.page.tmpl",
		"./ui/html/base.layout.tmpl",
		"./ui/html/footer.partial.tmpl",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.errorLog.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

	err = ts.Execute(w, mtd)
	if err != nil {
		app.errorLog.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
}

func (app *application) moviesView(w http.ResponseWriter, r *http.Request) {
	// Get id from incoming url
	vars := mux.Vars(r)
	movieID := vars["id"]
	var mtd movieTemplateData
	// Get movies list from API
	app.infoLog.Println("Calling movies API...")
	url := *app.apis.movies + movieID
	err := app.getAPIContent(url, &mtd.Movie)
	if err != nil {
		app.errorLog.Println(err.Error())
		return
	}
	app.infoLog.Println(mtd.Movie)

	// Load template files
	files := []string{
		"./ui/html/movies/view.page.tmpl",
		"./ui/html/base.layout.tmpl",
		"./ui/html/footer.partial.tmpl",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.errorLog.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

	err = ts.Execute(w, mtd.Movie)
	if err != nil {
		app.errorLog.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
}
