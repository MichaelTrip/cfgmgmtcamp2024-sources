package main

import (
	"fmt"
	"io"
	"log"
	"math/rand"
	"net"
	"net/http"
	"os"
	"os/signal"

	// "path/filepath"
	"syscall"
	"time"
)

const (
	message = "Demo app"
	logoURL = "https://raw.githubusercontent.com/kubernetes/kubernetes/master/logo/logo.png"
	// Logosize =  height=\"200\" width=\"200\"
	// logoURL = "https://cfgmgmtcamp.eu/images/logo.png" #
	// Logo size =  height=\"129\" width=\"355\"
)

var (
	addr     = ":8080"
	hostname string
	env      string
	function string
	commit   string
)

var quotes = []string{
	"Software is like sex: it's better when it's free. - Linus Torvalds",
	"Programs must be written for people to read, and only incidentally for machines to execute. - Harold Abelson",
	"Ethical software is software that respects users' freedom and community. To be ethical, software must be free as in freedom. - Richard Stallman",
	"Your work is going to fill a large part of your life, and the only way to be truly satisfied is to do what you believe is great work. And the only way to do great work is to love what you do. - Steve Jobs",
	"Good design is as little design as possible. - Dieter Rams",
	"Any sufficiently advanced technology is indistinguishable from magic. - Arthur C. Clarke",
	"Innovation distinguishes between a leader and a follower. - Steve Jobs",
	"Without requirements or design, programming is the art of adding bugs to an empty text file. - Louis Srygley",
	"To me, mathematics, computer science, and the arts are insanely related. They're all creative expressions. - Sebastian Thrun",
	"This is awesome! - Michael Trip",
	"Ingress testing! - Michael Trip",
	"Added some more quotes for testing - Michael Trip",
}

func init() {
	rand.Seed(time.Now().UnixNano())

	hostname, _ = os.Hostname()
	env = os.Getenv("ENV")
	if env == "" {
		env = "default"
	}
	function = os.Getenv("FUNCTION")
	if function == "" {
		function = "not specified"
	}

	// Create the static directory if it doesn't exist
	if _, err := os.Stat("static"); os.IsNotExist(err) {
		if err := os.Mkdir("static", 0755); err != nil {
			log.Fatal(err)
		}
	}
}

func clientIP(r *http.Request) string {
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return ""
	}
	userIP := net.ParseIP(ip)
	if userIP == nil {
		return ""
	}
	return userIP.String()
}

func downloadFile(url string, filepath string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}

func deleteStaticDir() {
	if _, err := os.Stat("static"); !os.IsNotExist(err) {
		if err := os.RemoveAll("static"); err != nil {
			log.Println(err)
		}
	}
}

func getVersion() string {
	if commit == "" {
		return "No version found"
	}
	return commit
}

func main() {
	log.SetOutput(os.Stdout) // Set log output to stdout

	// Download the Kubernetes logo
	if err := downloadFile(logoURL, "static/logo.png"); err != nil {
		log.Fatal(err)
	}
	// Handle stop signal to delete the static directory
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigCh
		log.Println("Received stop signal. Deleting static directory...")
		deleteStaticDir()
		os.Exit(0)
	}()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")

		quote := quotes[rand.Intn(len(quotes))]
		log.Printf("%s - %s %s %s", clientIP(r), r.Method, r.URL.Path, r.Proto)

		fmt.Fprintf(w, "<html><head><link rel=\"icon\" href=\"/favicon.ico\"><title>Golang Application version %s</title></head><body><img src=\"/static/logo.png\" alt=\"Kubernetes logo\" height=\"200\" width=\"200\"><h1>%s</h1><h2>Version %s</h2><p>Hostname: %s</p><p>Environment: %s</p><p>Function: %s</p><p>Client IP: %s</p><p>Headers:</p><ul>", getVersion(), message, getVersion(), hostname, env, function, clientIP(r))
		for name, values := range r.Header {
			for _, value := range values {
				fmt.Fprintf(w, "<li>%s: %s</li>", name, value)
			}
		}
		fmt.Fprintf(w, "</ul><p>Quote of the day:</p><blockquote>%s</blockquote></body></html>", quote)
	})

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/logo.png")
	})

	log.Printf("Starting server on %s...\n", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal(err)
	}
}
