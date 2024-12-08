package main
import "fmt"
import "os"
import "strings"

type p struct {
	freq byte
	y int
	x int
}

type pt struct {
	x int
	y int
}

func star1() {
	inb,_ := os.ReadFile("08.txt")
	in := string(inb)

	ps := make([]p, 0)
	maxx := 0
	maxy := 0
	
	for y,line := range strings.Split(in, "\n") {
		if line == "" { continue }
		for x := 0; x < len(line); x++ {
			if line[x] != '.' {
				ps = append(ps, p{freq: byte(line[x]), y: y, x: x})
			}
			maxx = x
		}
		maxy = y
	}

	antinodes := make(map[pt]bool)
	for i := 0; i < len(ps); i++ {
		p1 := ps[i]
		for j := i+1; j < len(ps); j++ {
			p2 := ps[j]
			if p1.freq != p2.freq { continue }
			dx, dy := p2.x - p1.x, p2.y - p1.y
			x1, y1 := p1.x - dx, p1.y - dy
			if x1 >= 0 && x1 <= maxx && y1 >= 0 && y1 <= maxy {
				antinodes[pt{x:x1, y:y1}] = true
			}
			x2, y2 := p2.x + dx, p2.y + dy
			if x2 >= 0 && x2 <= maxx && y2 >= 0 && y2 <= maxy {
				antinodes[pt{x:x2, y:y2}] = true
			}
		}
	}

	fmt.Printf("= %v\n", len(antinodes))
}

func star2() {
	inb,_ := os.ReadFile("08.txt")
	in := string(inb)

	ps := make([]p, 0)
	maxx := 0
	maxy := 0
	
	for y,line := range strings.Split(in, "\n") {
		if line == "" { continue }
		for x := 0; x < len(line); x++ {
			if line[x] != '.' {
				ps = append(ps, p{freq: byte(line[x]), y: y, x: x})
			}
			maxx = x
		}
		maxy = y
	}

	antinodes := make(map[pt]bool)
	for i := 0; i < len(ps); i++ {
		p1 := ps[i]
		for j := i+1; j < len(ps); j++ {
			p2 := ps[j]
			if p1.freq != p2.freq { continue }
			dx, dy := p2.x - p1.x, p2.y - p1.y
			x1, y1 := p1.x, p1.y
			for x1 >= 0 && x1 <= maxx && y1 >= 0 && y1 <= maxy {
				antinodes[pt{x:x1, y:y1}] = true
				x1 -= dx
				y1 -= dy
			}
			x2, y2 := p2.x, p2.y
			for x2 >= 0 && x2 <= maxx && y2 >= 0 && y2 <= maxy {
				antinodes[pt{x:x2, y:y2}] = true
				x2 += dx
				y2 += dy
			}
		}
	}

	fmt.Printf("= %v\n", len(antinodes))
}

func main() {
	star1()
	star2()
}
