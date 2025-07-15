// myip is a an HTTP server that responds with client's IP.
package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

func main() {
	port := flag.Int("port", 8080, "Application port")
	flag.Parse()

	ctx, cancel := signal.NotifyContext(
		context.Background(),
		os.Interrupt,
		syscall.SIGTERM,
	)
	defer cancel()

	// Create HTTP server
	addr := fmt.Sprintf(":%d", *port)
	mux := http.NewServeMux()
	mux.HandleFunc("/", handle)
	//nolint:gosec
	srv := http.Server{
		Addr:    addr,
		Handler: mux,
	}

	// Run HTTP server
	go func() {
		log.Printf("Listening on %s...", addr)
		err := srv.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Printf("HTTP server error: %v", err)
		}
	}()

	// Wait for SIGTERM/SIGINT
	<-ctx.Done()

	// Shutdown gracefully
	ctx, cancel = context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Printf("HTTP server shutdown error: %v", err)
	}
	log.Print("Shutdown gracefully")
}

func handle(w http.ResponseWriter, r *http.Request) {
	headers := []string{
		"Forwarded",
		"X-Client-IP",
		"X-Forwarded-For",
		"X-Real-IP",
	}

	var ip string
	for _, h := range headers {
		val := r.Header.Get(h)
		if val != "" {
			// Headers may contain multiple IPs
			ip = strings.Split(val, ",")[0]
			break
		}
	}
	// Fallback to RemoteAddr
	if ip == "" {
		host, _, err := net.SplitHostPort(r.RemoteAddr)
		if err == nil {
			ip = host
		} else {
			ip = r.RemoteAddr
		}
	}

	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte(ip + "\n")) //nolint:errcheck,gosec
}
