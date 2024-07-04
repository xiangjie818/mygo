package types

type CubeFSCluster struct {
	LockIndex   int64  `json:"LockIndex"`
	Key         string `json:"Key"`
	Flags       int64  `json:"Flags"`
	Value       string `json:"Value"`
	CreateIndex int64  `json:"CreateIndex"`
	ModifyIndex int64  `json:"ModifyIndex"`
}

type CubeFSMasterAddr struct {
	MasterAddr []string `json:"masterAddr"`
}

type ServiceRegistration struct {
	ID      string       `json:"ID"`
	Name    string       `json:"Name"`
	Tags    []string     `json:"Tags,omitempty"`
	Address string       `json:"Address"`
	Port    int64        `json:"Port"`
	Check   ServiceCheck `json:"Check,omitempty"`
}

type ServiceCheck struct {
	HTTP                           string `json:"HTTP,omitempty"`
	TCP                            string `json:"TCP,omitempty"`
	Interval                       string `json:"Interval"`
	Timeout                        string `json:"Timeout"`
	DeregisterCriticalServiceAfter string `json:"DeregisterCriticalServiceAfter"`
}

const (
	RegisterServicePath   = "/v1/agent/service/register"
	DeregisterServicePath = "/v1/agent/service/deregister"
)
