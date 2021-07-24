// hashy, https://tpkn.me
package hashy

import (
	"strings"
)

// Options is a basic structure required for the Hashy to work
type Options struct {
	Input             string // Input data, a raw csv or a path to csv file
	KeyColumns        []int  // Columns indexes that should be used as a hash keys (order matters: {1,2} and {2,1} are two different keys)
	Delimiter         rune   // Csv columns delimiter
	IncludeKeysValues bool   // Include key columns values in hash content
	SkipHeader        bool   // ...
	LazyQuotes        bool   // ...
}

// File makes hash from csv file
func File(o Options) (map[string][][]string, error) {
	var hash = make(map[string][][]string)
	
	err := csvFileReader(o, func(line []string) {
		key := makeHashKey(line, o.KeyColumns)
		
		if !o.IncludeKeysValues {
			removeKeyColumns(&line, o.KeyColumns)
		}
		
		hash[key] = append(hash[key], line)
	})
	
	return hash, err
}

// FileFlat makes hash from csv file with a single flat value
func FileFlat(o Options) (map[string]string, error) {
	var hash = make(map[string]string)
	
	err := csvFileReader(o, func(line []string) {
		key := makeHashKey(line, o.KeyColumns)
		
		if !o.IncludeKeysValues {
			removeKeyColumns(&line, o.KeyColumns)
		}
		
		// "There can be only one" value for each key
		if _, found_one := hash[key]; !found_one {
			hash[key] = strings.Join(line, string(o.Delimiter))
		}
	})
	
	return hash, err
}
