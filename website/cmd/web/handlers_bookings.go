package app

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/sayedppqq/go-microservices-project/bookings/pkg/models"
	modelShowTime "github.com/sayedppqq/go-microservices-project/showtimes/pkg/models"
	modelUser "github.com/sayedppqq/go-microservices-project/users/pkg/models"
	"net/http"
	"text/template"
)

type bookingTemplateData struct {
	Booking      models.Booking
	Bookings     []models.Booking
	BookingData  bookingData
	BookingsData []bookingData
}
type bookingData struct {
	ID           string
	UserFullName string
	ShowTimeDate string
}

func (app *application) loadBookingData(btd *bookingTemplateData, isList bool) {
	btd.BookingData = bookingData{}
	btd.BookingsData = []bookingData{}
	if isList {
		for _, b := range btd.Bookings {
			// Load user data
			userURL := *app.apis.users + b.UserID
			var user modelUser.User
			err := app.getAPIContent(userURL, &user)
			if err != nil {
				app.errorLog.Println(err.Error())
			}

			// Load showtime data
			showtimeURL := fmt.Sprintf("%s/%s", app.apis.showtimes, b.ShowtimeID)
			var showtime modelShowTime.ShowTime
			err = app.getAPIContent(showtimeURL, &showtime)
			if err != nil {
				app.errorLog.Println(err.Error())
			}

			bookingData := bookingData{
				ID:           b.ID.Hex(),
				UserFullName: fmt.Sprintf("%s %s", user.Name, user.LastName),
				ShowTimeDate: showtime.Date,
			}
			btd.BookingsData = append(btd.BookingsData, bookingData)
			app.infoLog.Println(b.UserID)
		}
	} else {
		b := btd.Booking

		// Load user data
		userURL := *app.apis.users + b.UserID
		var user modelUser.User
		err := app.getAPIContent(userURL, &user)
		if err != nil {
			app.errorLog.Println(err.Error())
			return
		}

		// Load showtime data
		showtimeURL := *app.apis.showtimes + b.ShowtimeID
		var showtime modelShowTime.ShowTime
		err = app.getAPIContent(showtimeURL, &showtime)
		if err != nil {
			app.errorLog.Println(err.Error())
			return
		}

		btd.BookingData = bookingData{
			ID:           b.ID.Hex(),
			UserFullName: fmt.Sprintf("%s %s", user.Name, user.LastName),
			ShowTimeDate: showtime.Date,
		}
	}
}

func (app *application) bookingsList(w http.ResponseWriter, r *http.Request) {
	var td bookingTemplateData
	app.infoLog.Println("Calling bookings API...")
	err := app.getAPIContent(*app.apis.bookings, &td.Bookings)
	if err != nil {
		app.errorLog.Println(err.Error())
		return
	}
	app.infoLog.Println(td.Bookings)

	app.loadBookingData(&td, true)

	// Load template files
	files := []string{
		"./ui/html/bookings/list.page.tmpl",
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

func (app *application) bookingsView(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookingID := vars["id"]

	var td bookingTemplateData
	app.infoLog.Println("Calling bookings API...")
	url := *app.apis.bookings + bookingID

	err := app.getAPIContent(url, &td.Booking)
	if err != nil {
		app.errorLog.Println(err.Error())
		return
	}
	app.infoLog.Println(td.Booking)
	app.infoLog.Println(url)

	app.loadBookingData(&td, false)

	// Load template files
	files := []string{
		"./ui/html/bookings/view.page.tmpl",
		"./ui/html/base.layout.tmpl",
		"./ui/html/footer.partial.tmpl",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.errorLog.Println(err.Error())
		http.Error(w, "Internal server error", 500)
		return
	}
	err = ts.Execute(w, td)
	if err != nil {
		app.errorLog.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
}
