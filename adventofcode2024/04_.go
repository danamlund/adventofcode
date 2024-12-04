package main

import "fmt"
import "os"
import "strings"

func star1() {
	inBytes, _ := os.ReadFile("04.txt")
	in := string(inBytes)

	grid := make([]string, 0)
	for _,line := range strings.Split(in, "\n") {
		if strings.TrimSpace(line) != "" {
			grid = append(grid, strings.TrimSpace(line))
		}
	}

	s := "XMAS"
	founds := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			for dx := -1; dx <= 1; dx++ {
				for dy := -1; dy <= 1; dy++ {
					y := i
					x := j
					found := true
					for _,c := range s {
						if x < 0 || x >= len(grid[i]) { found = false; break }
						if y < 0 || y >= len(grid) { found = false; break }
						if rune(grid[y][x]) != c { found = false; break }
						x += dx
						y += dy
					}
					if found {
						founds++;
					}
				}
			}
		}
	}
		
	
	fmt.Printf("= %v\n", founds)
}

func star2() {
	inBytes, _ := os.ReadFile("04.txt")
	in := string(inBytes)

	grid := make([]string, 0)
	for _,line := range strings.Split(in, "\n") {
		if strings.TrimSpace(line) != "" {
			grid = append(grid, strings.TrimSpace(line))
		}
	}

	founds := 0
	for y := 0; y < len(grid)-2; y++ {
		for x := 0; x < len(grid[y])-2; x++ {
			if ((grid[y][x] == 'M' && grid[y+1][x+1] == 'A' && grid[y+2][x+2] == 'S') ||
				(grid[y][x] == 'S' && grid[y+1][x+1] == 'A' && grid[y+2][x+2] == 'M')) &&
				((grid[y][x+2] == 'M' && grid[y+1][x+1] == 'A' && grid[y+2][x] == 'S') ||
					(grid[y][x+2] == 'S' && grid[y+1][x+1] == 'A' && grid[y+2][x] == 'M')) {
				founds++;
			}
		}
	}
		
	
	fmt.Printf("= %v\n", founds)
	// 1955 too low
}

func main() {
	star1()
	star2()
}
