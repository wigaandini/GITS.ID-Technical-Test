package main
import ("fmt")

func main() {
	var num int
	fmt.Scanf("%d", &num)

	for i := 0; i < num; i++ {
		fmt.Print((i*i + i + 2)/2)
		if i != num-1 {
			fmt.Print("-")
		} else {
			fmt.Println()
		}
	}
}