package cm

import (
	"encoding/json"
	"fmt"
	"github.com/shawn-tranter/go-basics/app"
	"github.com/shawn-tranter/go-basics/api"
	"github.com/shawn-tranter/go-basics/bld"
	"log"
)

type ConfigManagementRecord struct {
	Application *app.ApplicationRecord `json:"application"`
	API       *api.APIRecord       `json:"api"`
	Build       *bld.BuildRecord       `json:"build"`
}

var cmRecord *ConfigManagementRecord

func (cm *ConfigManagementRecord) String() string {
	if api.IsAPI() {
		return fmt.Sprintf("Application -> [%s], API -> [%s], Build -> [%s]", cm.Application.String(), cm.API.String(), cm.Build.String())

	} else {
		return fmt.Sprintf("Application -> [%s], Build -> [%s]", cm.Application.String(), cm.Build.String())
	}
}

func (cm *ConfigManagementRecord) JSON() (result string) {
	if cmMap, err := json.Marshal(cm); err == nil {
		result = fmt.Sprintf(string(cmMap))
	} else {
		log.Fatal(err)
	}

	return
}

func GetApplication() *app.ApplicationRecord {
	return Get().Application
}

func GetAPI() *bld.BuildRecord {
	return Get().Build
}

func GetBuild() *bld.BuildRecord {
	return Get().Build
}

func Get() *ConfigManagementRecord {
	if cmRecord == nil {
		if api.IsAPI() {
			cmRecord = &ConfigManagementRecord{app.Get(), api.Get(), bld.Get()}

		} else {
			cmRecord = &ConfigManagementRecord{app.Get(), nil, bld.Get()}
		}
	}

	return cmRecord
}
