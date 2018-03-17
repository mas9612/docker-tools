package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

type imageInfo struct {
	ID      string
	Created int64
	Name    string
}

func main() {
	imageName, generation, force := parseArgs()
	var imageInfos []imageInfo

	ctx := context.Background()
	client, err := client.NewEnvClient()
	if err != nil {
		log.Fatalf("[ERROR] client.NewEnvClient(): %s\n", err)
	}

	images, err := client.ImageList(ctx, types.ImageListOptions{})
	if err != nil {
		log.Fatalf("[ERROR] client.ImageList(): %s\n", err)
	}

	for _, image := range images {
		for _, repotag := range image.RepoTags {
			repository := strings.Split(repotag, ":")
			if repository[0] == *imageName {
				imageInfos = append(imageInfos, imageInfo{
					ID:      image.ID,
					Created: image.Created,
					Name:    repotag,
				})
			}
		}
	}

	sort.Slice(imageInfos, func(i, j int) bool { return imageInfos[i].Created > imageInfos[j].Created })

	if *generation > len(imageInfos) {
		*generation = len(imageInfos)
	}
	removeOptions := types.ImageRemoveOptions{
		Force: *force,
	}
	for _, image := range imageInfos[*generation:] {
		_, err := client.ImageRemove(ctx, image.ID, removeOptions)
		if err != nil {
			log.Fatalf("[ERROR] client.ImageRemove(): %s\n", err)
		}
		fmt.Printf("Image %s was deleted.\n", image.Name)
	}
}

func parseArgs() (*string, *int, *bool) {
	imageName := flag.String("name", "", "Image name that you want to delete")
	generation := flag.Int("generation", 0, "Delete images older than this generation")
	force := flag.Bool("f", false, "Force removal of the image")
	flag.Parse()

	if *imageName == "" {
		fmt.Fprint(os.Stderr, "Few arguments: You must specify image name that you want to delete.\nPlease see help.\n")
		os.Exit(1)
	}

	return imageName, generation, force
}
