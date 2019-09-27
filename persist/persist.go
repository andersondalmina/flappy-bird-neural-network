package persist

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"os"
	"path/filepath"
	"sync"
)

var lock sync.Mutex

func marshal(v interface{}) (io.Reader, error) {
	b, err := json.MarshalIndent(v, "", "\t")
	if err != nil {
		return nil, err
	}

	return bytes.NewReader(b), nil
}

func unmarshal(r io.Reader, v interface{}) error {
	return json.NewDecoder(r).Decode(v)
}

// GetRootDir return root directory path
func GetRootDir() string {
	dir, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}

	return filepath.Dir(dir)
}

// CheckFileExists verify if file exists
func CheckFileExists(path string) error {
	_, err := os.Stat(GetRootDir() + "/" + path)
	return err
}

// Save saves to the file
func Save(path string, v interface{}) error {
	lock.Lock()
	defer lock.Unlock()

	f, err := os.Create(GetRootDir() + "/" + path)
	if err != nil {
		return err
	}

	defer f.Close()

	r, err := marshal(v)
	if err != nil {
		return err
	}

	_, err = io.Copy(f, r)
	return err
}

// Load loads the file at path
func Load(path string, v interface{}) error {
	lock.Lock()
	defer lock.Unlock()

	f, err := os.Open(GetRootDir() + "/" + path)
	if err != nil {
		return err
	}

	defer f.Close()

	return unmarshal(f, v)
}
