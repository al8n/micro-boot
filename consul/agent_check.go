package consul

// AgentCheck represents a check known to the agent
type AgentCheck struct {
	Node        string
	CheckID     string
	Name        string
	Status      string
	Notes       string
	Output      string
	ServiceID   string
	ServiceName string
	Type        string
	Definition  HealthCheckDefinition
	Namespace   string `json:",omitempty"`
}
