// Copyright 2014 Penlook Development Team. All rights reserved.
// Use of this source code is governed by
// license that can be found in the LICENSE file.
// Author : Loi Nguyen <loint@penlook.com>

package neo

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNeo(t *testing.T) {
	assert.New(t)

	neo := Neo{}.Connect()

	n0, _ := neo.CreateNode(.Props{"name": kirk})
}
