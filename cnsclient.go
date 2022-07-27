package main

import (
	"os"
	"os/signal"
	"syscall"
	
	"github.com/Moonyongjung/cns-client/server"	
	"github.com/Moonyongjung/cns-client/util"
)

var configPath = "./config/config.json"

func init() {
	util.GetConfig().Read(configPath)	
}

func main() {	
	go server.RunHttpServer()
	
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop

	util.LogGw("Shutting down the client...")
	util.LogGw("Client gracefully stopped")	
}