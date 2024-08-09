package server

import (
	"bytes"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"

	"github.com/VesuvyX/2-image-generator-server/configs"
	"github.com/VesuvyX/2-image-generator-server/pkg/img"
)

func rend(w http.ResponseWriter, msg string) {
	_, err := w.Write([]byte(msg))
	if err != nil {
		log.Println(err)
	}
}

func rendImg(w http.ResponseWriter, buffer *bytes.Buffer) {
	w.Header().Set("Content-Type", "image/jpeg")
	w.Header().Set("Content-Length", strconv.Itoa(len(buffer.Bytes()))) // Content-Length - длина нашей картинки в байтах
	if _, err := w.Write(buffer.Bytes()); err != nil {
		log.Fatalln(err)
	}
}

func imgHandler(w http.ResponseWriter, r *http.Request) {
	buffer, err := img.Generate(strings.Split(r.URL.Path, "/"))
	if err != nil {
		log.Println(err)
	}
	rendImg(w, buffer)
}

func faviconHandler(w http.ResponseWriter, r *http.Request) {
	buffer, err := img.GenerateFavicon()
	if err != nil {
		log.Fatalln(err)
	}
	rendImg(w, buffer)
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	rend(w, "PONG")
}

func robotsHandler(w http.ResponseWriter, r *http.Request) {
	rend(w, "robots")
}

func Run(conf configs.ConfI) {
	http.HandleFunc("/", imgHandler)
	http.HandleFunc("/favicon.ico", faviconHandler)
	http.HandleFunc("/ping", pingHandler)
	http.HandleFunc("/robots.txt", robotsHandler)

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	// горутина
	go func() {
		log.Println("Server starting ...")
		if err := http.ListenAndServe(":"+conf.GetPort(), nil); err != nil {
			log.Fatalln(err)
		}
	}()

	// Принимаем из канала sigs
	signalValue := <-sigs
	signal.Stop(sigs)

	// Будет понятно, почему сервер остановился
	log.Println("stop signal", signalValue)
}
