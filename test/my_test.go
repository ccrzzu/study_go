package test

import (
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

func TestSort(t *testing.T)  {
	intList := [] int {2, 4, 3, 5, 7, 6, 9, 8, 1, 0}
	float8List := [] float64 {4.2, 5.9, 12.3, 10.0, 50.4, 99.9, 31.4, 27.81828, 3.14}
	stringList := [] string {"a", "c", "b", "d", "f", "i", "z", "x", "w", "y"}

	sort.Ints(intList)
	sort.Float64s(float8List)
	sort.Strings(stringList)

	fmt.Printf("%v\n%v\n%v\n", intList, float8List, stringList)
}

func TestSortReverse(t *testing.T)  {
	intList := [] int {2, 4, 3, 5, 7, 6, 9, 8, 1, 0}
	float8List := [] float64 {4.2, 5.9, 12.3, 10.0, 50.4, 99.9, 31.4, 27.81828, 3.14}
	stringList := [] string {"a", "c", "b", "d", "f", "i", "z", "x", "w", "y"}

	sort.Sort(sort.Reverse(sort.IntSlice(intList)))
	sort.Sort(sort.Reverse(sort.Float64Slice(float8List)))
	sort.Sort(sort.Reverse(sort.StringSlice(stringList)))

	fmt.Printf("%v\n%v\n%v\n", intList, float8List, stringList)
}

func TestSortDoubleDimensionalArray(t *testing.T)  {
	intervals := [][]int{{1,4},{1,5},{3,6},{2,8}}
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] == intervals[j][0]{
			return intervals[j][1] < intervals[i][1]
		}
		return intervals[i][0] < intervals[j][0]
	})
	fmt.Println(intervals)
}

func TestStringToLower(t *testing.T)  {
	str := "##Quentome"
	fmt.Println(strings.ToLower(str))
	fmt.Println(strconv.Itoa(2)+""+strconv.Itoa(6))
}
