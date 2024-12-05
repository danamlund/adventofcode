package main
import "fmt"
import "os"

func star1() {
	inBytes,_ := os.ReadFile("03.txt")
	in := string(inBytes)

	sum := 0
	for i := 0; i < len(in); i++ {
		var num1, num2 int;
		n,err := fmt.Sscanf(in[i:], "mul(%d,%d)", &num1, &num2);
		if n == 2 && err == nil && num1 >= 0 && num1 <= 999 && num2 >= 0 && num2 <= 999 {
			//fmt.Printf("i=%v mul(%v,%v)\n", i, num1, num2)
			sum += num1*num2;
		}
	}
	fmt.Printf("= %v\n", sum)
	// 4431252 too low
}

func star2() {
	inBytes,_ := os.ReadFile("03.txt")
	in := string(inBytes)

	sum := 0
	enabled := true
	for i := 0; i < len(in); i++ {
		_,err := fmt.Sscanf(in[i:], "do()");
		//fmt.Println("##1 ", i, err, in[i:])
		if err == nil {
			enabled = true
		}
		_,err = fmt.Sscanf(in[i:], "don't()");
		//fmt.Println("##2 ", i, err, in[1:])
		if err == nil {
			enabled = false
		}
		if enabled {
			var num1, num2 int;
			n,err := fmt.Sscanf(in[i:], "mul(%d,%d)", &num1, &num2);
			if n == 2 && err == nil && num1 >= 0 && num1 <= 999 && num2 >= 0 && num2 <= 999 {
				//fmt.Printf("i=%v mul(%v,%v)\n", i, num1, num2)
				sum += num1*num2;
			}
		}
	}
	fmt.Printf("= %v\n", sum)
	// 4431252 too low
}

func main() {
	//star1()
	star2()
}
