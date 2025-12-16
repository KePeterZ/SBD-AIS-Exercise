package mapred

import (
	"strings"
)

// MapReduce implements MapReduceInterface
type MapReduce struct {
}

// Ensure MapReduce implements the interface
var _ MapReduceInterface = (*MapReduce)(nil)

// Run executes the MapReduce word count
func (mr *MapReduce) Run(input []string) map[string]int {
	// Map phase
	intermediate := make(map[string][]int)

	for _, text := range input {
		kvs := mr.wordCountMapper(text)
		for _, kv := range kvs {
			intermediate[kv.Key] = append(intermediate[kv.Key], kv.Value)
		}
	}

	// Reduce phase
	result := make(map[string]int)
	for key, values := range intermediate {
		kv := mr.wordCountReducer(key, values)
		result[kv.Key] = kv.Value
	}

	return result
}

// wordCountMapper splits text into words and emits (word, 1)
func (mr *MapReduce) wordCountMapper(text string) []KeyValue {
	words := strings.Fields(text)
	kvs := make([]KeyValue, 0, len(words))

	for _, word := range words {
		word = strings.ToLower(word)
		kvs = append(kvs, KeyValue{
			Key:   word,
			Value: 1,
		})
	}

	return kvs
}

// wordCountReducer sums all counts for a given word
func (mr *MapReduce) wordCountReducer(key string, values []int) KeyValue {
	sum := 0
	for _, v := range values {
		sum += v
	}

	return KeyValue{
		Key:   key,
		Value: sum,
	}
}
