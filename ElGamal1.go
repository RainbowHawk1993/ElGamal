package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func main() {

	p := big.NewInt(593) //prime number
	g := big.NewInt(3)
	x := big.NewInt(9)
	m := big.NewInt(1912) // if this is higher than p, then essentially m = m % p
	//k := big.NewInt(111)
	one := big.NewInt(1)

	fmt.Println("p = ", p)
	fmt.Println("g = ", g)
	fmt.Println("x = ", x)
	fmt.Println("m = ", m)

	// generating random k beetween 0 and p - 1 (should be between 1 and p-1). maybe will fix this later
	k, err := rand.Int(rand.Reader, p)
	if err != nil {
		//error handling
	}
	fmt.Println("k = ", k)

	// y = g ** x % p
	y := (new(big.Int).Exp(g, x, p))
	fmt.Println("y = ", y)

	// c1 = g ** k % p
	c1 := (new(big.Int).Exp(g, k, p))
	fmt.Println("c1 = ", c1)

	// c2 = (m * y ** k) % p
	c2 := (new(big.Int).Exp(y, k, p))
	c2 = c2.Mul(c2, m) //multiplication
	c2 = c2.Rem(c2, p) // remainder-modulus-%
	fmt.Println("c2 = ", c2)

	Psub1 := (new(big.Int).Sub(p, one))
	Psub1subX := (new(big.Int).Sub(Psub1, x)) // p - 1 - x

	// 	decrypt = (c2 * (c1 ** (p - 1 - x))) % p
	decrypt := (new(big.Int).Exp(c1, Psub1subX, nil))
	decrypt = decrypt.Mul(c2, decrypt) //Multiplication
	decrypt = decrypt.Rem(decrypt, p)  // remainder-modulus-%

	fmt.Println("Decryption result ", decrypt) // == 1912 % 593

}
