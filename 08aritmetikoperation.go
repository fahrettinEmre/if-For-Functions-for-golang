//addition 10 ve 15 operand ismini alır, + ise operator ismini alır.

package main

import "fmt"

func main() {
	x, y := 15, 10 //int olduğu için sonuclar hep int deönecektir.
	fmt.Printf("%T, %v\n", x, x)
	fmt.Printf("%T, %v\n", y, y)
	fmt.Printf("%T, %v\n", (x + y), (x + y))
	fmt.Printf("%T, %v\n", (x - y), (x - y))
	fmt.Printf("%T, %v\n", (x * y), (x * y))
	fmt.Printf("%T, %v\n", (x / y), (x / y))

	// x:=50 x/=10 dersek sonuc 5 döner.

	z := 5 / 2 // int olduğu için sonuc 2 olarak döner .
	fmt.Printf("%T, %v\n", z, z)

	z2 := float64(5 / 2) // burada tip float64 olur ama sonuc olarak 2 verir.
	fmt.Printf("%T, %v\n", z2, z2)

	z3 := 5.0 / 2 // değişken olarak yazmadığımız için float cinsinden değeri verir.
	fmt.Printf("%T, %v\n", z3, z3)

	//a, b := 15.0, 10     burada değişken olarak atadığımız için farklı tiplerde aritmetik işlemler yapılmaz.
	//ama yukarıda direk değer verdiğimiz için işlem sonucunu yazdırdı.
	//fmt.Printf("%T, %v\n", (a / b), (a / b))

	// reaminder operatörü  % kalan değeri verir.

	fmt.Printf("%T, %v\n", (x % y), (x % y)) //15in 10 ile bölümünden kalan sonucu verir.

}
