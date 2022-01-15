# github.com/mitchallen/coin

## Usage

### Initialize your module

```
$ go mod init example.com/my-demo
```

### Get the module

Note that you need to include the **v** in the version tag.

```
$ go get github.com/mitchallen/coin@v0.2.0
```

### coin.Flip()

```go
package main

import (
	"fmt"
	"github.com/mitchallen/coin"
)

func main() {
	fmt.Println(coin.Flip())
}
```

### coin.Weighted(w)

```go
package main

import (
	"fmt"
	"github.com/mitchallen/coin"
)

func main() {
	// Will return true ~75% of the time
	fmt.Println(coin.Weighted(0.75))
}
```

## Init

How this repo was intialized as a **go** (golang) module

```
$ go mod init github.com/mitchallen/coin
```

## Testing

```
$ go test
```

```
$ cd (package)
$ go test
```

## Tagging

Substitute the light weight tag value with the latest tag to publish

```
$ git tag v0.1.0
$ git push origin --tags
```

## Publishing instructions

* https://go.dev/doc/modules/publishing

## Package

* https://pkg.go.dev/search?q=mitchallen/
