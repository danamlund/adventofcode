package main

import "fmt"
import "os"
import "bufio"
import "strings"

func star1() {
	f,_ := os.Open("02.txt")
	fscan := bufio.NewScanner(f)
	fscan.Split(bufio.ScanLines)

	safes := 0
	for fscan.Scan() {
		line := fscan.Text()
		count := 0
		prevnum := -1
		prevprevnum := -1
		safe := true
		for _,e  := range strings.Split(line, " ") {
			var num int
			fmt.Sscanf(e, "%d", &num)
			diff1 := num - prevnum
			diff2 := prevnum - prevprevnum
			if count >= 2 && (diff1 > 0) != (diff2 > 0) {
				safe = false
			}
			if count >= 1 && (diff1 > 3 || diff1 < -3 || diff1 == 0) {
				safe = false
			}
			//fmt.Printf("## %v: %v %v %v diff1=%v diff2=%v safe=%v\n", count, prevprevnum, prevnum, num, diff1, diff2, safe)
			prevprevnum = prevnum
			prevnum = num
			count++;
		}
		if safe {
			safes++
		}
		//fmt.Printf("\n")
	}
	
	fmt.Printf("= %v\n", safes);
	// 6 is wrong
}

func issafe(nums []int) bool {
	for i,_ := range nums {
		if i >= 1 {
			diff1 := nums[i] - nums[i-1]
			if diff1 >3 || diff1 < -3 || diff1 == 0 {
				return false
			}
			if i >= 2 {
				diff2 := nums[i-1] - nums[i-2]
				if (diff1 > 0) != (diff2 > 0) {
					return false
				}
			}
		}
	}
	return true
}

func star2() {
	f,_ := os.Open("02.txt")
	fscan := bufio.NewScanner(f)
	fscan.Split(bufio.ScanLines)

	safes := 0
	for fscan.Scan() {
		line := fscan.Text()
		list := make([]int, 0)
		for _,e  := range strings.Split(line, " ") {
			var num int
			fmt.Sscanf(e, "%d", &num)
			list = append(list, num)
		}
		if issafe(list) {
			safes++;
			//fmt.Printf("## %v safe\n", list)
		} else {
			for i := 0; i < len(list); i++ {
				list2 := append([]int(nil), list...)
				list2 = append(list2[:i], list2[i+1:]...)
				//fmt.Printf("##3 %v -> %v\n", list, list2)
				if issafe(list2) {
					safes++;
					//fmt.Printf("##2 %v safe\n", list2)
					break
				}
			}
		}
	}
	
	fmt.Printf("= %v\n", safes);
	// 311 is wrong
}

func main() {
	star1()
	star2()
}
