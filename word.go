package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

var word []string

func LoadWord() {
	word = ReadLines("word.txt")
}

func GetWord() string {
	return word[rand.Intn(len(word))]
}

func ReadLines(filePath string) []string {
	file, err := os.Open(filePath)
	if err != nil {
		return make([]string, 0)
	}
	defer file.Close()                // 关闭文本流
	scanner := bufio.NewScanner(file) // 读取文本数据
	file_data := make([]string, 0)
	for scanner.Scan() {
		file_data = append(file_data, scanner.Text())
	}

	return file_data
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyz0123456789")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func solve(s string) string {
	repl := strings.NewReplacer("之和", "", "之差", "", "加", "+", "减", "-", "乘以", "*", "除以", "/", "等于", "=", "一", "1", "二", "2", "三", "3", "四", "4", "五", "5", "六", "6", "七", "7", "八", "8", "九", "9", "零", "0", "十", "10")
	s = strings.Split(repl.Replace(s), "=")[0]

	calc := func(f1 string, f2 string, op string) string {
		n1, err := strconv.Atoi(f1)
		n2, err1 := strconv.Atoi(f2)
		if err != nil || err1 != nil {
			fmt.Println("calc:", f1, f2, op, err, err1)
			return "1"
		}
		switch op {
		case "+":
			return strconv.Itoa(n1 + n2)
		case "-":
			return strconv.Itoa(n1 - n2)
		case "*":
			return strconv.Itoa(n1 * n2)
		case "/":
			return strconv.Itoa(n1 / n2)
		}
		return "1"
	}

	// calculate for answer, example: 1+5*3=?
	field1 := ""
	field2 := ""
	op := ""
	for _, v := range s {
		if v == ' ' || v == '\t' {
			continue
		}
		if v == '+' || v == '-' || v == '*' || v == '/' {
			if op != "" {
				field1 = calc(field1, field2, op)
				field2 = ""
			}
			op = string(v)
		} else {
			if op == "" {
				field1 += string(v)
			} else {
				field2 += string(v)
			}
		}
	}

	res := calc(field1, field2, op)
	fmt.Println(s + "=" + res)
	return res
}
