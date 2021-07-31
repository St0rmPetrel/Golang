package core

type Result struct {
	mean	float32
	median	float32
	mode	float32
	sd		float32
}

func Calculate(ch chan int) result {
	var ret result

	for {
		res, ok := <-ch
		if ok == false {
			println("Channel Close ")
			break
		}
		println("Channel Open :", res)
	}
	return ret
}
