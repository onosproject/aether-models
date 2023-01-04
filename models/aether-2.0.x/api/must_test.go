// SPDX-FileCopyrightText: 2019-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package api

import (
	"github.com/onosproject/config-models/pkg/xpath/navigator"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"strings"
	"testing"
)

func Test_WalkAndValidatePortsMust(t *testing.T) {
	sampleConfig, err := ioutil.ReadFile("../testdata/sample-aether2-config.json")
	if err != nil {
		assert.NoError(t, err)
	}
	device := new(Device)

	schema, err := Schema()
	if err := schema.Unmarshal(sampleConfig, device); err != nil {
		assert.NoError(t, err)
	}
	schema.Root = device
	assert.NotNil(t, device)
	nn := navigator.NewYangNodeNavigator(schema.RootSchema(), device, false)
	assert.NotNil(t, nn)

	ynn, ynnOk := nn.(*navigator.YangNodeNavigator)
	assert.True(t, ynnOk)
	validateErr := ynn.WalkAndValidateMust()
	assert.NoError(t, validateErr)
}

func Test_WalkAndValidateMustPortsWrong(t *testing.T) {
	sampleConfig, err := ioutil.ReadFile("../testdata/sample-must-port-start-end-wrong.json")
	if err != nil {
		assert.NoError(t, err)
	}
	device := new(Device)

	schema, err := Schema()
	if err := schema.Unmarshal(sampleConfig, device); err != nil {
		assert.NoError(t, err)
	}
	schema.Root = device
	assert.NotNil(t, device)
	nn := navigator.NewYangNodeNavigator(schema.RootSchema(), device, false)
	assert.NotNil(t, nn)

	ynn, ynnOk := nn.(*navigator.YangNodeNavigator)
	assert.True(t, ynnOk)
	validateErr := ynn.WalkAndValidateMust()
	assert.True(t, strings.HasPrefix(validateErr.Error(), "port-start must be less than or equal to port-end. Must statement 'number(./ent:port-start) <= number(./ent:port-end)' to true. Container(s):"))
	assert.True(t, strings.HasSuffix(validateErr.Error(), "endpoint-id=da2]"))
}

func Test_WalkAndValidateMustIpDomainSlice(t *testing.T) {
	sampleConfig, err := ioutil.ReadFile("../testdata/sample-slice-dnn-notreuse.json")
	if err != nil {
		assert.NoError(t, err)
	}
	device := new(Device)

	schema, err := Schema()
	if err := schema.Unmarshal(sampleConfig, device); err != nil {
		assert.NoError(t, err)
	}
	schema.Root = device
	assert.NotNil(t, device)
	nn := navigator.NewYangNodeNavigator(schema.RootSchema(), device, false)
	assert.NotNil(t, nn)

	ynn, ynnOk := nn.(*navigator.YangNodeNavigator)
	assert.True(t, ynnOk)
	validateErr := ynn.WalkAndValidateMust()
	assert.True(t, strings.HasPrefix(validateErr.Error(), `a device group cannot be used in more than one slice in a site. Must statement 'count(ent:slice[set-contains(following-sibling::ent:slice/ent:device-group/@ent:device-group, ent:device-group/@ent:device-group)]/@ent:slice-id) = 0' to true. Container(s):`))
	assert.True(t, strings.HasSuffix(validateErr.Error(), `slice-id=acme-dallas-robots]`))
}
