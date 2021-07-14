package hashy

import (
	"fmt"
	"testing"
)

var config = Options{
	// Input:     "./_/short.csv",
	Input:     "./_/long.csv",
	// Input:     "./_/1m.csv",
	KeyColumns: []int{ 1,2,9 },
	SkipHeader: true,
	// Delimiter: ',',
}

func TestModuleName(t *testing.T) {
	hash, err := File(config)
	if err != nil {
		fmt.Println(err)
	}
	
	fmt.Println("done:", len(hash))
	
	for key, val := range hash {
		fmt.Println(">", key)
		for _, f := range val {
			fmt.Println("\t",f)
		}
	}
}

// -benchmem
func BenchmarkModuleName(b *testing.B) {
	for i := 0; i < b.N; i++ {
		File(config)
	}
}
