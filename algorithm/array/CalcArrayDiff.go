package array

func CalculateDifferences(arr []int) []int {
	differences := make([]int, 0)
	for i := 1; i < len(arr); i++ {
		diff := arr[i] - arr[i-1]
		differences = append(differences, diff)
	}
	return differences
}
