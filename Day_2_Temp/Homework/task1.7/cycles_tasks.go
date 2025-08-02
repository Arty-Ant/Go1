/*
=======
Задачи:
=======

3.7 Вводим строку без знаков препинания(5 слов).

	Найти самое длинное слово в строке и вывести сколько в нем букв.

Пример:
In: Скажи как дела в учебе
Out: учебе, 5

In: Закрепляем циклы в языке Golang
Out: Закрепляем, 10
*/
package main

// Ваш код

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	var max int = 0
	var word, maxWord, sentence string
	fmt.Println("Введите предложение из 5-ти слов без знаков препинания:")
	for i := 0; i < 5; i++ {
		fmt.Scan(&word)
		sentence = sentence + " " + word
	}
	wordsArray := strings.Split(sentence, " ")
	for _, word := range wordsArray {
		len := utf8.RuneCountInString(word)
		if len < max {
			break
		}
		max = len
		maxWord = word
	}
	fmt.Println(maxWord, ", ", max)
}
