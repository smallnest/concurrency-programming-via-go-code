package main

func main() {
	var m = make(map[int]int, 10) // 初始化一个map
	go func() {
		for {
			m[1] = 1 //设置key
		}
	}()

	go func() {
		for {
			for k, v := range m {
				_, _ = k, v
			}
		}
	}()

	select {}
}
