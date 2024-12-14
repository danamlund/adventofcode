package main
import "fmt"
import "os"
import "strings"

type robot struct {
	x, y, dx, dy int
}

func star1() {
	inBytes,_ := os.ReadFile("14.txt")
	in := string(inBytes)

	xlen, ylen := 101, 103
	//xlen, ylen := 11, 7

	robots := make([]robot, 0)
	
	for _,line := range strings.Split(in, "\n") {
		if line == "" { continue }
		var robot robot
		fmt.Sscanf(line, "p=%d,%d v=%d,%d", &robot.x, &robot.y, &robot.dx, &robot.dy)
		robots = append(robots, robot)
	}

	// for _,robot := range robots {
	// 	fmt.Printf("## %v,%v  %v,%v\n", robot.x, robot.y, robot.dx, robot.dy)
	// }

	mod := func(num, mod int) int {
		for num < 0 { num += mod }
		return num % mod
	}

	for second := 0; second < 100; second++ {
		for i := 0; i < len(robots); i++ {
			robot := &robots[i]
			//fmt.Printf("## %d,%d  %d,%d  %d,%d\n", robot.x, robot.y, robot.dx, robot.dy, mod(robot.x+robot.dx, xlen), mod(robot.y+robot.dy, ylen))
			robot.x = mod(robot.x+robot.dx, xlen)
			robot.y = mod(robot.y+robot.dy, ylen)
		}

		// fmt.Printf("\n")
		// for y := 0; y < ylen; y++ {
		// 	for x := 0; x < xlen; x++ {
		// 		amount := 0
		// 		for _,robot := range robots {
		// 			if robot.x == x && robot.y == y { amount++ }
		// 		}
		// 		if amount == 0 { fmt.Printf(".")
		// 		} else { fmt.Printf("%d", amount) }
		// 	}
		// 	fmt.Printf("\n")
		// }
	}

	sum := 0
	qs := make([]int, 4)
	for _,robot := range robots {
		if robot.x < xlen/2 && robot.y < ylen/2 { qs[0]++
		} else if robot.x > xlen/2 && robot.y < ylen/2 { qs[1]++
		} else if robot.x < xlen/2 && robot.y > ylen/2 { qs[2]++
		} else if robot.x > xlen/2 && robot.y > ylen/2 { qs[3]++
		}
	}
	sum = qs[0] * qs[1] * qs[2] * qs[3]
	
	fmt.Printf("= %d\n", sum)
}

func star2() {
	inBytes,_ := os.ReadFile("14.txt")
	in := string(inBytes)

	xlen, ylen := 101, 103
	//xlen, ylen := 11, 7

	robots := make([]robot, 0)
	
	for _,line := range strings.Split(in, "\n") {
		if line == "" { continue }
		var robot robot
		fmt.Sscanf(line, "p=%d,%d v=%d,%d", &robot.x, &robot.y, &robot.dx, &robot.dy)
		robots = append(robots, robot)
	}

	// for _,robot := range robots {
	// 	fmt.Printf("## %v,%v  %v,%v\n", robot.x, robot.y, robot.dx, robot.dy)
	// }

	mod := func(num, mod int) int {
		for num < 0 { num += mod }
		return num % mod
	}

	for second := 0; second < 100000; second++ {
		for i := 0; i < len(robots); i++ {
			robot := &robots[i]
			//fmt.Printf("## %d,%d  %d,%d  %d,%d\n", robot.x, robot.y, robot.dx, robot.dy, mod(robot.x+robot.dx, xlen), mod(robot.y+robot.dy, ylen))
			robot.x = mod(robot.x+robot.dx, xlen)
			robot.y = mod(robot.y+robot.dy, ylen)
		}

		//if second < 1000 { continue }

		nearmidx := 0
		for _,robot := range robots {
			if (robot.x > 2*ylen/10 && robot.x < 8*ylen/10) { nearmidx++ }
		}
		if nearmidx < 87*len(robots)/100 { continue }
		

		
		grid := make([][]byte, ylen)
		for y := 0; y < ylen; y++ { grid[y] = make([]byte, xlen) }
		for _,robot := range robots {
			grid[robot.y][robot.x]++
		}
		fmt.Printf("\n\nsecond=%d\n", second+1)
		for y := 0; y < ylen; y++ {
			for x := 0; x < xlen; x++ {
				if grid[y][x] == 0 { fmt.Printf(".")
				} else { fmt.Printf("%d", grid[y][x]) }
			}
			fmt.Printf("\n")
		}
	}
	// 7092 too low
	// 7093

	/*
...........................1.......................
...................................................
...................................................
...1.1.....1111111111111111111111111111111.........
...........1.............................1.........
...........1.............................1.........
...........1.............................1.........
.1.........1.............................1.........
...........1..............1..............1.........
......1....1.............111.............1.........
...........1............11111............1.........
...........1...........1111111...........1.........
...........1..........111111111..........1.........
...........1............11111............1........1
...........1...........1111111...........1.........
.........1.1..........111111111..........1.1.......
...........1.........11111111111.........1.........
...........1........1111111111111........1.........
...........1..........111111111..........1.1.......
...........1.........11111111111.........1.........
...........1........1111111111111........1.........
...........1.......111111111111111.......1.........
...........1......11111111111111111......1.........
...........1........1111111111111........1.........
...........1.......111111111111111.......1.........
...........1......11111111111111111......1.....1...
...........1.....1111111111111111111.....1.........
...........1....111111111111111111111....1.........
.........1.1.............111.............1.........
...........1.............111.............1......1..
...........1.............111.............1.........
...........1.............................1.........
...1.......1.............................1.........
...........1.............................1.........
...........1.............................1.........
...........1111111111111111111111111111111.........
..1........................................1.......
.1................1..........................1.....
...1...............................................
	*/
}


func main() {
	//star1()
	star2()
}
