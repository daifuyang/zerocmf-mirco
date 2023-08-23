package authentication

type KeyAuth struct {
	Meta   Meta   `json:"_meta,omitempty"`
	Key    string `json:"key,omitempty"`
	Header string `json:"header,omitempty"`
}
