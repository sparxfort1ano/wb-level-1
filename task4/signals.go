package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"time"
)

func worker(i int, sendCh <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for x := range sendCh {
		fmt.Println("Воркер номер", i+1, "вывел", x)
	}
}

func producer(rcvCh chan<- int, ctx context.Context) {
	ticker := time.NewTicker(500 * time.Millisecond) // чтобы в момент времени не выводилось на консоль очень много сообщений
	defer ticker.Stop()

	var x int = 1
	for {
		select {
		case <-ctx.Done(): // ждет SIGINT
			return
		case <-ticker.C: // тикает каждые полсекунды
			rcvCh <- x
			x++
		}
	}
}

func main() {
	var n int
	fmt.Print("Введите кол-во горутин-воркеров: ")
	fmt.Scan(&n)

	ch := make(chan int, n)

	var wg sync.WaitGroup // программа не завершится, пока не отработают все воркеры
	for i := 0; i < n; i++ {
		wg.Add(1)
		go worker(i, ch, &wg) // n воркеров читают данные из канала
	}

	ctx, _ := signal.NotifyContext(context.Background(), os.Interrupt)
	/* был выбран контекст как самый удобный handler сигналов (ctrl + C => <-ctx.Done())
	os.Interrupt в отличие от syscall.SIGINT кроссплатформенный*/

	producer(ch, ctx) // кладем данные в канал до Ctrl + C

	close(ch)

	wg.Wait() // ждем воркеров
	log.Println("graceful shutdown has got made")
}
