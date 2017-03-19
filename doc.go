package semaphore

/*
A semaphore implementation based on https://pauladamsmith.com/blog/2016/04/max-clients-go-net-http.html

Example

```
func main() {
	s := NewSemaphore(4)
	for i := 0; i < 22; i++ {
		s.Add()
		go func(i int) {
			defer func() {
				s.Done()
				fmt.Printf("finished %d\n", i)
			}()
			fmt.Printf("running: %d\n", i)
			time.Sleep(time.Second * time.Duration(i))
		}(i)
	}
	s.Wait()
}

```

*/
