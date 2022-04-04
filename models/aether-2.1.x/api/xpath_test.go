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
	ynn := navigator.NewYangNodeNavigator(schema.RootSchema(), device)
	assert.NotNil(t, ynn)

	tests := []navigator.XpathSelect{
		{
			Name: "test applications",
			Path: "/app:application",
			Expected: []string{
				"Iter Value: application: value of application",
			},
		},
		{
			Name: "test application acme",
			Path: "/app:application[@app:application-id='acme-dataacquisition']/@app:application-id",
			Expected: []string{
				"Iter Value: application-id: acme-dataacquisition",
			},
		},
		{
			Name: "test application acme id",
			Path: "/app:application[@app:application-id='acme-dataacquisition']/@app:application-id",
			Expected: []string{
				"Iter Value: application-id: acme-dataacquisition",
			},
		},
		{
			Name: "test application acme address",
			Path: "/app:application[@app:application-id='acme-dataacquisition']/app:address",
			Expected: []string{
				"Iter Value: address: da.acme.com",
			},
		},
		{
			Name: "test application acme endpoints",
			Path: "/app:application[@app:application-id='acme-dataacquisition']/app:endpoint",
			Expected: []string{
				"Iter Value: endpoint: value of endpoint",
				"Iter Value: endpoint: value of endpoint",
			},
		},
		{
			Name: "test application acme endpoints ids",
			Path: "/app:application[@app:application-id='acme-dataacquisition']/app:endpoint/@app:endpoint-id",
			Expected: []string{
				"Iter Value: endpoint-id: da1",
				"Iter Value: endpoint-id: da2",
			},
		},
		{
			Name: "test application acme endpoint da1 port-start",
			Path: "/app:application[@app:application-id='acme-dataacquisition']/app:endpoint[@app:endpoint-id='da1']/app:port-start",
			Expected: []string{
				"Iter Value: port-start: 1230",
			},
		},
		{
			Name: "test application acme endpoint da2 port-end",
			Path: "/app:application[@app:application-id='acme-dataacquisition']/app:endpoint[@app:endpoint-id='da2']/app:port-end",
			Expected: []string{
				"Iter Value: port-end: 1010",
			},
		},
		{
			Name: "test application acme endpoint da2 mbr downlink",
			Path: "/app:application[@app:application-id='acme-dataacquisition']/app:endpoint[@app:endpoint-id='da2']/app:mbr/app:downlink",
			Expected: []string{
				"Iter Value: downlink: 3000000",
			},
		},
		{
			Name: "test slice acme-chicago-robots slice-id",
			Path: "/site:site[@site:site-id='acme-chicago']/site:slice[@site:slice-id='acme-chicago-robots']/@site:slice-id",
			Expected: []string{
				"Iter Value: slice-id: acme-chicago-robots",
			},
		},
		{
			Name: "test slice acme-chicago-robots dg acme-chicago-production-robots-front ref",
			Path: "/site:site[@site:site-id='acme-chicago']/site:slice[@site:slice-id='acme-chicago-robots']/site:device-group[@site:device-group='acme-chicago-production-robots']/@site:device-group",
			Expected: []string{
				"Iter Value: device-group: acme-chicago-production-robots",
			},
		},
		{
			Name: "test slice acme-chicago-robots dg acme-chicago-production-robots enable",
			Path: "/site:site[@site:site-id='acme-chicago']/site:device-group[@site:device-group-id='acme-chicago-production-robots']/@site:device-group-id",
			Expected: []string{
				"Iter Value: device-group-id: acme-chicago-production-robots",
			},
		},
		{
			Name: "test slice acme-chicago-robots dg acme-chicago-production-robots enable ip-domain",
			Path: "/site:site[@site:site-id='acme-chicago']/site:device-group[@site:device-group-id='acme-chicago-production-robots']/site:ip-domain",
			Expected: []string{
				"Iter Value: ip-domain: acme-chicago",
			},
		},
		{
			Name: "test ip-domain acme-chicago dg ip-domain by ref dg",
			Path: "/site:site[@site:site-id='acme-chicago']/site:ip-domain[@site:ip-domain-id=../site:device-group[@site:device-group-id='acme-chicago-production-robots']/site:ip-domain]/@site:ip-domain-id",
			Expected: []string{
				"Iter Value: ip-domain-id: acme-chicago",
			},
		},
		{
			Name: "set of device-groups used in slices",
			// At this point we are in the ip-domain-id
			Path: "../../site:device-group[set-contains(../site:slice/site:device-group/@site:device-group, ./@site:device-group-id)]/@site:device-group-id",
			Expected: []string{
				"Iter Value: device-group-id: acme-chicago-production-robots",
			},
		},
		{
			Name:     "repeated device-groups in slices of this site",
			Path:     "./site:slice[set-contains(following-sibling::site:slice/site:device-group/@site:device-group, site:device-group/@site:device-group)]/@site:slice-id",
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
	ynn := navigator.NewYangNodeNavigator(schema.RootSchema(), device)
	assert.NotNil(t, ynn)

	tests := []navigator.XpathSelect{
		{
			Name: "ip-domain of current device-group",
			Path: "./site:device-group[@site:device-group-id='acme-dallas-default']/site:ip-domain",
			Expected: []string{
				"Iter Value: ip-domain: acme-dallas",
			},
		},
		{
			Name: "ip domain chicago ip-domain-id",
			Path: "./site:ip-domain[@site:ip-domain-id='acme-chicago']/@site:ip-domain-id",
			Expected: []string{
				"Iter Value: ip-domain-id: acme-chicago",
			},
		},
		{
			Name: "the ip domain of chicago device-group",
			Path: "./site:device-group[@site:device-group-id='acme-chicago-default']/site:ip-domain",
			Expected: []string{
				"Iter Value: ip-domain: acme-chicago",
			},
		},
		{
			Name: "the ip domain of dallas-default device-group",
			Path: "./site:device-group[@site:device-group-id='acme-dallas-default']/site:ip-domain",
			Expected: []string{
				"Iter Value: ip-domain: acme-dallas",
			},
		},
		{
			Name: "the DNN of ip-domain related to dallas-default device-group",
			Path: "./site:ip-domain[@site:ip-domain-id=../site:device-group[@site:device-group-id='acme-dallas-default']/site:ip-domain]/site:dnn",
			Expected: []string{
				"Iter Value: dnn: dnndallas",
			},
		},
		{
			Name: "the set of all DG ip-domain references",
			Path: "./site:device-group/site:ip-domain",
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
			Path: "./site:device-group[@site:device-group-id='acme-dallas-default']/following-sibling::node()",
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
			Path: "./site:device-group[set-contains(following-sibling::site:device-group/site:ip-domain, site:ip-domain)]/site:ip-domain",
			Expected: []string{
				"Iter Value: ip-domain: acme-chicago",
				"Iter Value: ip-domain: acme-dallas",
			},
		},
		{
			Name: "non unique dnns used in ip-domains",
			Path: "./site:ip-domain[set-contains(following-sibling::site:ip-domain/site:dnn, site:dnn)]/site:dnn",
			Expected: []string{
				"Iter Value: dnn: dnndallas",
			},
		},
		{
			Name: "set of dnns of ipdomains used in device-groups",
			Path: "site:ip-domain[set-contains(../site:device-group/site:ip-domain, @site:ip-domain-id)]/site:dnn",
			Expected: []string{
				"Iter Value: dnn: dnnacme",
				"Iter Value: dnn: dnndallas",
				"Iter Value: dnn: dnndallas",
			},
		},
		{
			Name: "set of device-groups used in slices",
			Path: "site:device-group[set-contains(../site:slice/site:device-group/@site:device-group, ./@site:device-group-id)]/@site:device-group-id",
			Expected: []string{
				"Iter Value: device-group-id: acme-chicago-default",
				"Iter Value: device-group-id: acme-dallas-extra",
				"Iter Value: device-group-id: acme-dallas-more",
			},
		},
		{
			Name: "set of ip-domains used in device-groups used slices",
			Path: "site:device-group[set-contains(../site:slice/site:device-group/@site:device-group, ./@site:device-group-id)]/site:ip-domain",
			Expected: []string{
				"Iter Value: ip-domain: acme-chicago",
				"Iter Value: ip-domain: acme-chicago",
				"Iter Value: ip-domain: acme-dallas",
			},
		},
		{
			Name: "set of dnns of ip-domains used in device-groups used slices",
			Path: "site:ip-domain[set-contains(../site:device-group[set-contains(../site:slice/site:device-group/@site:device-group, ./@site:device-group-id)]/site:ip-domain, @site:ip-domain-id)]/site:dnn",
			Expected: []string{
				"Iter Value: dnn: dnnacme",
				"Iter Value: dnn: dnndallas",
			},
		},
		{
			Name:     "set of dnns of ip-domains used in device-groups used slices 2",
			Path:     "site:ip-domain[set-contains(../site:device-group[set-contains(../site:slice[set-contains(following-sibling::site:slice/@site:slice-id, @site:slice-id)]/site:device-group/@site:device-group, ./@site:device-group-id)]/site:ip-domain, @site:ip-domain-id)]/site:dnn",
			Expected: []string{
				//"Iter Value: dnn: dnnacme",
				//"Iter Value: dnn: dnndallas",
			},
		},
		{
			Name: "set of dnns of ip-domains used in device-groups used in 1 named slice - acme-chicago-robots",
			Path: "site:ip-domain[set-contains(../site:device-group[set-contains(../site:slice[@site:slice-id='acme-chicago-robots']/site:device-group/@site:device-group, ./@site:device-group-id)]/site:ip-domain, @site:ip-domain-id)]/site:dnn",
			Expected: []string{
				"Iter Value: dnn: dnnacme",
			},
		},
		{
			Name: "set of device-groups used in 1 named slice - acme-chicago-robots",
			Path: "site:slice[@site:slice-id='acme-chicago-robots']/site:device-group/@site:device-group",
			Expected: []string{
				"Iter Value: device-group: acme-chicago-default",
				"Iter Value: device-group: acme-dallas-extra",
			},
		},
		{
			Name: "set of device-groups used in 1 named slice - acme-dallas-robots",
			Path: "site:slice[@site:slice-id='acme-dallas-robots']/site:device-group/@site:device-group",
			Expected: []string{
				"Iter Value: device-group: acme-dallas-extra",
				"Iter Value: device-group: acme-dallas-more",
			},
		},
		{
			Name: "repeated device-groups in slices of this site",
			Path: "./site:slice[set-contains(following-sibling::site:slice/site:device-group/@site:device-group, site:device-group/@site:device-group)]/@site:slice-id",
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
		assert.True(t, ynn.MoveToChild()) // first site

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
	ynn := navigator.NewYangNodeNavigator(schema.RootSchema(), device)
	assert.NotNil(t, ynn)

	tests := []navigator.XpathEvaluate{
		{
			Name:     "test get ip-domain",
			Path:     "string(./site:ip-domain)",
			Expected: "acme-dallas2",
		},
		{
			Name:     "test get again device-group",
			Path:     "string(./@site:device-group-id)",
			Expected: "acme-dallas-bis",
		},
		{
			Name:     "check more than 1 device-groups use 'acme-dallas' ip-domain",
			Path:     "count(../site:device-group[site:ip-domain='acme-dallas']) <= 1",
			Expected: false,
		},
		/*
			{
				Name:     "check that no more than 1 device-groups use 'acme-chicago' ip-domain",
				Path:     "count(../site:device-group[site:ip-domain='acme-chicago']) <= 1",
				Expected: true,
			},
		*/
		{
			Name:     "check that no more than 1 device-groups use 'acme-houston' ip-domain",
			Path:     "count(../site:device-group[site:ip-domain='acme-houston']) <= 1",
			Expected: true,
		},
		{
			Name:     "see that string of set of DG gives only first one",
			Path:     "string(//site:device-group/site:ip-domain)",
			Expected: "acme-chicago",
		},
		{
			Name:     "see that ipdomain is not reused by device-group",
			Path:     "boolean(../site:device-group[set-contains(following-sibling::site:device-group/site:ip-domain, site:ip-domain)])",
			Expected: true, // acme-dallas is used twice
		},
		{
			Name:     "see that ipdomain is not reused by device-group",
			Path:     "boolean(../site:device-group[set-contains(following-sibling::site:device-group/site:ip-domain, site:ip-domain)])",
			Expected: true, // acme-dallas is used twice
		},
		{
			Name:     "count of times a device-group is repeated in slices of this site",
			Path:     "count(../site:slice[set-contains(following-sibling::site:slice/site:device-group/@site:device-group, site:device-group/@site:device-group)]/@site:slice-id)",
			Expected: float64(1), // acme-dallas-extra is used twice
		},
	}

	for _, test := range tests {
		expr, testErr := xpath.Compile(test.Path)
		assert.NoError(t, testErr, test.Name)
		assert.NotNil(t, expr, test.Name)

		ynn.MoveToRoot()
		assert.True(t, ynn.MoveToChild()) // first site
		assert.True(t, ynn.MoveToChild()) // first device-group
		assert.True(t, ynn.MoveToNext())  // The device group 'acme-dallas-default'

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
			Path:     "count(site:ip-domain)",
			Expected: float64(4),
		},
		{
			Name:     "test count device-group",
			Path:     "count(site:device-group)",
			Expected: float64(5),
		},
		{
			Name:     "test count device-group whose following node has same ip-domain",
			Path:     "string(site:device-group[site:ip-domain=following::site:device-group/site:ip-domain])",
			Expected: "",
		},
		{
			Name:     "count of DNNs used in all the slices in this site",
			Path:     "count(site:ip-domain[set-contains(../site:device-group[set-contains(../site:slice/site:device-group/@site:device-group, ./@site:device-group-id)]/site:ip-domain, @site:ip-domain-id)]/site:dnn)",
			Expected: float64(2),
		},
		{
			Name:     "count of device-groups used in the slices in this site",
			Path:     "count(site:device-group[set-contains(../site:slice/site:device-group/@site:device-group, ./@site:device-group-id)]/@site:device-group-id)",
			Expected: float64(3),
		},
		{
			Name: "compare count of unique dnns to the count of device groups used in the slices in this site",
			Path: `count(site:device-group[set-contains(../site:slice/site:device-group/@site:device-group, ./@site:device-group-id)]/@site:device-group-id) =
count(../site:ip-domain[set-contains(../site:device-group[set-contains(../site:slice/site:device-group/@site:device-group, ./@site:device-group-id)]/site:ip-domain, @site:ip-domain-id)]/site:dnn)`,
			Expected: false,
		},
	}

	for _, test := range tests {
		expr, testErr := xpath.Compile(test.Path)
		assert.NoError(t, testErr, test.Name)
		assert.NotNil(t, expr, test.Name)

		ynn.MoveToRoot()
		assert.True(t, ynn.MoveToChild()) // First site

		result := expr.Evaluate(ynn)
		assert.Equal(t, test.Expected, result, test.Name)
	}
}
