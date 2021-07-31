package core

type Result struct {
	Mean	float32
	Median	float32
	Mode	float32
	Sd		float32
}

func Calculate(ch chan int) Result {
	var ret Result

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
