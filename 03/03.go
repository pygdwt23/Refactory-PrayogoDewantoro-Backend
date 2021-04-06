package main

import (
	"fmt"
	"strings"
)

func main() {
    text := `Aku ingin begini
		Aku ingin begitu
		Ingin ini itu banyak sekali
		
		Semua semua semua
		Dapat dikabulkan
		Dapat dikabulkan
		Dengan kantong ajaib
		
		Aku ingin terbang bebas
		Di angkasa
		Hei… baling baling bambu
		
		La... la... la...
		Aku sayang sekali
		Doraemon
		
		La... la... la...
		Aku sayang sekali`
    fields := strings.FieldsFunc(text, func(r rune) bool {
        return !('a' <= r && r <= 'z' || 'A' <= r && r <= 'Z')
    })
    words := make(map[string]int)
    for _, field := range fields {
        words[strings.ToLower(field)]++
    }
    fmt.Println(words)
}