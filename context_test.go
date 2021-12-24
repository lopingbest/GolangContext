package GolangContext

import (
	"context"
	"fmt"
	"runtime"
	"testing"
	"time"
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
func CreateCounter(ctx context.Context) chan int {
	destination := make(chan int)
	//isi dari channel
	go func() {
		//memastkan untuk close
		defer close(destination)
		counter := 1
		//perulangan
		for {
			//jika context nya selesai, maka berhenti "done"
			select {
			case <-ctx.Done():
				return
			default:
				//nilai default jika context belum selesai
				destination <- counter
				counter++
				time.Sleep(1 * time.Second) //simulasi slow
			}
		}
	}()
	return destination
}

func TestContextWithCancel(t *testing.T) {
	fmt.Println("Total Goroutine", runtime.NumGoroutine())
	parent := context.Background()
	ctx, cancel := context.WithCancel(parent)

	destination := CreateCounter(ctx)

	for n := range destination {
		fmt.Println("Counter", n)
		//menghentikan perulangan jika sudah sampai 10
		if n == 10 {
			break
		}
	}
	cancel() //pengiriman sinyal cancel kecontext
	//untuk memberi jeda waktu, agar cancel bekerja, karena goroutine bekerja secara asynchronous
	time.Sleep(2 * time.Second)
	fmt.Println("Total Goroutine", runtime.NumGoroutine())

}

func TestContextWithTimeOut(t *testing.T) {
	fmt.Println("Total Goroutine", runtime.NumGoroutine())
	parent := context.Background()
	ctx, cancel := context.WithTimeout(parent, 5*time.Second)
	defer cancel()
	destination := CreateCounter(ctx)

	//looping sesuai waktu yang diberikan untuk timeout
	for n := range destination {
		fmt.Println("Counter", n)
	}
	//untuk memberi jeda waktu, agar cancel bekerja, karena goroutine bekerja secara asynchronous
	time.Sleep(2 * time.Second)
	fmt.Println("Total Goroutine", runtime.NumGoroutine())

}
