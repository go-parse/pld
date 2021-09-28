# PLD

Parser of the list of HTTP proxy servers from the site - https://www.proxy-list.download


## Installation

Simple install the package to your [$GOPATH](https://github.com/golang/go/wiki/GOPATH "GOPATH") with the [go tool](https://golang.org/cmd/go/ "go command") from shell:
```bash
$ go get -u github.com/go-parse/pld
```
Make sure [Git is installed](https://git-scm.com/downloads) on your machine and in your system's `PATH`.


## Usage

To get a list of proxies execute the function Get() - this function is return the slice of the array contains items structure:
```go
type Proxy struct {
	Address   string
	Port      string
	Anonymity string
	Country   string
}
```


## Example

```go
import (
	"fmt"

	"github.com/go-parse/pld"
)

// ...

proxies := pld.Get()

for _, proxy := range proxies {

	fmt.Println("Address:", proxy.Address)
	fmt.Println("Port:", proxy.Port)
	fmt.Println("Anonymity:", proxy.Anonymity)
	fmt.Println("Country:", proxy.Country)
	fmt.Println("---------------------------")
}

```


## License

Copyright 2021 Vasiliy Vdovin

Licensed under the Apache License, Version 2.0 (the "License")\
you may not use this file except in compliance with the License.\
You may obtain a copy of the License at
http://www.apache.org/licenses/LICENSE-2.0\

Unless required by applicable law or agreed to in writing, software\
distributed under the License is distributed on an "AS IS" BASIS,\
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.\
See the License for the specific language governing permissions and\
limitations under the License.

Please read the [Licensing FAQ](https://www.apache.org/foundation/license-faq.html) if you have further questions regarding the license.
