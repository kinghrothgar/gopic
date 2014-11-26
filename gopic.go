package main

import (
	//"github.com/bmizerany/pat"
	"github.com/kinghrothgar/gopic/conf"
	"github.com/kinghrothgar/gopic/handler"
	"github.com/grooveshark/golib/gslog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

// To be set at build
var buildCommit string
var buildDate string

func listenAndServer(addr string, c chan error) {
	err := http.ListenAndServe(addr, nil)
	c <- err
}

func main() {
	//if conf.ShowVers {
	//	println("Commit: " + buildCommit)
	//	println("Date:   " + buildDate)
	//	os.Exit(0)
	//}

	gslog.Info("gopic started [build commit: %s, build date: %s]", buildCommit, buildDate)
	config, err := conf.Parse()
	if err != nil {
		gslog.Fatal("MAIN: failed to parse conf with error: %s", err.Error())
	}

	gslog.SetMinimumLevel(config.GetStr("loglevel"))
	if logFile := config.GetStr("logfile"); logFile != "" {
		gslog.SetLogFile(logFile)
	}

	cachePath := config.GetStr("cachepath")
	imgPath := config.GetStr("imagepath")

	// Setup route handlers
	//mux := pat.New()
	//mux.Get("/", http.HandlerFunc(handler.GetRoot))
	//mux.Get("/:album", http.HandlerFunc(handler.GetAlbum))

	http.Handle("/", http.HandlerFunc(handler.GetPage))

	// Server raw images
	http.Handle("/i/", http.StripPrefix("/i", http.FileServer(http.Dir(imgPath))))
	http.Handle("/c/", http.StripPrefix("/c", http.FileServer(http.Dir(cachePath))))

	listenOn := config.GetStr("listen")
	gslog.Info("MAIN: Listening on %s...", listenOn)
	c := make(chan error)
	go listenAndServer(listenOn, c)

	// Set up listening for os signals
	shutdownCh := make(chan os.Signal, 5)
	// TODO: What signals for Windows if any?
	signal.Notify(shutdownCh, syscall.SIGINT, syscall.SIGKILL)
	// Set up listening for os signals
	for {
		select {
		case <-shutdownCh:
			gslog.Info("MAIN: Syscall recieved, shutting down...")
			gslog.Flush()
			os.Exit(0)
		case err := <-c:
			gslog.Error("MAIN: ListenAndServe: %s", err)
			gslog.Fatal("MAIN: Failed to start server, exiting...")
		}
	}
}
