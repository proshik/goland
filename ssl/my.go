package main

import (
	"flag"
	"net/http"
	"fmt"
	"log"
	"time"
	"crypto/tls"
	"golang.org/x/crypto/acme/autocert"
)

var (
	httpPort = flag.String("addr", "8020", "TCP address to listen to for http")
	tlsDir   = flag.String("tlsDir", "", "Path to TLS certificate file and key ")
)

func redirectToHttps(w http.ResponseWriter, r *http.Request) {
	newURI := "https://" + r.Host + r.URL.String()
	http.Redirect(w, r, newURI, http.StatusFound)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Yeeees!11")
}

func main() {
	flag.Parse()

	http.HandleFunc("/", handler)
	// Запуск HTTPS сервера в отдельной go-рутине
	startHttpsServer()
	// Запуск HTTP сервера и редирект всех входящих запросов на HTTPS
	fmt.Printf("Starting HTTP server on port %s\n", httpPort)
	http.ListenAndServe(":"+*httpPort, http.HandlerFunc(redirectToHttps))
}

func startHttpsServer() {
	if len(*tlsDir) == 0 {
		log.Printf("-tlsAddr is empty, so skip serving https")
		return
	}

	httpsServer := &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		IdleTimeout:  120 * time.Second,
		Handler:      http.DefaultServeMux,
	}

	m := autocert.Manager{
		Prompt: autocert.AcceptTOS,
		Cache:  autocert.DirCache(*tlsDir),
	}

	httpsServer.Addr = ":443"
	httpsServer.TLSConfig = &tls.Config{GetCertificate: m.GetCertificate}

	go func() {
		fmt.Printf("Starting HTTPS server on %s\n", httpsServer.Addr)
		err := httpsServer.ListenAndServeTLS("", "")
		if err != nil {
			log.Fatalf("httpsSrv.ListendAndServeTLS() failed with %s", err)
		}
	}()
}
