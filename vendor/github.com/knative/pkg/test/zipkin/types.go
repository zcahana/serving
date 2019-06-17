package zipkin

type Span struct {
	TraceId  string `json:"trace_id,omitempty"`
	ParentId string `json:"parent_id,omitempty"`
	Id       string `json:"id,omitempty"`

	Kind string `json:"kind,omitempty"`
	Name string `json:"name,omitempty"`

	Timestamp uint64 `json:"timestamp,omitempty"`
	Duration  uint64 `json:"duration,omitempty"`

	LocalEndpoint  *Endpoint `json:"local_endpoint,omitempty"`
	RemoteEndpoint *Endpoint `json:"remote_endpoint,omitempty"`

	Annotations []*Annotation     `json:"annotations,omitempty"`
	Tags        map[string]string `json:"tags,omitempty"`

	Debug  bool `json:"debug,omitempty"`
	Shared bool `json:"shared,omitempty"`
}

type Endpoint struct {
	ServiceName string `json:"service_name,omitempty"`
	Ipv4        string `json:"ipv4,omitempty"`
	Ipv6        string `json:"ipv6,omitempty"`
	Port        int32  `json:"port,omitempty"`
}

type Annotation struct {
	Timestamp uint64 `json:"timestamp,omitempty"`
	Value     string `json:"value,omitempty"`
}
