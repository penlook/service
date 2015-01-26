// Copyright 2014 Penlook Development Team. All rights reserved.
// Use of this source code is governed by
// license that can be found in the LICENSE file.
// Author : Tin Nguyen <tinntt@penlook.com>

package main

import (
	"github.com/penlook/daemon"
	// "github.com/penlook/service/modules/s3"
)

func main() {

	service := daemon.Service{
		Name:        "s3",
		Description: "Penlook API Service",
		Process:     S3,
	}

	service.Initialize()
}

func S3() {
	// S3 service in here
}