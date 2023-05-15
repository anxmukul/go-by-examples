package main

import ("fmt"
		"unicode/utf8"
)

func main(){
	const str = "HelloWorld"
	for i:=0; i<len(str); i++{
		fmt.Println("%x ",str[i]);
	}
	// Counting how many runes are present in string str we can use utf8 package
	fmt.Println("Rune count:", utf8.RuneCountInString(str))


	for idx, runeValue := range str {
        fmt.Printf("%#U starts at %d\n", runeValue, idx)
    }
}