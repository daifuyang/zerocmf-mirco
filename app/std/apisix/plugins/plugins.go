package plugins

import (
	"app/std/apisix"
	"app/std/apisix/plugins/authentication"
)

func RoutePlugin(host string) apisix.RoutePlugins {
	return apisix.RoutePlugins{
		ForwardAuth: &authentication.ForwardAuth{
			Meta: authentication.Meta{
				Disable: false,
			},
			Uri:             "http://" + host + ":9080/api/authn/validation",
			RequestMethod:   "POST",
			RequestHeaders:  []string{"Authorization"},
			UpstreamHeaders: []string{"X-User-ID"},
			ClientHeaders:   []string{},
		},
	}
}
