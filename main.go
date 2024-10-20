package main

import (
	"fmt"
	"strings"
)

// Function to split text into words
func wordSplitter(data interface{}) interface{} {
	text := data.(string)
	words := strings.Fields(text) // Split by spaces (basic splitting)
	wordMap := make(map[string]int)
	for _, word := range words {
		wordMap[word] = 1 // Initialize each word's count as 1
	}
	return wordMap
}

// Function to combine word counts
func wordCounter(a, b interface{}) interface{} {
	mapA := a.(map[string]int)
	mapB := b.(map[string]int)

	// Merge two maps, summing word counts
	for word, countB := range mapB {
		if countA, exists := mapA[word]; exists {
			mapA[word] = countA + countB
		} else {
			mapA[word] = countB
		}
	}
	return mapA
}

func main() {
	// Example dataset
	data := []string{
		"hello world",
		"this is a simple word count example",
		"hello hello world",
		"word count example in golang",
	}

	// Step 1: Partition the dataset
	partitions := partitionData(data, 3) // Split data into 3 partitions

	// Step 2: Create an RDD from the partitions
	rdd := &RDD{Partitions: partitions}

	// Step 3: Map operation to split words
	mappedRDD := rdd.Map(wordSplitter)

	// Step 4: Reduce operation to aggregate word counts
	wordCounts := mappedRDD.Reduce(wordCounter).(map[string]int)

	// Output the word counts
	fmt.Println("Word Count Results:")
	for word, count := range wordCounts {
		fmt.Printf("%s: %d\n", word, count)
	}
}
