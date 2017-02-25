# mboxparser

[![Build Status](https://travis-ci.org/zeny-io/mboxparser.svg?branch=master)](https://travis-ci.org/zeny-io/mboxparser)
[![Coverage Status](https://coveralls.io/repos/github/zeny-io/mboxparser/badge.svg?branch=master)](https://coveralls.io/github/zeny-io/mboxparser?branch=master)

mbox file parser for golang.

# Usage

```golang
package main

import (
	"fmt"
	"github.com/zeny-io/mboxparser"
	"io"
	"os"
)

func main() {
	for n, arg := range os.Args {
		if n == 0 {
			continue
		}

		if mbox, err := mboxparser.ReadFile(arg); err == nil {
			for _, mail := range mbox.Messages {
				for k, vs := range mail.Header {
					for _, v := range vs {
						fmt.Printf("%s: %s\n", k, v)
					}
				}
				for _, body := range mail.Bodies {
					fmt.Println("====================================================")
					for k, vs := range body.Header {
						for _, v := range vs {
							fmt.Printf("%s: %s\n", k, v)
						}
					}
					fmt.Println("")
					io.Copy(os.Stdout, body.Content)
					fmt.Println("")
				}
				fmt.Println("====================================================\n\n")
			}
		} else {
			fmt.Printf("%s\n", err.Error())
		}
	}
}
```

# License

MIT

# Author

[@rosylilly: Sho Kusano](https://github.com/rosylilly) / <sho-kusano@zeny.io>
