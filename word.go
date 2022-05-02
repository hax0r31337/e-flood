package main

import (
	"bufio"
	"math/rand"
	"os"
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
