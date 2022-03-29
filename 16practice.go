package main

import "fmt"

func main() {
	x := 0
	for x < 10 { // diğer programlardaki while yapısı ile aynı.

		fmt.Println(x)
		x++
	}

	// Switch fallthrough ifadesi

	switch b := 25; {
	case b < 20:
		fmt.Printf("%d küçüktür 20\n", b)
		fallthrough // switch yapısında diğer case lere bakmayı sağlar.

	case b < 50:
		fmt.Printf("%d küçüktür 50\n", b)
		fallthrough

	case b < 100:
		fmt.Printf("%d küçüktür 100\n", b)
		fallthrough

	case b < 200:
		fmt.Printf("%d küçüktür 200\n", b)
	}
	// aşağıdaki if yapısını daha ideomatic halde yazınız..
	// ideomatic : sade,basit görünüm,bağımsız kod,kolay okunabilirlik
	/* if c := 20; c%2 == 0 {
		fmt.Println(x, "çifttir.")
	} else {
		fmt.Println(x, "tektir.")
	} */
	c := 20 // gereksiz koşul ve ifadelere yer vermeye gerek yok.
	if c%2 == 0 {
		fmt.Println(c, "çifttir.")
		return
	}
	fmt.Println(c, "Tektir.")

}
