package app

import (
	"encoding/json"
	"fmt"
	"log"
)

var (
	name    string = "-"
	version string = "-"
)

type ApplicationRecord struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

var application *ApplicationRecord

func (app *ApplicationRecord) String() string {
	return fmt.Sprintf("Name: %s, Version: %s", app.Name, app.Version)
}

func (app *ApplicationRecord) JSON() (result string) {
	if appMap, err := json.Marshal(app); err == nil {
		result = fmt.Sprintf(string(appMap))
	} else {
		log.Fatal(err)
	}

	return
}

func Get() *ApplicationRecord {
	if application == nil {
		application = &ApplicationRecord{name, version}
	}

	return application
}
