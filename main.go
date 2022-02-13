package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	//1.do panics
	doPanic()
	//2.do files
	doFiles()
}
func doPanic() {

	defer func() {
		value := recover()
		err, ok := value.(error)
		if !ok {
			fmt.Println("Its panic")
		}

		if err != nil {
			log.Printf("recoveres %s", err)
		}
	}()
	p1 := createFirstPanic()
	p2 := createSecondPanic()
	fmt.Println(p1)
	fmt.Println(p2)
}
func createFirstPanic() int {

	arr := [3]int{0, 1, 2}

	var index = len(arr) + 1
	return arr[index]

}
func createSecondPanic() int {

	x := 5 / 0
	return x
}

func doFiles() error {
	for i := 1; i < 1000000; i++ {
		fileName := "file" + strconv.Itoa(i) + ".txt"

		file, err := os.Create(fileName)

		func() {
			defer file.Close()
		}()
		if err != nil {
			return err
		}

	}
	return nil
}
