// incremet ++   ,   decrement --   POSTFIX x++    PREFIX ++x

package main

import "fmt"

func main() {
	x := 10
	fmt.Println(x)

	x = x + 1 // Sağdaki değer soldaki değişkene atandı.

	fmt.Println(x)

	x++ // burada postfıx sadece değişkenden sonra yazılır. Önce yazılmaz.
	// ++x yani prefıx ifedesi şeklinde kullanılmaz.
	fmt.Println(x)
	//fmt.Println(x++)  yazamam hata verir. Bir satırda iki tane statement bulunmaz.

	//Statement   ve Expression
	// Statement : Programlamada bişey yapılmasını belirten en küçük ifadedir.
	// Print "hello" ,  x=1

	// Expression ise örneğin 5*5 gibi ifadelerdir.

	//Print(5*5) dersek eğer bu statement expression olur.

	y := 10

	y = y - 1
	fmt.Println(y)
	y-- // decrement
	fmt.Println(y)
}
