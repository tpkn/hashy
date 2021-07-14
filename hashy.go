// hashy, https://tpkn.me
package hashy

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

// Options is a basic structure required for the Hashy to work
type Options struct {
	Input             string // Input data, a raw csv or a path to csv file
	KeyColumns        []int  // Columns indexes that should be used as a hash keys (order matters: {1,2} and {2,1} are two different keys)
	SkipHeader        bool   // Yep
	IncludeKeysValues bool   // Include key columns values in hash content
	Delimiter         rune   // Csv columns delimiter (optional)
}

// File makes hash from csv file
func File(options Options) (map[string][][]string, error) {
	if len(options.KeyColumns) == 0 {
		return nil, errors.New("Key column is not set")
	}
	// Default delimiter value is ','
	if options.Delimiter == 0 {
		options.Delimiter = ','
	}
	
	hash := make(map[string][][]string)
	
	checked := false
	header_skipped := false
	
	file, err := os.Open(options.Input)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	
	reader := csv.NewReader(file)
	reader.Comma = options.Delimiter
	reader.TrimLeadingSpace = true
	for {
		
		line, e := reader.Read()
		if e != nil || e == io.EOF {
			break
		}
		
		// Let's check really quick, if there are all the columns for hash key we need
		if !checked {
			max_key_index := getColumnsMaxIndex(options.KeyColumns)
			if max_key_index > len(line) - 1 {
				return nil, errors.New(fmt.Sprintf("Key column index '%v' is beyond the slice length!\nKeep in mind, index count starts at 0.", max_key_index))
			}
			checked = true
		}
		
		// Skip header
		if !header_skipped && options.SkipHeader {
			header_skipped = true
			continue
		}
		
		// Add data
		key := makeKey(line, options.KeyColumns)
		
		if !options.IncludeKeysValues {
			removeKeyColumnsValues(&line, options.KeyColumns)
		}
		
		hash[key] = append(hash[key], line)
	}
	
	return hash, nil
}

// Concatenate columns to make hash key
func makeKey(line []string, columns []int) string {
	if len(columns) == 1 {
		return line[columns[0]]
	}
	
	var key []string
	for _, v := range columns {
		key = append(key, line[v])
	}
	
	return strings.Join(key, ",")
}

// Removes key columns from the line
func removeKeyColumnsValues(line *[]string, key_columns []int) {
	for _, i := range key_columns {
		*line = append((*line)[ :i ], (*line)[ i + 1: ]...)
	}
}

// Find column with highest index
func getColumnsMaxIndex(c []int) int {
	max := 0
	for _, v := range c {
		if v > max {
			max = v
		}
	}
	return max
}