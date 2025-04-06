package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	inputFile := "input.txt"
	outputFile := "output.txt"

	// Чтение входного файла
	data, err := os.ReadFile(inputFile)
	if err != nil {
		panic("Ошибка при открытии входного файла: " + err.Error())
	}

	// Компиляция регулярного выражения
	re := regexp.MustCompile(`^(\d+)([+-/*/])(\d+)=\?$`)

	// Создание/очистка выходного файла
	f, err := os.OpenFile(outputFile, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		panic("Ошибка при создании выходного файла: " + err.Error())
	}
	defer f.Close()

	writer := bufio.NewWriter(f)
	defer writer.Flush()

	// Обработка каждой строки
	for _, line := range strings.Split(string(data), "\n") {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		matches := re.FindStringSubmatch(line)
		if matches == nil {
			continue
		}

		a, _ := strconv.Atoi(matches[1])
		b, _ := strconv.Atoi(matches[3])
		operator := matches[2]

		var result int
		switch operator {
		case "+":
			result = a + b
		case "-":
			result = a - b
		case "*":
			result = a * b
		case "/":
			result = a / b
		default:
			continue
		}

		// Формирование и запись результата
		output := fmt.Sprintf("%s%s%s=%d\n", matches[1], operator, matches[3], result)
		_, err = writer.WriteString(output)
		if err != nil {
			panic("Ошибка записи: " + err.Error())
		}
	}
}
