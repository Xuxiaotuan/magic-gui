package util

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/wader/goutubedl"
)

func DownloadViedo(ytdlpPatch string, source string, target string) {
	goutubedl.Path = ytdlpPatch
	result, err := goutubedl.New(context.Background(), source, goutubedl.Options{})
	if err != nil {
		log.Fatal(err)
	}
	downloadResult, err := result.Download(context.Background(), "best")
	if err != nil {
		log.Fatal(err)
	}
	defer downloadResult.Close()
	f, err := os.Create(target)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	io.Copy(f, downloadResult)
}
