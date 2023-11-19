package app

import (
	"fmt"
	"log"
	"os"
)

func Start() {
	opt := connectionOptions{}
	opt.pasrseCommandLineFlags()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime|log.Lshortfile)
	errLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	app := opt.GetApplication(errLog, infoLog)

	serverURI := fmt.Sprintf("%s:%d", *opt.serverAddr, *opt.serverPort)
	srv := opt.GetHTTPServer(errLog, serverURI, app)

	infoLog.Printf("Starting server on %s", serverURI)
	err := srv.ListenAndServe()
	errLog.Fatal(err)
}
