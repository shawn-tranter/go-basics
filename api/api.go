package api

import (
	"encoding/json"
	"fmt"
	"log"
)

var (
	name    string = "-"
	version string = "-"
)

type APIRecord struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

var api *APIRecord

func (api *APIRecord) String() string {
	return fmt.Sprintf("Name: %s, Version: %s", api.Name, api.Version)
}

func (api *APIRecord) JSON() (result string) {
	if apiMap, err := json.Marshal(api); err == nil {
		result = fmt.Sprintf(string(apiMap))
	} else {
		log.Fatal(err)
	}

	return
}

func Get() *APIRecord {
	if api == nil {
		api = &APIRecord{name, version}
	}

	return api
}
