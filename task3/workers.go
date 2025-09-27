package main

import (
	"fmt"
	"time"
)

func worker(i int, sendCh <-chan int) {
	for {
		x := <-sendCh
		fmt.Println("Воркер номер", i, "вывел", x)
	}
}

func producer(rcvCH chan<- int) {
	var x int = 1
	for {
		rcvCH <- x
		x++
		time.Sleep(500 * time.Millisecond) // чтобы в момент времени не выводилось на консоль очень много сообщений
	}
}

func main() {
	var n int
	fmt.Print("Введите кол-во горутин-воркеров: ")
	fmt.Scan(&n)

	ch := make(chan int, n)

	for i := 0; i < n; i++ {
		go worker(i, ch) // n воркеров читают данные из канала
	}

	producer(ch) // бесконечно кладем данные в канал

	// Ctrl + C для выхода
}
