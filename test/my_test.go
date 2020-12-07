package test

import (
	"container/heap"
	"container/list"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"testing"
	"time"
	"unicode/utf8"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // 把 sum 发送到通道 c
}

func TestMainAll(t *testing.T) {
	var balance = [5]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
	fmt.Println(balance)

	bs := balance[2:5]
	var bs2 []float32
	//bs2 = append(bs2, 0)
	copy(bs2, bs)
	//bs2[1] = 8.0
	fmt.Println(balance)
	fmt.Println(bs2)

	var s1 []int
	var s2 = []int{}
	fmt.Println(s1 == nil)
	fmt.Println(s2 == nil)

	fmt.Printf("%#v\n", s1)
	fmt.Printf("%#v\n", s2)

	//slice := append([]byte("hello "), "world"...)

	a := "wo爱你"
	fmt.Println(len(a))
	fmt.Println(len([]byte(a)))
	fmt.Println(len([]rune(a)))
	fmt.Println(utf8.RuneCountInString(a))
	fmt.Println(string([]rune(a)[2:3]))

	list := list.New()
	list.Init()
	list.PushBack(1)
	list.PushBack(2)
	//ring.New(3)

	fmt.Printf("len: %v\n", list.Len())
	fmt.Println(list.Back(), list.Front())
}

type MyStruct struct {
}

func (me MyStruct) Print() {
	fmt.Println("MyStruct print")
}

func TestStrings(t *testing.T) {
	left := strings.TrimLeft("hllhllo Tom", "hl")
	fmt.Println("strings.TrimLeft:", left)
	prefix := strings.TrimPrefix("hello Tom", "hl")
	fmt.Println("strings.TrimPrefix:", prefix)
}

func TestBase(t *testing.T) {
	var s1, s2 []string
	s2 = []string{}
	fmt.Println(s1 == nil, s2 == nil)
	fmt.Println(len(s1), len(s2))
}

func TestSort(t *testing.T) {
	intList := []int{2, 4, 3, 5, 7, 6, 9, 8, 1, 0}
	float8List := []float64{4.2, 5.9, 12.3, 10.0, 50.4, 99.9, 31.4, 27.81828, 3.14}
	stringList := []string{"a", "c", "b", "d", "f", "i", "z", "x", "w", "y"}

	sort.Ints(intList)
	sort.Float64s(float8List)
	sort.Strings(stringList)

	fmt.Printf("%v\n%v\n%v\n", intList, float8List, stringList)
}

func TestSortReverse(t *testing.T) {
	intList := []int{2, 4, 3, 5, 7, 6, 9, 8, 1, 0}
	float8List := []float64{4.2, 5.9, 12.3, 10.0, 50.4, 99.9, 31.4, 27.81828, 3.14}
	stringList := []string{"a", "c", "b", "d", "f", "i", "z", "x", "w", "y"}

	sort.Sort(sort.Reverse(sort.IntSlice(intList)))
	sort.Sort(sort.Reverse(sort.Float64Slice(float8List)))
	sort.Sort(sort.Reverse(sort.StringSlice(stringList)))

	fmt.Printf("%v\n%v\n%v\n", intList, float8List, stringList)
}

func TestSortDoubleDimensionalArray(t *testing.T) {
	intervals := [][]int{{1, 4}, {1, 5}, {3, 6}, {2, 8}}
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] == intervals[j][0] {
			return intervals[j][1] < intervals[i][1]
		}
		return intervals[i][0] < intervals[j][0]
	})
	fmt.Println(intervals)
}

func TestStringToLower(t *testing.T) {
	str := "##Quentome"
	fmt.Println(strings.ToLower(str))
	fmt.Println(strconv.Itoa(2) + "" + strconv.Itoa(6))
}

func TestSplitStr(t *testing.T) {
	fmt.Println(strings.Split("a", " "))
	split := strings.Split("a ", " ")
	fmt.Println(split)
	fmt.Println(len(split[len(split)-1]))
}

func TestAppend(t *testing.T) {
	a := make(map[string][]int)

	//value := []int{1}
	a["1"] = []int{1}

	for k, v := range a {
		fmt.Println("=== k:", k)
		fmt.Print("=== v:", v)
		fmt.Printf("##### %p\n", v)
	}

	val := a["1"]
	fmt.Printf("##### %p\n", val)
	val = append(val, 100)
	fmt.Printf("##### %p\n", val)

	//a["1"] = val

	for k, v := range a {
		fmt.Println("=== k:", k)
		fmt.Print("=== v:", v)
		fmt.Printf("##### %p\n", v)
	}
}

// An IntHeap is a min-heap of ints.
type IntHeap []int

func (h IntHeap) Len() int { return len(h) }

// 这里决定 大小顶堆 现在是小顶堆
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func TestNativeHeap(t *testing.T) {
	h := &IntHeap{2, 1, 5, 6, 4, 3, 7, 9, 8, 0} // 创建slice
	heap.Init(h)                                //堆化
	fmt.Println(*h)                             // [0 1 3 6 2 5 7 9 8 4] 由Less方法可控制小顶堆
	fmt.Println(heap.Pop(h))                    // 调用pop 0 返回移除的顶部最小元素
	heap.Push(h, 6)                             // 调用push [1 2 3 6 4 5 7 9 8] 添加一个元素进入堆中进行堆化
	fmt.Println("new: ", *h)                    // [1 2 3 6 4 5 7 9 8 6]
	for len(*h) > 0 {                           // 持续推出顶部最小元素
		fmt.Printf("%d \n ", heap.Pop(h))
	}
	h1 := &IntHeap{2, 1, 5, 6, 4, 3, 7, 9, 8, 0} // 创建slice
	sort.Sort(h1)
	fmt.Println(h1)
}
