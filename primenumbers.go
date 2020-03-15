package algorithm

//PrimeNumbersBuilder 质数筛
func PrimeNumbersBuilder(max int) []int {
	ans := make([]int, 0)
	isPrime := make([]bool, max+1)
	for i := 2; i <= max; i++ {
		isPrime[i] = true
	}
	for i := 2; i*i <= max; i++ {
		if isPrime[i] {
			for j := i * i; j <= max; j += i {
				isPrime[j] = false
			}
		}
	}
	for i := 2; i <= max; i++ {
		if isPrime[i] {
			ans = append(ans, i)
		}
	}
	return ans
}

// PrimeNumbersChanBuilder 管道筛选质数
// 纯属好玩，无法提升运行时间，因为管道是串联的
// 甚至因为有协程的切换而降低效率
func PrimeNumbersChanBuilder(max int) []int {
	ans := make([]int, 0)
	ch, exit := genarateChan(max)
	for {
		select {
		case prime := <-ch:
			ans = append(ans, prime)
			ch, exit = primeFilter(ch, exit, prime)
		case <-exit:
			return ans
		}
	}
}

func genarateChan(max int) (<-chan int, <-chan struct{}) {
	ch := make(chan int)
	exit := make(chan struct{})
	go func() {
		for i := 2; i <= max; i++ {
			ch <- i
		}
		close(exit)
	}()
	return ch, exit
}

func primeFilter(ch <-chan int, exit <-chan struct{}, prime int) (<-chan int, <-chan struct{}) {
	filterCh := make(chan int)
	filterExit := make(chan struct{})
	go func() {
		for {
			select {
			case i, ok := <-ch:
				if ok {
					if i%prime != 0 {
						filterCh <- i
					}
				}
			case <-exit:
				close(filterCh)
				close(filterExit)
				return
			}
		}
	}()
	return filterCh, filterExit
}
