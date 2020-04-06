package main

import (
	"io/ioutil"
	"net/http"
	"time"

	"github.com/cryptag/gosecure/content"
	"github.com/cryptag/gosecure/csp"
	"github.com/cryptag/gosecure/frame"
	"github.com/cryptag/gosecure/hsts"
	"github.com/cryptag/gosecure/referrer"
	"github.com/cryptag/gosecure/xss"

	log "github.com/sirupsen/logrus"
	"github.com/gorilla/mux"
	"github.com/justinas/alice"
	"golang.org/x/crypto/acme/autocert"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()

	// Hack to make up for the fact that
	//   r.NotFoundHandler = http.HandlerFunc(GetIndex)
	// doesn't do anything, since the below
	//   r.PathPrefix("/").Handler(...)
	// call returns its own 404, ignoring the value of
	//   r.NotFoundHandler

	urlPrefixes := []string{
		"/about",
		"/step",
	}
	for _, prefix := range urlPrefixes {
		r.PathPrefix(prefix).HandlerFunc(GetIndex)
	}

	handleBuildDir := http.FileServer(http.Dir("./public"))

	r.PathPrefix("/").Handler(gzipHandler(handleBuildDir)).Methods("GET")

	http.Handle("/", r)
	return r
}

func NewServer(listenAddr, domain string) *http.Server {
	r := NewRouter()

	middleware := alice.New(
		csp.GetCustomHandlerStyleUnsafeInline(domain, "www."+domain),
		hsts.PreloadHandler,
		frame.GetHandler("www."+domain),
		content.GetHandler,
		xss.GetHandler,
		referrer.NoHandler,
	)

	return &http.Server{
		Addr:         listenAddr,
		ReadTimeout:  1000 * time.Second,
		WriteTimeout: 1000 * time.Second,
		IdleTimeout:  120 * time.Second,
		Handler:      middleware.Then(r),
	}
}

func ProductionServer(httpsAddr string, domain string, manager *autocert.Manager) *http.Server {
	srv := NewServer(httpsAddr, domain)

	// Enable TLS
	srv.Handler = manager.HTTPHandler(srv.Handler)
	srv.TLSConfig = manager.TLSConfig()

	return srv
}

func GetIndex(w http.ResponseWriter, req *http.Request) {
	contents, err := ioutil.ReadFile("public/index.html")
	if err != nil {
		log.Errorf("Error serving index.html: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error: couldn't serve you index.html!"))
		return
	}
	w.Write(contents)
}

func redirectToHTTPS(httpAddr, myDomain, httpsPort string) {
	srv := &http.Server{
		Addr:         httpAddr,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		Handler: http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			w.Header().Set("Connection", "close")
			url := "https://" + myDomain + ":" + httpsPort + req.URL.Path
			http.Redirect(w, req, url, http.StatusFound)
		}),
	}
	log.Infof("Listening on %v", httpAddr)
	log.Fatal(srv.ListenAndServe())
}

func getAutocertManager(domain string) *autocert.Manager {
	return &autocert.Manager{
		Prompt:     autocert.AcceptTOS,
		HostPolicy: autocert.HostWhitelist(domain),
		Cache:      autocert.DirCache("./" + domain),
	}
}
