package main

import "fmt"

func main() {
	// case 1) for 문에 초기화 구문, 조건식, 후속 작업 정의
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	// case 2) for 문에 조건식만 사용
	sum2, i2 := 0, 0
	for i2 < 10 {
		sum2 += i2
		i2++
	}
	fmt.Println(sum2)

	// case 3) for 문에 조건식 생략
	sum3, i3 := 0, 0
	for {
		if i3 >= 10 {
			break
		}
		sum3 += i3
		i3++
	}

	fmt.Println(sum3)
}
