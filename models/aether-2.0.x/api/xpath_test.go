// SPDX-FileCopyrightText: 2019-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package api

import (
	"fmt"
	"github.com/antchfx/xpath"
	"github.com/onosproject/config-models/pkg/xpath/navigator"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

func Test_XPathSelect(t *testing.T) {
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
	ynn := navigator.NewYangNodeNavigator(schema.RootSchema(), device)
	assert.NotNil(t, ynn)

	tests := []navigator.XpathSelect{
		{
			Name: "test applications",
			Path: "/ent:enterprises/ent:enterprise/ent:application",
			Expected: []string{
				"Iter Value: application: value of application",
				"Iter Value: application: value of application",
				"Iter Value: application: value of application",
			},
		},
		{
			Name: "test application acme",
			Path: "/ent:enterprises/ent:enterprise/ent:application[@ent:application-id='acme-dataacquisition']/@ent:application-id",
			Expected: []string{
				"Iter Value: application-id: acme-dataacquisition",
			},
		},
		{
			Name: "test application acme id",
			Path: "/ent:enterprises/ent:enterprise/ent:application[@ent:application-id='acme-dataacquisition']/@ent:application-id",
			Expected: []string{
				"Iter Value: application-id: acme-dataacquisition",
			},
		},
		{
			Name: "test application acme address",
			Path: "/ent:enterprises/ent:enterprise/ent:application[@ent:application-id='acme-dataacquisition']/ent:address",
			Expected: []string{
				"Iter Value: address: da.acme.com",
			},
		},
		{
			Name: "test application acme endpoints",
			Path: "/ent:enterprises/ent:enterprise/ent:application[@ent:application-id='acme-dataacquisition']/ent:endpoint",
			Expected: []string{
				"Iter Value: endpoint: value of endpoint",
				"Iter Value: endpoint: value of endpoint",
			},
		},
		{
			Name: "test application acme endpoints ids",
			Path: "/ent:enterprises/ent:enterprise/ent:application[@ent:application-id='acme-dataacquisition']/ent:endpoint/@ent:endpoint-id",
			Expected: []string{
				"Iter Value: endpoint-id: da1",
				"Iter Value: endpoint-id: da2",
			},
		},
		{
			Name: "test application acme endpoint da1 port-start",
			Path: "/ent:enterprises/ent:enterprise/ent:application[@ent:application-id='acme-dataacquisition']/ent:endpoint[@ent:endpoint-id='da1']/ent:port-start",
			Expected: []string{
				"Iter Value: port-start: 1230",
			},
		},
		{
			Name: "test application acme endpoint da2 port-end",
			Path: "/ent:enterprises/ent:enterprise/ent:application[@ent:application-id='acme-dataacquisition']/ent:endpoint[@ent:endpoint-id='da2']/ent:port-end",
			Expected: []string{
				"Iter Value: port-end: 1010",
			},
		},
		{
			Name: "test application acme endpoint da2 mbr downlink",
			Path: "/ent:enterprises/ent:enterprise/ent:application[@ent:application-id='acme-dataacquisition']/ent:endpoint[@ent:endpoint-id='da2']/ent:mbr/ent:downlink",
			Expected: []string{
				"Iter Value: downlink: 3000000",
			},
		},
		{
			Name: "test slice starbucks-newyork-cameras slice-id",
			Path: "/ent:enterprises/ent:enterprise/ent:site[@ent:site-id='starbucks-newyork']/ent:slice[@ent:slice-id='starbucks-newyork-cameras']/@ent:slice-id",
			Expected: []string{
				"Iter Value: slice-id: starbucks-newyork-cameras",
			},
		},
		{
			Name: "test slice starbucks-newyork-cameras dg starbucks-newyork-cameras-front ref",
			Path: "/ent:enterprises/ent:enterprise/ent:site[@ent:site-id='starbucks-newyork']/ent:slice[@ent:slice-id='starbucks-newyork-cameras']/ent:device-group[@ent:device-group='starbucks-newyork-cameras-front']/@ent:device-group",
			Expected: []string{
				"Iter Value: device-group: starbucks-newyork-cameras-front",
			},
		},
		{
			Name: "test slice starbucks-newyork-cameras dg starbucks-newyork-cameras-front enable",
			Path: "/ent:enterprises/ent:enterprise/ent:site[@ent:site-id='starbucks-newyork']/ent:device-group[@ent:device-group-id='starbucks-newyork-cameras-front']/@ent:device-group-id",
			Expected: []string{
				"Iter Value: device-group-id: starbucks-newyork-cameras-front",
			},
		},
		{
			Name: "test slice starbucks-newyork-cameras dg starbucks-newyork-cameras-front enable",
			Path: "/ent:enterprises/ent:enterprise/ent:site[@ent:site-id='starbucks-newyork']/ent:device-group[@ent:device-group-id='starbucks-newyork-cameras-front']/ent:ip-domain",
			Expected: []string{
				"Iter Value: ip-domain: starbucks-newyork",
			},
		},
		{
			Name: "test ip-domain starbucks-newyork dg ip-domain by ref dg",
			Path: "/ent:enterprises/ent:enterprise/ent:site[@ent:site-id='starbucks-newyork']/ent:ip-domain[@ent:ip-domain-id=../ent:device-group[@ent:device-group-id='starbucks-newyork-pos']/ent:ip-domain]/@ent:ip-domain-id",
			Expected: []string{
				"Iter Value: ip-domain-id: starbucks-newyork",
			},
		},
	}

	for _, test := range tests {
		expr, err := xpath.Compile(test.Path)
		assert.NoError(t, err, test.Name)
		assert.NotNil(t, expr, test.Name)

		iter := expr.Select(ynn)
		resultCount := 0
		for iter.MoveNext() {
			assert.LessOrEqual(t, resultCount, len(test.Expected)-1, test.Name, ". More results than expected")
			assert.Equal(t, test.Expected[resultCount], fmt.Sprintf("Iter Value: %s: %s",
				iter.Current().LocalName(), iter.Current().Value()), test.Name)
			resultCount++
		}
		assert.Equal(t, len(test.Expected), resultCount, "%s. Did not receive all the expected results", test.Name)
		if resultCount < len(test.Expected) {
			t.FailNow()
		}
	}
}

func Test_XPathSelectIPDomain(t *testing.T) {
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
	ynn := navigator.NewYangNodeNavigator(schema.RootSchema(), device)
	assert.NotNil(t, ynn)

	tests := []navigator.XpathSelect{
		{
			Name: "ip-domain of current device-group",
			Path: "./ent:ip-domain",
			Expected: []string{
				"Iter Value: ip-domain: acme-dallas",
			},
		},
		{
			Name: "verify again that we can get ip-domain of current device-group",
			Path: "./ent:ip-domain",
			Expected: []string{
				"Iter Value: ip-domain: acme-dallas",
			},
		},
		{
			Name: "ip domain chicago ip-domain-id",
			Path: "../ent:ip-domain[@ent:ip-domain-id='acme-chicago']/@ent:ip-domain-id",
			Expected: []string{
				"Iter Value: ip-domain-id: acme-chicago",
			},
		},
		{
			Name: "the ip domain of chicago device-group",
			Path: "../ent:device-group[@ent:device-group-id='acme-chicago-default']/ent:ip-domain",
			Expected: []string{
				"Iter Value: ip-domain: acme-chicago",
			},
		},
		{
			Name: "the ip domain of dallas device-group",
			Path: "../ent:device-group[@ent:device-group-id='acme-dallas-default']/ent:ip-domain",
			Expected: []string{
				"Iter Value: ip-domain: acme-dallas",
			},
		},
		{
			Name: "the DNN of ip-domain related to dallas device-group",
			Path: "../ent:ip-domain[@ent:ip-domain-id=../ent:device-group[@ent:device-group-id='acme-dallas-default']/ent:ip-domain]/ent:dnn",
			Expected: []string{
				"Iter Value: dnn: dnndallas",
			},
		},
		{
			Name: "the set of all DG ip-domain references",
			Path: "//ent:device-group/ent:ip-domain",
			Expected: []string{
				"Iter Value: ip-domain: acme-chicago",
				"Iter Value: ip-domain: acme-dallas",
				"Iter Value: ip-domain: acme-chicago",
				"Iter Value: ip-domain: acme-dallas",
			},
		},
		{
			Name: "the set of all DG following nodes",
			Path: "../following::node()",
			Expected: []string{
				"Iter Value: traffic-class: value of traffic-class",
				"Iter Value: traffic-class-id: class-1",
			},
		},
		{
			Name:     "the set of all DG preceding-sibling nodes",
			Path:     "./preceding-sibling::node()",
			Expected: []string{
				//"Iter Value: device-group: value of device-group",
			},
		},
		{
			Name: "the set of all DG following-sibling nodes",
			Path: "./following-sibling::node()",
			Expected: []string{
				"Iter Value: device-group: value of device-group",
				"Iter Value: device-group: value of device-group",
				"Iter Value: ip-domain: value of ip-domain",
				"Iter Value: ip-domain: value of ip-domain",
				"Iter Value: ip-domain: value of ip-domain",
				"Iter Value: ip-domain: value of ip-domain",
				"Iter Value: site-id: acme-chicago",
				"Iter Value: slice: value of slice",
			},
		},
		{
			Name: "the id of ip-domain related to all device-groups",
			Path: "../ent:ip-domain[contains(@ent:ip-domain-id,string(//ent:device-group/ent:ip-domain))]/@ent:ip-domain-id",
			Expected: []string{
				"Iter Value: ip-domain-id: acme-chicago", // Should be more than this
			},
		},
	}

	for _, test := range tests {
		expr, err := xpath.Compile(test.Path)
		assert.NoError(t, err, test.Name)
		assert.NotNil(t, expr, test.Name)

		ynn.MoveToRoot()
		assert.True(t, ynn.MoveToChild())
		assert.True(t, ynn.MoveToChild())
		assert.True(t, ynn.MoveToChild())
		assert.True(t, ynn.MoveToNext())
		assert.True(t, ynn.MoveToChild())
		assert.True(t, ynn.MoveToNext()) // Device-group list

		assert.Equal(t, "device-group", ynn.LocalName())
		assert.Equal(t, xpath.ElementNode, ynn.NodeType())

		iter := expr.Select(ynn)
		resultCount := 0
		for iter.MoveNext() {
			assert.LessOrEqual(t, resultCount, len(test.Expected)-1, test.Name, ". More results than expected")
			assert.Equal(t, test.Expected[resultCount], fmt.Sprintf("Iter Value: %s: %s",
				iter.Current().LocalName(), iter.Current().Value()), test.Name)
			resultCount++
		}
		assert.Equal(t, len(test.Expected), resultCount, "%s. Did not receive all the expected results", test.Name)
		if resultCount < len(test.Expected) {
			t.FailNow()
		}
	}
}

func Test_XPathEvaluateDeviceGroup(t *testing.T) {
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
	ynn := navigator.NewYangNodeNavigator(schema.RootSchema(), device)
	assert.NotNil(t, ynn)

	tests := []navigator.XpathEvaluate{
		{
			Name:     "test get ip-domain",
			Path:     "string(./ent:ip-domain)",
			Expected: "acme-dallas",
		},
		{
			Name:     "test get again ip-domain",
			Path:     "string(./@ent:device-group-id)",
			Expected: "acme-dallas-default", // Why is it different?
		},
		{
			Name:     "check more than 1 device-groups use 'acme-dallas' ip-domain",
			Path:     "count(../ent:device-group[ent:ip-domain='acme-dallas']) <= 1",
			Expected: false,
		},
		{
			Name:     "check that no more than 1 device-groups use 'acme-chicago' ip-domain",
			Path:     "count(../ent:device-group[ent:ip-domain='acme-chicago']) <= 1",
			Expected: true,
		},
		{
			Name:     "check that no more than 1 device-groups use 'acme-houston' ip-domain",
			Path:     "count(../ent:device-group[ent:ip-domain='acme-houston']) <= 1",
			Expected: true,
		},
		{
			Name:     "see that string of set of DG gives only first one",
			Path:     "string(//ent:device-group/ent:ip-domain)",
			Expected: "acme-chicago",
		},
	}

	for _, test := range tests {
		expr, testErr := xpath.Compile(test.Path)
		assert.NoError(t, testErr, test.Name)
		assert.NotNil(t, expr, test.Name)

		ynn.MoveToRoot()
		assert.True(t, ynn.MoveToChild())
		assert.True(t, ynn.MoveToChild())
		assert.True(t, ynn.MoveToChild())
		assert.True(t, ynn.MoveToNext())
		assert.True(t, ynn.MoveToChild())
		assert.True(t, ynn.MoveToNext()) // The device group 'acme-dallas-default'

		result := expr.Evaluate(ynn)
		assert.Equal(t, test.Expected, result, test.Name)
	}
}

func Test_XPathEvaluateSlice(t *testing.T) {
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
	ynn := navigator.NewYangNodeNavigator(schema.RootSchema(), device)
	assert.NotNil(t, ynn)

	tests := []navigator.XpathEvaluate{
		{
			Name:     "test count ip-domain",
			Path:     "count(ent:ip-domain)",
			Expected: float64(4),
		},
		{
			Name:     "test count device-group",
			Path:     "count(ent:device-group)",
			Expected: float64(4),
		},
		{
			Name:     "test count device-group whose following node has same ip-domain",
			Path:     "string(ent:device-group[ent:ip-domain=following::ent:device-group/ent:ip-domain])",
			Expected: "",
		},
	}

	for _, test := range tests {
		expr, testErr := xpath.Compile(test.Path)
		assert.NoError(t, testErr, test.Name)
		assert.NotNil(t, expr, test.Name)

		ynn.MoveToRoot()
		assert.True(t, ynn.MoveToChild())
		assert.True(t, ynn.MoveToChild())
		assert.True(t, ynn.MoveToChild())
		assert.True(t, ynn.MoveToNext())

		result := expr.Evaluate(ynn)
		assert.Equal(t, test.Expected, result, test.Name)
	}
}
