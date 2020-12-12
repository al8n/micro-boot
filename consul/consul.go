package consul

var (
	defaultConsulFlagsPrefix = "consul"
)

type Config struct {
	Client ClientConfig `json:"client" yaml:"client"`
}
