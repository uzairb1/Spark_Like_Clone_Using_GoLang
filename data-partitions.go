package main

// Partition the data across workers
type Partition struct {
	Data []interface{}
}

func partitionData(data []string, numPartitions int) []Partition {
	partitions := make([]Partition, numPartitions)
	for i, val := range data {
		idx := i % numPartitions
		partitions[idx].Data = append(partitions[idx].Data, val)
	}
	return partitions
}
