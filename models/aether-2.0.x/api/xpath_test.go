// SPDX-FileCopyrightText: 2019-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package api

import (
	"fmt"
	"github.com/SeanCondon/xpath"
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
	ynn := navigator.NewYangNodeNavigator(schema.RootSchema(), device, false)
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
				"Iter Value: ip-domain: starbucks-newyork-1",
			},
		},
		{
			Name: "test ip-domain starbucks-newyork dg ip-domain by ref dg",
			Path: "/ent:enterprises/ent:enterprise/ent:site[@ent:site-id='starbucks-newyork']/ent:ip-domain[@ent:ip-domain-id=../ent:device-group[@ent:device-group-id='starbucks-newyork-pos']/ent:ip-domain]/@ent:ip-domain-id",
			Expected: []string{
				"Iter Value: ip-domain-id: starbucks-newyork-1",
			},
		},
		{
			Name: "set of device-groups used in slices",
			Path: "ent:device-group[set-contains(../ent:slice/ent:device-group/@ent:device-group, ./@ent:device-group-id)]/@ent:device-group-id",
			Expected: []string{
				"Iter Value: device-group-id: starbucks-seattle-cameras-cter",
				"Iter Value: device-group-id: starbucks-seattle-cameras-store",
				"Iter Value: device-group-id: starbucks-seattle-pos-tills",
			},
		},
		{
			Name:     "repeated device-groups in slices of this site",
			Path:     "./ent:slice[set-contains(following-sibling::ent:slice/ent:device-group/@ent:device-group, ent:device-group/@ent:device-group)]/@ent:slice-id",
			Expected: []string{},
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

func Test_XPathSelectSite(t *testing.T) {
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
	ynn := navigator.NewYangNodeNavigator(schema.RootSchema(), device, false)
	assert.NotNil(t, ynn)

	tests := []navigator.XpathSelect{
		{
			Name: "ip-domain of current device-group",
			Path: "./ent:device-group[@ent:device-group-id='acme-dallas-default']/ent:ip-domain",
			Expected: []string{
				"Iter Value: ip-domain: acme-dallas",
			},
		},
		{
			Name: "ip domain chicago ip-domain-id",
			Path: "./ent:ip-domain[@ent:ip-domain-id='acme-chicago']/@ent:ip-domain-id",
			Expected: []string{
				"Iter Value: ip-domain-id: acme-chicago",
			},
		},
		{
			Name: "the ip domain of chicago device-group",
			Path: "./ent:device-group[@ent:device-group-id='acme-chicago-default']/ent:ip-domain",
			Expected: []string{
				"Iter Value: ip-domain: acme-chicago",
			},
		},
		{
			Name: "the ip domain of dallas-default device-group",
			Path: "./ent:device-group[@ent:device-group-id='acme-dallas-default']/ent:ip-domain",
			Expected: []string{
				"Iter Value: ip-domain: acme-dallas",
			},
		},
		{
			Name: "the DNN of ip-domain related to dallas-default device-group",
			Path: "./ent:ip-domain[@ent:ip-domain-id=../ent:device-group[@ent:device-group-id='acme-dallas-default']/ent:ip-domain]/ent:dnn",
			Expected: []string{
				"Iter Value: dnn: dnndallas",
			},
		},
		{
			Name: "the set of all DG ip-domain references",
			Path: "./ent:device-group/ent:ip-domain",
			Expected: []string{
				"Iter Value: ip-domain: acme-chicago",
				"Iter Value: ip-domain: acme-dallas2",
				"Iter Value: ip-domain: acme-dallas",
				"Iter Value: ip-domain: acme-chicago",
				"Iter Value: ip-domain: acme-dallas",
			},
		},
		{
			Name: "the set of all following-sibling nodes of first DG",
			Path: "./ent:device-group[@ent:device-group-id='acme-dallas-default']/following-sibling::node()",
			Expected: []string{
				"Iter Value: device-group: value of device-group",
				"Iter Value: device-group: value of device-group",
				"Iter Value: ip-domain: value of ip-domain",
				"Iter Value: ip-domain: value of ip-domain",
				"Iter Value: ip-domain: value of ip-domain",
				"Iter Value: ip-domain: value of ip-domain",
				"Iter Value: site-id: acme-chicago",
				"Iter Value: slice: value of slice",
				"Iter Value: slice: value of slice",
			},
		},
		{
			Name: "non unique ip-domains used in device-groups",
			Path: "./ent:device-group[set-contains(following-sibling::ent:device-group/ent:ip-domain, ent:ip-domain)]/ent:ip-domain",
			Expected: []string{
				"Iter Value: ip-domain: acme-chicago",
				"Iter Value: ip-domain: acme-dallas",
			},
		},
		{
			Name: "non unique dnns used in ip-domains",
			Path: "./ent:ip-domain[set-contains(following-sibling::ent:ip-domain/ent:dnn, ent:dnn)]/ent:dnn",
			Expected: []string{
				"Iter Value: dnn: dnndallas",
			},
		},
		{
			Name: "set of dnns of ipdomains used in device-groups",
			Path: "ent:ip-domain[set-contains(../ent:device-group/ent:ip-domain, @ent:ip-domain-id)]/ent:dnn",
			Expected: []string{
				"Iter Value: dnn: dnnacme",
				"Iter Value: dnn: dnndallas",
				"Iter Value: dnn: dnndallas",
			},
		},
		{
			Name: "set of device-groups used in slices",
			Path: "ent:device-group[set-contains(../ent:slice/ent:device-group/@ent:device-group, ./@ent:device-group-id)]/@ent:device-group-id",
			Expected: []string{
				"Iter Value: device-group-id: acme-chicago-default",
				"Iter Value: device-group-id: acme-dallas-extra",
				"Iter Value: device-group-id: acme-dallas-more",
			},
		},
		{
			Name: "set of ip-domains used in device-groups used slices",
			Path: "ent:device-group[set-contains(../ent:slice/ent:device-group/@ent:device-group, ./@ent:device-group-id)]/ent:ip-domain",
			Expected: []string{
				"Iter Value: ip-domain: acme-chicago",
				"Iter Value: ip-domain: acme-chicago",
				"Iter Value: ip-domain: acme-dallas",
			},
		},
		{
			Name: "set of dnns of ip-domains used in device-groups used slices",
			Path: "ent:ip-domain[set-contains(../ent:device-group[set-contains(../ent:slice/ent:device-group/@ent:device-group, ./@ent:device-group-id)]/ent:ip-domain, @ent:ip-domain-id)]/ent:dnn",
			Expected: []string{
				"Iter Value: dnn: dnnacme",
				"Iter Value: dnn: dnndallas",
			},
		},
		{
			Name:     "set of dnns of ip-domains used in device-groups used slices 2",
			Path:     "ent:ip-domain[set-contains(../ent:device-group[set-contains(../ent:slice[set-contains(following-sibling::ent:slice/@ent:slice-id, @ent:slice-id)]/ent:device-group/@ent:device-group, ./@ent:device-group-id)]/ent:ip-domain, @ent:ip-domain-id)]/ent:dnn",
			Expected: []string{
				//"Iter Value: dnn: dnnacme",
				//"Iter Value: dnn: dnndallas",
			},
		},
		{
			Name: "set of dnns of ip-domains used in device-groups used in 1 named slice - acme-chicago-robots",
			Path: "ent:ip-domain[set-contains(../ent:device-group[set-contains(../ent:slice[@ent:slice-id='acme-chicago-robots']/ent:device-group/@ent:device-group, ./@ent:device-group-id)]/ent:ip-domain, @ent:ip-domain-id)]/ent:dnn",
			Expected: []string{
				"Iter Value: dnn: dnnacme",
			},
		},
		{
			Name: "set of device-groups used in 1 named slice - acme-chicago-robots",
			Path: "ent:slice[@ent:slice-id='acme-chicago-robots']/ent:device-group/@ent:device-group",
			Expected: []string{
				"Iter Value: device-group: acme-chicago-default",
				"Iter Value: device-group: acme-dallas-extra",
			},
		},
		{
			Name: "set of device-groups used in 1 named slice - acme-dallas-robots",
			Path: "ent:slice[@ent:slice-id='acme-dallas-robots']/ent:device-group/@ent:device-group",
			Expected: []string{
				"Iter Value: device-group: acme-dallas-extra",
				"Iter Value: device-group: acme-dallas-more",
			},
		},
		{
			Name: "repeated device-groups in slices of this site",
			Path: "./ent:slice[set-contains(following-sibling::ent:slice/ent:device-group/@ent:device-group, ent:device-group/@ent:device-group)]/@ent:slice-id",
			Expected: []string{
				"Iter Value: slice-id: acme-chicago-robots",
			},
		},
	}

	for _, test := range tests {
		expr, err := xpath.Compile(test.Path)
		assert.NoError(t, err, test.Name)
		assert.NotNil(t, expr, test.Name)

		ynn.MoveToRoot()
		assert.True(t, ynn.MoveToChild()) // enterprises
		assert.True(t, ynn.MoveToChild()) // enterprise acme
		assert.True(t, ynn.MoveToChild()) // enterprise-id acme
		assert.True(t, ynn.MoveToNext())  // site acme-chicago

		assert.Equal(t, "site", ynn.LocalName())
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
	ynn := navigator.NewYangNodeNavigator(schema.RootSchema(), device, false)
	assert.NotNil(t, ynn)

	tests := []navigator.XpathEvaluate{
		{
			Name:     "test get ip-domain",
			Path:     "string(./ent:ip-domain)",
			Expected: "acme-dallas2",
		},
		{
			Name:     "test get again device-group",
			Path:     "string(./@ent:device-group-id)",
			Expected: "acme-dallas-bis",
		},
		{
			Name:     "check more than 1 device-groups use 'acme-dallas' ip-domain",
			Path:     "count(../ent:device-group[ent:ip-domain='acme-dallas']) <= 1",
			Expected: false,
		},
		/*
			{
				Name:     "check that no more than 1 device-groups use 'acme-chicago' ip-domain",
				XPath:     "count(../ent:device-group[ent:ip-domain='acme-chicago']) <= 1",
				Expected: true,
			},
		*/
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
		{
			Name:     "see that ipdomain is not reused by device-group",
			Path:     "boolean(../ent:device-group[set-contains(following-sibling::ent:device-group/ent:ip-domain, ent:ip-domain)])",
			Expected: true, // acme-dallas is used twice
		},
		{
			Name:     "see that ipdomain is not reused by device-group",
			Path:     "boolean(../ent:device-group[set-contains(following-sibling::ent:device-group/ent:ip-domain, ent:ip-domain)])",
			Expected: true, // acme-dallas is used twice
		},
		{
			Name:     "count of times a device-group is repeated in slices of this site",
			Path:     "count(../ent:slice[set-contains(following-sibling::ent:slice/ent:device-group/@ent:device-group, ent:device-group/@ent:device-group)]/@ent:slice-id)",
			Expected: float64(1), // acme-dallas-extra is used twice
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
	ynn := navigator.NewYangNodeNavigator(schema.RootSchema(), device, false)
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
			Expected: float64(5),
		},
		{
			Name:     "test count device-group whose following node has same ip-domain",
			Path:     "string(ent:device-group[ent:ip-domain=following::ent:device-group/ent:ip-domain])",
			Expected: "",
		},
		{
			Name:     "count of DNNs used in all the slices in this site",
			Path:     "count(ent:ip-domain[set-contains(../ent:device-group[set-contains(../ent:slice/ent:device-group/@ent:device-group, ./@ent:device-group-id)]/ent:ip-domain, @ent:ip-domain-id)]/ent:dnn)",
			Expected: float64(2),
		},
		{
			Name:     "count of device-groups used in the slices in this site",
			Path:     "count(ent:device-group[set-contains(../ent:slice/ent:device-group/@ent:device-group, ./@ent:device-group-id)]/@ent:device-group-id)",
			Expected: float64(3),
		},
		{
			Name: "compare count of unique dnns to the count of device groups used in the slices in this site",
			Path: `count(ent:device-group[set-contains(../ent:slice/ent:device-group/@ent:device-group, ./@ent:device-group-id)]/@ent:device-group-id) =
count(../ent:ip-domain[set-contains(../ent:device-group[set-contains(../ent:slice/ent:device-group/@ent:device-group, ./@ent:device-group-id)]/ent:ip-domain, @ent:ip-domain-id)]/ent:dnn)`,
			Expected: false,
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
