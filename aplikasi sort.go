// Quick Sort in Golang
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	data := generatedata(20)
	fmt.Println("\n--- Unsorted --- \n\n", data)
	quicksort(data)
	fmt.Println("\n--- Sorted by Quick Sort---\n\n", data, "\n")
    selectionsort(data)
    fmt.Println("\n--- Sorted by Selection Sort---\n\n", data, "\n")
}

// Generates a slice of size, size filled with random numbers
func generatedata(size int) []int {

	data := make([]int, size, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		data[i] = rand.Intn(999) - rand.Intn(999)
	}
	return data
}
 
func quicksort(a []int) []int {
    if len(a) < 2 {
        return a
    }
     
    left, right := 0, len(a)-1
     
    pivot := rand.Int() % len(a)
     
    a[pivot], a[right] = a[right], a[pivot]
     
    for i, _ := range a {
        if a[i] < a[right] {
            a[left], a[i] = a[i], a[left]
            left++
        }
    }
     
    a[left], a[right] = a[right], a[left]
     
    quicksort(a[:left])
    quicksort(a[left+1:])
     
    return a
}

func selectionsort(data []int) []int{
    var temp, j int
    for i:=1; i < len(data);i++{
        j=i
        for j>0{
            if data[j-1]>data[j]{
                temp= data[j]
                data[j]=data[j-1]
                data[j-1]=temp
            }
            j=j-1
        }
    }
    return data   
}