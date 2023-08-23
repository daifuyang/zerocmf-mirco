package apisix

import (
	"app/std/apisix/plugins/authentication"
	"app/std/util"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
)

type ProxyRewrite struct {
	RegexURI []string `json:"regex_uri,omitempty"`
}

type RoutePlugins struct {
	JWTAuth      *authentication.JwtAuth     `json:"jwt-auth,omitempty"`
	KeyAuth      *authentication.KeyAuth     `json:"key-auth,omitempty"`
	ForwardAuth  *authentication.ForwardAuth `json:"forward-auth,omitempty"`
	ProxyRewrite *ProxyRewrite               `json:"proxy-rewrite,omitempty"`
}

type Route struct {
	URI       string       `json:"uri"`
	Name      string       `json:"name"`
	Methods   []string     `json:"methods,omitempty"`
	ServiceID string       `json:"service_id"`
	Plugins   RoutePlugins `json:"plugins"`
	Status    int          `json:"status"`
}

func (r *Route) Register(host string, apiKey string, routes []Route) (err error) {
	for _, v := range routes {
		var body bytes.Buffer
		err = json.NewEncoder(&body).Encode(&v)
		if err != nil {
			return
		}

		header := map[string]string{"X-API-KEY": apiKey}
		code, response := util.Request("PUT", "http://"+host+":9180/apisix/admin/routes/"+v.Name, &body, header)
		if !(code == 201 || code == 200) {
			err = errors.New(v.Name + ":" + string(response))
			return
		}
		fmt.Println("register routes "+v.Name+":", string(response))
	}
	return
}
