/*
Дана последовательность, состоящая из целых чисел. Напишите программу, которая подсчитывает количество положительных чисел среди элементов последовательности.
*/

package main
import(
    "fmt"
)

func main(){
    var a int
    //var result []int
    fmt.Scan(&a)
    array := make([]int, a, a)
    if a >= 0 && a <= 100 {
        for i, _:= range array {
            fmt.Scan(&array[i])
            if i % 2 ==0 {
                fmt.Printf("%d ", array[i])
            }
        }
    }  
}
