package main
import "fmt"
import "strings"
import "os"

func cansum(result int, nums []int) bool {
	return cansum2(result, nums[0], nums[1:])
}

func cansum2(result int, cur int, nums []int) bool {
	if len(nums) == 0 { return result == cur }
	if cansum2(result, cur + nums[0], nums[1:]) { return true }
	if cansum2(result, cur * nums[0], nums[1:]) { return true }
	return false
}

func star() {
	inBytes,_ := os.ReadFile("07.txt")
	in := string(inBytes)

	sum := 0
	for _,line := range strings.Split(in, "\n") {
		if strings.TrimSpace(line) == "" { continue }
		splitcol := strings.Split(line, ":")
		var result int;
		fmt.Sscanf(splitcol[0], "%d", &result)
		nums := make([]int, 0)
		for _,numstr := range strings.Split(splitcol[1], " ") {
			if strings.TrimSpace(numstr) == "" { continue }
			var num int;
			fmt.Sscanf(numstr, "%d", &num)
			nums = append(nums, num)
		}
		if cansum(result, nums) { sum += result }
		//fmt.Printf("%d = %v =  %v\n", result, nums, cansum(result, nums))
	}

	fmt.Printf("= %v\n", sum)
}

func star2_cansum(result int, nums []int) bool {
	return star2_cansum2(result, nums[0], nums[1:])
}

func star2_cansum2(result int, cur int, nums []int) bool {
	if len(nums) == 0 { return result == cur }
	if star2_cansum2(result, cur + nums[0], nums[1:]) { return true }
	if star2_cansum2(result, cur * nums[0], nums[1:]) { return true }
	var cur2 int
	fmt.Sscanf(fmt.Sprintf("%d%d", cur, nums[0]), "%d", &cur2)
	//fmt.Printf("## %d%d = %d\n", cur, nums[0], cur2)
	if star2_cansum2(result, cur2, nums[1:]) { return true }
	return false
}

func star2() {
	inBytes,_ := os.ReadFile("07.txt")
	in := string(inBytes)

	sum := 0
	for _,line := range strings.Split(in, "\n") {
		if strings.TrimSpace(line) == "" { continue }
		splitcol := strings.Split(line, ":")
		var result int;
		fmt.Sscanf(splitcol[0], "%d", &result)
		nums := make([]int, 0)
		for _,numstr := range strings.Split(splitcol[1], " ") {
			if strings.TrimSpace(numstr) == "" { continue }
			var num int;
			fmt.Sscanf(numstr, "%d", &num)
			nums = append(nums, num)
		}
		if star2_cansum(result, nums) { sum += result }
		//fmt.Printf("%d = %v =  %v\n", result, nums, cansum(result, nums))
	}

	fmt.Printf("= %v\n", sum)
}

func main() {
	//star1()
	star2()
}
