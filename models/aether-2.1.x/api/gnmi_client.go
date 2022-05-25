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









func (c *GnmiClient) GetSite_Item(ctx context.Context, target string, 
key string,
) (*OnfSite_Site, error) {
    gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
    defer cancel()


    
    
    path :=  []*gnmi.Path{
        {
            Elem: []*gnmi.PathElem{
            {
                    Name: "site",
                    Key: map[string]string{
                        
                        "site-id": string(key),
                        
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

    if res, ok := st.Site[key]; ok {
        return res, nil
    }

    return nil, status.Error(codes.NotFound, "OnfSite_Site-not-found")
    }








func (c *GnmiClient) DeleteSite_Item(ctx context.Context, target string, 
key string,
) (*gnmi.SetResponse, error) {
    gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
    defer cancel()


    
    
    path :=  []*gnmi.Path{
        {
            Elem: []*gnmi.PathElem{
            {
                    Name: "site",
                    Key: map[string]string{
                        
                        "site-id": string(key),
                        
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








func (c *GnmiClient) UpdateSite_Item(ctx context.Context, target string,  data OnfSite_Site,
) (*gnmi.SetResponse, error) {
    gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
    defer cancel()


    
    

    
    path :=  []*gnmi.Path{
        {
            Elem: []*gnmi.PathElem{
            {
                    Name: "site",
                    Key: map[string]string{
                        "site-id": string(*data.SiteId),
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








func (c *GnmiClient) GetTemplate_Item(ctx context.Context, target string, 
key string,
) (*OnfTemplate_Template, error) {
    gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
    defer cancel()


    
    
    path :=  []*gnmi.Path{
        {
            Elem: []*gnmi.PathElem{
            {
                    Name: "template",
                    Key: map[string]string{
                        
                        "template-id": string(key),
                        
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

    if res, ok := st.Template[key]; ok {
        return res, nil
    }

    return nil, status.Error(codes.NotFound, "OnfTemplate_Template-not-found")
    }








func (c *GnmiClient) DeleteTemplate_Item(ctx context.Context, target string, 
key string,
) (*gnmi.SetResponse, error) {
    gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
    defer cancel()


    
    
    path :=  []*gnmi.Path{
        {
            Elem: []*gnmi.PathElem{
            {
                    Name: "template",
                    Key: map[string]string{
                        
                        "template-id": string(key),
                        
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








func (c *GnmiClient) UpdateTemplate_Item(ctx context.Context, target string,  data OnfTemplate_Template,
) (*gnmi.SetResponse, error) {
    gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
    defer cancel()


    
    

    
    path :=  []*gnmi.Path{
        {
            Elem: []*gnmi.PathElem{
            {
                    Name: "template",
                    Key: map[string]string{
                        "template-id": string(*data.TemplateId),
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








func (c *GnmiClient) GetTrafficClass_Item(ctx context.Context, target string, 
key string,
) (*OnfTrafficClass_TrafficClass, error) {
    gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
    defer cancel()


    
    
    path :=  []*gnmi.Path{
        {
            Elem: []*gnmi.PathElem{
            {
                    Name: "traffic-class",
                    Key: map[string]string{
                        
                        "traffic-class-id": string(key),
                        
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

    if res, ok := st.TrafficClass[key]; ok {
        return res, nil
    }

    return nil, status.Error(codes.NotFound, "OnfTrafficClass_TrafficClass-not-found")
    }








func (c *GnmiClient) DeleteTrafficClass_Item(ctx context.Context, target string, 
key string,
) (*gnmi.SetResponse, error) {
    gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
    defer cancel()


    
    
    path :=  []*gnmi.Path{
        {
            Elem: []*gnmi.PathElem{
            {
                    Name: "traffic-class",
                    Key: map[string]string{
                        
                        "traffic-class-id": string(key),
                        
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








func (c *GnmiClient) UpdateTrafficClass_Item(ctx context.Context, target string,  data OnfTrafficClass_TrafficClass,
) (*gnmi.SetResponse, error) {
    gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
    defer cancel()


    
    

    
    path :=  []*gnmi.Path{
        {
            Elem: []*gnmi.PathElem{
            {
                    Name: "traffic-class",
                    Key: map[string]string{
                        "traffic-class-id": string(*data.TrafficClassId),
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








func (c *GnmiClient) GetApplication_Item(ctx context.Context, target string, 
key string,
) (*OnfApplication_Application, error) {
    gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
    defer cancel()


    
    
    path :=  []*gnmi.Path{
        {
            Elem: []*gnmi.PathElem{
            {
                    Name: "application",
                    Key: map[string]string{
                        
                        "application-id": string(key),
                        
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

    if res, ok := st.Application[key]; ok {
        return res, nil
    }

    return nil, status.Error(codes.NotFound, "OnfApplication_Application-not-found")
    }








func (c *GnmiClient) DeleteApplication_Item(ctx context.Context, target string, 
key string,
) (*gnmi.SetResponse, error) {
    gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
    defer cancel()


    
    
    path :=  []*gnmi.Path{
        {
            Elem: []*gnmi.PathElem{
            {
                    Name: "application",
                    Key: map[string]string{
                        
                        "application-id": string(key),
                        
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








func (c *GnmiClient) UpdateApplication_Item(ctx context.Context, target string,  data OnfApplication_Application,
) (*gnmi.SetResponse, error) {
    gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
    defer cancel()


    
    

    
    path :=  []*gnmi.Path{
        {
            Elem: []*gnmi.PathElem{
            {
                    Name: "application",
                    Key: map[string]string{
                        "application-id": string(*data.ApplicationId),
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





func (c *GnmiClient) GetSite(ctx context.Context, target string, 
) (map[string]*OnfSite_Site, error) {
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





func (c *GnmiClient) UpdateSite(ctx context.Context, target string,  list map[string]*OnfSite_Site,
) (*gnmi.SetResponse, error) {
    gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
    defer cancel()

    
    basePathElems :=  []*gnmi.PathElem{
    }
    req := &gnmi.SetRequest{
        Update: []*gnmi.Update{},
    }
    for _, item := range list {

        path := &gnmi.Path{
            Elem: append(basePathElems, &gnmi.PathElem{
                Name: "list2a",
                Key: map[string]string{
                    "site-id": string(*item.SiteId),
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





func (c *GnmiClient) GetTemplate(ctx context.Context, target string, 
) (map[string]*OnfTemplate_Template, error) {
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





func (c *GnmiClient) UpdateTemplate(ctx context.Context, target string,  list map[string]*OnfTemplate_Template,
) (*gnmi.SetResponse, error) {
    gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
    defer cancel()

    
    basePathElems :=  []*gnmi.PathElem{
    }
    req := &gnmi.SetRequest{
        Update: []*gnmi.Update{},
    }
    for _, item := range list {

        path := &gnmi.Path{
            Elem: append(basePathElems, &gnmi.PathElem{
                Name: "list2a",
                Key: map[string]string{
                    "template-id": string(*item.TemplateId),
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





func (c *GnmiClient) GetTrafficClass(ctx context.Context, target string, 
) (map[string]*OnfTrafficClass_TrafficClass, error) {
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





func (c *GnmiClient) UpdateTrafficClass(ctx context.Context, target string,  list map[string]*OnfTrafficClass_TrafficClass,
) (*gnmi.SetResponse, error) {
    gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
    defer cancel()

    
    basePathElems :=  []*gnmi.PathElem{
    }
    req := &gnmi.SetRequest{
        Update: []*gnmi.Update{},
    }
    for _, item := range list {

        path := &gnmi.Path{
            Elem: append(basePathElems, &gnmi.PathElem{
                Name: "list2a",
                Key: map[string]string{
                    "traffic-class-id": string(*item.TrafficClassId),
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





func (c *GnmiClient) GetApplication(ctx context.Context, target string, 
) (map[string]*OnfApplication_Application, error) {
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





func (c *GnmiClient) UpdateApplication(ctx context.Context, target string,  list map[string]*OnfApplication_Application,
) (*gnmi.SetResponse, error) {
    gnmiCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
    defer cancel()

    
    basePathElems :=  []*gnmi.PathElem{
    }
    req := &gnmi.SetRequest{
        Update: []*gnmi.Update{},
    }
    for _, item := range list {

        path := &gnmi.Path{
            Elem: append(basePathElems, &gnmi.PathElem{
                Name: "list2a",
                Key: map[string]string{
                    "application-id": string(*item.ApplicationId),
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


