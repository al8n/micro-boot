package main

import (
	"errors"
	bootflag "github.com/ALiuGuanyan/uni/micro-boot/flag"
	"strconv"
)

// CustomConfigNeedToCheckValidity is a custom config structure do not need to check the validity
type CustomConfigNeedToCheckValidity struct {
	Key string `json:"key" yaml:"key"`
	Port string `json:"port" yaml:"port"`
}

// BindFlags binds the command line flags with CustomConfigNeedToCheckValidity
func (n *CustomConfigNeedToCheckValidity) BindFlags(fs *bootflag.FlagSet)  {
	fs.StringVar(&n.Key, "key", "value", "key value pair")
	fs.StringVarP(&n.Port, "port", "p", "8080", "ipv4 port")
}

// Parse is used to check the validity and avoid weird values 
func (n *CustomConfigNeedToCheckValidity) Parse() (err error) {
	p, err := strconv.ParseUint(n.Port, 10, 64)
	if err != nil {
		return err
	}
	if p > 65535 {
		return errors.New("not a valid port")
	}
	return nil
}

func main()  {}



