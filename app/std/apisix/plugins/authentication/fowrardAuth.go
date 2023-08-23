package authentication

type ForwardAuth struct {
	Meta            Meta     `json:"_meta,omitempty"`
	Uri             string   `json:"uri"`
	RequestMethod   string   `json:"request_method,omitempty"`
	RequestHeaders  []string `json:"request_headers,omitempty"`
	UpstreamHeaders []string `json:"upstream_headers,omitempty"`
	ClientHeaders   []string `json:"client_headers,omitempty"`
}
