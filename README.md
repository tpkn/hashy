# Hashy
Hash tables from csv data with ease

## Usage

Convert csv to a hash map:

```go
import (
   "fmt"
   "github.com/tpkn/hashy"
)

var options = hashy.Options{
   Input:             "./test/file.csv",
   KeyColumns:        []int{ 0, 12, 45 },
   IncludeKeysValues: false,
   Delimiter:         ',',
   SkipHeader:        true,
   LazyQuotes:        false,
}

var hash, err = hashy.File(options)
if err != nil {
   panic(err)
}

for key, val := range hash {
   fmt.Println(key)
   for _, f := range val {
      fmt.Println("\t", fmt.Sprintf("%q", f))
   }
}
```



To make a flat hash map (one key - one value):

```go
var hash, err = hashy.FileFlat(options)
```





