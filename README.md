<h1>speedy</h1>
<p>
  <img alt="GitHub go.mod Go version" src="https://img.shields.io/github/go-mod/go-version/tsirysndr/speedy">
  <a href="https://pkg.go.dev/github.com/tsirysndr/speedy" target="_blank">
    <img alt="GoDoc" src="https://godoc.org/github.com/tsirysndr/speedy?status.svg">
  </a>
  <a href="https://github.com/tsirysndr/speedy/tree/master/_examples" target="_blank">
    <img alt="view examples" src="https://img.shields.io/badge/learn%20by-examples-0077b3.svg">
  </a>
  <a href="https://travis-ci.com/tsirysndr/speedy" target="_blank">
    <img alt="build status" src="https://img.shields.io/travis/tsirysndr/speedy/master.svg">
  </a>
  <a href="https://codecov.io/github/tsirysndr/speedy?branch=master" target="_blank">
    <img alt="Code Coverage" src="https://img.shields.io/codecov/c/github/tsirysndr/speedy/master.svg" />
  </a>
  <a href="https://github.com/tsirysndr/speedy/blob/master/LICENSE" target="_blank">
    <img alt="License: BSD" src="https://img.shields.io/badge/License-BSD-yellow.svg" />
  </a>
  <a href="https://twitter.com/tsiry\_sndr" target="_blank">
    <img alt="Twitter: tsiry\_sndr" src="https://img.shields.io/twitter/follow/tsiry_sndr.svg?style=social" />
  </a>
</p>

Go library to test internet speed using [fast.com](https://fast.com) and [speedtest.net](https://speedtest.net)

## 🚚 Install

```sh
go get -u github.com/tsirysndr/speedy
```

## 🚀 Usage

```go
import (
	"fmt"

	"github.com/tsirysndr/speedy"
)

func main() {
	r, _ := speedy.Start(speedy.OOKLA) // Or speedy.NETFLIX
	fmt.Printf("Download: %f Mbps | Upload: %f Mbps", r.Download, r.Upload)
}
```

## Run tests

```sh
go test
```

## Author

👤 **Tsiry Sandratraina**

* Twitter: [@tsiry\_sndr](https://twitter.com/tsiry\_sndr)
* Github: [@tsirysndr](https://github.com/tsirysndr)

## Show your support

Give a ⭐️ if this project helped you!

## 📝 License

Copyright © 2021 [Tsiry Sandratraina](https://github.com/tsirysndr).<br />
This project is [BSD](https://github.com/tsirysndr/speedy/blob/master/LICENSE) licensed.
