package helloChannel

func HelloFromChannel(mc chan string, ec chan bool, s string) {
	for i := 0; i < 1000; i++ {
		mc <- s
	}
	ec <- true
}
