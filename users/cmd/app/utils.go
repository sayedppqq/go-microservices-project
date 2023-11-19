package app

import (
	"context"
	"flag"
	"fmt"
	"github.com/sayedppqq/go-microservices-project/users/pkg/models/mongodb"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
	"os"
	"runtime/debug"
	"time"
)

type connectionOptions struct {
	serverAddr    *string
	serverPort    *int
	mongoURI      *string
	mongoDatabase *string
	enableCred    *bool
}
type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
	users    *mongodb.UserModel
}

func (opt *connectionOptions) GetMongoClient(ctx context.Context) (*mongo.Client, error) {
	co := options.Client().ApplyURI(*opt.mongoURI)
	if *opt.enableCred {
		co.Auth = &options.Credential{
			Username: os.Getenv("MONGODB_USERNAME"),
			Password: os.Getenv("MONGODB_PASSWORD"),
		}
	}
	client, err := mongo.Connect(ctx, co)
	if err != nil {
		return nil, err
	}
	return client, err
}

func (opt *connectionOptions) GetApplication(errorLog, infoLog *log.Logger, client *mongo.Client) *application {
	// Initialize a new instance of application containing the dependencies.
	return &application{
		errorLog: errorLog,
		infoLog:  infoLog,
		users: &mongodb.UserModel{
			C: client.Database(*opt.mongoDatabase).Collection("users"),
		},
	}
}

func (opt *connectionOptions) GetHTTPServer(errorLog *log.Logger, serverURI string, app *application) *http.Server {

	return &http.Server{
		Addr:         serverURI,
		Handler:      app.routes(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  time.Minute,
		ErrorLog:     errorLog,
	}
}

func (app *application) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.errorLog.Output(2, trace)

	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (app *application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

func (opt *connectionOptions) pasrseCommandLineFlags() {
	// Define command-line flags
	opt.serverAddr = flag.String("serverAddr", "", "HTTP server network address")
	opt.serverPort = flag.Int("serverPort", 4003, "HTTP server network port")
	opt.mongoURI = flag.String("mongoURI", "mongodb://localhost:27017", "Database hostname url")
	opt.mongoDatabase = flag.String("mongoDatabase", "users", "Database name")
	opt.enableCred = flag.Bool("enableCredentials", false, "Enable the use of credentials for mongo connection")
	flag.Parse()
}
