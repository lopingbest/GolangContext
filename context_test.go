package GolangContext

import (
	"context"
	"fmt"
	"runtime"
	"testing"
)

//contet bersifat imutable (tidak bisa diubah). Jika diubah maka akan membuat child baru
func TestContext(t *testing.T) {
	background := context.Background()
	fmt.Println(background)

	todo := context.TODO()
	fmt.Println(todo)
}

func TestContextWithValue(t *testing.T) {
	//parent context
	contextA := context.Background()

	//Child context A
	contextB := context.WithValue(contextA, "b", "B")
	contextC := context.WithValue(contextA, "c", "C")

	contextD := context.WithValue(contextB, "d", "D")
	contextE := context.WithValue(contextB, "e", "E")

	contextF := context.WithValue(contextC, "f", "F")

	fmt.Println(contextA)
	fmt.Println(contextB)
	fmt.Println(contextC)
	fmt.Println(contextD)
	fmt.Println(contextE)
	fmt.Println(contextF)

	fmt.Println("")

	fmt.Println(contextF.Value("f"))
	fmt.Println(contextF.Value("c"))
	fmt.Println(contextF.Value("b"))
	fmt.Println(contextA.Value("b"))

}

//membuat goroutine untuk mengirim data terus terusan kedalam sebuah channel
func CreateCounter() chan int {
	destination := make(chan int)
	//isi dari channel
	go func() {
		//memastkan untuk close
		defer close(destination)
		counter := 1
		//perulangan untuk
		for {
			destination <- counter
			counter++
		}
	}()
	return destination
}

func TestContextWithCancel(t *testing.T) {
	fmt.Println("Total Goroutine", runtime.NumGoroutine())
	destination := CreateCounter()

	for n := range destination {
		fmt.Println("Counter", n)
		//menghentikan perulangan jika sudah sampai 10
		if n == 10 {
			break
		}

	}
	fmt.Println("Total Goroutine", runtime.NumGoroutine())

}
