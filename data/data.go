package main

import (
	"encoding/json"
	"errors"
	"log/slog"
	"os"
	"path/filepath"

	"github.com/go-drift/drift/pkg/platform"
)

type Data struct {
	Items []string
}

func (data *Data) Load() error {
	slog.Info("Data: Load()")
	cachePath, err := platform.Storage.GetAppDirectory(platform.AppDirectoryCache)
	if err != nil {
		slog.Info("Data: Error: " + err.Error())
		return err
	}

	fname := filepath.Join(cachePath, "data.json")
	f, err := os.Open(fname)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			slog.Info("Data: No existing data to load")
			return nil
		}
		slog.Info("Data: Error: " + err.Error())
		return err
	}
	defer f.Close()

	var d Data
	dec := json.NewDecoder(f)
	err = dec.Decode(&d)
	if err != nil {
		slog.Info("Data: Error: " + err.Error())
		return err
	}
	data.Items = d.Items
	return nil
}

func (data *Data) Save() error {
	cachePath, err := platform.Storage.GetAppDirectory(platform.AppDirectoryCache)
	if err != nil {
		return err
	}

	f, err := os.CreateTemp(cachePath, "data.")
	if err != nil {
		return err
	}
	defer func() {
		f.Close()
		os.Remove(f.Name())
	}()

	enc := json.NewEncoder(f)
	err = enc.Encode(data)
	if err != nil {
		return err
	}
	fname := filepath.Join(cachePath, "data.json")
	return os.Rename(f.Name(), fname)

}
