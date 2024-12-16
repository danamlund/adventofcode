// 09:47
package main
import "fmt"
import "os"
import "strings"

type xy struct {
	x, y int
}

func star1() {
	inBytes,_ := os.ReadFile("16.txt")
	in := string(inBytes)
	map1 := make([][]byte, 0)
	for _,line := range strings.Split(in, "\n") {
		if line == "" { continue }
		row := make([]byte, len(line))
		for i := 0; i < len(line); i++ {
			row[i] = line[i]
		}
		map1 = append(map1, row)
	}
	xlen, ylen := len(map1[0]), len(map1)

	posses := 0
	
	s, e := xy{}, xy{}
	for y := 0; y < ylen; y++ {
		for x := 0; x < xlen; x++ {
			if map1[y][x] == 'S' { s = xy{x,y} }
			if map1[y][x] == 'E' { e = xy{x,y} }
			//fmt.Printf("%c", map1[y][x])
			if map1[y][x] != '#' { posses++ }
		}
		//fmt.Printf("\n")
	}

	dirs := [...]xy {{1,0},{0,1},{-1,0},{0,-1}}

	copymap := func(in map[xy]bool) map[xy]bool {
		copy := make(map[xy]bool)
		for k,v := range in {
			copy[k] = v
		}
		return copy
	}

	rotateleft := func(d xy) xy { return xy{ -d.y, d.x } }
	rotateright := func(d xy) xy { return xy{ d.y, -d.x } }

	posseens := make(map[xy]int)

	posbestscore := make(map[xy]int)
	
	var path func(p xy, d xy, seens map[xy]bool, score int) int
	path = func(p xy, d xy, seens map[xy]bool, score int) int {
		if _,ok := posbestscore[p]; ok && score > posbestscore[p] {
			return 99999999999999
		}
		posbestscore[p] = score
		oldlen := len(posseens)
		posseens[p]++
		if len(posseens) != oldlen {
			//fmt.Printf("### seens %v\n", posseens)
			fmt.Printf("### seens %v/%v\n", len(posseens), posses)
		}
		if p == e { return score }
		if len(posseens) >= 1393 {
			//fmt.Printf("## %v %v score=%v\n", p, d, score)
		}
		scorebest := 99999999
		seenscopy := copymap(seens)
		seenscopy[p] = true
		for _, d2 := range dirs {
			n := xy{p.x+d2.x, p.y+d2.y}
			if map1[n.y][n.x]== '#' { continue }
			if d != d2 && rotateleft(d) != d2 && rotateright(d) != d2 { continue }
			if _,ok := seens[n]; ok { continue }
			scorenew := score
			if d2 != d { scorenew += 1000 }
			scorenew += 1
			score2 := path(n, d2, seenscopy, scorenew)
			if score2 < scorebest { scorebest = score2 }
		}
		return scorebest
	}

	score := path(s, xy{1,0}, make(map[xy]bool), 0)
	fmt.Printf("= %v\n", score)
}

func star2() {
	inBytes,_ := os.ReadFile("16.txt")
	in := string(inBytes)
	map1 := make([][]byte, 0)
	for _,line := range strings.Split(in, "\n") {
		if line == "" { continue }
		row := make([]byte, len(line))
		for i := 0; i < len(line); i++ {
			row[i] = line[i]
		}
		map1 = append(map1, row)
	}
	xlen, ylen := len(map1[0]), len(map1)

	posses := 0
	
	s, e := xy{}, xy{}
	for y := 0; y < ylen; y++ {
		for x := 0; x < xlen; x++ {
			if map1[y][x] == 'S' { s = xy{x,y} }
			if map1[y][x] == 'E' { e = xy{x,y} }
			//fmt.Printf("%c", map1[y][x])
			if map1[y][x] != '#' { posses++ }
		}
		//fmt.Printf("\n")
	}

	dirs := [...]xy {{1,0},{0,1},{-1,0},{0,-1}}

	// copymap := func(in map[xy]bool) map[xy]bool {
	// 	copy := make(map[xy]bool)
	// 	for k,v := range in {
	// 		copy[k] = v
	// 	}
	// 	return copy
	// }

	rotateleft := func(d xy) xy { return xy{ -d.y, d.x } }
	rotateright := func(d xy) xy { return xy{ d.y, -d.x } }

	posseens := make(map[xy]int)
	type state struct { p, d xy }
	statebestscores := make(map[state]int)

	bestseens := make(map[xy]bool)
	bestscore := 99999999999999

	var path func(p xy, d xy, seens map[xy]bool, score int) int
	path = func(p xy, d xy, seens map[xy]bool, score int) int {
		st := state{p, d}
		if _,ok := statebestscores[st]; ok && score > statebestscores[st] {
			return 99999999999999
		}
		statebestscores[st] = score
		oldlen := len(posseens)
		posseens[p]++
		if len(posseens) != oldlen {
			//fmt.Printf("### seens %v\n", posseens)
			//fmt.Printf("### seens %v/%v\n", len(posseens), posses)
		}
		if p == e {
			if score == bestscore {
				//fmt.Printf("## %v %v score=%v\n", p, d, score)
				for k,v := range seens {
					bestseens[k] = v
				}
			}
			return score
		}
		if len(posseens) >= 1393 {
			//fmt.Printf("## %v %v score=%v\n", p, d, score)
		}
		scorebest := 99999999
		seens[p] = true
		for _, d2 := range dirs {
			n := xy{p.x+d2.x, p.y+d2.y}
			if map1[n.y][n.x]== '#' { continue }
			if d != d2 && rotateleft(d) != d2 && rotateright(d) != d2 { continue }
			if _,ok := seens[n]; ok { continue }
			scorenew := score
			if d2 != d { scorenew += 1000 }
			scorenew += 1
			score2 := path(n, d2, seens, scorenew)
			if score2 < scorebest { scorebest = score2 }
		}
		delete(seens, p)
		return scorebest
	}

	bestscore = path(s, xy{1,0}, make(map[xy]bool), 0)
	path(s, xy{1,0}, make(map[xy]bool), 0)
	
	fmt.Printf("= %v\n", len(bestseens) + 1)
}

func main() {
	star1()
	star2()
}
