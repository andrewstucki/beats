// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package http

import (
	"github.com/elastic/beats/v7/libbeat/asset"
)

func init() {
	if err := asset.SetFields("packetbeat", "http", asset.ModuleFieldsPri, AssetHttp); err != nil {
		panic(err)
	}
}

// AssetHttp returns asset data.
// This is the base64 encoded gzipped contents of ../../packetbeat/protos/http.
func AssetHttp() string {
	return "eJzUVMuOlUAQ3fMVJ7N2+AAWJm6MszEubuJyUkAB7dCP6a72yt+b5nEdoKMxNy5kBfU451T1aR7xwlOFQcQVgCgZucLDp8vly0MBtBwar5woayqk4GNw3KhONeDvbASd4rENZYH1rSoA4BGGNN9Q0yOT4wq9t3GL7LCfTGe9pvQBqm0UyMAzIzy/Rg4CMi08B2dN4HLFeEv6lnjtucUzk2RqzhpzHLsBmVr2YZfbcGz9jRs5pJbg81LxwtPV+vZQslP6/pAEPkCTQ2ONkDLK9POiGnISPberoFUzOm/1nF9nLU9oXwfVDNsYELshQYXE0ak+eqpHLvHU3cquSoYZNpDmE+QqIS0I5BnOc0hWUWbu0RwC9fwufUy4qnFEzQjsyJNwi3o6ITZWawplkT2C1KfzJ0CjomNGq97Tslvx8ajekQwVoh/L18h+KjLGWgz4B2ediv7eWkFIYnh2g6dwlLmjvGwXZenA0nE8av5B2qW7/dkKPtpo2vw6/x9H7/8Ev567LH1CSxa/09InzHr6vaUb2x6F3Gfo9B8ubxtbnTWz5K9UznP/QMHK8zMAAP//lgPSsQ=="
}
