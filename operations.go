package main

type RDD struct {
	Partitions []Partition
}

// Map applies a function to each partition and returns a new RDD
func (rdd *RDD) Map(f func(interface{}) interface{}) *RDD {
	newPartitions := make([]Partition, len(rdd.Partitions))
	for i, partition := range rdd.Partitions {
		newPartitions[i].Data = make([]interface{}, len(partition.Data))
		for j, val := range partition.Data {
			newPartitions[i].Data[j] = f(val)
		}
	}
	return &RDD{Partitions: newPartitions}
}

// Reduce aggregates the partitions by applying a function to them
func (rdd *RDD) Reduce(f func(interface{}, interface{}) interface{}) interface{} {
	var result interface{}
	for _, partition := range rdd.Partitions {
		for _, val := range partition.Data {
			if result == nil {
				result = val
			} else {
				result = f(result, val)
			}
		}
	}
	return result
}
