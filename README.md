defaults
========
initialize structs with default values

how to get
------------

    go get github.com/i/defaults

how to do
---------

```go
package main

import (
  "fmt"

  "github.com/i/defaults"
)

type Person struct {
  Name            string `default:"Fred Durst"`
  Age             int    `default:"12"`
  ConvictedFelon  bool   `default:"true"`
}

func main() {
  var p person
  if err := defaults.Set(&p); err != nil {
    panic(err)
  }

  fmt.Println(p.Name)           // prints Fred Durst
  fmt.Println(p.Age)            // prints 12
  fmt.Println(p.ConvictedFelon) // prints true
}
```
