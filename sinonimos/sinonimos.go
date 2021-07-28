package sinonimos

import (
	"fmt"
)

//Función de Try
func Iguales() string {
	var text1 string
	var text2 string
	var aument int = 0
	var certif int = 0
	var i int = 0
	var cont int = 0

	var result string

	//-------Introducción de Datos por usuario
	fmt.Println("Enter first word: ")
	fmt.Scanln(&text1)
	fmt.Println("Enter second word: ")
	fmt.Scanln(&text2)

	var textNumberA int = len(text1)
	var textNumberB int = len(text2)

	if textNumberA == textNumberB {
		for cont < textNumberA {
			if text1[aument] == text2[i] {
				aument++
				certif++
				cont++
				i = 0

			} else {
				i++
			}

			if i >= textNumberA {
				break
			}
		}
	}

	if certif == textNumberA {
		result = "Claramente tiene las mismas letras"
		//fmt.Println(result)
	} else {
		result = "Claramente No tiene las mismas letras"
		//fmt.Println(result)
	}
	return result
}
