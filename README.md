# go-loadavg
A cross-platform golang library for retrieving load average.

## How to use

```go
package main

import (
	"github.com/mikoim/go-loadavg"
	"log"
)

func main() {
	loadavg, err := loadavg.Parse()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("load averages: %.2f %.2f %.2f", loadavg.LoadAverage1, loadavg.LoadAverage5, loadavg.LoadAverage10)
}

```

## Supported platform

|    GOOS   |     Source     |   Support   |
|:---------:|:--------------:|:-----------:|
|   darwin  |  getloadavg()  | coming soon |
| dragonfly |  getloadavg()  | coming soon |
|  freebsd  |  /proc/loadavg | coming soon |
|   linux   |  /proc/loadavg |      O      |
|   netbsd  |  getloadavg()  | coming soon |
|  openbsd  |  getloadavg()  | coming soon |
|   plan9   | /dev/sysstat ? |      ?      |
|  solaris  |        ?       |      ?      |
|  windows  |        ?       |      ?      |

## License
[![WTFPL](http://www.wtfpl.net/wp-content/uploads/2012/12/wtfpl-badge-4.png)](http://www.wtfpl.net/)
