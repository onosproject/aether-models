// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package api

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// aStr facilitates easy declaring of pointers to strings
func aStr(s string) *string {
     return &s
}

func TestApplicationAddress(t *testing.T) {
	app := OnfEnterprise_Enterprises_Enterprise_Application {}

	app.Address = aStr("my.host.name")
	err := app.Validate()
	assert.Nil(t, err)

	app.Address = aStr("1.2.3.4")
	err = app.Validate()
	assert.Nil(t, err)

	app.Address = aStr("1.2.3.4/1")
	err = app.Validate()
	assert.Nil(t, err)

	app.Address = aStr("1.2.3.4/32")
	err = app.Validate()
	assert.Nil(t, err)

	// Error - malformed ip address
	app.Address = aStr("1..2.3.4")
	err = app.Validate()
	assert.Error(t, err)

	// Error - missing subnet bits
	app.Address = aStr("1.2.3.4/")
	err = app.Validate()
	assert.Error(t, err)

	// Error - subnet bits are too many
	app.Address = aStr("1.2.3.4/33")
	err = app.Validate()
	assert.Error(t, err)

	// Error - garbage in subnet section
	app.Address = aStr("1.2.3.4/x33")
	err = app.Validate()
	assert.Error(t, err)

	// Error - subnet not a number
	app.Address = aStr("1.2.3.4/x")
	err = app.Validate()
	assert.Error(t, err)
}
