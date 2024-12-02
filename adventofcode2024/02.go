/*
https://adventofcode.com/2024/day/2
--- Day 2: Red-Nosed Reports ---

Fortunately, the first location The Historians want to search isn't a long walk from the Chief Historian's office.

While the Red-Nosed Reindeer nuclear fusion/fission plant appears to contain no sign of the Chief Historian, the engineers there run up to you as soon as they see you. Apparently, they still talk about the time Rudolph was saved through molecular synthesis from a single electron.

They're quick to add that - since you're already here - they'd really appreciate your help analyzing some unusual data from the Red-Nosed reactor. You turn to check if The Historians are waiting for you, but they seem to have already divided into groups that are currently searching every corner of the facility. You offer to help with the unusual data.

The unusual data (your puzzle input) consists of many reports, one report per line. Each report is a list of numbers called levels that are separated by spaces. For example:

7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9

This example data contains six reports each containing five levels.

The engineers are trying to figure out which reports are safe. The Red-Nosed reactor safety systems can only tolerate levels that are either gradually increasing or gradually decreasing. So, a report only counts as safe if both of the following are true:

    The levels are either all increasing or all decreasing.
    Any two adjacent levels differ by at least one and at most three.

In the example above, the reports can be found safe or unsafe by checking those rules:

    7 6 4 2 1: Safe because the levels are all decreasing by 1 or 2.
    1 2 7 8 9: Unsafe because 2 7 is an increase of 5.
    9 7 6 2 1: Unsafe because 6 2 is a decrease of 4.
    1 3 2 4 5: Unsafe because 1 3 is increasing but 3 2 is decreasing.
    8 6 4 4 1: Unsafe because 4 4 is neither an increase or a decrease.
    1 3 6 7 9: Safe because the levels are all increasing by 1, 2, or 3.

So, in this example, 2 reports are safe.

Analyze the unusual data from the engineers. How many reports are safe?

Your puzzle answer was 246.
--- Part Two ---

The engineers are surprised by the low number of safe reports until they realize they forgot to tell you about the Problem Dampener.

The Problem Dampener is a reactor-mounted module that lets the reactor safety systems tolerate a single bad level in what would otherwise be a safe report. It's like the bad level never happened!

Now, the same rules apply as before, except if removing a single level from an unsafe report would make it safe, the report instead counts as safe.

More of the above example's reports are now safe:

    7 6 4 2 1: Safe without removing any level.
    1 2 7 8 9: Unsafe regardless of which level is removed.
    9 7 6 2 1: Unsafe regardless of which level is removed.
    1 3 2 4 5: Safe by removing the second level, 3.
    8 6 4 4 1: Safe by removing the third level, 4.
    1 3 6 7 9: Safe without removing any level.

Thanks to the Problem Dampener, 4 reports are actually safe!

Update your analysis by handling situations where the Problem Dampener can remove a single level from unsafe reports. How many reports are now safe?

Your puzzle answer was 318.

Both parts of this puzzle are complete! They provide two gold stars: **
*/
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
