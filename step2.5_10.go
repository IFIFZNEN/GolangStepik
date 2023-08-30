/*
На вход дается строка, из нее нужно сделать другую строку, оставив только нечетные символы (считая с нуля)

Sample Input:
ihgewlqlkot

Sample Output:
hello

*/

package main

import(
    "fmt"

)

func main(){
    var text, newText string
    fmt.Scan(&text)
    for i,v := range text {
        if i % 2 != 0{ 
            newText = newText + string(v)
        }
    }
    fmt.Println(newText)
}

