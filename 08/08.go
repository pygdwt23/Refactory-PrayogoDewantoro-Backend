package main

import (
	"fmt"
)

func masking(s string) string {
rs := [] rune(s)
for i := 0; i < len(rs)-3; i++ {
	rs[i] = '*'
}
return string(rs)

}

func main() {
	fmt.Println(masking("23dn3ir30fd2eddd"))
}