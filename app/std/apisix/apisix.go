package apisix

import (
	"app/std/util"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
)

type Apisix struct {
	ApiKey   string   `json:"apiKey,env=APISIX_APIKEY"`
	Host     string   `json:"host,env=APISIX_HOST"`
	Name     string   `json:"name,optional"`
	Upstream Upstream `json:"upstream,optional"`
}

func (a Apisix) Register(routes []Route) (err error) {
	name := a.Name
	apikey := a.ApiKey
	host := a.Host

	a.Upstream.Name = name
	err = a.Upstream.Register(host, apikey)
	if err != nil {
		return
	}

	service := Service{
		Name:       name,
		UpstreamId: name,
	}
	err = service.Register(host, apikey)
	if err != nil {
		return err
	}

	//err = a.jas()
	//if err != nil {
	//	return err
	//}

	err = new(Route).Register(host, apikey, routes)
	if err != nil {
		return err
	}
	return
}

func (a Apisix) jas() (err error) {
	apiKey := a.ApiKey
	host := a.Host
	var params struct {
		Uri     string `json:"uri"`
		Plugins struct {
			PublicApi struct{} `json:"public-api"`
		} `json:"plugins"`
	}

	params.Uri = "/apisix/plugin/jwt/sign"
	var body bytes.Buffer
	err = json.NewEncoder(&body).Encode(&params)
	if err != nil {
		return
	}

	header := map[string]string{"X-API-KEY": apiKey}
	code, response := util.Request("PUT", "http://"+host+":9180/apisix/admin/routes/jas", &body, header)
	if code == 201 || code == 200 {
		fmt.Println("register routes jas:", string(response))
		return
	}
	err = errors.New("jas" + ":" + string(response))
	return
}
