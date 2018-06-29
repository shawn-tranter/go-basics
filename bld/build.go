package bld

import (
	"encoding/json"
	"fmt"
	"log"
)

var (
	buildTime string = "-"
	gitHash   string = "-"
)

type BuildRecord struct {
	Time    string `json:"time"`
	GitHash string `json:"git_hash"`
}

var build *BuildRecord

func (bld *BuildRecord) String() string {
	return fmt.Sprintf("Time %s, Git Hash: %s", bld.Time, bld.GitHash)
}

func (bld *BuildRecord) JSON() (result string) {
	if buildMap, err := json.Marshal(bld); err == nil {
		result = fmt.Sprintf(string(buildMap))
	} else {
		log.Fatal(err)
	}

	return
}

func Get() *BuildRecord {
	if build == nil {
		build = &BuildRecord{buildTime, gitHash}
	}

	return build
}
