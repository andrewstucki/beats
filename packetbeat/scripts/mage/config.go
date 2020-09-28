// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package mage

import (
	devtools "github.com/elastic/beats/v7/dev-tools/mage"
	"github.com/magefile/mage/mg"
)

const (
	// configTemplateGlob matches Packetbeat protocol config file templates.
	configTemplateGlob = "protos/*/_meta/config*.yml.tmpl"
)

var defaultDevice = map[string]string{
	"darwin":  "en0",
	"linux":   "any",
	"windows": "0",
}

func device(goos string) string {
	dev, found := defaultDevice[goos]
	if found {
		return dev
	}
	return "any"
}

// ConfigFileParams returns the default ConfigFileParams for generating
// packetbeat*.yml files.
func ConfigFileParams() devtools.ConfigFileParams {
	p := devtools.DefaultConfigFileParams()
	p.Templates = append(p.Templates, devtools.OSSBeatDir("_meta/config/*.tmpl"))
	p.ExtraVars = map[string]interface{}{
		"device": device,
	}
	return p
}

// Fields generates fields.yml and fields.go files for the Beat.
func Fields() {
	mg.Deps(libbeatAndPacketbeatCommonFieldsGo, protosFieldsGo, fieldsYML)
}

// libbeatAndPacketbeatCommonFieldsGo generates a fields.go containing both
// libbeat and packetbeat's common fields.
func libbeatAndPacketbeatCommonFieldsGo() error {
	if err := devtools.GenerateFieldsYAML(); err != nil {
		return err
	}
	return devtools.GenerateAllInOneFieldsGo()
}

// protosFieldsGo generates a fields.go for each protocol.
func protosFieldsGo() error {
	return devtools.GenerateModuleFieldsGo(devtools.OSSBeatDir("protos"))
}

// fieldsYML generates the fields.yml file containing all fields.
func fieldsYML() error {
	return devtools.GenerateFieldsYAML(devtools.OSSBeatDir("protos"))
}
