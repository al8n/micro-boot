package main

import bootflag "github.com/ALiuGuanyan/micro-boot/flag"

// CustomConfigNeedNotToCheckValidity is a custom config structure do not need to check the validity
type CustomConfigNeedNotToCheckValidity struct {
	Key string `json:"key" yaml:"key"`
	Description string `json:"description" yaml:"description"`
}

// BindFlags binds the command line flags with CustomConfigNeedNotToCheckValidity
func (n *CustomConfigNeedNotToCheckValidity) BindFlags(fs *bootflag.FlagSet)  {
	fs.StringVar(&n.Key, "key", "value", "key value pair")
	fs.StringVarP(&n.Description, "description", "d", "some description", "the description")
}

// Parse leave it returns nil, because CustomConfigNeedNotToCheckValidity does not need to check the validity
func (n *CustomConfigNeedNotToCheckValidity) Parse() (err error) {
	return nil
}

func main()  {}


