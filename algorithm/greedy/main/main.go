package main
import "study_go/algorithm/greedy"

func main(){
	n := 300
	dis := []int{150,180,120,100,280,160}
	greedy.MinCarRefuelOil(n,dis)
}