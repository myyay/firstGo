package runner

import (
	"errors"
	"log"
	"testing"
	"time"
)

func TestRunner(t *testing.T) {
	producer := func(data dataChan) error {
		for i := 0; i < 30; i++ {
			data <- i
			log.Printf("producer sent %v ", i)
		}

		return nil
	}

	consumer := func(data dataChan) error {
	forLoop:
		for {
			select {
			case c := <-data:
				log.Printf("consume : %v ", c)
			default:
				break forLoop
			}
		}

		return errors.New("finish Error")
	}

	r := NewRunner(false, 30, producer, consumer)
	go r.startAll()

	time.Sleep(3 * time.Second)

}
