package main

import "fmt"
import "os"
import "strings"

func star1() {
	inBytes,_ := os.ReadFile("05.txt")
	in := string(inBytes)

	var befores [100][100]bool

	sum := 0
	for _,line := range strings.Split(in, "\n") {
		var a,b int;
		_, err := fmt.Sscanf(line, "%d|%d", &a, &b)
		if err == nil {
			//fmt.Printf("%v, %v\n", a, b)
			befores[a][b] = true
			// for i := 0; i < 100; i++ {
			// 	if befores[i][a] { befores[i][b] = true }
			// }
			// for i := 0; i < 100; i++ {
			// 	if befores[b][i] { befores[a][i] = true }
			// }
		} else if line != "" {
			pages := make([]int, 0);
			for _,page := range strings.Split(line, ",") {
				var pagenum int;
				fmt.Sscanf(page, "%d", &pagenum)
				pages = append(pages, pagenum)
			}

			valid := true
			for i := 0; i < len(pages); i++ {
				for j := i+1; j < len(pages); j++ {
					if befores[pages[j]][pages[i]] { valid = false }
				}
			}
			//fmt.Printf("a %v %v\n", pages, valid)
			if valid {
				sum += pages[len(pages)/2]
			}
		}
	}
	fmt.Printf("= %v\n", sum)
}

func star2() {
	inBytes,_ := os.ReadFile("05.txt")
	in := string(inBytes)

	var befores [100][100]bool

	sum := 0
	for _,line := range strings.Split(in, "\n") {
		var a,b int;
		_, err := fmt.Sscanf(line, "%d|%d", &a, &b)
		if err == nil {
			//fmt.Printf("%v, %v\n", a, b)
			befores[a][b] = true
			// for i := 0; i < 100; i++ {
			// 	if befores[i][a] { befores[i][b] = true }
			// }
			// for i := 0; i < 100; i++ {
			// 	if befores[b][i] { befores[a][i] = true }
			// }
		} else if line != "" {
			pages := make([]int, 0);
			for _,page := range strings.Split(line, ",") {
				var pagenum int;
				fmt.Sscanf(page, "%d", &pagenum)
				pages = append(pages, pagenum)
			}

			valid := true
			for i := 0; i < len(pages); i++ {
				for j := i+1; j < len(pages); j++ {
					if befores[pages[j]][pages[i]] { valid = false }
				}
			}
			//fmt.Printf("a %v %v\n", pages, valid)
			if !valid {
				//fmt.Printf("bad=%v\n", pages)
				for i := 0; i < len(pages); i++ {
					for j := i+1; j < len(pages); j++ {
						if befores[pages[j]][pages[i]] {
							tmp := pages[i]
							pages[i] = pages[j]
							pages[j] = tmp
							i = -1
							break
						}
					}
				}
				//fmt.Printf("fix=%v\n", pages)
				sum += pages[len(pages)/2]
			}
		}
	}
	fmt.Printf("= %v\n", sum)
}

func main() {
	star1()
	star2()
}
