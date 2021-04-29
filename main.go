package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Hello World")
	fmt.Println("3 + 5 = ", sum(3, 5))
	fmt.Println("3 * 5 = ", multiply(3, 5))
	playArray()

	messageGreeting := "Hello API: "
	counterHit := 0
	http.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		counterHit += 1
		message := fmt.Sprintf("%s: %d", messageGreeting, counterHit)
		fmt.Fprintf(w, message)
		fmt.Println("API hit")
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
func playArray() {
	arr1 := make([]int, 10, 10)
	arr1[0] = 1
	arr1[1] = 11
	arr1[2] = 22
	arr1[3] = 33
	arr1[4] = 41
	printSlice(arr1)

	arr2 := arr1[2:4]
	printSlice(arr2)

	arr3 := make([]int, 3, 3)
	copy(arr3, arr1)
	printSlice(arr3)

	//
	arr3[2] = 999
	printSlice(arr1)
	printSlice(arr2)
	printSlice(arr3)
}
func printSlice(arr []int) {
	fmt.Println(arr, "len:", len(arr), cap(arr))
}
func sum(n1, n2 int) int {
	return n1 + n2
}
func multiply(n1, n2 int) int {
	return n1 * n2
}
