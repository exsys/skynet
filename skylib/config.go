//Copyright (c) 2011 Brian Ketelsen

//Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

//The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

//THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

package skylib

import (
	"flag"
	"log"
  "time"
)

type BindAddr struct {
	IPAddress string
	Port      int
}

type ServiceConfig struct {
	Log         *log.Logger `json:"-"`
	Name        string
	Version     string
	Region      string
	ServiceAddr *BindAddr
	AdminAddr   *BindAddr
	DoozerConfig *DoozerConfig `json:"-"`
}

type ClientConfig struct {
	Log         *log.Logger `json:"-"`
	DoozerConfig *DoozerConfig `json:"-"`
  ConnectionPoolSize int
  IdleTimeout time.Duration
}

func GetServiceConfigFromFlags() *ServiceConfig {
	var (
		bindPort       *int    = flag.Int("port", 9999, "tcp port to listen")
		bindAddr       *string = flag.String("address", "127.0.0.1", "address to bind")
		region         *string = flag.String("region", "unknown", "region service is located in")
		doozer         *string = flag.String("doozer", "127.0.0.1:8046", "initial doozer instance to connect to")
		doozerBoot     *string = flag.String("doozerboot", "127.0.0.1:8046", "initial doozer instance to connect to")
		doozerDiscover *bool   = flag.Bool("autodiscover", true, "auto discover new doozer instances")
	)

	flag.Parse()

	return &ServiceConfig{
		Region: *region,
		ServiceAddr: &BindAddr{
			IPAddress: *bindAddr,
			Port:      *bindPort,
		},
    DoozerConfig: &DoozerConfig {
      Uri: *doozer,
      BootUri: *doozerBoot,
      AutoDiscover: *doozerDiscover,
    },
	}
}
