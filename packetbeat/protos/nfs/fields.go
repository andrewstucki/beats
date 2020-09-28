// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package nfs

import (
	"github.com/elastic/beats/v7/libbeat/asset"
)

func init() {
	if err := asset.SetFields("packetbeat", "nfs", asset.ModuleFieldsPri, AssetNfs); err != nil {
		panic(err)
	}
}

// AssetNfs returns asset data.
// This is the base64 encoded gzipped contents of ../../packetbeat/protos/nfs.
func AssetNfs() string {
	return "eJykkz1v2zAQhnf/ikM6ZKndoZk0FAgcBOhQO2ja2aCpk3QoRRK8kxP11xeUJVuGPgLXnowT73ne48cS/mCdgM14ASAkBhO42zy/3i0AUmQdyAs5m8Dm+RUOD1++AnvUlJEGPKAVyAhNyqsFtP+SBQDAEqwqsePGn9QeE8iDq3xb6a/v9xwwMDl7qne9xtm8Vxyk88GJ0850ALBVucewWgwMJVkXdjd7GszHNlGzuMMDrLc/Xra/N0/gPAYVP8WmEZTz2qU4RfvW+wAN+8yL/Z/BBSgV2UGdLGjFCC47ZblgaWUMj+RhUVLx3HRnU0Bv6rajRXWY4PXkLbkgbjdr+Pmynr6Dc/fqndKppJFZIrPKESQoy0o3mSlFK5TR6LnOD99HjozeB6lKil1m1MGFOVpcFuPo434eG0ZwOmC6qi6G/fhiR0E8ZQz3DBVjAEov7kW0LytL71PG/BZjc+L/o5zd/xOeUTubqlCfRHydiUWV/qrxHsOeJETl9yd4K0gXIAW2kaBUuiCLUKoacrTxjeCUu10bK1OyXwU2HXGWs+WeO08P/QkeDSlGHsqUMTumvziYU8WOXtUrKRJgVwWNq30t2P9YUn588QlIqCYjd+cD0QhvJAWokFclWhnZiOYFXRMuRRayTY5bErYvdyTivwAAAP//FgAaSA=="
}
