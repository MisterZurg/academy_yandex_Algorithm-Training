/*
Телефонные номера в адресной книге мобильного телефона имеют один из следующих форматов:
+7<код><номер>
8<код><номер>
<номер>
где <номер> — это семь цифр, а <код> — это три цифры или три цифры в круглых скобках.
Если код не указан, то считается, что он равен 495.
Кроме того, в записи телефонного номера может стоять знак “-” между любыми двумя цифрами(см. пример).
На данный момент в адресной книге телефона Васи записано всего три телефонных номера,
и он хочет записать туда еще один. Но он не может понять,
не записан ли уже такой номер в телефонной книге.
Помогите ему! Два телефонных номера совпадают, если у них равны коды и равны номера.
Например, +7(916)0123456 и 89160123456 — это один и тот же номер.

Формат ввода
В первой строке входных данных записан номер телефона, который Вася хочет добавить в адресную книгу своего телефона. В следующих трех строках записаны три номера телефонов, которые уже находятся в адресной книге телефона Васи. Гарантируется, что каждая из записей соответствует одному из трех приведенных в условии форматов.

Формат вывода
Для каждого телефонного номера в адресной книге выведите YES (заглавными буквами), если он совпадает с тем телефонным номером, который Вася хочет добавить в адресную книгу или NO (заглавными буквами) в противном случае.
*/
package main

import (
	"fmt"
	"strings"
)

func main() {
	var willingToAddThisNumber string
	fmt.Scan(&willingToAddThisNumber)
	// fmt.Println(parseNumber(willingToAddThisNumber))
	phonebook := make([]string, 3)

	for phoneInd := range phonebook {
		fmt.Scan(&phonebook[phoneInd])
	}
	for _, phone := range phonebook {
		cadAddPhone(willingToAddThisNumber, phone)
	}

}

func cadAddPhone(phoneNumber, phoneFromBook string) {
	parsedPhone := parseNumber(phoneNumber)
	parsedPhoneFromBook := parseNumber(phoneFromBook)
	//fmt.Println(parsedPhone , parsedPhoneFromBook)
	if parsedPhone == parsedPhoneFromBook {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

func parseNumber(phoneNumber string) string {
	// Remove -
	number := strings.ReplaceAll(phoneNumber, "-", "")
	// Remove ( )
	number = strings.Replace(number, "(", "", 1)
	number = strings.Replace(number, ")", "", 1)
	// fmt.Println(number)
	// +7<код><номер>
	if len(number) == 12 {
		return number
	} else if len(number) == 11 { // 8<код><номер>
		return "+7" + number[1:]
	}
	return "+7495" + number
}
