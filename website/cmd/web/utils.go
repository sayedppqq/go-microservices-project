package main

import (
	"flag"
	"log"
	"net/http"
	"time"
)

type connectionOptions struct {
	serverAddr *string
	serverPort *int

	usersAPI     *string
	moviesAPI    *string
	showtimesAPI *string
	bookingsAPI  *string
}
type apis struct {
	users     *string
	movies    *string
	showtimes *string
	bookings  *string
}

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
	apis     apis
}

func (opt *connectionOptions) GetApplication(errorLog, infoLog *log.Logger) *application {
	return &application{
		errorLog: errorLog,
		infoLog:  infoLog,
		apis: apis{
			users:     opt.usersAPI,
			movies:    opt.moviesAPI,
			showtimes: opt.showtimesAPI,
			bookings:  opt.bookingsAPI,
		},
	}
}
func (opt *connectionOptions) GetHTTPServer(errorLog *log.Logger, serverURI string, app *application) *http.Server {
	srv := &http.Server{
		Addr:         serverURI,
		ErrorLog:     errorLog,
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
}
func (opt *connectionOptions) pasrseCommandLineFlags() {
	// Define command-line flags
	opt.serverAddr = flag.String("serverAddr", "", "HTTP server network address")
	opt.serverPort = flag.Int("serverPort", 8080, "HTTP server network port")

	opt.usersAPI = flag.String("usersAPI", "http://localhost:4000/api/users/", "Users API")
	opt.moviesAPI = flag.String("moviesAPI", "http://localhost:4000/api/movies/", "Movies API")
	opt.showtimesAPI = flag.String("showtimesAPI", "http://localhost:4000/api/showtimes/", "Showtimes API")
	opt.bookingsAPI = flag.String("bookingsAPI", "http://localhost:4000/api/bookings/", "Bookings API")
	flag.Parse()
}
