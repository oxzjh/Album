package main

import (
	"api/album"
	"flag"
	"golib/server"
	"golib/server/http"
	"log"
	"time"
)

func main() {
	var (
		geoURL   string
		geoToken string
	)
	flag.StringVar(&geoURL, "g", "", "GEO URL")
	flag.StringVar(&geoToken, "t", "", "GEO token")
	flag.Parse()

	if err := album.Initialize("album.db", 0x92, "models/det.bin", "models/rec.bin", "models/vocab.txt", "models/txt.bin", "models/img.bin", 4, geoURL, geoToken); err != nil {
		log.Fatal(err)
	}
	http.MaxLength = 20 << 20
	http.Domains = []string{"*"}
	http.AllowHeaders = "*"
	http.ReturnErr = true
	log.Fatal(server.ServeHTTP(":3000", http.NewServer(), 5*time.Second))
}
