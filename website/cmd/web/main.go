package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	opt := connectionOptions{}
	opt.pasrseCommandLineFlags()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	app := opt.GetApplication(infoLog, errLog)

	serverURI := fmt.Sprintf("%s:%d", *opt.serverAddr, *serverPort)

}
