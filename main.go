package main

import (
	"flag"
	"fmt"

	log "github.com/sirupsen/logrus"
)

func main() {
	listenPort := flag.Int("l", 5001,
		"HTTP port to listen on when not in standalone mode; useful for development and when receiving proxied traffic from Nginx/Tor daemon/etc")
	standalone := flag.Bool("standalone", false, "Run in standalone mode: listen on :443, redirect from :80 to 443, auto-renew SSL certs")
	domain := flag.String("domain", "",
		"Domain of this service; required when running in standalone mode")
	flag.Parse()

	if *standalone {
		log.SetLevel(log.ErrorLevel)

		if *domain == "" {
			log.Fatal("You must specify a -domain; e.g., -domain runarelay.org")
		}

		go redirectToHTTPS(":80", *domain, "443")
		go httpsServer(*domain)
	}

	localServer(*listenPort)
}

func httpsServer(domain string) {
	manager := getAutocertManager(domain)
	srv := ProductionServer(":443", domain, manager)
	log.Infof("Listening for HTTPS traffic @ %v")
	log.Fatal(srv.ListenAndServeTLS("", ""))
}

func localServer(port int) {
	listenAddr := fmt.Sprintf("127.0.0.1:%d", port)
	srv := NewServer(listenAddr, "")
	log.Infof("Listening on %s", listenAddr)
	log.Fatal(srv.ListenAndServe())
}
