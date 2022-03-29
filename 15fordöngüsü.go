/* Kontrol akış diyagramları (programın ilerleme şekli)
1- Sequential  -  Sıralı olarak gitmektedir.
2- Conditional - Herhangi bir koşula bağlı olarak ilerlemektedir.
3- Loop - İteration (Tekrarlama)  Çevrim içerisine girerek ilerler. */

package main

import "fmt"

func main() {

	for i := 0; i <= 10; i += 5 {
		fmt.Println(i)
	}
	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}
	//fmt.Println(i) //i for loop içinde declare edilmiştir.

	/* 	for { // İnfinite Loop (sonsuz değer)
		fmt.Println("fed")  // control+c ile cıkılır.
	} */

	/* 	for j := 0; true; j += 5 {  //false olursa işlem yapılmaz.
		fmt.Println(j)

	} */
	k := 10
	for k >= 0 { // sadece conditial (şart) durumu ile yazılabilir.
		fmt.Println(k)
		k--
	}

	for t := 0; t <= 10; t++ {
		if t%3 == 0 {

			continue // burada eşitlik sağlanıyorsa tekrar for döngüsünün başına git diyor.
		}
		fmt.Println(t)
	}

	for m := 0; m <= 10; m++ {
		if m == 3 {
			break // eşitlik sağlandığı anda for döngüsünden çıkar.
		}
		fmt.Println(m)
	}

}

// for <init statement> ; < condition>; <post statement> {----}
// başlangıç degeri
