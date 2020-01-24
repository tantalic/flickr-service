package main

import (
	"fmt"
	"os"

	"github.com/kelseyhightower/envconfig"
)

const (
	Version = "1.1.1"
)

type config struct {
	// Flickr
	PhotoAlbum      string `envconfig:"FLICKR_ALBUM" required:"true"`
	Key             string `envconfig:"FLICKR_KEY" required:"true"`
	RefreshInterval int    `envconfig:"REFRESH_INTERVAL" default:15"` // in minutes

	// API Server
	Host string `envconfig:"HOST" default:""`
	Port int    `envconfig:"PORT" default:"3000"`
}

func main() {
	c, err := getConfig()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading configuration from environment (%s).\n", err)
		os.Exit(1)
	}

	StartUpdatingPhotos(c)
	err = StartApiServer(c)
	fmt.Fprintf(os.Stderr, "Error starting http server: %s\n", err)
}

func getConfig() (config, error) {
	var c config
	err := envconfig.Process("PHOTO_SERVICE", &c)
	return c, err
}
