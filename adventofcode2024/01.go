package main

import "fmt"
import "os"
import "bufio"
import "sort"
import "strings"
import "strconv"

func Absi (x int) int {
	if x >= 0 {
		return x
	} else {
		return -x
	}
}

func star1() {
	f, _ := os.Open("01.txt")
	fscan := bufio.NewScanner(f)
	fscan.Split(bufio.ScanLines)
	
	list1 := make([]int, 0)
	list2 := make([]int, 0)

	for fscan.Scan() {
		line := fscan.Text()
		split := strings.Split(line, "   ")
		num1,_ := strconv.Atoi(split[0])
		num2,_ := strconv.Atoi(split[1])
		list1 = append(list1, num1)
		list2 = append(list2, num2)
	}

	sort.Ints(list1)
	sort.Ints(list2)

	diffs := 0
	for i := 0; i < len(list1); i++ {
		diffs += Absi(list1[i] - list2[i])
		//fmt.Printf("%v - %v = %v (%v)\n", list1[i], list2[i], Absi(list1[i] - list2[i]), diffs)
	}
	
	fmt.Printf("= %v\n", diffs);
}

func star2() {
	f, _ := os.Open("01.txt")
	fscan := bufio.NewScanner(f)
	fscan.Split(bufio.ScanLines)
	
	list1 := make([]int, 0)
	list2 := make([]int, 0)

	for fscan.Scan() {
		line := fscan.Text()
		split := strings.Split(line, "   ")
		num1,_ := strconv.Atoi(split[0])
		num2,_ := strconv.Atoi(split[1])
		list1 = append(list1, num1)
		list2 = append(list2, num2)
	}

	sort.Ints(list1)
	sort.Ints(list2)

	sum := 0
	for i := 0; i < len(list1); i++	{
		sames := 0
		for j := 0; j < len(list2); j++ {
			if list1[i] == list2[j] {
				sames++;
			}
		}
		sum += list1[i] * sames;
	}
	
	fmt.Printf("= %v\n", sum);
}

func main() {
	star1()
	star2()
}
