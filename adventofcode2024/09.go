package main
import "fmt"
import "os"
import "strings"

func star1() {
	inBytes,_ := os.ReadFile("09.txt")
	in := string(inBytes)

	for _,line := range strings.Split(in, "\n") {
		if line == "" { continue }

		disk := make([]int, 0)
		for i := 0; i < len(line); i++ {
			id := i/2;
			if i % 2 == 1 { id = -1 }
			len := int(line[i] - '0');

			for j := 0; j < len; j++ {
				disk = append(disk, id);
			}
		}

		// for _,v := range disk {
		// 	if v == -1 { fmt.Printf(".")
		// 	} else { fmt.Printf("%v", v) }
		// }
		// fmt.Printf("\n")

		s, e := 0, len(disk)-1
		for s < e {
			if disk[s] != -1 {
				s++
			} else if disk[e] == -1 {
				e--
			} else {
				disk[s], disk[e] = disk[e], -1
				s++
				e--
			}
		}
		// for _,v := range disk {
		// 	if v == -1 { fmt.Printf(".")
		// 	} else { fmt.Printf("%v", v) }
		// }
		// fmt.Printf("\n")

		sum := 0
		for i,v := range disk {
			if v != -1 { sum += i*v }
		}
		fmt.Printf("= %v\n", sum)
	}
}

func star2() {
	inBytes,_ := os.ReadFile("09.txt")
	in := string(inBytes)

	for _,line := range strings.Split(in, "\n") {
		if line == "" { continue }

		maxId := 0
		
		disk := make([]int, 0)
		for i := 0; i < len(line); i++ {
			id := i/2;
			if i % 2 == 1 { id = -1 }
			len := int(line[i] - '0');

			for j := 0; j < len; j++ {
				disk = append(disk, id);
			}
			if id > maxId { maxId = id }
		}

		// for _,v := range disk {
		// 	if v == -1 { fmt.Printf(".")
		// 	} else { fmt.Printf("%v", v) }
		// }
		// fmt.Printf("\n")

		for fileId := maxId; fileId >= 0; fileId-- {
			e := len(disk)-1
			for e > 0 && disk[e] != fileId { e-- }
			fileEnd := e
			fileId := disk[fileEnd]
			for e > 0 && disk[e] == fileId { e-- }
			fileLen := fileEnd - e
			fileStart := fileEnd - fileLen + 1
			
			for s := 0; s < len(disk); {
				for s < len(disk) && disk[s] != -1 { s++ }
				freeStart := s
				for s < len(disk) && disk[s] == -1 { s++ }
				free := s - freeStart
				// fmt.Printf("s=%v,%v e=%v,%v  id=%v\n", freeStart, free, fileStart, fileLen, fileId)
				if fileLen <= free && freeStart < fileStart {
					for i := 0; i < fileLen; i++ {
						disk[freeStart+i], disk[fileStart+i] = fileId, -1
					}
					break
				}
			}
			// for _,v := range disk {
			// 	if v == -1 { fmt.Printf(".")
			// 	} else { fmt.Printf("%v", v) }
			// }
			// fmt.Printf("\n")
		}
		
		// for _,v := range disk {
		// 	if v == -1 { fmt.Printf(".")
		// 	} else { fmt.Printf("%v", v) }
		// }
		// fmt.Printf("\n")

		sum := 0
		for i,v := range disk {
			if v != -1 { sum += i*v }
		}
		fmt.Printf("= %v\n", sum)
	}
}

func main() {
	star1()
	star2()
}
