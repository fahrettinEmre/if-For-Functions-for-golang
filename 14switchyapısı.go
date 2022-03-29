package main

import "fmt"

func main() {

	// switch

	grade := 10
	switch grade := 4; grade { // switch yapısı if/else yapısının daha basit görünmesi için kullanılır.
	case 5: // if grade==5 ise
		fmt.Println("Pekiyi")
	case 4:
		y := 100 // y burada sadece case 4 için geçerlidir.
		fmt.Println("İyi")
		fmt.Println(y)
	case 3:
		fmt.Println("Orta") // burdan sonraki case'lere bakmaz direk switch yapısının dışına çıkar.
		// break  burda gizli break vardır.
	case 2:
		fmt.Println("Zayıf")
	case 1:
		fmt.Println("Kötü")
	default: // else yapısı gibidir. // switch içerisinde herhangi bir yere yazılabilir.
		fmt.Println("Geçersiz not.")

	}
	fmt.Println(grade)

	if grade == 5 {
		fmt.Println("Pekiyi")
	}
	if grade == 4 {
		fmt.Println("İyi")
	}
	if grade == 3 {
		fmt.Println("Orta")
	}
	if grade == 2 {
		fmt.Println("Zayıf")
	}
	if grade == 1 {
		fmt.Println("Kötü")
	} else {
		fmt.Println("Geçersiz not.")
	}
}
