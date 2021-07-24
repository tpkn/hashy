package hashy

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

// csvFileReader reads csv file and pass it's lines into callback function
func csvFileReader(o Options, callback func(line []string)) error {
	if len(o.KeyColumns) == 0 {
		return errors.New("key column is not set")
	}
	
	var max_key_index = 0
	var columns_checked = false
	var header_skipped = false
	
	file, err := os.Open(o.Input)
	if err != nil {
		return err
	}
	defer file.Close()
	
	var reader = csv.NewReader(file)
	reader.Comma = o.Delimiter
	reader.LazyQuotes = o.LazyQuotes
	reader.TrimLeadingSpace = true
	for {
		
		line, e := reader.Read()
		if e != nil {
			if e == io.EOF {
				break
			} else {
				if strings.Contains(e.Error(), "wrong number of fields") {
					// Process line if it has wrong number of field, but it's length is >= max_key_index
					if len(line) - 1 < max_key_index {
						continue
					}
				} else if strings.Contains(e.Error(), "index out of range") {
					// Skip this one
					continue
				} else {
					return e
				}
			}
		}
		
		// Skip header if needed
		if !header_skipped && o.SkipHeader {
			header_skipped = true
			continue
		}
		
		// Let's check really quick, if there are all the columns for hash key we need
		if !columns_checked {
			max_key_index = getColumnsMaxIndex(o.KeyColumns)
			if max_key_index > len(line) - 1 {
				return errors.New(fmt.Sprintf("Key column index '%v' is beyond the slice length! Keep in mind, index count starts at 0.", max_key_index))
			}
			columns_checked = true
		}
		
		callback(line)
		
	}
	
	return nil
}

// makeHashKey concats columns to make a hash key
func makeHashKey(line []string, columns []int) string {
	if len(columns) == 1 {
		return line[columns[0]]
	}
	
	var key []string
	for _, v := range columns {
		key = append(key, line[v])
	}
	
	return strings.Join(key, ",")
}

// removeKeyColumns removes key columns from the line
func removeKeyColumns(line *[]string, key_columns []int) {
	if len(key_columns) == len(*line) {
		*line = []string{ "" }
	} else {
		for _, i := range key_columns {
			*line = append((*line)[ :i ], (*line)[ i + 1: ]...)
		}
	}
}

// getColumnsMaxIndex finds the highest index of columns
func getColumnsMaxIndex(c []int) int {
	max := 0
	for _, v := range c {
		if v > max {
			max = v
		}
	}
	return max
}
