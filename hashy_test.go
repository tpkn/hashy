package hashy

import (
	"fmt"
	"testing"
)

var config = Options{
	// Input:             "./_/short.csv",
	// Input:             "./_/long.csv",
	// Input:             "./_/no_valid.csv",
	Input:             "./_/1gb.csv",
	KeyColumns:        []int{1},
	SkipHeader:        false,
	Delimiter:         ',',
	IncludeKeysValues: false,
}

func TestFile(t *testing.T) {
	hash, err := File(config)
	if err != nil {
		fmt.Println("(!) csv file error:", err)
	}
	
	fmt.Println("total:", len(hash))
	
	for key, val := range hash {
		fmt.Println(key)
		for _, f := range val {
			fmt.Println("\t", fmt.Sprintf("%q", f))
		}
		
		break
	}
}

func TestFileFlat(t *testing.T) {
	hash, err := FileFlat(config)
	if err != nil {
		fmt.Println("(!) csv file error:", err)
	}
	
	fmt.Println("total:", len(hash))
	
	for key, val := range hash {
		fmt.Println(key, ":", val)
		
		break
	}
}

// -benchmem
func BenchmarkModuleName(b *testing.B) {
	for i := 0; i < b.N; i++ {
		File(config)
	}
}
