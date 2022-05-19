package sortKM

// TODO: answer here
import "sort"

//concrete strategy implementation
type DescendingSort struct{}

func (ds *DescendingSort) Sort(array []int) {
	//choose any sort algo you want
	// TODO: answer here
	sort.Slice(array, func(i, j int) bool {
		return array[i] > array[j]
	})
}
