/*
Задача №2

Вход:
Пользователь должен ввести "правильный пароль", состоящий из:
цифр, букв латинского алфавита(строчные и прописные) и
специальных символов  special = "_!@#$%^&"

Всего 4 набора различных символов.
В пароле обязательно должен быть хотя бы один символ из каждого набора.
Длина пароля от 8(мин) до 15(макс) символов.
Максимальное количество попыток ввода неправильного пароля - 5.
Каждый раз выводим номер попытки.
*Желательно выводить пояснение, почему пароль не принят и что нужно исправить.

digits = "0123456789"
lowercase = "abcdefghiklmnopqrstvxyz"
uppercase = "ABCDEFGHIKLMNOPQRSTVXYZ"
special = "_!@#$%^&"

Выход:
Написать, что пароль правильный и он принят.

Пример:
хороший пароль -> o58anuahaunH!
хороший пароль -> aaaAAA111!!!
плохой пароль -> saucacAusacu8

Реализацию оформить через функцию:
1. checkPassword(pass string) (bool, errors <- на усмотрение)
*/
package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Ваш код

func main() {
	var pass string
	for i := range 5 {
		fmt.Scan(&pass)
		result, errorMsg := checkPassword(pass)
		if result != true {
			fmt.Println("У вас сталось", 4-i, "попытки")
			fmt.Println(errorMsg)
		} else {
			fmt.Println("Пароль верный и он принят")
			break
		}
	}
}

// функция проверки пароля на соотвествие условиям надёжного пароля
func checkPassword(pass string) (result bool, errorMsg string) {
	var digits, lwLetters, upLetters, specials bool
	var digitError, lettersLwError, lettersUpError, specialError string
	var digitsArrey []string
	var lettersLwArrey []string
	var lettersUpArrey []string
	specialArrey := strings.Split("_!@#$%^&", "")
	for i := range 10 {
		digitsArrey = append(digitsArrey, strconv.Itoa(i))
	}
	for a, b := 'a', 'A'; a <= 'z'; a, b = a+1, b+1 {
		lettersLwArrey = append(lettersLwArrey, string(a))
		lettersUpArrey = append(lettersUpArrey, string(b))
	}
	passArrey := strings.Split(pass, "")
	for _, symbol := range passArrey {
		for _, digit := range digitsArrey {
			if digits == true {
				break
			}
			if digit == symbol {
				digits = true
				digitError = ""
				break
			} else {
				digitError = "Не хватает цифр"
			}
		}
		for _, lwLetter := range lettersLwArrey {
			if lwLetters == true {
				break
			}
			if lwLetter == symbol {
				lwLetters = true
				lettersLwError = ""
				break
			} else {
				lettersLwError = "Не хватает прописных букв"
			}
		}
		for _, upLetter := range lettersUpArrey {
			if upLetters == true {
				break
			}
			if upLetter == symbol {
				upLetters = true
				lettersUpError = ""
				break
			} else {
				lettersUpError = "Не хватает заглавных букв"
			}
		}
		for _, special := range specialArrey {
			if specials == true {
				break
			}
			if special == symbol {
				specials = true
				specialError = ""
				break
			} else {
				specialError = "Не хватает символов _!@#$%^&"
			}
		}
	}
	if len(passArrey) > 7 && len(passArrey) < 16 && digits && lwLetters && upLetters && specials {
		result = true
	} else if len(passArrey) < 8 {
		result = false
		errorMsg = "Пароль должен быть не меньше 8-ми символов"
	} else if len(passArrey) > 15 {
		result = false
		errorMsg = "Пароль должен быть не более 15-ти символов"
	} else {
		result = false
		errorMsg = fmt.Sprintf("%s\n%s\n%s\n%s", digitError, lettersLwError, lettersUpError, specialError)
	}
	return result, errorMsg
}
