# gocovercounter

in development ...

gocovercounter lets you to generate coverage-annotated source.
code is Roughly mirror of golang/go/src/cmd/cover/cover.go.

## Quick Start

use package as CLI.

```bash
$ go get -u github.com/po3rin/gocovercounter/cmd/coco
$ coco ./testdata/testfile.go
```

coco outputs following. this code is coverage-annotated source.

```go
//line ./testdata/testfile.go:1
package testfile

import (
	"bufio"
	"fmt"
	"os"
)

func Testfile() {GoCover.Count[0] = 1;
	fmt.Println("text: ")
	stdin := bufio.NewScanner(os.Stdin)
	stdin.Scan()
	text := stdin.Text()

	if text != "hello" {GoCover.Count[2] = 1;
		return
	}
	GoCover.Count[1] = 1;fmt.Println("hello")
}

var GoCover = struct {
	Count     [3]uint32
	Pos       [3 * 3]uint32
	NumStmt   [3]uint16
} {
	Pos: [3 * 3]uint32{
		9, 15, 0x150011, // [0]
		18, 18, 0x160002, // [1]
		15, 17, 0x30015, // [2]
	},
	NumStmt: [3]uint16{
		5, // 0
		1, // 1
		1, // 2
	},
}

```

If you use ```-o``` flag, coco creates output file.

```
$ coco -o out.txt ./testdata/testfile.go
```
