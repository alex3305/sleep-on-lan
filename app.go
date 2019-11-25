package main

import (
	"os"
	"fmt"
	"net/http"
	"log"
	"runtime"
	"github.com/julienschmidt/httprouter"
)

var ps PowerRequester

func determineOS() {
	if runtime.GOOS == "linux" {
		if os.Getegid() != 0 {
			log.Fatalln("This application has to be ran as root")
		}
		ps = newSystemCtl()
	} else if runtime.GOOS == "windows" {
		log.Fatalln("Windows isn't currently supported")
	} else {
		log.Fatalln("Other OS's aren't supported")
	}
}

func info(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "Application.json")
	fmt.Fprintf(w, "{'Loaded': '%s'}", ps.getApplication())
}

func requestSuspend(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "Application.json")
	fmt.Fprint(w, "{'Request': 'OK'}")
	ps.suspend()
}

func requestHibernate(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "Application.json")
	fmt.Fprint(w, "{'Request': 'OK'}")
	ps.hibernate()
}

func requestReboot(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "Application.json")
	fmt.Fprint(w, "{'Request': 'OK'}")
	ps.reboot()
}

func requestPoweroff(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "Application.json")
	fmt.Fprint(w, "{'Request': 'OK'}")
	ps.poweroff()
}

func main() {
	log.Println("Starting Sleep on LAN application...")
	determineOS()
	log.Println(fmt.Printf("Determined power system as %s", ps.getApplication()))

	router := httprouter.New()
	
	router.GET("/", info)
	router.GET("/suspend", requestSuspend)
	router.GET("/hibernate", requestHibernate)
	router.GET("/reboot", requestReboot)
	router.GET("/poweroff", requestPoweroff)

	log.Println("Starting web server on 0.0.0.0:53305")
	log.Fatal(http.ListenAndServe(":53305", router))
}
