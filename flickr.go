package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"sync"
	"time"

	"github.com/azer/go-flickr"
)

var (
	photosMutex sync.RWMutex
	photos      []Photo
)

func StartUpdatingPhotos(c config) {
	client := &flickr.Client{
		Key: c.Key,
	}

	go func() {
		updatePhotos(client, c.PhotoAlbum)

		ticker := time.NewTicker(time.Minute * time.Duration(c.RefreshInterval))
		for _ = range ticker.C {
			updatePhotos(client, c.PhotoAlbum)
		}
	}()
}

func updatePhotos(client *flickr.Client, albumId string) {

	log.Println("Fetching latest photos from flickr")

	album, err := client.Album(albumId)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error fetching album (%s).\n", err)
		return
	}

	log.Printf("Fetched %d photos", len(album.Photos))

	newPhotos := make([]Photo, len(album.Photos))
	for i, photo := range album.Photos {
		newPhotos[i] = Photo{
			Title:  photo.Title,
			URL:    fmt.Sprintf("https://www.flickr.com/photos/%s/%s", album.Owner, photo.Id),
			Square: flickr.GenerateURL("q", photo.Id, photo.Farm, photo.Secret, photo.Server, "jpg"),
			Small:  flickr.GenerateURL("z", photo.Id, photo.Farm, photo.Secret, photo.Server, "jpg"),
			Medium: flickr.GenerateURL("b", photo.Id, photo.Farm, photo.Secret, photo.Server, "jpg"),
			Large:  flickr.GenerateURL("h", photo.Id, photo.Farm, photo.Secret, photo.Server, "jpg"),
		}
	}

	setPhotos(newPhotos)
}

func GetPhotos(count int) []Photo {
	results := make([]Photo, count)

	photosMutex.RLock()
	defer photosMutex.RUnlock()

	randomIndexes := rand.Perm(len(photos))
	for i := range results {
		randomIndex := randomIndexes[i]
		results[i] = photos[randomIndex]
	}

	return results
}

func setPhotos(newPhotos []Photo) {
	photosMutex.Lock()
	defer photosMutex.Unlock()

	photos = newPhotos
}
