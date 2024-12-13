package main
import "fmt"
import "os"
import "strings"

func star1() {
	inBytes,_ := os.ReadFile("13.txt")
	in := string(inBytes)

	sum := 0
	ax, ay, bx, by, px, py := 0, 0, 0, 0, 0, 0
	for _,line := range strings.Split(in, "\n") {
		fmt.Sscanf(line, "Button A: X%d, Y%d", &ax, &ay);
		fmt.Sscanf(line, "Button B: X%d, Y%d", &bx, &by);
		fmt.Sscanf(line, "Prize: X=%d, Y=%d", &px, &py);
		if line == "" {
			//fmt.Printf("! %d,%d %d,%d %d,%d\n", ax, ay, bx, by, px, py)

			best := 9999
			for as := 0; as < 100; as++ {
				for bs := 0; bs < 100; bs++ {
					if as == 80 && bs == 40 {
						// fmt.Printf("!! %d == %d && %d == %d \n", as*ax + bs*bx, px, as*ay + bs*by, py)
					}
					if as*ax + bs*bx == px && as*ay + bs*by == py {
						cost := 3*as + 1*bs
						if cost < best { best = cost }
					}
				}
			}
			if best < 9999 {
				sum += best
			}
		}
	}

	fmt.Printf("= %d\n", sum);
}

func star2() {
	inBytes,_ := os.ReadFile("13.txt")
	in := string(inBytes)

	sum := 0
	ax, ay, bx, by, px, py := 0, 0, 0, 0, 0, 0
	for _,line := range strings.Split(in, "\n") {
		fmt.Sscanf(line, "Button A: X%d, Y%d", &ax, &ay);
		fmt.Sscanf(line, "Button B: X%d, Y%d", &bx, &by);
		fmt.Sscanf(line, "Prize: X=%d, Y=%d", &px, &py);
		if line == "" {
			px += 10000000000000
			py += 10000000000000
			//fmt.Printf("! %d,%d %d,%d %d,%d\n", ax, ay, bx, by, px, py)

			// as*ax + bs*bx == px
			// as*ay + bs*by == py

			// as = (px - bs*bx)/ax
			// ((px - bs*bx)/ax)*ay + bs*by == py
			// ay*px/ax - ay*bs*bx/ax + bs*by = py
			// ay*px/ax - py = ay*bs*bx/ax - bs*by
			// ay*px/ax - py = bs*(ay*bx/ax) - bs*by
			// ay*px/ax - py = bs*((ay*bx/ax) - by)
			// (ay*px/ax - py) / ((ay*bx/ax) - by) = bs

			bs1 := (float64(ay)*float64(px)/float64(ax) - float64(py)) / ((float64(ay)*float64(bx)/float64(ax)) - float64(by))
			//as1 := (px - bs*bx)/ax

			// bsmin := 0
			// bsmax := 1000
			// for {
			// 	bs := bsmax
			// 	foo := (ax-(bs*bx))/ax - (py-(bs*by))/ay
			// 	if foo < 0 { bsmax *= 10
			// 	} else { break }
			// }

			// bs := 0
			// for {
			// 	bs = (bsmin + bsmax)/2
			// 	foo := (ax-(bs*bx))/ax - (py-(bs*by))/ay
			// 	fmt.Printf("## %d..%d  %d = %d\n", bsmin, bsmax, bs, foo)
			// 	if foo == 0 || bsmin <= bsmax - 1 {
			// 		break
			// 	} else if foo < 0 { bsmin = bs
			// 	} else if foo > 0 { bsmax = bs 
			// 	}
			// }

			// as := px/ax - (bs*bx)/ax
			//fmt.Printf("##1 a=%d b=%d\n", as, bs)

			//			8757495063522

			for badd := -100000; badd <= 100000; badd++ {
				bs := int(bs1) + badd
				as := (px - bs*bx)/ax
				if as*ax + bs*bx == px && as*ay + bs*by == py {
					//fmt.Printf("##2 a=%d b=%d\n", as, bs)
					sum += 3*as + 1*bs
					break;
				}
			}
			
			
			// best := 9999
			// for as1 := 0; as1 < 100000; as1++ {
			// 	for bs1 := 0; bs1 < 100000; bs1++ {
			// 		as := as1*gcda
			// 		bs := bs1*gcdb
			// 		if as*ax + bs*bx == px && as*ay + bs*by == py {
			// 			cost := 3*as + 1*bs
			// 			if cost < best { best = cost }
			// 			fmt.Printf("a=%d b=%d\n", as, bs)
			// 		}
			// 	}
			// }
			// if best < 9999 {
			// 	sum += best
			// }
		}
	}

	fmt.Printf("= %d\n", sum);
	// 8757495063522 too low
}

func main() {
	star1()
	star2()
}
