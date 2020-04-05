package algorithm

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestPubSub(t *testing.T) {
	pub := NewPublisher(10, 3*time.Second)
	defer pub.Close()

	strSub := pub.Subcrib(func(v interface{}) bool {
		_, ok := v.(string)
		return ok
	})
	intSub := pub.Subcrib(func(v interface{}) bool {
		_, ok := v.(int)
		return ok
	})
	allSub := pub.Subcrib(nil)
	var wg sync.WaitGroup
	wg.Add(3)
	go func() {
		defer wg.Done()
		v := <-strSub
		fmt.Println("str subcrib:", v)
	}()
	go func() {
		defer wg.Done()
		v := <-intSub
		fmt.Println("int subcrib:", v)
	}()
	go func() {
		defer wg.Done()
		for i := 0; i < 2; i++ {
			v := <-allSub
			fmt.Println("all subcrib:", v)
		}
	}()
	pub.Publish("hello string")
	pub.Publish(123)
	wg.Wait()
}
