package main

import (
	"flag"
	"fmt"
	"strings"

	log "github.com/sirupsen/logrus"
)

func main() {
	httpAddr := flag.String("http", ":80",
		"Address to listen on HTTP (will redirect to HTTPS port)")
	httpsAddr := flag.String("https", ":443",
		"Address to listen on HTTPS")
	domain := flag.String("domain", "",
		"Domain of this service")
	torPort := flag.Int("tor-port", 5002,
		"HTTP port to listen on for Tor Onion service traffic")
	torDomain := flag.String("tor-domain", "",
		"Domain to listen on for Onion service you set up on localhost")
	flag.Parse()

	log.SetLevel(log.ErrorLevel)

	if *domain == "" {
		log.Fatal("You must specify a -domain; e.g., -domain runarelay.com")
	}

	// Start HTTP server listening for incoming HTTP traffic to *torDomain
	go torServer(*torPort, *torDomain)

	// Set up HTTP -> HTTPS redirection
	httpsPort := strings.SplitN(*httpsAddr, ":", 2)[1]
	go redirectToHTTPS(*httpAddr, *domain, httpsPort)

	// HTTPS server
	manager := getAutocertManager(*domain)
	srv := ProductionServer(*httpsAddr, *domain, manager)
	log.Infof("Listening for HTTPS traffic @ %v", *httpsAddr)
	log.Fatal(srv.ListenAndServeTLS("", ""))
}

func torServer(torPort int, torDomain string) {
	if torDomain == "" {
		log.Errorln(`HTTP server listening for .onion traffic NOT started! Consider setting up an Onion service on this server, then adding

    -tor-domain myoniondomainabc123xyzabc123xyzabc123xyz.onion`)
		return
	}

	listenAddr := fmt.Sprintf("127.0.0.1:%d", torPort)
	torSrv := NewServer(listenAddr, torDomain)
	log.Infof("Listening on %s for traffic from %v", listenAddr, torDomain)
	log.Fatal(torSrv.ListenAndServe())
}
