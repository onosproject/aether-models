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









func (c *GnmiClient) GetConnectivityServices_ConnectivityService_Item(ctx context.Context, target string, 
key string,
) (*OnfConnectivityService_ConnectivityServices_ConnectivityService, error) {
    gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
    defer cancel()


    
    
    path :=  []*gnmi.Path{
        {
            Elem: []*gnmi.PathElem{
            {
                    Name: "connectivity-services",
                },
            {
                    Name: "connectivity-service",
                    Key: map[string]string{
                        
                        "connectivity-service-id": string(key),
                        
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

    if reflect.ValueOf(st.ConnectivityServices).Kind() == reflect.Ptr && reflect.ValueOf(st.ConnectivityServices).IsNil() {
        return nil, status.Error(codes.NotFound, "OnfConnectivityService_ConnectivityServices_ConnectivityService-not-found")
    }
    if res, ok := st.ConnectivityServices.ConnectivityService[key]; ok {
        return res, nil
    }

    return nil, status.Error(codes.NotFound, "OnfConnectivityService_ConnectivityServices_ConnectivityService-not-found")
    }








func (c *GnmiClient) DeleteConnectivityServices_ConnectivityService_Item(ctx context.Context, target string, 
key string,
) (*gnmi.SetResponse, error) {
    gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
    defer cancel()


    
    
    path :=  []*gnmi.Path{
        {
            Elem: []*gnmi.PathElem{
            {
                    Name: "connectivity-services",
                },
            {
                    Name: "connectivity-service",
                    Key: map[string]string{
                        
                        "connectivity-service-id": string(key),
                        
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








func (c *GnmiClient) UpdateConnectivityServices_ConnectivityService_Item(ctx context.Context, target string,  data OnfConnectivityService_ConnectivityServices_ConnectivityService,
) (*gnmi.SetResponse, error) {
    gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
    defer cancel()


    
    

    
    path :=  []*gnmi.Path{
        {
            Elem: []*gnmi.PathElem{
            {
                    Name: "connectivity-services",
                },
            {
                    Name: "connectivity-service",
                    Key: map[string]string{
                        "connectivity-service-id": string(*data.ConnectivityServiceId),
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








func (c *GnmiClient) GetEnterprises_Enterprise_Item(ctx context.Context, target string, 
key string,
) (*OnfEnterprise_Enterprises_Enterprise, error) {
    gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
    defer cancel()


    
    
    path :=  []*gnmi.Path{
        {
            Elem: []*gnmi.PathElem{
            {
                    Name: "enterprises",
                },
            {
                    Name: "enterprise",
                    Key: map[string]string{
                        
                        "enterprise-id": string(key),
                        
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

    if reflect.ValueOf(st.Enterprises).Kind() == reflect.Ptr && reflect.ValueOf(st.Enterprises).IsNil() {
        return nil, status.Error(codes.NotFound, "OnfEnterprise_Enterprises_Enterprise-not-found")
    }
    if res, ok := st.Enterprises.Enterprise[key]; ok {
        return res, nil
    }

    return nil, status.Error(codes.NotFound, "OnfEnterprise_Enterprises_Enterprise-not-found")
    }








func (c *GnmiClient) DeleteEnterprises_Enterprise_Item(ctx context.Context, target string, 
key string,
) (*gnmi.SetResponse, error) {
    gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
    defer cancel()


    
    
    path :=  []*gnmi.Path{
        {
            Elem: []*gnmi.PathElem{
            {
                    Name: "enterprises",
                },
            {
                    Name: "enterprise",
                    Key: map[string]string{
                        
                        "enterprise-id": string(key),
                        
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








func (c *GnmiClient) UpdateEnterprises_Enterprise_Item(ctx context.Context, target string,  data OnfEnterprise_Enterprises_Enterprise,
) (*gnmi.SetResponse, error) {
    gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
    defer cancel()


    
    

    
    path :=  []*gnmi.Path{
        {
            Elem: []*gnmi.PathElem{
            {
                    Name: "enterprises",
                },
            {
                    Name: "enterprise",
                    Key: map[string]string{
                        "enterprise-id": string(*data.EnterpriseId),
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





func (c *GnmiClient) GetConnectivityServices_ConnectivityService(ctx context.Context, target string, 
) (map[string]*OnfConnectivityService_ConnectivityServices_ConnectivityService, error) {
    gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
    defer cancel()

    

path :=  []*gnmi.Path{
    {
        Elem: []*gnmi.PathElem{
        {
                Name: "connectivity-services",
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

    if reflect.ValueOf(st.ConnectivityServices).Kind() == reflect.Ptr && reflect.ValueOf(st.ConnectivityServices).IsNil() {
        return nil, status.Error(codes.NotFound, "OnfConnectivityService_ConnectivityServices_ConnectivityService-not-found")
    }
    if reflect.ValueOf(st.ConnectivityServices.ConnectivityService).Kind() == reflect.Ptr && reflect.ValueOf(st.ConnectivityServices.ConnectivityService).IsNil() {
        return nil, status.Error(codes.NotFound, "OnfConnectivityService_ConnectivityServices_ConnectivityService-not-found")
    }

    return st.ConnectivityServices.ConnectivityService, nil
}





func (c *GnmiClient) DeleteConnectivityServices_ConnectivityService(ctx context.Context, target string, 
) (*gnmi.SetResponse, error) {
    gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
    defer cancel()

    

path :=  []*gnmi.Path{
    {
        Elem: []*gnmi.PathElem{
        {
                Name: "connectivity-services",
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





func (c *GnmiClient) UpdateConnectivityServices_ConnectivityService(ctx context.Context, target string,  list map[string]*OnfConnectivityService_ConnectivityServices_ConnectivityService,
) (*gnmi.SetResponse, error) {
    gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
    defer cancel()

    
    basePathElems :=  []*gnmi.PathElem{
    {
        Name: "connectivity-services",
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
                    "connectivity-service-id": string(*item.ConnectivityServiceId),
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





func (c *GnmiClient) GetEnterprises_Enterprise(ctx context.Context, target string, 
) (map[string]*OnfEnterprise_Enterprises_Enterprise, error) {
    gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
    defer cancel()

    

path :=  []*gnmi.Path{
    {
        Elem: []*gnmi.PathElem{
        {
                Name: "enterprises",
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

    if reflect.ValueOf(st.Enterprises).Kind() == reflect.Ptr && reflect.ValueOf(st.Enterprises).IsNil() {
        return nil, status.Error(codes.NotFound, "OnfEnterprise_Enterprises_Enterprise-not-found")
    }
    if reflect.ValueOf(st.Enterprises.Enterprise).Kind() == reflect.Ptr && reflect.ValueOf(st.Enterprises.Enterprise).IsNil() {
        return nil, status.Error(codes.NotFound, "OnfEnterprise_Enterprises_Enterprise-not-found")
    }

    return st.Enterprises.Enterprise, nil
}





func (c *GnmiClient) DeleteEnterprises_Enterprise(ctx context.Context, target string, 
) (*gnmi.SetResponse, error) {
    gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
    defer cancel()

    

path :=  []*gnmi.Path{
    {
        Elem: []*gnmi.PathElem{
        {
                Name: "enterprises",
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





func (c *GnmiClient) UpdateEnterprises_Enterprise(ctx context.Context, target string,  list map[string]*OnfEnterprise_Enterprises_Enterprise,
) (*gnmi.SetResponse, error) {
    gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
    defer cancel()

    
    basePathElems :=  []*gnmi.PathElem{
    {
        Name: "enterprises",
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
                    "enterprise-id": string(*item.EnterpriseId),
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



func (c *GnmiClient) GetConnectivityServices(ctx context.Context, target string, 
) (*OnfConnectivityService_ConnectivityServices, error) {
gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
defer cancel()



path :=  []*gnmi.Path{
    {
        Elem: []*gnmi.PathElem{
        {
                Name: "connectivity-services",
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

    if reflect.ValueOf(st.ConnectivityServices).Kind() == reflect.Ptr && reflect.ValueOf(st.ConnectivityServices).IsNil() {
    return nil, status.Error(codes.NotFound, "OnfConnectivityService_ConnectivityServices-not-found")
    }

    return st.ConnectivityServices, nil

}



func (c *GnmiClient) DeleteConnectivityServices(ctx context.Context, target string, 
) (*gnmi.SetResponse, error) {
gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
defer cancel()



path :=  []*gnmi.Path{
    {
        Elem: []*gnmi.PathElem{
        {
                Name: "connectivity-services",
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



func (c *GnmiClient) UpdateConnectivityServices(ctx context.Context, target string,  data OnfConnectivityService_ConnectivityServices,
) (*gnmi.SetResponse, error) {
gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
defer cancel()



path :=  []*gnmi.Path{
    {
        Elem: []*gnmi.PathElem{
        {
                Name: "connectivity-services",
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



func (c *GnmiClient) GetEnterprises(ctx context.Context, target string, 
) (*OnfEnterprise_Enterprises, error) {
gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
defer cancel()



path :=  []*gnmi.Path{
    {
        Elem: []*gnmi.PathElem{
        {
                Name: "enterprises",
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

    if reflect.ValueOf(st.Enterprises).Kind() == reflect.Ptr && reflect.ValueOf(st.Enterprises).IsNil() {
    return nil, status.Error(codes.NotFound, "OnfEnterprise_Enterprises-not-found")
    }

    return st.Enterprises, nil

}



func (c *GnmiClient) DeleteEnterprises(ctx context.Context, target string, 
) (*gnmi.SetResponse, error) {
gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
defer cancel()



path :=  []*gnmi.Path{
    {
        Elem: []*gnmi.PathElem{
        {
                Name: "enterprises",
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



func (c *GnmiClient) UpdateEnterprises(ctx context.Context, target string,  data OnfEnterprise_Enterprises,
) (*gnmi.SetResponse, error) {
gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
defer cancel()



path :=  []*gnmi.Path{
    {
        Elem: []*gnmi.PathElem{
        {
                Name: "enterprises",
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


