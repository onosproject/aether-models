/*
* SPDX-FileCopyrightText: 2022-present Intel Corporation
*
* SPDX-License-Identifier: Apache-2.0
*/

// Generated via gnmi-gen.go, do NOT edit

package api

import (
    "context"
    "github.com/onosproject/config-models/pkg/gnmi-client-gen/gnmi_utils"
    "github.com/openconfig/gnmi/proto/gnmi"
    "google.golang.org/grpc"
    "google.golang.org/grpc/codes"
    "google.golang.org/grpc/status"
    "reflect"
    "time"
)

type GnmiClient struct {
    client gnmi.GNMIClient
}

func NewAetherGnmiClient(conn *grpc.ClientConn) *GnmiClient {
    gnmi_client := gnmi.NewGNMIClient(conn)
    return &GnmiClient{client: gnmi_client}
}









func (c *GnmiClient) GetConnectivityService_ConnectivityService_Item(ctx context.Context, target string, 
key string,
) (*OnfConnectivityService_ConnectivityService_ConnectivityService, error) {
    gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
    defer cancel()


    
    
    path :=  []*gnmi.Path{
        {
            Elem: []*gnmi.PathElem{
            {
                    Name: "connectivity-service",
                },
            {
                    Name: "connectivity-service",
                    Key: map[string]string{
                        
                        "id": string(key),
                        
                        },
                },
            },
            Target: target,
        },
    }

    req := &gnmi.GetRequest{
        Encoding: gnmi.Encoding_JSON,
        Path:     path,
    }
    res, err := c.client.Get(gnmiCtx, req)

    if err != nil {
        return nil, err
    }

    val, err := gnmi_utils.GetResponseUpdate(res)

    if err != nil {
        return nil, err
    }

    json := val.GetJsonVal()
    st := Device{}
    Unmarshal(json, &st)

    if reflect.ValueOf(st.ConnectivityService).Kind() == reflect.Ptr && reflect.ValueOf(st.ConnectivityService).IsNil() {
        return nil, status.Error(codes.NotFound, "OnfConnectivityService_ConnectivityService_ConnectivityService-not-found")
    }
    if res, ok := st.ConnectivityService.ConnectivityService[key]; ok {
        return res, nil
    }

    return nil, status.Error(codes.NotFound, "OnfConnectivityService_ConnectivityService_ConnectivityService-not-found")
    }








func (c *GnmiClient) DeleteConnectivityService_ConnectivityService_Item(ctx context.Context, target string, 
key string,
) (*gnmi.SetResponse, error) {
    gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
    defer cancel()


    
    
    path :=  []*gnmi.Path{
        {
            Elem: []*gnmi.PathElem{
            {
                    Name: "connectivity-service",
                },
            {
                    Name: "connectivity-service",
                    Key: map[string]string{
                        
                        "id": string(key),
                        
                        },
                },
            },
            Target: target,
        },
    }

    req := &gnmi.SetRequest{
        Delete: []*gnmi.Path{
            {
            Elem:   path[0].Elem,
            Target: target,
            },
        },
    }
    return c.client.Set(gnmiCtx, req)
    }








func (c *GnmiClient) UpdateConnectivityService_ConnectivityService_Item(ctx context.Context, target string,  data OnfConnectivityService_ConnectivityService_ConnectivityService,
) (*gnmi.SetResponse, error) {
    gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
    defer cancel()


    
    

    
    path :=  []*gnmi.Path{
        {
            Elem: []*gnmi.PathElem{
            {
                    Name: "connectivity-service",
                },
            {
                    Name: "connectivity-service",
                    Key: map[string]string{
                        "id": string(*data.Id),
                        },
                },
            },
            Target: target,
        },
    }

    req, err := gnmi_utils.CreateGnmiSetForContainer(ctx, data, path[0], target)
    if err != nil {
        return nil, err
    }

    return c.client.Set(gnmiCtx, req)
    }








func (c *GnmiClient) GetEnterprise_Enterprise_Item(ctx context.Context, target string, 
key string,
) (*OnfEnterprise_Enterprise_Enterprise, error) {
    gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
    defer cancel()


    
    
    path :=  []*gnmi.Path{
        {
            Elem: []*gnmi.PathElem{
            {
                    Name: "enterprise",
                },
            {
                    Name: "enterprise",
                    Key: map[string]string{
                        
                        "id": string(key),
                        
                        },
                },
            },
            Target: target,
        },
    }

    req := &gnmi.GetRequest{
        Encoding: gnmi.Encoding_JSON,
        Path:     path,
    }
    res, err := c.client.Get(gnmiCtx, req)

    if err != nil {
        return nil, err
    }

    val, err := gnmi_utils.GetResponseUpdate(res)

    if err != nil {
        return nil, err
    }

    json := val.GetJsonVal()
    st := Device{}
    Unmarshal(json, &st)

    if reflect.ValueOf(st.Enterprise).Kind() == reflect.Ptr && reflect.ValueOf(st.Enterprise).IsNil() {
        return nil, status.Error(codes.NotFound, "OnfEnterprise_Enterprise_Enterprise-not-found")
    }
    if res, ok := st.Enterprise.Enterprise[key]; ok {
        return res, nil
    }

    return nil, status.Error(codes.NotFound, "OnfEnterprise_Enterprise_Enterprise-not-found")
    }








func (c *GnmiClient) DeleteEnterprise_Enterprise_Item(ctx context.Context, target string, 
key string,
) (*gnmi.SetResponse, error) {
    gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
    defer cancel()


    
    
    path :=  []*gnmi.Path{
        {
            Elem: []*gnmi.PathElem{
            {
                    Name: "enterprise",
                },
            {
                    Name: "enterprise",
                    Key: map[string]string{
                        
                        "id": string(key),
                        
                        },
                },
            },
            Target: target,
        },
    }

    req := &gnmi.SetRequest{
        Delete: []*gnmi.Path{
            {
            Elem:   path[0].Elem,
            Target: target,
            },
        },
    }
    return c.client.Set(gnmiCtx, req)
    }








func (c *GnmiClient) UpdateEnterprise_Enterprise_Item(ctx context.Context, target string,  data OnfEnterprise_Enterprise_Enterprise,
) (*gnmi.SetResponse, error) {
    gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
    defer cancel()


    
    

    
    path :=  []*gnmi.Path{
        {
            Elem: []*gnmi.PathElem{
            {
                    Name: "enterprise",
                },
            {
                    Name: "enterprise",
                    Key: map[string]string{
                        "id": string(*data.Id),
                        },
                },
            },
            Target: target,
        },
    }

    req, err := gnmi_utils.CreateGnmiSetForContainer(ctx, data, path[0], target)
    if err != nil {
        return nil, err
    }

    return c.client.Set(gnmiCtx, req)
    }








func (c *GnmiClient) GetSite_Site_Item(ctx context.Context, target string, 
key string,
) (*OnfSite_Site_Site, error) {
    gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
    defer cancel()


    
    
    path :=  []*gnmi.Path{
        {
            Elem: []*gnmi.PathElem{
            {
                    Name: "site",
                },
            {
                    Name: "site",
                    Key: map[string]string{
                        
                        "id": string(key),
                        
                        },
                },
            },
            Target: target,
        },
    }

    req := &gnmi.GetRequest{
        Encoding: gnmi.Encoding_JSON,
        Path:     path,
    }
    res, err := c.client.Get(gnmiCtx, req)

    if err != nil {
        return nil, err
    }

    val, err := gnmi_utils.GetResponseUpdate(res)

    if err != nil {
        return nil, err
    }

    json := val.GetJsonVal()
    st := Device{}
    Unmarshal(json, &st)

    if reflect.ValueOf(st.Site).Kind() == reflect.Ptr && reflect.ValueOf(st.Site).IsNil() {
        return nil, status.Error(codes.NotFound, "OnfSite_Site_Site-not-found")
    }
    if res, ok := st.Site.Site[key]; ok {
        return res, nil
    }

    return nil, status.Error(codes.NotFound, "OnfSite_Site_Site-not-found")
    }








func (c *GnmiClient) DeleteSite_Site_Item(ctx context.Context, target string, 
key string,
) (*gnmi.SetResponse, error) {
    gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
    defer cancel()


    
    
    path :=  []*gnmi.Path{
        {
            Elem: []*gnmi.PathElem{
            {
                    Name: "site",
                },
            {
                    Name: "site",
                    Key: map[string]string{
                        
                        "id": string(key),
                        
                        },
                },
            },
            Target: target,
        },
    }

    req := &gnmi.SetRequest{
        Delete: []*gnmi.Path{
            {
            Elem:   path[0].Elem,
            Target: target,
            },
        },
    }
    return c.client.Set(gnmiCtx, req)
    }








func (c *GnmiClient) UpdateSite_Site_Item(ctx context.Context, target string,  data OnfSite_Site_Site,
) (*gnmi.SetResponse, error) {
    gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
    defer cancel()


    
    

    
    path :=  []*gnmi.Path{
        {
            Elem: []*gnmi.PathElem{
            {
                    Name: "site",
                },
            {
                    Name: "site",
                    Key: map[string]string{
                        "id": string(*data.Id),
                        },
                },
            },
            Target: target,
        },
    }

    req, err := gnmi_utils.CreateGnmiSetForContainer(ctx, data, path[0], target)
    if err != nil {
        return nil, err
    }

    return c.client.Set(gnmiCtx, req)
    }








func (c *GnmiClient) GetTrafficClass_TrafficClass_Item(ctx context.Context, target string, 
key string,
) (*OnfTrafficClass_TrafficClass_TrafficClass, error) {
    gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
    defer cancel()


    
    
    path :=  []*gnmi.Path{
        {
            Elem: []*gnmi.PathElem{
            {
                    Name: "traffic-class",
                },
            {
                    Name: "traffic-class",
                    Key: map[string]string{
                        
                        "id": string(key),
                        
                        },
                },
            },
            Target: target,
        },
    }

    req := &gnmi.GetRequest{
        Encoding: gnmi.Encoding_JSON,
        Path:     path,
    }
    res, err := c.client.Get(gnmiCtx, req)

    if err != nil {
        return nil, err
    }

    val, err := gnmi_utils.GetResponseUpdate(res)

    if err != nil {
        return nil, err
    }

    json := val.GetJsonVal()
    st := Device{}
    Unmarshal(json, &st)

    if reflect.ValueOf(st.TrafficClass).Kind() == reflect.Ptr && reflect.ValueOf(st.TrafficClass).IsNil() {
        return nil, status.Error(codes.NotFound, "OnfTrafficClass_TrafficClass_TrafficClass-not-found")
    }
    if res, ok := st.TrafficClass.TrafficClass[key]; ok {
        return res, nil
    }

    return nil, status.Error(codes.NotFound, "OnfTrafficClass_TrafficClass_TrafficClass-not-found")
    }








func (c *GnmiClient) DeleteTrafficClass_TrafficClass_Item(ctx context.Context, target string, 
key string,
) (*gnmi.SetResponse, error) {
    gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
    defer cancel()


    
    
    path :=  []*gnmi.Path{
        {
            Elem: []*gnmi.PathElem{
            {
                    Name: "traffic-class",
                },
            {
                    Name: "traffic-class",
                    Key: map[string]string{
                        
                        "id": string(key),
                        
                        },
                },
            },
            Target: target,
        },
    }

    req := &gnmi.SetRequest{
        Delete: []*gnmi.Path{
            {
            Elem:   path[0].Elem,
            Target: target,
            },
        },
    }
    return c.client.Set(gnmiCtx, req)
    }








func (c *GnmiClient) UpdateTrafficClass_TrafficClass_Item(ctx context.Context, target string,  data OnfTrafficClass_TrafficClass_TrafficClass,
) (*gnmi.SetResponse, error) {
    gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
    defer cancel()


    
    

    
    path :=  []*gnmi.Path{
        {
            Elem: []*gnmi.PathElem{
            {
                    Name: "traffic-class",
                },
            {
                    Name: "traffic-class",
                    Key: map[string]string{
                        "id": string(*data.Id),
                        },
                },
            },
            Target: target,
        },
    }

    req, err := gnmi_utils.CreateGnmiSetForContainer(ctx, data, path[0], target)
    if err != nil {
        return nil, err
    }

    return c.client.Set(gnmiCtx, req)
    }








func (c *GnmiClient) GetVcs_Vcs_Item(ctx context.Context, target string, 
key string,
) (*OnfVcs_Vcs_Vcs, error) {
    gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
    defer cancel()


    
    
    path :=  []*gnmi.Path{
        {
            Elem: []*gnmi.PathElem{
            {
                    Name: "vcs",
                },
            {
                    Name: "vcs",
                    Key: map[string]string{
                        
                        "id": string(key),
                        
                        },
                },
            },
            Target: target,
        },
    }

    req := &gnmi.GetRequest{
        Encoding: gnmi.Encoding_JSON,
        Path:     path,
    }
    res, err := c.client.Get(gnmiCtx, req)

    if err != nil {
        return nil, err
    }

    val, err := gnmi_utils.GetResponseUpdate(res)

    if err != nil {
        return nil, err
    }

    json := val.GetJsonVal()
    st := Device{}
    Unmarshal(json, &st)

    if reflect.ValueOf(st.Vcs).Kind() == reflect.Ptr && reflect.ValueOf(st.Vcs).IsNil() {
        return nil, status.Error(codes.NotFound, "OnfVcs_Vcs_Vcs-not-found")
    }
    if res, ok := st.Vcs.Vcs[key]; ok {
        return res, nil
    }

    return nil, status.Error(codes.NotFound, "OnfVcs_Vcs_Vcs-not-found")
    }








func (c *GnmiClient) DeleteVcs_Vcs_Item(ctx context.Context, target string, 
key string,
) (*gnmi.SetResponse, error) {
    gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
    defer cancel()


    
    
    path :=  []*gnmi.Path{
        {
            Elem: []*gnmi.PathElem{
            {
                    Name: "vcs",
                },
            {
                    Name: "vcs",
                    Key: map[string]string{
                        
                        "id": string(key),
                        
                        },
                },
            },
            Target: target,
        },
    }

    req := &gnmi.SetRequest{
        Delete: []*gnmi.Path{
            {
            Elem:   path[0].Elem,
            Target: target,
            },
        },
    }
    return c.client.Set(gnmiCtx, req)
    }








func (c *GnmiClient) UpdateVcs_Vcs_Item(ctx context.Context, target string,  data OnfVcs_Vcs_Vcs,
) (*gnmi.SetResponse, error) {
    gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
    defer cancel()


    
    

    
    path :=  []*gnmi.Path{
        {
            Elem: []*gnmi.PathElem{
            {
                    Name: "vcs",
                },
            {
                    Name: "vcs",
                    Key: map[string]string{
                        "id": string(*data.Id),
                        },
                },
            },
            Target: target,
        },
    }

    req, err := gnmi_utils.CreateGnmiSetForContainer(ctx, data, path[0], target)
    if err != nil {
        return nil, err
    }

    return c.client.Set(gnmiCtx, req)
    }








func (c *GnmiClient) GetApplication_Application_Item(ctx context.Context, target string, 
key string,
) (*OnfApplication_Application_Application, error) {
    gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
    defer cancel()


    
    
    path :=  []*gnmi.Path{
        {
            Elem: []*gnmi.PathElem{
            {
                    Name: "application",
                },
            {
                    Name: "application",
                    Key: map[string]string{
                        
                        "id": string(key),
                        
                        },
                },
            },
            Target: target,
        },
    }

    req := &gnmi.GetRequest{
        Encoding: gnmi.Encoding_JSON,
        Path:     path,
    }
    res, err := c.client.Get(gnmiCtx, req)

    if err != nil {
        return nil, err
    }

    val, err := gnmi_utils.GetResponseUpdate(res)

    if err != nil {
        return nil, err
    }

    json := val.GetJsonVal()
    st := Device{}
    Unmarshal(json, &st)

    if reflect.ValueOf(st.Application).Kind() == reflect.Ptr && reflect.ValueOf(st.Application).IsNil() {
        return nil, status.Error(codes.NotFound, "OnfApplication_Application_Application-not-found")
    }
    if res, ok := st.Application.Application[key]; ok {
        return res, nil
    }

    return nil, status.Error(codes.NotFound, "OnfApplication_Application_Application-not-found")
    }








func (c *GnmiClient) DeleteApplication_Application_Item(ctx context.Context, target string, 
key string,
) (*gnmi.SetResponse, error) {
    gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
    defer cancel()


    
    
    path :=  []*gnmi.Path{
        {
            Elem: []*gnmi.PathElem{
            {
                    Name: "application",
                },
            {
                    Name: "application",
                    Key: map[string]string{
                        
                        "id": string(key),
                        
                        },
                },
            },
            Target: target,
        },
    }

    req := &gnmi.SetRequest{
        Delete: []*gnmi.Path{
            {
            Elem:   path[0].Elem,
            Target: target,
            },
        },
    }
    return c.client.Set(gnmiCtx, req)
    }








func (c *GnmiClient) UpdateApplication_Application_Item(ctx context.Context, target string,  data OnfApplication_Application_Application,
) (*gnmi.SetResponse, error) {
    gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
    defer cancel()


    
    

    
    path :=  []*gnmi.Path{
        {
            Elem: []*gnmi.PathElem{
            {
                    Name: "application",
                },
            {
                    Name: "application",
                    Key: map[string]string{
                        "id": string(*data.Id),
                        },
                },
            },
            Target: target,
        },
    }

    req, err := gnmi_utils.CreateGnmiSetForContainer(ctx, data, path[0], target)
    if err != nil {
        return nil, err
    }

    return c.client.Set(gnmiCtx, req)
    }








func (c *GnmiClient) GetIpDomain_IpDomain_Item(ctx context.Context, target string, 
key string,
) (*OnfIpDomain_IpDomain_IpDomain, error) {
    gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
    defer cancel()


    
    
    path :=  []*gnmi.Path{
        {
            Elem: []*gnmi.PathElem{
            {
                    Name: "ip-domain",
                },
            {
                    Name: "ip-domain",
                    Key: map[string]string{
                        
                        "id": string(key),
                        
                        },
                },
            },
            Target: target,
        },
    }

    req := &gnmi.GetRequest{
        Encoding: gnmi.Encoding_JSON,
        Path:     path,
    }
    res, err := c.client.Get(gnmiCtx, req)

    if err != nil {
        return nil, err
    }

    val, err := gnmi_utils.GetResponseUpdate(res)

    if err != nil {
        return nil, err
    }

    json := val.GetJsonVal()
    st := Device{}
    Unmarshal(json, &st)

    if reflect.ValueOf(st.IpDomain).Kind() == reflect.Ptr && reflect.ValueOf(st.IpDomain).IsNil() {
        return nil, status.Error(codes.NotFound, "OnfIpDomain_IpDomain_IpDomain-not-found")
    }
    if res, ok := st.IpDomain.IpDomain[key]; ok {
        return res, nil
    }

    return nil, status.Error(codes.NotFound, "OnfIpDomain_IpDomain_IpDomain-not-found")
    }








func (c *GnmiClient) DeleteIpDomain_IpDomain_Item(ctx context.Context, target string, 
key string,
) (*gnmi.SetResponse, error) {
    gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
    defer cancel()


    
    
    path :=  []*gnmi.Path{
        {
            Elem: []*gnmi.PathElem{
            {
                    Name: "ip-domain",
                },
            {
                    Name: "ip-domain",
                    Key: map[string]string{
                        
                        "id": string(key),
                        
                        },
                },
            },
            Target: target,
        },
    }

    req := &gnmi.SetRequest{
        Delete: []*gnmi.Path{
            {
            Elem:   path[0].Elem,
            Target: target,
            },
        },
    }
    return c.client.Set(gnmiCtx, req)
    }








func (c *GnmiClient) UpdateIpDomain_IpDomain_Item(ctx context.Context, target string,  data OnfIpDomain_IpDomain_IpDomain,
) (*gnmi.SetResponse, error) {
    gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
    defer cancel()


    
    

    
    path :=  []*gnmi.Path{
        {
            Elem: []*gnmi.PathElem{
            {
                    Name: "ip-domain",
                },
            {
                    Name: "ip-domain",
                    Key: map[string]string{
                        "id": string(*data.Id),
                        },
                },
            },
            Target: target,
        },
    }

    req, err := gnmi_utils.CreateGnmiSetForContainer(ctx, data, path[0], target)
    if err != nil {
        return nil, err
    }

    return c.client.Set(gnmiCtx, req)
    }








func (c *GnmiClient) GetTemplate_Template_Item(ctx context.Context, target string, 
key string,
) (*OnfTemplate_Template_Template, error) {
    gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
    defer cancel()


    
    
    path :=  []*gnmi.Path{
        {
            Elem: []*gnmi.PathElem{
            {
                    Name: "template",
                },
            {
                    Name: "template",
                    Key: map[string]string{
                        
                        "id": string(key),
                        
                        },
                },
            },
            Target: target,
        },
    }

    req := &gnmi.GetRequest{
        Encoding: gnmi.Encoding_JSON,
        Path:     path,
    }
    res, err := c.client.Get(gnmiCtx, req)

    if err != nil {
        return nil, err
    }

    val, err := gnmi_utils.GetResponseUpdate(res)

    if err != nil {
        return nil, err
    }

    json := val.GetJsonVal()
    st := Device{}
    Unmarshal(json, &st)

    if reflect.ValueOf(st.Template).Kind() == reflect.Ptr && reflect.ValueOf(st.Template).IsNil() {
        return nil, status.Error(codes.NotFound, "OnfTemplate_Template_Template-not-found")
    }
    if res, ok := st.Template.Template[key]; ok {
        return res, nil
    }

    return nil, status.Error(codes.NotFound, "OnfTemplate_Template_Template-not-found")
    }








func (c *GnmiClient) DeleteTemplate_Template_Item(ctx context.Context, target string, 
key string,
) (*gnmi.SetResponse, error) {
    gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
    defer cancel()


    
    
    path :=  []*gnmi.Path{
        {
            Elem: []*gnmi.PathElem{
            {
                    Name: "template",
                },
            {
                    Name: "template",
                    Key: map[string]string{
                        
                        "id": string(key),
                        
                        },
                },
            },
            Target: target,
        },
    }

    req := &gnmi.SetRequest{
        Delete: []*gnmi.Path{
            {
            Elem:   path[0].Elem,
            Target: target,
            },
        },
    }
    return c.client.Set(gnmiCtx, req)
    }








func (c *GnmiClient) UpdateTemplate_Template_Item(ctx context.Context, target string,  data OnfTemplate_Template_Template,
) (*gnmi.SetResponse, error) {
    gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
    defer cancel()


    
    

    
    path :=  []*gnmi.Path{
        {
            Elem: []*gnmi.PathElem{
            {
                    Name: "template",
                },
            {
                    Name: "template",
                    Key: map[string]string{
                        "id": string(*data.Id),
                        },
                },
            },
            Target: target,
        },
    }

    req, err := gnmi_utils.CreateGnmiSetForContainer(ctx, data, path[0], target)
    if err != nil {
        return nil, err
    }

    return c.client.Set(gnmiCtx, req)
    }








func (c *GnmiClient) GetUpf_Upf_Item(ctx context.Context, target string, 
key string,
) (*OnfUpf_Upf_Upf, error) {
    gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
    defer cancel()


    
    
    path :=  []*gnmi.Path{
        {
            Elem: []*gnmi.PathElem{
            {
                    Name: "upf",
                },
            {
                    Name: "upf",
                    Key: map[string]string{
                        
                        "id": string(key),
                        
                        },
                },
            },
            Target: target,
        },
    }

    req := &gnmi.GetRequest{
        Encoding: gnmi.Encoding_JSON,
        Path:     path,
    }
    res, err := c.client.Get(gnmiCtx, req)

    if err != nil {
        return nil, err
    }

    val, err := gnmi_utils.GetResponseUpdate(res)

    if err != nil {
        return nil, err
    }

    json := val.GetJsonVal()
    st := Device{}
    Unmarshal(json, &st)

    if reflect.ValueOf(st.Upf).Kind() == reflect.Ptr && reflect.ValueOf(st.Upf).IsNil() {
        return nil, status.Error(codes.NotFound, "OnfUpf_Upf_Upf-not-found")
    }
    if res, ok := st.Upf.Upf[key]; ok {
        return res, nil
    }

    return nil, status.Error(codes.NotFound, "OnfUpf_Upf_Upf-not-found")
    }








func (c *GnmiClient) DeleteUpf_Upf_Item(ctx context.Context, target string, 
key string,
) (*gnmi.SetResponse, error) {
    gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
    defer cancel()


    
    
    path :=  []*gnmi.Path{
        {
            Elem: []*gnmi.PathElem{
            {
                    Name: "upf",
                },
            {
                    Name: "upf",
                    Key: map[string]string{
                        
                        "id": string(key),
                        
                        },
                },
            },
            Target: target,
        },
    }

    req := &gnmi.SetRequest{
        Delete: []*gnmi.Path{
            {
            Elem:   path[0].Elem,
            Target: target,
            },
        },
    }
    return c.client.Set(gnmiCtx, req)
    }








func (c *GnmiClient) UpdateUpf_Upf_Item(ctx context.Context, target string,  data OnfUpf_Upf_Upf,
) (*gnmi.SetResponse, error) {
    gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
    defer cancel()


    
    

    
    path :=  []*gnmi.Path{
        {
            Elem: []*gnmi.PathElem{
            {
                    Name: "upf",
                },
            {
                    Name: "upf",
                    Key: map[string]string{
                        "id": string(*data.Id),
                        },
                },
            },
            Target: target,
        },
    }

    req, err := gnmi_utils.CreateGnmiSetForContainer(ctx, data, path[0], target)
    if err != nil {
        return nil, err
    }

    return c.client.Set(gnmiCtx, req)
    }








func (c *GnmiClient) GetDeviceGroup_DeviceGroup_Item(ctx context.Context, target string, 
key string,
) (*OnfDeviceGroup_DeviceGroup_DeviceGroup, error) {
    gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
    defer cancel()


    
    
    path :=  []*gnmi.Path{
        {
            Elem: []*gnmi.PathElem{
            {
                    Name: "device-group",
                },
            {
                    Name: "device-group",
                    Key: map[string]string{
                        
                        "id": string(key),
                        
                        },
                },
            },
            Target: target,
        },
    }

    req := &gnmi.GetRequest{
        Encoding: gnmi.Encoding_JSON,
        Path:     path,
    }
    res, err := c.client.Get(gnmiCtx, req)

    if err != nil {
        return nil, err
    }

    val, err := gnmi_utils.GetResponseUpdate(res)

    if err != nil {
        return nil, err
    }

    json := val.GetJsonVal()
    st := Device{}
    Unmarshal(json, &st)

    if reflect.ValueOf(st.DeviceGroup).Kind() == reflect.Ptr && reflect.ValueOf(st.DeviceGroup).IsNil() {
        return nil, status.Error(codes.NotFound, "OnfDeviceGroup_DeviceGroup_DeviceGroup-not-found")
    }
    if res, ok := st.DeviceGroup.DeviceGroup[key]; ok {
        return res, nil
    }

    return nil, status.Error(codes.NotFound, "OnfDeviceGroup_DeviceGroup_DeviceGroup-not-found")
    }








func (c *GnmiClient) DeleteDeviceGroup_DeviceGroup_Item(ctx context.Context, target string, 
key string,
) (*gnmi.SetResponse, error) {
    gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
    defer cancel()


    
    
    path :=  []*gnmi.Path{
        {
            Elem: []*gnmi.PathElem{
            {
                    Name: "device-group",
                },
            {
                    Name: "device-group",
                    Key: map[string]string{
                        
                        "id": string(key),
                        
                        },
                },
            },
            Target: target,
        },
    }

    req := &gnmi.SetRequest{
        Delete: []*gnmi.Path{
            {
            Elem:   path[0].Elem,
            Target: target,
            },
        },
    }
    return c.client.Set(gnmiCtx, req)
    }








func (c *GnmiClient) UpdateDeviceGroup_DeviceGroup_Item(ctx context.Context, target string,  data OnfDeviceGroup_DeviceGroup_DeviceGroup,
) (*gnmi.SetResponse, error) {
    gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
    defer cancel()


    
    

    
    path :=  []*gnmi.Path{
        {
            Elem: []*gnmi.PathElem{
            {
                    Name: "device-group",
                },
            {
                    Name: "device-group",
                    Key: map[string]string{
                        "id": string(*data.Id),
                        },
                },
            },
            Target: target,
        },
    }

    req, err := gnmi_utils.CreateGnmiSetForContainer(ctx, data, path[0], target)
    if err != nil {
        return nil, err
    }

    return c.client.Set(gnmiCtx, req)
    }





func (c *GnmiClient) GetConnectivityService_ConnectivityService(ctx context.Context, target string, 
) (map[string]*OnfConnectivityService_ConnectivityService_ConnectivityService, error) {
    gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
    defer cancel()

    

path :=  []*gnmi.Path{
    {
        Elem: []*gnmi.PathElem{
        {
                Name: "connectivity-service",
            },
        {
                Name: "connectivity-service",
            },
        },
        Target: target,
    },
}
    req := &gnmi.GetRequest{
        Encoding: gnmi.Encoding_JSON,
        Path:     path,
    }
    res, err := c.client.Get(gnmiCtx, req)

    if err != nil {
        return nil, err
    }

    val, err := gnmi_utils.GetResponseUpdate(res)

    if err != nil {
        return nil, err
    }

    json := val.GetJsonVal()
    st := Device{}
    Unmarshal(json, &st)

    if reflect.ValueOf(st.ConnectivityService).Kind() == reflect.Ptr && reflect.ValueOf(st.ConnectivityService).IsNil() {
        return nil, status.Error(codes.NotFound, "OnfConnectivityService_ConnectivityService_ConnectivityService-not-found")
    }
    if reflect.ValueOf(st.ConnectivityService.ConnectivityService).Kind() == reflect.Ptr && reflect.ValueOf(st.ConnectivityService.ConnectivityService).IsNil() {
        return nil, status.Error(codes.NotFound, "OnfConnectivityService_ConnectivityService_ConnectivityService-not-found")
    }

    return st.ConnectivityService.ConnectivityService, nil
}





func (c *GnmiClient) DeleteConnectivityService_ConnectivityService(ctx context.Context, target string, 
) (*gnmi.SetResponse, error) {
    gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
    defer cancel()

    

path :=  []*gnmi.Path{
    {
        Elem: []*gnmi.PathElem{
        {
                Name: "connectivity-service",
            },
        {
                Name: "connectivity-service",
            },
        },
        Target: target,
    },
}
    req := &gnmi.SetRequest{
        Delete: []*gnmi.Path{
            {
                Elem:   path[0].Elem,
                Target: target,
            },
        },
    }
    return c.client.Set(gnmiCtx, req)
}





func (c *GnmiClient) UpdateConnectivityService_ConnectivityService(ctx context.Context, target string,  list map[string]*OnfConnectivityService_ConnectivityService_ConnectivityService,
) (*gnmi.SetResponse, error) {
    gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
    defer cancel()

    
    basePathElems :=  []*gnmi.PathElem{
    {
        Name: "connectivity-service",
        },
    }
    req := &gnmi.SetRequest{
        Update: []*gnmi.Update{},
    }
    for _, item := range list {

        path := &gnmi.Path{
            Elem: append(basePathElems, &gnmi.PathElem{
                Name: "list2a",
                Key: map[string]string{
                    "id": string(*item.Id),
                    },
            }),
            Target: target,
        }

        // TODO if it's pointer, pass the value
        // if it's a value pass it directly
        r, err := gnmi_utils.CreateGnmiSetForContainer(ctx, *item, path, target)
        if err != nil {
            return nil, err
        }
        req.Update = append(req.Update, r.Update...)
    }

    return c.client.Set(gnmiCtx, req)
}





func (c *GnmiClient) GetEnterprise_Enterprise(ctx context.Context, target string, 
) (map[string]*OnfEnterprise_Enterprise_Enterprise, error) {
    gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
    defer cancel()

    

path :=  []*gnmi.Path{
    {
        Elem: []*gnmi.PathElem{
        {
                Name: "enterprise",
            },
        {
                Name: "enterprise",
            },
        },
        Target: target,
    },
}
    req := &gnmi.GetRequest{
        Encoding: gnmi.Encoding_JSON,
        Path:     path,
    }
    res, err := c.client.Get(gnmiCtx, req)

    if err != nil {
        return nil, err
    }

    val, err := gnmi_utils.GetResponseUpdate(res)

    if err != nil {
        return nil, err
    }

    json := val.GetJsonVal()
    st := Device{}
    Unmarshal(json, &st)

    if reflect.ValueOf(st.Enterprise).Kind() == reflect.Ptr && reflect.ValueOf(st.Enterprise).IsNil() {
        return nil, status.Error(codes.NotFound, "OnfEnterprise_Enterprise_Enterprise-not-found")
    }
    if reflect.ValueOf(st.Enterprise.Enterprise).Kind() == reflect.Ptr && reflect.ValueOf(st.Enterprise.Enterprise).IsNil() {
        return nil, status.Error(codes.NotFound, "OnfEnterprise_Enterprise_Enterprise-not-found")
    }

    return st.Enterprise.Enterprise, nil
}





func (c *GnmiClient) DeleteEnterprise_Enterprise(ctx context.Context, target string, 
) (*gnmi.SetResponse, error) {
    gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
    defer cancel()

    

path :=  []*gnmi.Path{
    {
        Elem: []*gnmi.PathElem{
        {
                Name: "enterprise",
            },
        {
                Name: "enterprise",
            },
        },
        Target: target,
    },
}
    req := &gnmi.SetRequest{
        Delete: []*gnmi.Path{
            {
                Elem:   path[0].Elem,
                Target: target,
            },
        },
    }
    return c.client.Set(gnmiCtx, req)
}





func (c *GnmiClient) UpdateEnterprise_Enterprise(ctx context.Context, target string,  list map[string]*OnfEnterprise_Enterprise_Enterprise,
) (*gnmi.SetResponse, error) {
    gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
    defer cancel()

    
    basePathElems :=  []*gnmi.PathElem{
    {
        Name: "enterprise",
        },
    }
    req := &gnmi.SetRequest{
        Update: []*gnmi.Update{},
    }
    for _, item := range list {

        path := &gnmi.Path{
            Elem: append(basePathElems, &gnmi.PathElem{
                Name: "list2a",
                Key: map[string]string{
                    "id": string(*item.Id),
                    },
            }),
            Target: target,
        }

        // TODO if it's pointer, pass the value
        // if it's a value pass it directly
        r, err := gnmi_utils.CreateGnmiSetForContainer(ctx, *item, path, target)
        if err != nil {
            return nil, err
        }
        req.Update = append(req.Update, r.Update...)
    }

    return c.client.Set(gnmiCtx, req)
}





func (c *GnmiClient) GetSite_Site(ctx context.Context, target string, 
) (map[string]*OnfSite_Site_Site, error) {
    gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
    defer cancel()

    

path :=  []*gnmi.Path{
    {
        Elem: []*gnmi.PathElem{
        {
                Name: "site",
            },
        {
                Name: "site",
            },
        },
        Target: target,
    },
}
    req := &gnmi.GetRequest{
        Encoding: gnmi.Encoding_JSON,
        Path:     path,
    }
    res, err := c.client.Get(gnmiCtx, req)

    if err != nil {
        return nil, err
    }

    val, err := gnmi_utils.GetResponseUpdate(res)

    if err != nil {
        return nil, err
    }

    json := val.GetJsonVal()
    st := Device{}
    Unmarshal(json, &st)

    if reflect.ValueOf(st.Site).Kind() == reflect.Ptr && reflect.ValueOf(st.Site).IsNil() {
        return nil, status.Error(codes.NotFound, "OnfSite_Site_Site-not-found")
    }
    if reflect.ValueOf(st.Site.Site).Kind() == reflect.Ptr && reflect.ValueOf(st.Site.Site).IsNil() {
        return nil, status.Error(codes.NotFound, "OnfSite_Site_Site-not-found")
    }

    return st.Site.Site, nil
}





func (c *GnmiClient) DeleteSite_Site(ctx context.Context, target string, 
) (*gnmi.SetResponse, error) {
    gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
    defer cancel()

    

path :=  []*gnmi.Path{
    {
        Elem: []*gnmi.PathElem{
        {
                Name: "site",
            },
        {
                Name: "site",
            },
        },
        Target: target,
    },
}
    req := &gnmi.SetRequest{
        Delete: []*gnmi.Path{
            {
                Elem:   path[0].Elem,
                Target: target,
            },
        },
    }
    return c.client.Set(gnmiCtx, req)
}





func (c *GnmiClient) UpdateSite_Site(ctx context.Context, target string,  list map[string]*OnfSite_Site_Site,
) (*gnmi.SetResponse, error) {
    gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
    defer cancel()

    
    basePathElems :=  []*gnmi.PathElem{
    {
        Name: "site",
        },
    }
    req := &gnmi.SetRequest{
        Update: []*gnmi.Update{},
    }
    for _, item := range list {

        path := &gnmi.Path{
            Elem: append(basePathElems, &gnmi.PathElem{
                Name: "list2a",
                Key: map[string]string{
                    "id": string(*item.Id),
                    },
            }),
            Target: target,
        }

        // TODO if it's pointer, pass the value
        // if it's a value pass it directly
        r, err := gnmi_utils.CreateGnmiSetForContainer(ctx, *item, path, target)
        if err != nil {
            return nil, err
        }
        req.Update = append(req.Update, r.Update...)
    }

    return c.client.Set(gnmiCtx, req)
}





func (c *GnmiClient) GetTrafficClass_TrafficClass(ctx context.Context, target string, 
) (map[string]*OnfTrafficClass_TrafficClass_TrafficClass, error) {
    gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
    defer cancel()

    

path :=  []*gnmi.Path{
    {
        Elem: []*gnmi.PathElem{
        {
                Name: "traffic-class",
            },
        {
                Name: "traffic-class",
            },
        },
        Target: target,
    },
}
    req := &gnmi.GetRequest{
        Encoding: gnmi.Encoding_JSON,
        Path:     path,
    }
    res, err := c.client.Get(gnmiCtx, req)

    if err != nil {
        return nil, err
    }

    val, err := gnmi_utils.GetResponseUpdate(res)

    if err != nil {
        return nil, err
    }

    json := val.GetJsonVal()
    st := Device{}
    Unmarshal(json, &st)

    if reflect.ValueOf(st.TrafficClass).Kind() == reflect.Ptr && reflect.ValueOf(st.TrafficClass).IsNil() {
        return nil, status.Error(codes.NotFound, "OnfTrafficClass_TrafficClass_TrafficClass-not-found")
    }
    if reflect.ValueOf(st.TrafficClass.TrafficClass).Kind() == reflect.Ptr && reflect.ValueOf(st.TrafficClass.TrafficClass).IsNil() {
        return nil, status.Error(codes.NotFound, "OnfTrafficClass_TrafficClass_TrafficClass-not-found")
    }

    return st.TrafficClass.TrafficClass, nil
}





func (c *GnmiClient) DeleteTrafficClass_TrafficClass(ctx context.Context, target string, 
) (*gnmi.SetResponse, error) {
    gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
    defer cancel()

    

path :=  []*gnmi.Path{
    {
        Elem: []*gnmi.PathElem{
        {
                Name: "traffic-class",
            },
        {
                Name: "traffic-class",
            },
        },
        Target: target,
    },
}
    req := &gnmi.SetRequest{
        Delete: []*gnmi.Path{
            {
                Elem:   path[0].Elem,
                Target: target,
            },
        },
    }
    return c.client.Set(gnmiCtx, req)
}





func (c *GnmiClient) UpdateTrafficClass_TrafficClass(ctx context.Context, target string,  list map[string]*OnfTrafficClass_TrafficClass_TrafficClass,
) (*gnmi.SetResponse, error) {
    gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
    defer cancel()

    
    basePathElems :=  []*gnmi.PathElem{
    {
        Name: "traffic-class",
        },
    }
    req := &gnmi.SetRequest{
        Update: []*gnmi.Update{},
    }
    for _, item := range list {

        path := &gnmi.Path{
            Elem: append(basePathElems, &gnmi.PathElem{
                Name: "list2a",
                Key: map[string]string{
                    "id": string(*item.Id),
                    },
            }),
            Target: target,
        }

        // TODO if it's pointer, pass the value
        // if it's a value pass it directly
        r, err := gnmi_utils.CreateGnmiSetForContainer(ctx, *item, path, target)
        if err != nil {
            return nil, err
        }
        req.Update = append(req.Update, r.Update...)
    }

    return c.client.Set(gnmiCtx, req)
}





func (c *GnmiClient) GetVcs_Vcs(ctx context.Context, target string, 
) (map[string]*OnfVcs_Vcs_Vcs, error) {
    gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
    defer cancel()

    

path :=  []*gnmi.Path{
    {
        Elem: []*gnmi.PathElem{
        {
                Name: "vcs",
            },
        {
                Name: "vcs",
            },
        },
        Target: target,
    },
}
    req := &gnmi.GetRequest{
        Encoding: gnmi.Encoding_JSON,
        Path:     path,
    }
    res, err := c.client.Get(gnmiCtx, req)

    if err != nil {
        return nil, err
    }

    val, err := gnmi_utils.GetResponseUpdate(res)

    if err != nil {
        return nil, err
    }

    json := val.GetJsonVal()
    st := Device{}
    Unmarshal(json, &st)

    if reflect.ValueOf(st.Vcs).Kind() == reflect.Ptr && reflect.ValueOf(st.Vcs).IsNil() {
        return nil, status.Error(codes.NotFound, "OnfVcs_Vcs_Vcs-not-found")
    }
    if reflect.ValueOf(st.Vcs.Vcs).Kind() == reflect.Ptr && reflect.ValueOf(st.Vcs.Vcs).IsNil() {
        return nil, status.Error(codes.NotFound, "OnfVcs_Vcs_Vcs-not-found")
    }

    return st.Vcs.Vcs, nil
}





func (c *GnmiClient) DeleteVcs_Vcs(ctx context.Context, target string, 
) (*gnmi.SetResponse, error) {
    gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
    defer cancel()

    

path :=  []*gnmi.Path{
    {
        Elem: []*gnmi.PathElem{
        {
                Name: "vcs",
            },
        {
                Name: "vcs",
            },
        },
        Target: target,
    },
}
    req := &gnmi.SetRequest{
        Delete: []*gnmi.Path{
            {
                Elem:   path[0].Elem,
                Target: target,
            },
        },
    }
    return c.client.Set(gnmiCtx, req)
}





func (c *GnmiClient) UpdateVcs_Vcs(ctx context.Context, target string,  list map[string]*OnfVcs_Vcs_Vcs,
) (*gnmi.SetResponse, error) {
    gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
    defer cancel()

    
    basePathElems :=  []*gnmi.PathElem{
    {
        Name: "vcs",
        },
    }
    req := &gnmi.SetRequest{
        Update: []*gnmi.Update{},
    }
    for _, item := range list {

        path := &gnmi.Path{
            Elem: append(basePathElems, &gnmi.PathElem{
                Name: "list2a",
                Key: map[string]string{
                    "id": string(*item.Id),
                    },
            }),
            Target: target,
        }

        // TODO if it's pointer, pass the value
        // if it's a value pass it directly
        r, err := gnmi_utils.CreateGnmiSetForContainer(ctx, *item, path, target)
        if err != nil {
            return nil, err
        }
        req.Update = append(req.Update, r.Update...)
    }

    return c.client.Set(gnmiCtx, req)
}





func (c *GnmiClient) GetApplication_Application(ctx context.Context, target string, 
) (map[string]*OnfApplication_Application_Application, error) {
    gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
    defer cancel()

    

path :=  []*gnmi.Path{
    {
        Elem: []*gnmi.PathElem{
        {
                Name: "application",
            },
        {
                Name: "application",
            },
        },
        Target: target,
    },
}
    req := &gnmi.GetRequest{
        Encoding: gnmi.Encoding_JSON,
        Path:     path,
    }
    res, err := c.client.Get(gnmiCtx, req)

    if err != nil {
        return nil, err
    }

    val, err := gnmi_utils.GetResponseUpdate(res)

    if err != nil {
        return nil, err
    }

    json := val.GetJsonVal()
    st := Device{}
    Unmarshal(json, &st)

    if reflect.ValueOf(st.Application).Kind() == reflect.Ptr && reflect.ValueOf(st.Application).IsNil() {
        return nil, status.Error(codes.NotFound, "OnfApplication_Application_Application-not-found")
    }
    if reflect.ValueOf(st.Application.Application).Kind() == reflect.Ptr && reflect.ValueOf(st.Application.Application).IsNil() {
        return nil, status.Error(codes.NotFound, "OnfApplication_Application_Application-not-found")
    }

    return st.Application.Application, nil
}





func (c *GnmiClient) DeleteApplication_Application(ctx context.Context, target string, 
) (*gnmi.SetResponse, error) {
    gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
    defer cancel()

    

path :=  []*gnmi.Path{
    {
        Elem: []*gnmi.PathElem{
        {
                Name: "application",
            },
        {
                Name: "application",
            },
        },
        Target: target,
    },
}
    req := &gnmi.SetRequest{
        Delete: []*gnmi.Path{
            {
                Elem:   path[0].Elem,
                Target: target,
            },
        },
    }
    return c.client.Set(gnmiCtx, req)
}





func (c *GnmiClient) UpdateApplication_Application(ctx context.Context, target string,  list map[string]*OnfApplication_Application_Application,
) (*gnmi.SetResponse, error) {
    gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
    defer cancel()

    
    basePathElems :=  []*gnmi.PathElem{
    {
        Name: "application",
        },
    }
    req := &gnmi.SetRequest{
        Update: []*gnmi.Update{},
    }
    for _, item := range list {

        path := &gnmi.Path{
            Elem: append(basePathElems, &gnmi.PathElem{
                Name: "list2a",
                Key: map[string]string{
                    "id": string(*item.Id),
                    },
            }),
            Target: target,
        }

        // TODO if it's pointer, pass the value
        // if it's a value pass it directly
        r, err := gnmi_utils.CreateGnmiSetForContainer(ctx, *item, path, target)
        if err != nil {
            return nil, err
        }
        req.Update = append(req.Update, r.Update...)
    }

    return c.client.Set(gnmiCtx, req)
}





func (c *GnmiClient) GetIpDomain_IpDomain(ctx context.Context, target string, 
) (map[string]*OnfIpDomain_IpDomain_IpDomain, error) {
    gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
    defer cancel()

    

path :=  []*gnmi.Path{
    {
        Elem: []*gnmi.PathElem{
        {
                Name: "ip-domain",
            },
        {
                Name: "ip-domain",
            },
        },
        Target: target,
    },
}
    req := &gnmi.GetRequest{
        Encoding: gnmi.Encoding_JSON,
        Path:     path,
    }
    res, err := c.client.Get(gnmiCtx, req)

    if err != nil {
        return nil, err
    }

    val, err := gnmi_utils.GetResponseUpdate(res)

    if err != nil {
        return nil, err
    }

    json := val.GetJsonVal()
    st := Device{}
    Unmarshal(json, &st)

    if reflect.ValueOf(st.IpDomain).Kind() == reflect.Ptr && reflect.ValueOf(st.IpDomain).IsNil() {
        return nil, status.Error(codes.NotFound, "OnfIpDomain_IpDomain_IpDomain-not-found")
    }
    if reflect.ValueOf(st.IpDomain.IpDomain).Kind() == reflect.Ptr && reflect.ValueOf(st.IpDomain.IpDomain).IsNil() {
        return nil, status.Error(codes.NotFound, "OnfIpDomain_IpDomain_IpDomain-not-found")
    }

    return st.IpDomain.IpDomain, nil
}





func (c *GnmiClient) DeleteIpDomain_IpDomain(ctx context.Context, target string, 
) (*gnmi.SetResponse, error) {
    gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
    defer cancel()

    

path :=  []*gnmi.Path{
    {
        Elem: []*gnmi.PathElem{
        {
                Name: "ip-domain",
            },
        {
                Name: "ip-domain",
            },
        },
        Target: target,
    },
}
    req := &gnmi.SetRequest{
        Delete: []*gnmi.Path{
            {
                Elem:   path[0].Elem,
                Target: target,
            },
        },
    }
    return c.client.Set(gnmiCtx, req)
}





func (c *GnmiClient) UpdateIpDomain_IpDomain(ctx context.Context, target string,  list map[string]*OnfIpDomain_IpDomain_IpDomain,
) (*gnmi.SetResponse, error) {
    gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
    defer cancel()

    
    basePathElems :=  []*gnmi.PathElem{
    {
        Name: "ip-domain",
        },
    }
    req := &gnmi.SetRequest{
        Update: []*gnmi.Update{},
    }
    for _, item := range list {

        path := &gnmi.Path{
            Elem: append(basePathElems, &gnmi.PathElem{
                Name: "list2a",
                Key: map[string]string{
                    "id": string(*item.Id),
                    },
            }),
            Target: target,
        }

        // TODO if it's pointer, pass the value
        // if it's a value pass it directly
        r, err := gnmi_utils.CreateGnmiSetForContainer(ctx, *item, path, target)
        if err != nil {
            return nil, err
        }
        req.Update = append(req.Update, r.Update...)
    }

    return c.client.Set(gnmiCtx, req)
}





func (c *GnmiClient) GetTemplate_Template(ctx context.Context, target string, 
) (map[string]*OnfTemplate_Template_Template, error) {
    gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
    defer cancel()

    

path :=  []*gnmi.Path{
    {
        Elem: []*gnmi.PathElem{
        {
                Name: "template",
            },
        {
                Name: "template",
            },
        },
        Target: target,
    },
}
    req := &gnmi.GetRequest{
        Encoding: gnmi.Encoding_JSON,
        Path:     path,
    }
    res, err := c.client.Get(gnmiCtx, req)

    if err != nil {
        return nil, err
    }

    val, err := gnmi_utils.GetResponseUpdate(res)

    if err != nil {
        return nil, err
    }

    json := val.GetJsonVal()
    st := Device{}
    Unmarshal(json, &st)

    if reflect.ValueOf(st.Template).Kind() == reflect.Ptr && reflect.ValueOf(st.Template).IsNil() {
        return nil, status.Error(codes.NotFound, "OnfTemplate_Template_Template-not-found")
    }
    if reflect.ValueOf(st.Template.Template).Kind() == reflect.Ptr && reflect.ValueOf(st.Template.Template).IsNil() {
        return nil, status.Error(codes.NotFound, "OnfTemplate_Template_Template-not-found")
    }

    return st.Template.Template, nil
}





func (c *GnmiClient) DeleteTemplate_Template(ctx context.Context, target string, 
) (*gnmi.SetResponse, error) {
    gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
    defer cancel()

    

path :=  []*gnmi.Path{
    {
        Elem: []*gnmi.PathElem{
        {
                Name: "template",
            },
        {
                Name: "template",
            },
        },
        Target: target,
    },
}
    req := &gnmi.SetRequest{
        Delete: []*gnmi.Path{
            {
                Elem:   path[0].Elem,
                Target: target,
            },
        },
    }
    return c.client.Set(gnmiCtx, req)
}





func (c *GnmiClient) UpdateTemplate_Template(ctx context.Context, target string,  list map[string]*OnfTemplate_Template_Template,
) (*gnmi.SetResponse, error) {
    gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
    defer cancel()

    
    basePathElems :=  []*gnmi.PathElem{
    {
        Name: "template",
        },
    }
    req := &gnmi.SetRequest{
        Update: []*gnmi.Update{},
    }
    for _, item := range list {

        path := &gnmi.Path{
            Elem: append(basePathElems, &gnmi.PathElem{
                Name: "list2a",
                Key: map[string]string{
                    "id": string(*item.Id),
                    },
            }),
            Target: target,
        }

        // TODO if it's pointer, pass the value
        // if it's a value pass it directly
        r, err := gnmi_utils.CreateGnmiSetForContainer(ctx, *item, path, target)
        if err != nil {
            return nil, err
        }
        req.Update = append(req.Update, r.Update...)
    }

    return c.client.Set(gnmiCtx, req)
}





func (c *GnmiClient) GetUpf_Upf(ctx context.Context, target string, 
) (map[string]*OnfUpf_Upf_Upf, error) {
    gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
    defer cancel()

    

path :=  []*gnmi.Path{
    {
        Elem: []*gnmi.PathElem{
        {
                Name: "upf",
            },
        {
                Name: "upf",
            },
        },
        Target: target,
    },
}
    req := &gnmi.GetRequest{
        Encoding: gnmi.Encoding_JSON,
        Path:     path,
    }
    res, err := c.client.Get(gnmiCtx, req)

    if err != nil {
        return nil, err
    }

    val, err := gnmi_utils.GetResponseUpdate(res)

    if err != nil {
        return nil, err
    }

    json := val.GetJsonVal()
    st := Device{}
    Unmarshal(json, &st)

    if reflect.ValueOf(st.Upf).Kind() == reflect.Ptr && reflect.ValueOf(st.Upf).IsNil() {
        return nil, status.Error(codes.NotFound, "OnfUpf_Upf_Upf-not-found")
    }
    if reflect.ValueOf(st.Upf.Upf).Kind() == reflect.Ptr && reflect.ValueOf(st.Upf.Upf).IsNil() {
        return nil, status.Error(codes.NotFound, "OnfUpf_Upf_Upf-not-found")
    }

    return st.Upf.Upf, nil
}





func (c *GnmiClient) DeleteUpf_Upf(ctx context.Context, target string, 
) (*gnmi.SetResponse, error) {
    gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
    defer cancel()

    

path :=  []*gnmi.Path{
    {
        Elem: []*gnmi.PathElem{
        {
                Name: "upf",
            },
        {
                Name: "upf",
            },
        },
        Target: target,
    },
}
    req := &gnmi.SetRequest{
        Delete: []*gnmi.Path{
            {
                Elem:   path[0].Elem,
                Target: target,
            },
        },
    }
    return c.client.Set(gnmiCtx, req)
}





func (c *GnmiClient) UpdateUpf_Upf(ctx context.Context, target string,  list map[string]*OnfUpf_Upf_Upf,
) (*gnmi.SetResponse, error) {
    gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
    defer cancel()

    
    basePathElems :=  []*gnmi.PathElem{
    {
        Name: "upf",
        },
    }
    req := &gnmi.SetRequest{
        Update: []*gnmi.Update{},
    }
    for _, item := range list {

        path := &gnmi.Path{
            Elem: append(basePathElems, &gnmi.PathElem{
                Name: "list2a",
                Key: map[string]string{
                    "id": string(*item.Id),
                    },
            }),
            Target: target,
        }

        // TODO if it's pointer, pass the value
        // if it's a value pass it directly
        r, err := gnmi_utils.CreateGnmiSetForContainer(ctx, *item, path, target)
        if err != nil {
            return nil, err
        }
        req.Update = append(req.Update, r.Update...)
    }

    return c.client.Set(gnmiCtx, req)
}





func (c *GnmiClient) GetDeviceGroup_DeviceGroup(ctx context.Context, target string, 
) (map[string]*OnfDeviceGroup_DeviceGroup_DeviceGroup, error) {
    gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
    defer cancel()

    

path :=  []*gnmi.Path{
    {
        Elem: []*gnmi.PathElem{
        {
                Name: "device-group",
            },
        {
                Name: "device-group",
            },
        },
        Target: target,
    },
}
    req := &gnmi.GetRequest{
        Encoding: gnmi.Encoding_JSON,
        Path:     path,
    }
    res, err := c.client.Get(gnmiCtx, req)

    if err != nil {
        return nil, err
    }

    val, err := gnmi_utils.GetResponseUpdate(res)

    if err != nil {
        return nil, err
    }

    json := val.GetJsonVal()
    st := Device{}
    Unmarshal(json, &st)

    if reflect.ValueOf(st.DeviceGroup).Kind() == reflect.Ptr && reflect.ValueOf(st.DeviceGroup).IsNil() {
        return nil, status.Error(codes.NotFound, "OnfDeviceGroup_DeviceGroup_DeviceGroup-not-found")
    }
    if reflect.ValueOf(st.DeviceGroup.DeviceGroup).Kind() == reflect.Ptr && reflect.ValueOf(st.DeviceGroup.DeviceGroup).IsNil() {
        return nil, status.Error(codes.NotFound, "OnfDeviceGroup_DeviceGroup_DeviceGroup-not-found")
    }

    return st.DeviceGroup.DeviceGroup, nil
}





func (c *GnmiClient) DeleteDeviceGroup_DeviceGroup(ctx context.Context, target string, 
) (*gnmi.SetResponse, error) {
    gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
    defer cancel()

    

path :=  []*gnmi.Path{
    {
        Elem: []*gnmi.PathElem{
        {
                Name: "device-group",
            },
        {
                Name: "device-group",
            },
        },
        Target: target,
    },
}
    req := &gnmi.SetRequest{
        Delete: []*gnmi.Path{
            {
                Elem:   path[0].Elem,
                Target: target,
            },
        },
    }
    return c.client.Set(gnmiCtx, req)
}





func (c *GnmiClient) UpdateDeviceGroup_DeviceGroup(ctx context.Context, target string,  list map[string]*OnfDeviceGroup_DeviceGroup_DeviceGroup,
) (*gnmi.SetResponse, error) {
    gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
    defer cancel()

    
    basePathElems :=  []*gnmi.PathElem{
    {
        Name: "device-group",
        },
    }
    req := &gnmi.SetRequest{
        Update: []*gnmi.Update{},
    }
    for _, item := range list {

        path := &gnmi.Path{
            Elem: append(basePathElems, &gnmi.PathElem{
                Name: "list2a",
                Key: map[string]string{
                    "id": string(*item.Id),
                    },
            }),
            Target: target,
        }

        // TODO if it's pointer, pass the value
        // if it's a value pass it directly
        r, err := gnmi_utils.CreateGnmiSetForContainer(ctx, *item, path, target)
        if err != nil {
            return nil, err
        }
        req.Update = append(req.Update, r.Update...)
    }

    return c.client.Set(gnmiCtx, req)
}



func (c *GnmiClient) GetConnectivityService(ctx context.Context, target string, 
) (*OnfConnectivityService_ConnectivityService, error) {
gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
defer cancel()



path :=  []*gnmi.Path{
    {
        Elem: []*gnmi.PathElem{
        {
                Name: "connectivity-service",
            },
        },
        Target: target,
    },
}

req := &gnmi.GetRequest{
    Encoding: gnmi.Encoding_JSON,
    Path:     path,
    }
    res, err := c.client.Get(gnmiCtx, req)

    if err != nil {
    return nil, err
    }

    val, err := gnmi_utils.GetResponseUpdate(res)

    if err != nil {
    return nil, err
    }

    json := val.GetJsonVal()
    st := Device{}
    Unmarshal(json, &st)

    if reflect.ValueOf(st.ConnectivityService).Kind() == reflect.Ptr && reflect.ValueOf(st.ConnectivityService).IsNil() {
    return nil, status.Error(codes.NotFound, "OnfConnectivityService_ConnectivityService-not-found")
    }

    return st.ConnectivityService, nil

}



func (c *GnmiClient) DeleteConnectivityService(ctx context.Context, target string, 
) (*gnmi.SetResponse, error) {
gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
defer cancel()



path :=  []*gnmi.Path{
    {
        Elem: []*gnmi.PathElem{
        {
                Name: "connectivity-service",
            },
        },
        Target: target,
    },
}


req := &gnmi.SetRequest{
    Delete: []*gnmi.Path{
    {
    Elem:   path[0].Elem,
    Target: target,
    },
    },
    }
    return c.client.Set(gnmiCtx, req)
}



func (c *GnmiClient) UpdateConnectivityService(ctx context.Context, target string,  data OnfConnectivityService_ConnectivityService,
) (*gnmi.SetResponse, error) {
gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
defer cancel()



path :=  []*gnmi.Path{
    {
        Elem: []*gnmi.PathElem{
        {
                Name: "connectivity-service",
            },
        },
        Target: target,
    },
}


req, err := gnmi_utils.CreateGnmiSetForContainer(ctx, data, path[0], target)
    if err != nil {
    return nil, err
    }

    return c.client.Set(gnmiCtx, req)
}



func (c *GnmiClient) GetEnterprise(ctx context.Context, target string, 
) (*OnfEnterprise_Enterprise, error) {
gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
defer cancel()



path :=  []*gnmi.Path{
    {
        Elem: []*gnmi.PathElem{
        {
                Name: "enterprise",
            },
        },
        Target: target,
    },
}

req := &gnmi.GetRequest{
    Encoding: gnmi.Encoding_JSON,
    Path:     path,
    }
    res, err := c.client.Get(gnmiCtx, req)

    if err != nil {
    return nil, err
    }

    val, err := gnmi_utils.GetResponseUpdate(res)

    if err != nil {
    return nil, err
    }

    json := val.GetJsonVal()
    st := Device{}
    Unmarshal(json, &st)

    if reflect.ValueOf(st.Enterprise).Kind() == reflect.Ptr && reflect.ValueOf(st.Enterprise).IsNil() {
    return nil, status.Error(codes.NotFound, "OnfEnterprise_Enterprise-not-found")
    }

    return st.Enterprise, nil

}



func (c *GnmiClient) DeleteEnterprise(ctx context.Context, target string, 
) (*gnmi.SetResponse, error) {
gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
defer cancel()



path :=  []*gnmi.Path{
    {
        Elem: []*gnmi.PathElem{
        {
                Name: "enterprise",
            },
        },
        Target: target,
    },
}


req := &gnmi.SetRequest{
    Delete: []*gnmi.Path{
    {
    Elem:   path[0].Elem,
    Target: target,
    },
    },
    }
    return c.client.Set(gnmiCtx, req)
}



func (c *GnmiClient) UpdateEnterprise(ctx context.Context, target string,  data OnfEnterprise_Enterprise,
) (*gnmi.SetResponse, error) {
gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
defer cancel()



path :=  []*gnmi.Path{
    {
        Elem: []*gnmi.PathElem{
        {
                Name: "enterprise",
            },
        },
        Target: target,
    },
}


req, err := gnmi_utils.CreateGnmiSetForContainer(ctx, data, path[0], target)
    if err != nil {
    return nil, err
    }

    return c.client.Set(gnmiCtx, req)
}



func (c *GnmiClient) GetSite(ctx context.Context, target string, 
) (*OnfSite_Site, error) {
gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
defer cancel()



path :=  []*gnmi.Path{
    {
        Elem: []*gnmi.PathElem{
        {
                Name: "site",
            },
        },
        Target: target,
    },
}

req := &gnmi.GetRequest{
    Encoding: gnmi.Encoding_JSON,
    Path:     path,
    }
    res, err := c.client.Get(gnmiCtx, req)

    if err != nil {
    return nil, err
    }

    val, err := gnmi_utils.GetResponseUpdate(res)

    if err != nil {
    return nil, err
    }

    json := val.GetJsonVal()
    st := Device{}
    Unmarshal(json, &st)

    if reflect.ValueOf(st.Site).Kind() == reflect.Ptr && reflect.ValueOf(st.Site).IsNil() {
    return nil, status.Error(codes.NotFound, "OnfSite_Site-not-found")
    }

    return st.Site, nil

}



func (c *GnmiClient) DeleteSite(ctx context.Context, target string, 
) (*gnmi.SetResponse, error) {
gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
defer cancel()



path :=  []*gnmi.Path{
    {
        Elem: []*gnmi.PathElem{
        {
                Name: "site",
            },
        },
        Target: target,
    },
}


req := &gnmi.SetRequest{
    Delete: []*gnmi.Path{
    {
    Elem:   path[0].Elem,
    Target: target,
    },
    },
    }
    return c.client.Set(gnmiCtx, req)
}



func (c *GnmiClient) UpdateSite(ctx context.Context, target string,  data OnfSite_Site,
) (*gnmi.SetResponse, error) {
gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
defer cancel()



path :=  []*gnmi.Path{
    {
        Elem: []*gnmi.PathElem{
        {
                Name: "site",
            },
        },
        Target: target,
    },
}


req, err := gnmi_utils.CreateGnmiSetForContainer(ctx, data, path[0], target)
    if err != nil {
    return nil, err
    }

    return c.client.Set(gnmiCtx, req)
}



func (c *GnmiClient) GetTrafficClass(ctx context.Context, target string, 
) (*OnfTrafficClass_TrafficClass, error) {
gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
defer cancel()



path :=  []*gnmi.Path{
    {
        Elem: []*gnmi.PathElem{
        {
                Name: "traffic-class",
            },
        },
        Target: target,
    },
}

req := &gnmi.GetRequest{
    Encoding: gnmi.Encoding_JSON,
    Path:     path,
    }
    res, err := c.client.Get(gnmiCtx, req)

    if err != nil {
    return nil, err
    }

    val, err := gnmi_utils.GetResponseUpdate(res)

    if err != nil {
    return nil, err
    }

    json := val.GetJsonVal()
    st := Device{}
    Unmarshal(json, &st)

    if reflect.ValueOf(st.TrafficClass).Kind() == reflect.Ptr && reflect.ValueOf(st.TrafficClass).IsNil() {
    return nil, status.Error(codes.NotFound, "OnfTrafficClass_TrafficClass-not-found")
    }

    return st.TrafficClass, nil

}



func (c *GnmiClient) DeleteTrafficClass(ctx context.Context, target string, 
) (*gnmi.SetResponse, error) {
gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
defer cancel()



path :=  []*gnmi.Path{
    {
        Elem: []*gnmi.PathElem{
        {
                Name: "traffic-class",
            },
        },
        Target: target,
    },
}


req := &gnmi.SetRequest{
    Delete: []*gnmi.Path{
    {
    Elem:   path[0].Elem,
    Target: target,
    },
    },
    }
    return c.client.Set(gnmiCtx, req)
}



func (c *GnmiClient) UpdateTrafficClass(ctx context.Context, target string,  data OnfTrafficClass_TrafficClass,
) (*gnmi.SetResponse, error) {
gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
defer cancel()



path :=  []*gnmi.Path{
    {
        Elem: []*gnmi.PathElem{
        {
                Name: "traffic-class",
            },
        },
        Target: target,
    },
}


req, err := gnmi_utils.CreateGnmiSetForContainer(ctx, data, path[0], target)
    if err != nil {
    return nil, err
    }

    return c.client.Set(gnmiCtx, req)
}



func (c *GnmiClient) GetVcs(ctx context.Context, target string, 
) (*OnfVcs_Vcs, error) {
gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
defer cancel()



path :=  []*gnmi.Path{
    {
        Elem: []*gnmi.PathElem{
        {
                Name: "vcs",
            },
        },
        Target: target,
    },
}

req := &gnmi.GetRequest{
    Encoding: gnmi.Encoding_JSON,
    Path:     path,
    }
    res, err := c.client.Get(gnmiCtx, req)

    if err != nil {
    return nil, err
    }

    val, err := gnmi_utils.GetResponseUpdate(res)

    if err != nil {
    return nil, err
    }

    json := val.GetJsonVal()
    st := Device{}
    Unmarshal(json, &st)

    if reflect.ValueOf(st.Vcs).Kind() == reflect.Ptr && reflect.ValueOf(st.Vcs).IsNil() {
    return nil, status.Error(codes.NotFound, "OnfVcs_Vcs-not-found")
    }

    return st.Vcs, nil

}



func (c *GnmiClient) DeleteVcs(ctx context.Context, target string, 
) (*gnmi.SetResponse, error) {
gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
defer cancel()



path :=  []*gnmi.Path{
    {
        Elem: []*gnmi.PathElem{
        {
                Name: "vcs",
            },
        },
        Target: target,
    },
}


req := &gnmi.SetRequest{
    Delete: []*gnmi.Path{
    {
    Elem:   path[0].Elem,
    Target: target,
    },
    },
    }
    return c.client.Set(gnmiCtx, req)
}



func (c *GnmiClient) UpdateVcs(ctx context.Context, target string,  data OnfVcs_Vcs,
) (*gnmi.SetResponse, error) {
gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
defer cancel()



path :=  []*gnmi.Path{
    {
        Elem: []*gnmi.PathElem{
        {
                Name: "vcs",
            },
        },
        Target: target,
    },
}


req, err := gnmi_utils.CreateGnmiSetForContainer(ctx, data, path[0], target)
    if err != nil {
    return nil, err
    }

    return c.client.Set(gnmiCtx, req)
}



func (c *GnmiClient) GetApplication(ctx context.Context, target string, 
) (*OnfApplication_Application, error) {
gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
defer cancel()



path :=  []*gnmi.Path{
    {
        Elem: []*gnmi.PathElem{
        {
                Name: "application",
            },
        },
        Target: target,
    },
}

req := &gnmi.GetRequest{
    Encoding: gnmi.Encoding_JSON,
    Path:     path,
    }
    res, err := c.client.Get(gnmiCtx, req)

    if err != nil {
    return nil, err
    }

    val, err := gnmi_utils.GetResponseUpdate(res)

    if err != nil {
    return nil, err
    }

    json := val.GetJsonVal()
    st := Device{}
    Unmarshal(json, &st)

    if reflect.ValueOf(st.Application).Kind() == reflect.Ptr && reflect.ValueOf(st.Application).IsNil() {
    return nil, status.Error(codes.NotFound, "OnfApplication_Application-not-found")
    }

    return st.Application, nil

}



func (c *GnmiClient) DeleteApplication(ctx context.Context, target string, 
) (*gnmi.SetResponse, error) {
gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
defer cancel()



path :=  []*gnmi.Path{
    {
        Elem: []*gnmi.PathElem{
        {
                Name: "application",
            },
        },
        Target: target,
    },
}


req := &gnmi.SetRequest{
    Delete: []*gnmi.Path{
    {
    Elem:   path[0].Elem,
    Target: target,
    },
    },
    }
    return c.client.Set(gnmiCtx, req)
}



func (c *GnmiClient) UpdateApplication(ctx context.Context, target string,  data OnfApplication_Application,
) (*gnmi.SetResponse, error) {
gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
defer cancel()



path :=  []*gnmi.Path{
    {
        Elem: []*gnmi.PathElem{
        {
                Name: "application",
            },
        },
        Target: target,
    },
}


req, err := gnmi_utils.CreateGnmiSetForContainer(ctx, data, path[0], target)
    if err != nil {
    return nil, err
    }

    return c.client.Set(gnmiCtx, req)
}



func (c *GnmiClient) GetIpDomain(ctx context.Context, target string, 
) (*OnfIpDomain_IpDomain, error) {
gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
defer cancel()



path :=  []*gnmi.Path{
    {
        Elem: []*gnmi.PathElem{
        {
                Name: "ip-domain",
            },
        },
        Target: target,
    },
}

req := &gnmi.GetRequest{
    Encoding: gnmi.Encoding_JSON,
    Path:     path,
    }
    res, err := c.client.Get(gnmiCtx, req)

    if err != nil {
    return nil, err
    }

    val, err := gnmi_utils.GetResponseUpdate(res)

    if err != nil {
    return nil, err
    }

    json := val.GetJsonVal()
    st := Device{}
    Unmarshal(json, &st)

    if reflect.ValueOf(st.IpDomain).Kind() == reflect.Ptr && reflect.ValueOf(st.IpDomain).IsNil() {
    return nil, status.Error(codes.NotFound, "OnfIpDomain_IpDomain-not-found")
    }

    return st.IpDomain, nil

}



func (c *GnmiClient) DeleteIpDomain(ctx context.Context, target string, 
) (*gnmi.SetResponse, error) {
gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
defer cancel()



path :=  []*gnmi.Path{
    {
        Elem: []*gnmi.PathElem{
        {
                Name: "ip-domain",
            },
        },
        Target: target,
    },
}


req := &gnmi.SetRequest{
    Delete: []*gnmi.Path{
    {
    Elem:   path[0].Elem,
    Target: target,
    },
    },
    }
    return c.client.Set(gnmiCtx, req)
}



func (c *GnmiClient) UpdateIpDomain(ctx context.Context, target string,  data OnfIpDomain_IpDomain,
) (*gnmi.SetResponse, error) {
gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
defer cancel()



path :=  []*gnmi.Path{
    {
        Elem: []*gnmi.PathElem{
        {
                Name: "ip-domain",
            },
        },
        Target: target,
    },
}


req, err := gnmi_utils.CreateGnmiSetForContainer(ctx, data, path[0], target)
    if err != nil {
    return nil, err
    }

    return c.client.Set(gnmiCtx, req)
}



func (c *GnmiClient) GetTemplate(ctx context.Context, target string, 
) (*OnfTemplate_Template, error) {
gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
defer cancel()



path :=  []*gnmi.Path{
    {
        Elem: []*gnmi.PathElem{
        {
                Name: "template",
            },
        },
        Target: target,
    },
}

req := &gnmi.GetRequest{
    Encoding: gnmi.Encoding_JSON,
    Path:     path,
    }
    res, err := c.client.Get(gnmiCtx, req)

    if err != nil {
    return nil, err
    }

    val, err := gnmi_utils.GetResponseUpdate(res)

    if err != nil {
    return nil, err
    }

    json := val.GetJsonVal()
    st := Device{}
    Unmarshal(json, &st)

    if reflect.ValueOf(st.Template).Kind() == reflect.Ptr && reflect.ValueOf(st.Template).IsNil() {
    return nil, status.Error(codes.NotFound, "OnfTemplate_Template-not-found")
    }

    return st.Template, nil

}



func (c *GnmiClient) DeleteTemplate(ctx context.Context, target string, 
) (*gnmi.SetResponse, error) {
gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
defer cancel()



path :=  []*gnmi.Path{
    {
        Elem: []*gnmi.PathElem{
        {
                Name: "template",
            },
        },
        Target: target,
    },
}


req := &gnmi.SetRequest{
    Delete: []*gnmi.Path{
    {
    Elem:   path[0].Elem,
    Target: target,
    },
    },
    }
    return c.client.Set(gnmiCtx, req)
}



func (c *GnmiClient) UpdateTemplate(ctx context.Context, target string,  data OnfTemplate_Template,
) (*gnmi.SetResponse, error) {
gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
defer cancel()



path :=  []*gnmi.Path{
    {
        Elem: []*gnmi.PathElem{
        {
                Name: "template",
            },
        },
        Target: target,
    },
}


req, err := gnmi_utils.CreateGnmiSetForContainer(ctx, data, path[0], target)
    if err != nil {
    return nil, err
    }

    return c.client.Set(gnmiCtx, req)
}



func (c *GnmiClient) GetUpf(ctx context.Context, target string, 
) (*OnfUpf_Upf, error) {
gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
defer cancel()



path :=  []*gnmi.Path{
    {
        Elem: []*gnmi.PathElem{
        {
                Name: "upf",
            },
        },
        Target: target,
    },
}

req := &gnmi.GetRequest{
    Encoding: gnmi.Encoding_JSON,
    Path:     path,
    }
    res, err := c.client.Get(gnmiCtx, req)

    if err != nil {
    return nil, err
    }

    val, err := gnmi_utils.GetResponseUpdate(res)

    if err != nil {
    return nil, err
    }

    json := val.GetJsonVal()
    st := Device{}
    Unmarshal(json, &st)

    if reflect.ValueOf(st.Upf).Kind() == reflect.Ptr && reflect.ValueOf(st.Upf).IsNil() {
    return nil, status.Error(codes.NotFound, "OnfUpf_Upf-not-found")
    }

    return st.Upf, nil

}



func (c *GnmiClient) DeleteUpf(ctx context.Context, target string, 
) (*gnmi.SetResponse, error) {
gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
defer cancel()



path :=  []*gnmi.Path{
    {
        Elem: []*gnmi.PathElem{
        {
                Name: "upf",
            },
        },
        Target: target,
    },
}


req := &gnmi.SetRequest{
    Delete: []*gnmi.Path{
    {
    Elem:   path[0].Elem,
    Target: target,
    },
    },
    }
    return c.client.Set(gnmiCtx, req)
}



func (c *GnmiClient) UpdateUpf(ctx context.Context, target string,  data OnfUpf_Upf,
) (*gnmi.SetResponse, error) {
gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
defer cancel()



path :=  []*gnmi.Path{
    {
        Elem: []*gnmi.PathElem{
        {
                Name: "upf",
            },
        },
        Target: target,
    },
}


req, err := gnmi_utils.CreateGnmiSetForContainer(ctx, data, path[0], target)
    if err != nil {
    return nil, err
    }

    return c.client.Set(gnmiCtx, req)
}



func (c *GnmiClient) GetDeviceGroup(ctx context.Context, target string, 
) (*OnfDeviceGroup_DeviceGroup, error) {
gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
defer cancel()



path :=  []*gnmi.Path{
    {
        Elem: []*gnmi.PathElem{
        {
                Name: "device-group",
            },
        },
        Target: target,
    },
}

req := &gnmi.GetRequest{
    Encoding: gnmi.Encoding_JSON,
    Path:     path,
    }
    res, err := c.client.Get(gnmiCtx, req)

    if err != nil {
    return nil, err
    }

    val, err := gnmi_utils.GetResponseUpdate(res)

    if err != nil {
    return nil, err
    }

    json := val.GetJsonVal()
    st := Device{}
    Unmarshal(json, &st)

    if reflect.ValueOf(st.DeviceGroup).Kind() == reflect.Ptr && reflect.ValueOf(st.DeviceGroup).IsNil() {
    return nil, status.Error(codes.NotFound, "OnfDeviceGroup_DeviceGroup-not-found")
    }

    return st.DeviceGroup, nil

}



func (c *GnmiClient) DeleteDeviceGroup(ctx context.Context, target string, 
) (*gnmi.SetResponse, error) {
gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
defer cancel()



path :=  []*gnmi.Path{
    {
        Elem: []*gnmi.PathElem{
        {
                Name: "device-group",
            },
        },
        Target: target,
    },
}


req := &gnmi.SetRequest{
    Delete: []*gnmi.Path{
    {
    Elem:   path[0].Elem,
    Target: target,
    },
    },
    }
    return c.client.Set(gnmiCtx, req)
}



func (c *GnmiClient) UpdateDeviceGroup(ctx context.Context, target string,  data OnfDeviceGroup_DeviceGroup,
) (*gnmi.SetResponse, error) {
gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
defer cancel()



path :=  []*gnmi.Path{
    {
        Elem: []*gnmi.PathElem{
        {
                Name: "device-group",
            },
        },
        Target: target,
    },
}


req, err := gnmi_utils.CreateGnmiSetForContainer(ctx, data, path[0], target)
    if err != nil {
    return nil, err
    }

    return c.client.Set(gnmiCtx, req)
}


