package main
import "fmt"
import "os"
import "strings"

func star1() {
	inBytes,_ := os.ReadFile("12.txt")
	in := string(inBytes)
	garden := make([]string, 0)
	for _,line := range strings.Split(in, "\n") {
		if line == "" { continue }
		garden = append(garden, line)
	}
	xlen, ylen := len(garden[0]), len(garden)

	regions := make([][]int, 0)
	for y := 0; y < ylen; y++ {
		regions = append(regions, make([]int, xlen))
	}

	// for y := 0; y < ylen; y++ {
	// 	for x := 0; x < xlen; x++ {
	// 		fmt.Printf("%c", garden[y][x])
	// 	}
	// 	fmt.Printf("   ")
	// 	for x := 0; x < xlen; x++ {
	// 		fmt.Printf("%d", regions[y][x])
	// 	}
	// 	fmt.Printf("\n")
	// }

	regionIds := 1
	sum := 0
	for y := 0; y < ylen; y++ {
		for x := 0; x < xlen; x++ {
			if regions[y][x] != 0 { continue }
			area := 0
			perimeter := 0
			plant := garden[y][x]
			regionId := regionIds
			regionIds++;
			var f func(x, y int);
			f = func(x, y int) {
				if x < 0 || x >= xlen || y < 0 || y >= ylen { return }
				if regions[y][x] == regionId { return }
				if garden[y][x] != plant { return }
				area++
				regions[y][x] = regionId
				for dy := -1; dy <= 1; dy++ {
					for dx := -1; dx <= 1; dx++ {
						if dy == 0 && dx == 0 { continue }
						if dy != 0 && dx != 0 { continue }
						if x+dx < 0 || x+dx >= xlen {
							perimeter++
						} else if y+dy < 0 || y+dy >= ylen {
							perimeter++
						} else if garden[y+dy][x+dx] != plant {
							perimeter++
						} else {
							f(x+dx, y+dy)
						}
					}
				}
			}
			f(x, y)
			sum += area * perimeter
			// fmt.Printf("%v,%v  %c  %d*%d  %d\n", x,y, garden[y][x], area, perimeter, sum)
		}
	}
	
	

	fmt.Printf("= %d\n", sum)
}


func star2() {
	inBytes,_ := os.ReadFile("12.txt")
	in := string(inBytes)
	garden := make([]string, 0)
	for _,line := range strings.Split(in, "\n") {
		if line == "" { continue }
		garden = append(garden, line)
	}
	xlen, ylen := len(garden[0]), len(garden)

	regions := make([][]int, 0)
	for y := 0; y < ylen; y++ {
		regions = append(regions, make([]int, xlen))
	}

	// for y := 0; y < ylen; y++ {
	// 	for x := 0; x < xlen; x++ {
	// 		fmt.Printf("%c", garden[y][x])
	// 	}
	// 	fmt.Printf("   ")
	// 	for x := 0; x < xlen; x++ {
	// 		fmt.Printf("%d", regions[y][x])
	// 	}
	// 	fmt.Printf("\n")
	// }

	regionIds := 1
	sum := 0
	for y := 0; y < ylen; y++ {
		for x := 0; x < xlen; x++ {
			if regions[y][x] != 0 { continue }
			area := 0
			perimeter := 0
			plant := garden[y][x]
			regionId := regionIds
			regionIds++;
			var f func(x, y int);
			f = func(x, y int) {
				if x < 0 || x >= xlen || y < 0 || y >= ylen { return }
				if regions[y][x] == regionId { return }
				if garden[y][x] != plant { return }
				area++
				regions[y][x] = regionId
				for dy := -1; dy <= 1; dy++ {
					for dx := -1; dx <= 1; dx++ {
						if dy == 0 && dx == 0 { continue }
						if dy != 0 && dx != 0 { continue }
						if x+dx < 0 || x+dx >= xlen {
							perimeter++
						} else if y+dy < 0 || y+dy >= ylen {
							perimeter++
						} else if garden[y+dy][x+dx] != plant {
							perimeter++
						} else {
							f(x+dx, y+dy)
						}
					}
				}
			}
			f(x, y)

			isThisRegion := func(x, y int) bool {
				if x < 0 || x >= xlen { return false }
				if y < 0 || y >= ylen { return false }
				return regions[y][x] == regionId
			}
			
			sides := 0
			for y1 := 0; y1 <= ylen; y1++ {
				inSide := false
				inSideSide := true
				for x1 := 0; x1 <= xlen; x1++ {
					//fmt.Printf("####1 %d: %d,%d = %v,%v\n", regionId, x1, y1, isThisRegion(x1, y1), isThisRegion(x1, y1-1))
					if isThisRegion(x1, y1) != isThisRegion(x1, y1-1) {
						if !inSide || inSideSide != isThisRegion(x1, y1) {
							sides++
							//fmt.Printf("##1 %d: %d,%d = %d\n", regionId, x1, y1, sides)
						}
						inSide = true
						inSideSide = isThisRegion(x1, y1)
					} else {
						inSide = false
					}
				}
			}
			for x1 := 0; x1 <= xlen; x1++ {
				inSide := false
				inSideSide := true
				for y1 := 0; y1 <= ylen; y1++ {
					if isThisRegion(x1, y1) != isThisRegion(x1-1, y1) {
						if !inSide || inSideSide != isThisRegion(x1, y1) {
							sides++
							//fmt.Printf("##2 %d: %d,%d = %d\n", regionId, x1, y1, sides)
						}
						inSide = true
						inSideSide = isThisRegion(x1, y1)
					} else {
						inSide = false
					}
				}
			}
					
			sum += area * sides
			//fmt.Printf("%v,%v  %c  %d*%d  %d\n", x,y, garden[y][x], area, sides, sum)
		}
	}
	
	// for y := 0; y < ylen; y++ {
	// 	for x := 0; x < xlen; x++ {
	// 		fmt.Printf("%c", garden[y][x])
	// 	}
	// 	fmt.Printf("   ")
	// 	for x := 0; x < xlen; x++ {
	// 		fmt.Printf("%d", regions[y][x])
	// 	}
	// 	fmt.Printf("\n")
	// }
	

	fmt.Printf("= %d\n", sum)

	// 846196 too low
}

func main() {
	//star1()
	star2()
}
