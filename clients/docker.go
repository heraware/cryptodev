package clients

import (
	"context"
	"fmt"
	"log"

	"github.com/boltdb/bolt"
	"github.com/docker/docker/api/types"
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

func (d *Docker) deleteContainerInfo(containerName string) error {
	return DB.Update(func(tx *bolt.Tx) error {
		txBucket := tx.Bucket([]byte("containers"))
		if txBucket == nil {
			return fmt.Errorf("Bucket doesn't exists")
		}
		return txBucket.Delete([]byte(containerName))
	})
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
	value, err := d.getContainerInfo(containerName)
	if err != nil {
		return false
	}
	if len(value) == 0 {
		return false
	}
	return true
}

func (d *Docker) runContainer(containerID string) error {
	return d.Client.ContainerStart(context.Background(), containerID, types.ContainerStartOptions{})
}

func (d *Docker) createAndRunContainer(name string, image string) {
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
	if err := d.runContainer(containerBody.ID); err != nil {
		panic(err)
	}
	fmt.Println(fmt.Sprintf("Node: %s is created and running - Container ID: %s", containerName, containerBody.ID))
}

// CreateNode create a container with the cryptocurrency client
func (d *Docker) CreateNode(name string) {
	d.createAndRunContainer(name, images[name])
}

func (d *Docker) RunNode(name string) {
	containerName := fmt.Sprintf("cryptodev-%s", name)
	containerID, err := d.getContainerInfo(containerName)
	if err != nil {
		panic(err)
	}
	containerIDString := string(containerID)
	err = d.runContainer(containerIDString)
	if err != nil {
		fmt.Println(d.deleteContainerInfo(containerName))
		log.Fatalf("Container: %s doesn't exists, try to run `cryptodev create %s`", containerName, name)
	}
	fmt.Println(fmt.Sprintf("Node: %s is running - Container ID: %s", containerName, containerIDString))
}
