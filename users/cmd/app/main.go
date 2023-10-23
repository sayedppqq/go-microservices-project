package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	opt := connectionOptions{}
	opt.pasrseCommandLineFlags()
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	client, err := opt.GetMongoClient(ctx)
	if err != nil {
		errLog.Fatal(err)
	}
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	infoLog.Println("Database connection established")

	app := opt.GetApplication(errLog, infoLog, client)

	serverURI := fmt.Sprintf("%v:%d", *opt.serverAddr, *opt.serverPort)
	srv := opt.GetHTTPServer(errLog, serverURI, app)

	infoLog.Println("starting users server on ", serverURI)

	err = srv.ListenAndServe()
	errLog.Fatal(err)
}
