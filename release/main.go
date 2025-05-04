package main

import (
	"api/album"
	"context"
	"embed"
	"encoding/json"
	"fmt"
	"golib/chrome"
	"golib/server"
	"golib/server/http"
	"log"
	"os"
)

//go:embed dist
var dist embed.FS

type Config struct {
	Threads          int     `json:"threads"`
	ThumbnailSize    int     `json:"thumbnail_size"`
	FaceMargin       float64 `json:"face_margin"`
	FaceSize         int     `json:"face_size"`
	ClusterThreshold float32 `json:"cluster_threshold"`
	ClipMinScore     float32 `json:"clip_min_score"`
}

func main() {
	f, err := os.Open("config.json")
	if err != nil {
		log.Fatal(err)
	}
	var conf Config
	json.NewDecoder(f).Decode(&conf)
	f.Close()

	album.ThumbnailSize = conf.ThumbnailSize
	album.FaceMargin = conf.FaceMargin
	album.FaceSize = conf.FaceSize
	if err = album.Initialize("album.db", 0x92, "models/det.bin", "models/rec.bin", "models/vocab.txt", "models/txt.bin", "models/img.bin", conf.Threads, conf.ClusterThreshold, conf.ClipMinScore, geoURL, geoToken); err != nil {
		log.Fatal(err)
	}
	http.MaxLength = 20 << 20
	http.Domains = []string{"*"}
	http.AllowHeaders = "*"
	l1, port1, err := server.ServeLocal(http.NewServer())
	if err != nil {
		log.Fatal(err)
	}
	defer l1.Close()
	apiURL := fmt.Sprintf("http://127.0.0.1:%d", port1)

	c, err := chrome.New(context.Background(), "", "", 800, 600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()

	c.Bind("fetchURL", func() string {
		return apiURL
	})

	l2, port2, err := server.ServeLocalFile(dist)
	if err != nil {
		log.Fatal(err)
	}
	defer l2.Close()
	c.Load(fmt.Sprintf("http://127.0.0.1:%d/dist", port2))
	c.DisableContextMenu()
	c.DisableDevTools()
	<-c.Done()
}
