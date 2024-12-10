package main
import "fmt"
import "os"
import "strings"

func star1() {
	inBytes,_ := os.ReadFile("10.txt")
	in := string(inBytes)

	mp := make([][]byte, 0)
	ylen, xlen := 0, 0
	
	for _,line := range strings.Split(in, "\n") {
		if line == "" { continue }
		row := make([]byte, 0)
		for _,c := range line {
			row = append(row, byte(c - '0'))
		}
		mp = append(mp, row)
		ylen++;
		xlen = len(line)
	}

	// for y := 0; y < ylen; y++ {
	// 	for x := 0;x < xlen; x++ {
	// 		fmt.Printf("%v", mp[y][x])
	// 	}
	// 	fmt.Printf("\n")
	// }

	sum := 0
	for y := 0; y < ylen; y++ {
		for x := 0; x < xlen; x++ {
			if mp[y][x] != 0 { continue }

			seens := make([][]int, ylen)
			for y1 := 0; y1 < ylen; y1++ {
				seens[y1] = make([]int, xlen)
			}
			seens[y][x] = 1

			worked := true
			for worked {
				worked = false
				for y1 := 0; y1 < ylen; y1++ {
					for x1 := 0; x1 < xlen; x1++ {
						if seens[y1][x1] != 1 { continue }
						seens[y1][x1] = 2
						for dy := -1; dy <= 1; dy++ {
							for dx := -1; dx <= 1; dx++ {
								if dx == 0 && dy == 0 { continue }
								if dx != 0 && dy != 0 { continue }
								y2, x2 := y1+dy, x1+dx
								if y2 < 0 || y2 >= ylen { continue }
								if x2 < 0 || x2 >= xlen { continue }
								if seens[y2][x2] >= 1 { continue }
								if mp[y2][x2] == mp[y1][x1] + 1  {
									seens[y2][x2] = 1
									worked = true
								}
							}
						}
						
					}
				}
			}

			// for y1 := 0; y1 < ylen; y1++ {
			// 	for x1 := 0; x1 < xlen; x1++ {
			// 		if seens[y1][x1] < 2 {
			// 			fmt.Printf(".")
			// 		} else {
			// 			fmt.Printf("%v", mp[y1][x1])
			// 		}
			// 	}
			// 	fmt.Printf("\n")
			// }
			// fmt.Printf("\n")

			trailheads := 0
			for y1 := 0; y1 < ylen; y1++ {
				for x1 := 0; x1 < xlen; x1++ {
					if mp[y1][x1] == 9 && seens[y1][x1] == 2 { trailheads++; }
				}
			}
			sum += trailheads
			// fmt.Printf("%v,%v = %v\n", y, x, trailheads)
		}
	}
	fmt.Printf("= %v\n", sum)
}

func star2() {
	inBytes,_ := os.ReadFile("10.txt")
	in := string(inBytes)

	mp := make([][]byte, 0)
	ylen, xlen := 0, 0
	
	for _,line := range strings.Split(in, "\n") {
		if line == "" { continue }
		row := make([]byte, 0)
		for _,c := range line {
			row = append(row, byte(c - '0'))
		}
		mp = append(mp, row)
		ylen++;
		xlen = len(line)
	}

	// for y := 0; y < ylen; y++ {
	// 	for x := 0;x < xlen; x++ {
	// 		fmt.Printf("%v", mp[y][x])
	// 	}
	// 	fmt.Printf("\n")
	// }

	var pathsAmount func(x1,y1, x3,y3 int, seens *[][]bool) int;
	pathsAmount = func(x1,y1, x3,y3 int, seens *[][]bool) int {
		(*seens)[y1][x1] = true
		sum := 0
		for dy := -1; dy <= 1; dy++ {
			for dx := -1; dx <= 1; dx++ {
				if dx == 0 && dy == 0 { continue }
				if dx != 0 && dy != 0 { continue }
				y2, x2 := y1+dy, x1+dx
				if y2 < 0 || y2 >= ylen { continue }
				if x2 < 0 || x2 >= xlen { continue }
				if (*seens)[y2][x2] { continue }
				if mp[y2][x2] == mp[y1][x1] + 1  {
					if y2 == y3 && x2 == x3 {
						sum++
					} else {
						sum += pathsAmount(x2, y2, x3, y3, seens)
					}
				}
			}
		}
		(*seens)[y1][x1] = false
		return sum
	}

	pathsAmount0 := func(x1,y1, x3,y3 int) int {
		seens := make([][]bool, ylen)
		for y := 0; y < ylen; y++ {
			seens[y] = make([]bool, xlen)
		}
		return pathsAmount(x1,y1, x3,y3, &seens)
	}
	
	sum := 0
	for y := 0; y < ylen; y++ {
		for x := 0; x < xlen; x++ {
			if mp[y][x] != 0 { continue }

			seens := make([][]int, ylen)
			for y1 := 0; y1 < ylen; y1++ {
				seens[y1] = make([]int, xlen)
			}
			seens[y][x] = 1

			worked := true
			for worked {
				worked = false
				for y1 := 0; y1 < ylen; y1++ {
					for x1 := 0; x1 < xlen; x1++ {
						if seens[y1][x1] != 1 { continue }
						seens[y1][x1] = 2
						for dy := -1; dy <= 1; dy++ {
							for dx := -1; dx <= 1; dx++ {
								if dx == 0 && dy == 0 { continue }
								if dx != 0 && dy != 0 { continue }
								y2, x2 := y1+dy, x1+dx
								if y2 < 0 || y2 >= ylen { continue }
								if x2 < 0 || x2 >= xlen { continue }
								if seens[y2][x2] >= 1 { continue }
								if mp[y2][x2] == mp[y1][x1] + 1  {
									seens[y2][x2] = 1
									worked = true
								}
							}
						}
						
					}
				}
			}

			// for y1 := 0; y1 < ylen; y1++ {
			// 	for x1 := 0; x1 < xlen; x1++ {
			// 		if seens[y1][x1] < 2 {
			// 			fmt.Printf(".")
			// 		} else {
			// 			fmt.Printf("%v", mp[y1][x1])
			// 		}
			// 	}
			// 	fmt.Printf("\n")
			// }
			// fmt.Printf("\n")

			paths := 0
			for y1 := 0; y1 < ylen; y1++ {
				for x1 := 0; x1 < xlen; x1++ {
					if mp[y1][x1] == 9 && seens[y1][x1] == 2 {
						paths += pathsAmount0(x,y, x1,y1)
					}
				}
			}
			sum += paths
			// fmt.Printf("%v,%v = %v\n", y, x, trailheads)
		}
	}
	fmt.Printf("= %v\n", sum)
}

func main() {
	star1()
	star2()
}
