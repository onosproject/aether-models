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
	app := OnfApplication_Applications_Application{}

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

func TestSimCard(t *testing.T) {
	sim := OnfSite_Sites_Site_SimCard{}

	// 18 digit, starts with 0
	sim.Iccid = aStr("023456789012345678F")
	err := sim.Validate()
	assert.Nil(t, err)

	// 18 + check digit
	sim.Iccid = aStr("123456789012345678F")
	err = sim.Validate()
	assert.Nil(t, err)

	// 19 + check digit
	sim.Iccid = aStr("1234567890123456789F")
	err = sim.Validate()
	assert.Nil(t, err)

	// 20 + check digit
	sim.Iccid = aStr("12345678901234567890F")
	err = sim.Validate()
	assert.Nil(t, err)

	// 21 + check digit
	sim.Iccid = aStr("123456789012345678901F")
	err = sim.Validate()
	assert.Nil(t, err)

	// 17 + check digit -- too short
	sim.Iccid = aStr("12345678901234567F")
	err = sim.Validate()
	assert.Error(t, err)

	// 22 + check digit -- too long
	sim.Iccid = aStr("12345678901234567F")
	err = sim.Validate()
	assert.Error(t, err)

	// invalid check digit
	sim.Iccid = aStr("123456789012345678G")
	err = sim.Validate()
	assert.Error(t, err)

	// invalid digit
	sim.Iccid = aStr("1234567890123456A8F")
	err = sim.Validate()
	assert.Error(t, err)
}

func TestDevice(t *testing.T) {
	dev := OnfSite_Sites_Site_Device{}

	// 14 digit, starts with 0
	dev.Imei = aStr("02345678901234")
	err := dev.Validate()
	assert.Nil(t, err)

	// 14 digit
	dev.Imei = aStr("12345678901234")
	err = dev.Validate()
	assert.Nil(t, err)

	// 15 digit
	dev.Imei = aStr("123456789012345")
	err = dev.Validate()
	assert.Nil(t, err)

	// 16 digit
	dev.Imei = aStr("1234567890123456")
	err = dev.Validate()
	assert.Nil(t, err)

	// 13 digit - too short
	dev.Imei = aStr("123456789012343")
	err = dev.Validate()
	assert.Nil(t, err)

	// 17 digit - too long
	dev.Imei = aStr("12345678901234567")
	err = dev.Validate()
	assert.Error(t, err)

	// not a valid digit
	dev.Imei = aStr("123456789012345A")
	err = dev.Validate()
	assert.Error(t, err)
}

func TestIdentifier(t *testing.T) {
	ent := OnfSite_Sites_Site{}

	ent.SiteId = aStr("test")
	err := ent.Validate()
	assert.Nil(t, err)

	// Minimum length
	ent.SiteId = aStr("a")
	err = ent.Validate()
	assert.Nil(t, err)

	// Maximum length
	ent.SiteId = aStr("a12345678901234567890123456789012345678901234567890123456789012")
	err = ent.Validate()
	assert.Nil(t, err)

	ent.SiteId = aStr("test-1")
	err = ent.Validate()
	assert.Nil(t, err)

	// Two non-consecutive dashes are fine
	ent.SiteId = aStr("test-test-1")
	err = ent.Validate()
	assert.Nil(t, err)

	// Two consecutive dashes are not allow
	ent.SiteId = aStr("test--1")
	err = ent.Validate()
	assert.Error(t, err)

	// Uppercase is not allowed
	ent.SiteId = aStr("Uppercase")
	err = ent.Validate()
	assert.Error(t, err)

	// Uppercase is not allowed, even if you hide it in the middle
	ent.SiteId = aStr("upperCase")
	err = ent.Validate()
	assert.Error(t, err)

	// Starts with a number is not allowed
	ent.SiteId = aStr("1-test")
	err = ent.Validate()
	assert.Error(t, err)

	// Empty string is not allowed
	ent.SiteId = aStr("")
	err = ent.Validate()
	assert.Error(t, err)

	// Too Long
	ent.SiteId = aStr("a123456789012345678901234567890123456789012345678901234567890123")
	err = ent.Validate()
	assert.Error(t, err)

	// No periods
	ent.SiteId = aStr("test.two")
	err = ent.Validate()
	assert.Error(t, err)

	// No spaces
	ent.SiteId = aStr("test two")
	err = ent.Validate()
	assert.Error(t, err)
}
