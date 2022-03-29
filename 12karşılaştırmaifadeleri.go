/* +	Addition	x + y	Sum of x and y
-	Subtraction	x - y	Subtracts one value from another
*	Multiplication	x * y	Multiplies two values
/	Division	x / y	Quotient of x and y
%	Modulus	x % y	Remainder of x divided by y
++	Increment	x++	Increases the value of a variable by 1
--	Decrement	x--	Decreases the value of a variable by 1 */

/* x = y	Assign	x = y
x += y	Add and assign	x = x + y
x -= y	Subtract and assign	x = x - y
x *= y	Multiply and assign	x = x * y
x /= y	Divide and assign quotient	x = x / y
x %= y	Divide and assign modulus	x = x % y */

/* ==	Equal	x == y	True if x is equal to y
!=	Not equal	x != y	True if x is not equal to y
<	Less than	x < y	True if x is less than y
<=	Less than or equal to	x <= y	True if x is less than or equal to y
>	Greater than	x > y	True if x is greater than y
>=	Greater than or equal to	x >= y	True if x is greater than or equal to y */

/* &&	Logical And	Returns true if both statements are true	x < y && x > z
||	Logical Or	Returns true if one of the statements is true	x < y || x > z
!	Logical Not	Reverse the result, returns false if the result is true	!(x == y && x > z) */

package main

import "fmt"

func main() {
	x := 27
	y := 30
	if false { // false olduğu için  buraya girmeden program biter.
		fmt.Println(x, "tek sayıdır.")
	}

	if !false { //if'den sonra sonucu boolen olan bir yapı kullanmamız gerekir.
		fmt.Printf("%v tek sayıdır.\n", x)
	}
	if y%2 == 0 {
		fmt.Println(y, "çift sayıdır.")
	} else {
		fmt.Println(y, "tek sayıdır.")
	}
	if 5 > 3 {
		fmt.Println("Mesaj gösterilecek mi?")
	}

	a := -5 // branching (dallanma)
	if a < 0 {
		fmt.Println(a, "negatif sayıdır.")
	} else if a%2 == 0 {
		fmt.Println(a, "çift sayıdır.")
	} else {
		fmt.Println(a, "tek sayıdır.")
	}
}

// if <boolen expression> {code}else if <boolen expression> {code} else {code}
