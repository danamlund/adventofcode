package main
import "fmt"
import "strings"
import "os"

func star1() {
	inBytes,_ := os.ReadFile("06.txt")
	in := string(inBytes)

	area := make([][]byte, 0)
	
	for _,line := range strings.Split(in, "\n") {
		if strings.TrimSpace(line) != "" {
			area = append(area, []byte(strings.TrimSpace(line)))
		}
	}

	ylen := len(area)
	xlen := len(area[0])

	

	gx, gy := -1,-1
	for y := 0; y < ylen; y++ {
		for x := 0; x < xlen; x++ {
			//fmt.Printf("%c", rune(area[y][x]))
			if area[y][x] == '^' { gx, gy = x, y }
		}
		//fmt.Printf("\n")
	}
	gxd := 0
	gyd := -1
	area[gy][gx] = 'X'

	for gx+gxd >= 0 && gx+gxd < xlen && gy+gyd >= 0 && gy+gyd < ylen {
		if area[gy+gyd][gx+gxd] == '#' {
			gxd, gyd = -gyd, gxd
		} else {
			gx += gxd
			gy += gyd
			area[gy][gx] = 'X'
		}
	}

	// fmt.Printf("\n")
	// for y := 0; y < ylen; y++ {
	// 	for x := 0; x < xlen; x++ {
	// 		fmt.Printf("%c", rune(area[y][x]))
	// 	}
	// 	fmt.Printf("\n")
	// }

	sum := 0
	for y := 0; y < ylen; y++ {
		for x := 0; x < xlen; x++ {
			if area[y][x] == 'X' { sum++ }
		}
	}
	fmt.Printf("= %v\n", sum)
	// 5161 too low
}

func star2() {
	inBytes,_ := os.ReadFile("06.txt")
	in := string(inBytes)

	area := make([][]byte, 0)
	
	for _,line := range strings.Split(in, "\n") {
		if strings.TrimSpace(line) != "" {
			area = append(area, []byte(strings.TrimSpace(line)))
		}
	}

	ylen := len(area)
	xlen := len(area[0])

	

	sx, sy := -1,-1
	for y := 0; y < ylen; y++ {
		for x := 0; x < xlen; x++ {
			//fmt.Printf("%c", rune(area[y][x]))
			if area[y][x] == '^' { sx, sy = x, y }
		}
		//fmt.Printf("\n")
	}

	loops := 0
	for oy := 0; oy < ylen; oy++ {
		for ox := 0; ox < xlen; ox++ {
			if area[oy][ox] != '.' { continue }
			area[oy][ox] = '#'
			gx, gy := sx, sy
			gxd := 0
			gyd := -1
			area[gy][gx] = 'X'
			steps := 0
			for steps < ylen*xlen && gx+gxd >= 0 && gx+gxd < xlen && gy+gyd >= 0 && gy+gyd < ylen {
				if area[gy+gyd][gx+gxd] == '#' {
					gxd, gyd = -gyd, gxd
				} else {
					gx += gxd
					gy += gyd
					area[gy][gx] = 'X'
				}
				steps++
			}
			if steps == ylen*xlen { loops++ }
			for y := 0; y < ylen; y++ {
				for x := 0; x < xlen; x++ {
					if area[y][x] != '#' { area[y][x] = '.' }
				}
			}
			area[oy][ox] = '.'
		}
	}
	

	// fmt.Printf("\n")
	// for y := 0; y < ylen; y++ {
	// 	for x := 0; x < xlen; x++ {
	// 		fmt.Printf("%c", rune(area[y][x]))
	// 	}
	// 	fmt.Printf("\n")
	// }

	fmt.Printf("= %v\n", loops)
}

func main() {
	star1()
	star2()
}
