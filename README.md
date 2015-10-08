# cc-serialize
textutils.serialize from ComputerCraft implemented in Go. Can be used on string indexed maps, slices, structs, int, int64, float64 and strings.

## Usage

```
package main

import "github.com/1lann/cc-serialize"

type Test struct {
  Name string
  Age int
  SomeNumbers []int
}

func main() {
  test := Test{
    Name: "John Smith",
    Age: 28,
    SomeNumbers: []int{1, 2, 3, 4, 5},
  }

  ccserialize.Serialize(test)
  // Returns
  // {["name"] = "John Smith", ["age"] = 28, ["somenumbers"] = {[1] = 1, [2] = 2, [3] = 3, [4] = 4, [5] = 5, }, }
}
```
