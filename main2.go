package main

import (
	"fmt"
	"strconv"
	"net/http"
	// "golang.org/x/text/number"
)

// no "main()" here
func hello(a int, b int) int {
    fmt.Println("Adding numbers:", a, b)
    return a + b
}

func Add(numbers ...int) int {
	total := 0
	for i:= range numbers {
		total += numbers[i]
	}
	return total
}
func Subtract(numbers ...int) int {
	total := 50
	for i:= range numbers {
		total -= numbers[i]
	}
	return total
}
func swap(x, y string) (string, string) {
	return y, x
}
func convert(str string) (int, error){
	return  strconv.Atoi(str)
}

func openUrl(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("error opening %s: %v", url, err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK{
		return fmt.Errorf("bad status code for %s: %d", url , resp.StatusCode)
	}
	fmt.Println("Response status:", resp.Status)
	return nil 
}
func loops() error{
	 if y := 10; y > 5 {   // short statement before the semicolon
        fmt.Println(y)
		return nil
    }else{
		return fmt.Errorf("y is not greater than 5")
	}
}
func demo() {
    for i := 0; i < 3; i++ {
        y := i * 2
        fmt.Println("y inside loop:", y)
    }
    // fmt.Println(y) // This will cause a compile-time error
}
var packageVar = "Mein toh Bhraam hu Bhraam"

func other() {
    fmt.Println("Accessing in other():", packageVar)
}