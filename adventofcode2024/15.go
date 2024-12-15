package main
import "fmt"
import "os"
import "strings"
import "log"

func star1() {
	inBytes,_ := os.ReadFile("15.txt")
	in := string(inBytes)

	wh := make([][]byte, 0)
	movs := make([]byte, 0)
	inState := 0
	for _,line := range strings.Split(in, "\n") {
		if inState == 0 {
			if line == "" {
				inState = 1
				continue
			}
			row := make([]byte, len(line))
			for x := 0; x < len(line); x++ {
				row[x] = line[x]
			}
			wh = append(wh, row)
			continue
		}
		for i := 0; i < len(line); i++ {
			movs = append(movs, line[i])
		}
	}
	xlen, ylen := len(wh[0]), len(wh)

	rx, ry := 0, 0
	for y := 0; y < ylen; y++ {
		for x := 0; x < xlen; x++ {
			if wh[y][x] == '@' {
				rx,ry = x,y
			}
		}
	}
	for _,mov := range movs {
		// fmt.Printf("\n\nr=%d,%d %c\n", rx, ry, mov)
		// for y := 0; y < ylen; y++ {
		// 	for x := 0; x < xlen; x++ {
		// 		fmt.Printf("%c", wh[y][x])
		// 	}
		// 	fmt.Printf("\n")
		// }

		dx, dy := 0, 0
		if mov == '^' { dx,dy = 0,-1 }
		if mov == '>' { dx,dy = 1,0 }
		if mov == 'v' { dx,dy = 0,1 }
		if mov == '<' { dx,dy = -1,0 }
		x1, y1 := rx+dx, ry+dy
		for wh[y1][x1] == 'O' { x1, y1 = x1+dx, y1+dy }
		if wh[y1][x1] == '#' { continue }

		for {
			x2, y2 := x1-dx, y1-dy
			wh[y1][x1] = wh[y2][x2]
			if wh[y1][x1] == '@' {
				rx, ry = x1, y1
				wh[y2][x2] = '.'
				break
			}
			x1, y1 = x2, y2
		}
	}

	sum := 0
	for y := 0; y < ylen; y++ {
		for x := 0; x < xlen; x++ {
			if wh[y][x] == 'O' { sum += 100 * y + x }
		}
	}
	

	fmt.Printf("= %d\n", sum)
}

func star2() {
	inBytes,_ := os.ReadFile("15.txt")
	in := string(inBytes)

	wh := make([][]byte, 0)
	movs := make([]byte, 0)
	inState := 0
	for _,line := range strings.Split(in, "\n") {
		if inState == 0 {
			if line == "" {
				inState = 1
				continue
			}
			row := make([]byte, 0)
			for x := 0; x < len(line); x++ {
				if line[x] == '#' {
					row = append(row, '#')
					row = append(row, '#')
				}
				if line[x] == 'O' {
					row = append(row, '[')
					row = append(row, ']')
				}
				if line[x] == '.' {
					row = append(row, '.')
					row = append(row, '.')
				}
				if line[x] == '@' {
					row = append(row, '@')
					row = append(row, '.')
				}
			}
			wh = append(wh, row)
			continue
		}
		for i := 0; i < len(line); i++ {
			movs = append(movs, line[i])
		}
	}
	xlen, ylen := len(wh[0]), len(wh)

	rx, ry := 0, 0
	for y := 0; y < ylen; y++ {
		for x := 0; x < xlen; x++ {
			if wh[y][x] == '@' {
				rx,ry = x,y
			}
		}
	}


	copywh := func(wh [][]byte) [][]byte {
		copy := make([][]byte, ylen)
		for y := 0; y < ylen; y++ {
			rowcopy := make([]byte, xlen)
			for x := 0; x < xlen; x++ {
				rowcopy[x] = wh[y][x]
			}
			copy[y] = rowcopy
		}
		return copy
	}
	
	var domov func(x, y, dx, dy int, wh [][]byte) bool;
	domov = func(x, y, dx, dy int, wh [][]byte) bool {
		x1, y1 := x+dx, y+dy
		c1 := wh[y1][x1]
		if c1 == '#' { return false }
		if c1 == '.' {
			wh[y1][x1] = wh[y][x]
			wh[y][x] = '.'
			return true
		}
		if c1 == '[' || c1 == ']' {
			if !domov(x1, y1, dx, dy, wh) { return false }
			if c1 == '[' && dy != 0 && !domov(x1+1, y1, dx, dy, wh) { return false }
			if c1 == ']' && dy != 0 && !domov(x1-1, y1, dx, dy, wh) { return false }
			wh[y1][x1] = wh[y][x]
			wh[y][x] = '.'
			return true
		}
		log.Fatal("foo")
		return false
	}
	
	for i := 0; i < len(movs); i++ {
		mov := movs[i]
		// fmt.Printf("\n\nr=%d,%d %c\n", rx, ry, mov)
		// for y := 0; y < ylen; y++ {
		// 	for x := 0; x < xlen; x++ {
		// 		fmt.Printf("%c", wh[y][x])
		// 	}
		// 	fmt.Printf("\n")
		// }

		dx, dy := 0, 0
		if mov == '^' { dx,dy = 0,-1 }
		if mov == '>' { dx,dy = 1,0 }
		if mov == 'v' { dx,dy = 0,1 }
		if mov == '<' { dx,dy = -1,0 }

		wh2 := copywh(wh)
		if domov(rx, ry, dx, dy, wh2) {
			domov(rx, ry, dx, dy, wh)
			rx, ry = rx+dx, ry+dy
		}
	}
	
	// for y := 0; y < ylen; y++ {
	// 	for x := 0; x < xlen; x++ {
	// 		fmt.Printf("%c", wh[y][x])
	// 	}
	// 	fmt.Printf("\n")
	// }

	sum := 0
	for y := 0; y < ylen; y++ {
		for x := 0; x < xlen; x++ {
			if wh[y][x] == '[' { sum += 100 * y + x }
		}
	}
	

	fmt.Printf("= %d\n", sum)
}

func main() {
	star1()
	star2()
}
