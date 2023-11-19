package app

import (
	"github.com/gorilla/mux"
	"github.com/sayedppqq/go-microservices-project/showtimes/pkg/models"
	"net/http"
	"strings"
	"text/template"
)

type showtimeTemplateData struct {
	ShowTime  models.ShowTime
	ShowTimes []models.ShowTime
	Movies    string
}

func (app *application) showtimesList(w http.ResponseWriter, r *http.Request) {
	var td showtimeTemplateData
	app.infoLog.Println("Calling showtimes API...")
	err := app.getAPIContent(*app.apis.showtimes, &td.ShowTimes)
	if err != nil {
		app.errorLog.Println("error on getAPIContent of showtime", err.Error())
		return
	}
	app.infoLog.Println(td.ShowTime)

	// Load template files
	files := []string{
		"./ui/html/showtimes/list.page.tmpl",
		"./ui/html/base.layout.tmpl",
		"./ui/html/footer.partial.tmpl",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.errorLog.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

	err = ts.Execute(w, td)
	if err != nil {
		app.errorLog.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
}

func (app *application) showtimesView(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	showtimeID := vars["id"]
	var td showtimeTemplateData
	app.infoLog.Println("Calling showtimes API...")
	url := *app.apis.showtimes + showtimeID
	err := app.getAPIContent(url, &td.ShowTime)
	if err != nil {
		app.errorLog.Println("error on getAPIContent of showtime", err.Error())
		return
	}
	app.infoLog.Println(td.ShowTime)

	// Load movie names
	movies := td.ShowTime.Movies
	td.Movies = strings.Join(movies, ", ")

	// Load template files
	files := []string{
		"./ui/html/showtimes/view.page.tmpl",
		"./ui/html/base.layout.tmpl",
		"./ui/html/footer.partial.tmpl",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.errorLog.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

	err = ts.Execute(w, td)
	if err != nil {
		app.errorLog.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}

}
