/*
Задача №1
Написать функцию, которая расшифрует строку.
code = "220411112603141304"
Каждые две цифры - это либо буква латинского алфавита в нижнем регистре либо пробел.
Отчет с 00 -> 'a' и до 25 -> 'z', 26 -> ' '(пробел).
Вход: строка из цифр. Выход: Текст.
Проверка работы функции выполняется через другую строку.

Задача №1.1
Реализовать и функцию зашифровки

codeToString(code) -> "???????'

stringToCode("hello") -> "??????"
*/
package main

// Ваш код

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var code, result, word, newResult string
	fmt.Println("Введите код для расшифровки:")
	fmt.Scan(&code) //считываем код
	if len(code)%2 != 0 {
		fmt.Println("Введите чётное количество чисел")
		return
	}
	codeArrey := strings.Split(code, "") //преобразовываем код в массив
	//fmt.Println(codeArrey)
	mapper := make(map[string]string) //создание мапы
	//генерируем мапу в стиле [00:a] для перевода чисел в буквы
	for r, i := rune('a'), 0; i <= 26; r, i = r+1, i+1 {
		mapper[fmt.Sprintf("%02d", i)] = string(r)
		if i == 26 {
			mapper[fmt.Sprintf("%02d", i)] = " "
		}
	}
	//fmt.Println(mapper)
	//Берём первое значение мапы и сравниваем со всеми слайсами через цикл (поэлементно)
	for a, b := 0, 2; b <= len(codeArrey); a, b = a+2, b+2 { //цикл на генерацию 2-х значений для слайсов
		for key := range mapper { //цикл мапы с объявлением переменной ключа "00"
			var codeKey string
			var subResult string
			codeSlice := codeArrey[a:b] //взятия срезов из массива кода на базе значений из цикла
			//fmt.Println("Срез кода:", codeSlice)
			keyArrey := strings.Split(key, "") //преобразование ключа в массив
			//fmt.Println("Массив ключа:", keyArrey)
			for index, value := range keyArrey { //цикл сравнения значений слайса кода и массива ключа
				//fmt.Println("Значение ключа:", value)
				//fmt.Println("Значение среза кода:", codeSlice[index])
				if codeSlice[index] == value { //позначимое сравнение значений массива и слайса через индекс
					codeKey += value
					if len(strings.Split(codeKey, "")) == 2 {
						subResult = codeKey
					} //else {
					//fmt.Println("Нет совпадения")
					//}
				}
			}
			result += mapper[subResult]
		}
	}
	fmt.Println(result)

	fmt.Println("А теперь введите слово маленькими латинскими буквами для зашифровки")
	input := bufio.NewScanner(os.Stdin)
	if input.Scan() {
		word = input.Text()
	}
	wordArrey := strings.Split(word, "") //делаем массив из слова
	mapper2 := make(map[string]string)   //создание мапы для шифрования
	//генерируем мапу в стиле [a:00] для перевода букв в числа
	for r, i := rune('a'), 0; i <= 26; r, i = r+1, i+1 {
		mapper2[string(r)] = fmt.Sprintf("%02d", i)
		if i == 26 {
			mapper2[" "] = fmt.Sprintf("%02d", i)
		}
	}
	//fmt.Println(mapper2)
	//Берём первое значение массива букв и сравниваем со значениями мапы
	for _, letter := range wordArrey { //цикл вывода букв
		for key := range mapper2 { //цикл мапы с объявлением переменной значения "a"
			//fmt.Println(mapper2[key])
			//fmt.Println(key)
			//fmt.Println(letter)
			if letter == key {
				newResult += mapper2[key]
				//fmt.Println(newResult)
			}
		}
	}
	fmt.Println("Зашифрованное послание:", newResult)
}
