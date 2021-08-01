package core

import (
	"fmt"
	"math"
	"sort"
)

type data struct {
	size       uint
	sum        float64
	frequency  map[int]uint
	order_keys []int
}

type ErrOutOfBounds uint

func (e ErrOutOfBounds) Error() string {
	return fmt.Sprintf("Index %v is out of bounds", uint(e))
}

func (d *data) store_data(ch chan int) {
	d.frequency = make(map[int]uint)
	d.order_keys = make([]int, 0)
	for {
		res, ok := <-ch
		if ok == false {
			break
		}
		d.size++
		d.sum += float64(res)
		if d.frequency[res] == 0 {
			d.order_keys = append(d.order_keys, res)
		}
		d.frequency[res] += 1
	}
	sort.Ints(d.order_keys)
}

func (d *data) calculate_median() float64 {
	if d.size%2 != 0 {
		ret, _ := d.find_index((d.size - 1) / 2)
		return float64(ret)
	}
	left, _ := d.find_index((d.size - 1) / 2)
	right, _ := d.find_index(d.size / 2)
	return float64(left+right) / 2.0
}

func (d *data) find_index(id uint) (int, error) {
	e := ErrOutOfBounds(id)
	if id >= d.size {
		return 0, e
	}
	counter := uint(0)
	for _, key := range d.order_keys {
		counter += d.frequency[key]
		if counter > id {
			return key, nil
		}
	}
	return 0, e
}

func (d *data) calculate_mode() int {
	freq_max := uint(1)
	key_ret := d.order_keys[0]
	for key := range d.order_keys {
		if d.frequency[key] > freq_max {
			key_ret = key
			freq_max = d.frequency[key]
		}
	}
	return (key_ret)
}

func (d *data) calculate_sd() float64 {
	var sd float64 = 0

	mean := d.calculate_mean()
	for k, v := range d.frequency {
		sd += float64(v) * math.Pow(float64(k)-mean, 2)
	}
	return math.Sqrt(sd / float64(d.size))
}

func (d *data) calculate_mean() float64 {
	return (d.sum / float64(d.size))
}
