package main

import (
	"fmt"
	"time"
)

// Master-worker logic
type Task struct {
	id      int
	payload string
}

func worker(id int, tasks <-chan Task, results chan<- string) {
	for task := range tasks {
		time.Sleep(time.Second) // Simulate some processing delay
		results <- fmt.Sprintf("Worker %d processed task %d: %s", id, task.id, task.payload)
	}
}

func master() {
	numWorkers := 3
	tasks := make(chan Task, 10)
	results := make(chan string, 10)

	for w := 1; w <= numWorkers; w++ {
		go worker(w, tasks, results)
	}

	// Distribute tasks (You will use partitionData to distribute)
	data := []string{
		"hello world",
		"this is a simple word count example",
		"hello hello world",
		"word count example in golang",
	}

	// Call partitionData to split the data
	partitions := partitionData(data, numWorkers)

	// Assign tasks based on partitions
	for i, partition := range partitions {
		tasks <- Task{id: i, payload: fmt.Sprintf("%v", partition)}
	}
	close(tasks)

	for i := 0; i < len(partitions); i++ {
		fmt.Println(<-results)
	}
}
