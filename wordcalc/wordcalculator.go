package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"./cw"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Введите выражение в формате \"123\" + \"5\" или \"123\" * 5: ")
	text, err := reader.ReadString('\n')
	if err != nil {
		panic(fmt.Sprintf("ошибка при чтении ввода: %v", err))
	}
	text = strings.TrimSpace(text)

	var operator string
	if strings.Contains(text, " + ") {
		operator = " + "
	} else if strings.Contains(text, " - ") {
		operator = " - "
	} else if strings.Contains(text, " / ") {
		operator = " / "
	} else if strings.Contains(text, " * ") {
		operator = " * "
	} else {
		panic("оператор не найден")
	}

	zxc := strings.SplitN(text, operator, 2)
	if len(zxc) != 2 {
		panic("введено некорректное значение")
	}

	word1 := strings.TrimSpace(zxc[0])
	word2 := strings.TrimSpace(zxc[1])

	result := cw.CalcWord(word1, word2, operator)
	fmt.Printf("%q", result)
}
