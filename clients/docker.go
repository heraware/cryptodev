package clients

import (
	"context"
	"fmt"
	"log"

	"github.com/boltdb/bolt"
	"github.com/docker/docker/api/types/container"
	"github.com/moby/moby/client"
)

// Docker struct to manage all CryutoDev action with Docker
type Docker struct {
	Client *client.Client
}

var images = map[string]string{
	"bitcoin":  "heraware/bitcoin:latest",
	"litecoin": "heraware/litecoin:latest",
}

// NewDockerClient Return a Docker instance with Docker client
// configured from OS ENV
func NewDockerClient() *Docker {
	client, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}
	docker := Docker{Client: client}
	return &docker
}

func (d *Docker) getContainerInfo(containerName string) ([]byte, error) {
	var result []byte
	if err := DB.View(func(tx *bolt.Tx) error {
		txBucket := tx.Bucket([]byte("containers"))
		if txBucket == nil {
			return fmt.Errorf("Bucket doesn't exists")
		}
		result = txBucket.Get([]byte(containerName))
		return nil
	}); err != nil {
		return nil, err
	}
	return result, nil
}

func (d *Docker) saveContainerInfo(containerName string, containerID string) error {
	if err := DB.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte("containers"))
		if err != nil {
			return err
		}
		if err := b.Put([]byte(containerName), []byte(containerID)); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return err
	}
	return nil
}

func (d *Docker) containerExists(containerName string) bool {
	_, err := d.getContainerInfo(containerName)
	if err != nil {
		return false
	}
	return true
}

func (d *Docker) createAndRunContainer(name string, image string) *container.ContainerCreateCreatedBody {
	containerName := fmt.Sprintf("cryptodev-%s", name)
	if d.containerExists(containerName) {
		log.Fatalf("Container: %s already exists", containerName)
	}
	containerConfig := container.Config{
		Image: images[name],
	}
	containerBody, err := d.Client.ContainerCreate(context.Background(), &containerConfig, nil, nil, containerName)
	if err != nil {
		panic(err)
	}
	d.saveContainerInfo(containerName, containerBody.ID)
	return &containerBody
}

// CreateNode create a container with the cryptocurrency client
func (d *Docker) CreateNode(name string) {
	d.createAndRunContainer(name, images[name])
}
