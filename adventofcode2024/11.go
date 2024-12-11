package main
import "fmt"
import "os"
import "strings"
import "math"

func star1() {
	inBytes,_ := os.ReadFile("11.txt")
	in := string(inBytes)

	stones := make([]int, 0)
	
	for _,numstr := range strings.Split(in, " ") {
		var num int
		fmt.Sscanf(numstr, "%d", &num);
		stones = append(stones, num)
	}

	//fmt.Printf("%v\n", stones)
	for blinks := 0; blinks < 25; blinks++ {
		stones1 := make([]int, 0)
		for _,stone := range stones {
			if stone == 0 {
				stones1 = append(stones1, 1)
				continue
			}
			//fmt.Printf("## %v %v %v\n", stone, math.Log10(float64(stone)), int(math.Log10(float64(stone)))+1)
			if int(math.Log10(float64(stone)) + 1) % 2 == 0 {
				str := fmt.Sprintf("%d", stone)
				var num int
				fmt.Sscanf(str[:len(str)/2], "%d", &num)
				//fmt.Printf("## %v: %v = %v = %d\n", stone, str, str[:len(str)/2], num)
				stones1 = append(stones1, num)
				fmt.Sscanf(str[len(str)/2:], "%d", &num)
				//fmt.Printf("## %v: %v = %v = %d\n", stone, str, str[len(str)/2:], num)
				stones1 = append(stones1, num)
				continue
			}
			stones1 = append(stones1, stone * 2024)
		}
		stones = stones1
		//fmt.Printf("%v\n", stones)
	}

	fmt.Printf("= %v\n", len(stones))
}

func star2() {
	inBytes,_ := os.ReadFile("11.txt")
	in := string(inBytes)

	stones := make([]int, 0)
	
	for _,numstr := range strings.Split(in, " ") {
		var num int
		fmt.Sscanf(numstr, "%d", &num);
		stones = append(stones, num)
	}

	cache := make(map[string]int)

	var blinker func(stone, blinks int) int;
	
	blinker = func(stone, blinks int) int {
		str := fmt.Sprintf("%d_%d", stone, blinks)
		val,ok := cache[str]
		if ok { return val }
		if blinks == 0 {
			return 1
		}
		if stone == 0 {
			amount := blinker(1, blinks-1)
			cache[str] = amount
			return amount;
		}
		if int(math.Log10(float64(stone)) + 1) % 2 == 0 {
			str := fmt.Sprintf("%d", stone)
			var num1, num2 int
			fmt.Sscanf(str[:len(str)/2], "%d", &num1)
			fmt.Sscanf(str[len(str)/2:], "%d", &num2)
			amount := blinker(num1, blinks-1) + blinker(num2, blinks-1)
			cache[str] = amount
			return amount
		}
		amount := blinker(stone * 2024, blinks-1)
		cache[str] = amount
		return amount
	}

	//fmt.Printf("%v\n", stones)
	stonesamount := 0
	for _,stone := range stones {
		stonesamount += blinker(stone, 75)
	}

	fmt.Printf("= %v\n", stonesamount)
}

func main() {
	//star1()
	star2()
}
