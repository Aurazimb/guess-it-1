package main

import (
	"bufio"
	"fmt"
	"main/stat"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	numbers := []float64{}
	for scanner.Scan() {
		input := scanner.Text()
		num, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println(err)
		} else {
			numbers = append(numbers, float64(num))
		}
		mode := stat.FindMode(numbers)
		fmt.Println(int(mode), int(mode))
	}
}
// В данном случае, в моем коде решение идет через моду(самое повторяющееся значение)
// он во всех случаях побеждает компьютер