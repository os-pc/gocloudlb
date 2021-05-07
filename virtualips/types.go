package virtualips

type CreateOpts struct {
	// The ID for the IP address
	ID uint64 `json:"id,omitempty"`

	// The IP address type
	Type string `json:"type,omitempty"`

	// The IP version
	IPVersion string `json:"ipVersion,omitempty"`
}

type VirtualIp struct {
	// The ID for the IP address
	ID uint64 `json:"id"`

	// The IP address
	Address string `json:"address"`

	// The IP address type
	Type string `json:"type"`

	// The IP version
	IPVersion string `json:"ipVersion"`
}
