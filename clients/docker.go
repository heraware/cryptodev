package clients

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/boltdb/bolt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
)

// Docker struct to manage all CryutoDev action with Docker
type Docker struct {
	Client *client.Client
}

var images = map[string]string{
	"bitcoin":  "heraware/bitcoin:latest",
	"litecoin": "heraware/litecoin:latest",
}

var ports = map[string]nat.PortMap{
	"bitcoin": map[nat.Port][]nat.PortBinding{
		"20001/tcp": []nat.PortBinding{
			nat.PortBinding{HostIP: "0.0.0.0", HostPort: "20001"},
		},
		"20000/tcp": []nat.PortBinding{
			nat.PortBinding{HostIP: "0.0.0.0", HostPort: "20000"},
		},
	},
	"litecoin": map[nat.Port][]nat.PortBinding{
		"21001/tcp": []nat.PortBinding{
			nat.PortBinding{HostIP: "0.0.0.0", HostPort: "21001"},
		},
		"21000/tcp": []nat.PortBinding{
			nat.PortBinding{HostIP: "0.0.0.0", HostPort: "21000"},
		},
	},
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

func (d *Docker) stopContainer(containerID string) error {
	timeout := 10 * time.Second
	return d.Client.ContainerStop(context.Background(), containerID, &timeout)
}

func (d *Docker) createAndRunContainer(name string, image string) {
	containerName := fmt.Sprintf("cryptodev-%s", name)
	if d.containerExists(containerName) {
		log.Fatalf("Container: %s already exists", containerName)
	}
	containerConfig := container.Config{
		Image: images[name],
	}
	hostConfig := container.HostConfig{
		PortBindings: ports[name],
		Privileged:   false,
	}
	containerBody, err := d.Client.ContainerCreate(context.Background(), &containerConfig, &hostConfig, nil, containerName)
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

func (d *Docker) StopNode(name string) {
	containerName := fmt.Sprintf("cryptodev-%s", name)
	containerID, err := d.getContainerInfo(containerName)
	if err != nil {
		panic(err)
	}
	containerIDString := string(containerID)
	err = d.stopContainer(containerIDString)
	if err != nil {
		log.Fatalf("Container: %s doesn't exists, try to run `cryptodev create %s`", containerName, name)
	}
	fmt.Println(fmt.Sprintf("Node: %s was stopped", containerName))
}
