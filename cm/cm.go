package cm

import (
	"encoding/json"
	"fmt"
	"github.com/shawn-tranter/go-basics/app"
	"github.com/shawn-tranter/go-basics/bld"
	"log"
)

type ConfigManagementRecord struct {
	Application *app.ApplicationRecord `json:"application"`
	Build       *bld.BuildRecord       `json:"build"`
}

var cmRecord *ConfigManagementRecord

func (cm *ConfigManagementRecord) String() string {
	return fmt.Sprintf("Application -> %s, Build -> %s", cm.Application.String(), cm.Build.String())
}

func (cm *ConfigManagementRecord) JSON() (result string) {
	if cmMap, err := json.Marshal(cm); err == nil {
		result = fmt.Sprintf(string(cmMap))
	} else {
		log.Fatal(err)
	}

	return
}

func Get() *ConfigManagementRecord {
	if cmRecord == nil {
		cmRecord = &ConfigManagementRecord{app.Get(), bld.Get()}
	}

	return cmRecord
}
