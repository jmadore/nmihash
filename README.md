# nmi-hash

> Package nmihash is used as a source code example for a Go (golang) implementation of the examples found at:
// https://support.nmi.com/hc/en-gb/articles/115001683343-Implementation-Generating-a-Hash-Code

[![Go Report Card](https://goreportcard.com/badge/github.com/jmadore/nmihash)](https://goreportcard.com/report/github.com/jmadore/nmihash) [![Documentation](https://godoc.org/github.com/jmadore/nmihash?status.svg)](http://godoc.org/github.com/jmadore/nmihash)

---
## Prerequisites

- [Go](https://go.dev/) 1.14 or newer
- [Go Modules](https://github.com/golang/go/wiki/Modules)

## Installation

## Using `go get`

You can run the following command to install `nmihash`:

```
go get -u github.com/jmadore/nmihash
```
The code is self-contained and doesn't have any external dependencies.


## Build

> Build
```shell
> go build
```

## Tests

```shell
> go test -v
```

## Example


```go
package main

import (
	"fmt"

	"github.com/jmadore/nmihash"
)

func main() {

	hashKey := "trVxrnoz22bvwvnV"
	terminalID := "99999999"
	refNum := "0000000765"
	amount := "1.23"

	hash, err := nmihash.CalculateHash(hashKey, terminalID, refNum, amount)
	if err != nil {
		fmt.Printf("Hash failed with error %v", err.Error())
		return
	}

	fmt.Println("Base64 Encoded Hash =", hash)
	return
}

```


---

## Built With

* [Go](https://go.dev/) - Language used
* [VS CODE](https://code.visualstudio.com/) - IDE/Editor


## Contributing

Please read [CONTRIBUTING.md](https://gist.github.com/jmadore/b24679402957c63ec426) for details on our code of conduct, and the process for submitting pull requests to us.

## Versioning

We use [SemVer](http://semver.org/) for versioning. For the versions available, see the [tags on this repository](https://github.com/jmadore/nmihash/tags). 

## Authors

* **Jay Madore** - *Initial work* - [jmadore](https://github.com/jmadore)


## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details

## Acknowledgments

* [NMI Examples](https://support.nmi.com/hc/en-gb/articles/115001683343-Implementation-Generating-a-Hash-Code)
* [README.md template](https://gist.github.com/PurpleBooth/)
* [go-james](https://github.com/pieterclaerhout/go-james/)

