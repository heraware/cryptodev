package clients

import (
	"fmt"
	"log"
	"os"
	"runtime"

	"github.com/boltdb/bolt"
)

// DB pointer of bolt DB
var DB *bolt.DB

func userHomeDir() string {
	if runtime.GOOS == "windows" {
		home := os.Getenv("HOMEDRIVE") + os.Getenv("HOMEPATH")
		if home == "" {
			home = os.Getenv("USERPROFILE")
		}
		return home
	}
	return os.Getenv("HOME")
}

func initializeDB() {
	dir := userHomeDir()
	var err error
	DB, err = bolt.Open(fmt.Sprintf("%s/.crytodev.db", dir), 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func init() {
	initializeDB()
}
