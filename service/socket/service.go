// Copyright 2014 Penlook Development Team. All rights reserved.
// Use of this source code is governed by
// license that can be found in the LICENSE file.
// Author : Loi Nguyen <loint@penlook.com>

package main

import (
	// "github.com/penlook/socket"
	"fmt"
	"github.com/penlook/service/component/daemon"
)

func main() {

	service := daemon.Service{
		Process: Socket,
	}
	service.GetInfo("socket")
	service.Initialize()
}

func Socket() {
}