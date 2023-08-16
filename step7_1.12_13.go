/*
Напишите программу, принимающая на вход число 
�
(
�
≥
4
)
N(N≥4), а затем 
�
N целых чисел-элементов среза. На вывод нужно подать 4-ый (3 по индексу) элемент данного среза.
*/

package main
import "fmt"

func main() {
    var a int
    fmt.Scan(&a)
    n := make([]int, a, a)
    //fmt.Println(a)
    //fmt.Println(n)
    for i , _ := range(n){
        fmt.Scan(&n[i])
    }
    fmt.Printf("%d", n[3])
    
}
