package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

const workersAmount int = 3

func worker(i int, sendCh <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for x := range sendCh {
		fmt.Println("Воркер номер", i+1, "вывел", x)
	}
}

func producer(rcvCh chan<- int, n int) {
	ticker := time.NewTicker(250 * time.Millisecond) // тикер ограничивает частоту сообщений (1 раз в 250 мс)
	defer ticker.Stop()

	// Здесь используется time.After вместо context.WithTimeout или time.NewTimer.
	// Выбор в пользу time.After: нужен простой одноразовый таймер.
	timer := time.After(time.Duration(n) * time.Second)

	var x int = 1
	for {
		select {
		case <-timer: // этот кейс будет готов по истечению n секунд
			return
		case <-ticker.C: // выдаём значения каждые 250 мс
			rcvCh <- x
			x++
		}
	}
}

func main() {
	var n int
	fmt.Print("Введите кол-во секунд, в течение которого в канал будут поступать данные: ")
	fmt.Scan(&n)

	ch := make(chan int, workersAmount)

	var wg sync.WaitGroup // программа не завершится, пока не отработают все воркеры
	for i := 0; i < workersAmount; i++ {
		wg.Add(1)
		go worker(i, ch, &wg) // workersAmount воркеров читают данные из канала
	}

	producer(ch, n) //  завершаем producer по таймауту

	close(ch)

	wg.Wait() // ждем воркеров
	log.Println("Прошло", n, "секунд")
	log.Println("graceful shutdown complete")
}
