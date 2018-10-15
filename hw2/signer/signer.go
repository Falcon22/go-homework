package main

import (
	"fmt"
	"sort"
	"strconv"
	"sync"
)

type MyString []string

func (this MyString) Len() int {
	return len(this)
}

func (this MyString) Less(i, j int) bool {
	return this[i] < this[j]
}

func (this MyString) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

func ExecutePipeline(jobs ...job) {
	wg := &sync.WaitGroup{}
	defer wg.Wait()
	in := make(chan interface{})
	for i := range jobs {
		wg.Add(1)
		outTmp := make(chan interface{})
		go func(in, out chan interface{}, work job, wg *sync.WaitGroup) {
			defer wg.Done()
			work(in, out)
			close(out)
		}(in, outTmp, jobs[i], wg)
		in = outTmp
	}
}

func SingleHash(in, out chan interface{}) {
	wg, mu := &sync.WaitGroup{}, &sync.Mutex{}
	defer wg.Wait()
	for i := range in {
		wg.Add(1)
		go func(i interface{}, wg *sync.WaitGroup, mu *sync.Mutex) {
			defer wg.Done()
			data, ok := i.(int)
			if !ok {
				fmt.Printf("Cant convert result data to integer\n")
			}
			str := strconv.Itoa(data)
			var crc32, crc32md5 string

			mu.Lock()
			md5 := DataSignerMd5(str)
			mu.Unlock()

			wg1 := &sync.WaitGroup{}
			wg1.Add(2)
			go func(wg1 *sync.WaitGroup) {
				defer wg1.Done()
				crc32 = DataSignerCrc32(str)
			}(wg1)

			go func(wg1 *sync.WaitGroup) {
				defer wg1.Done()
				crc32md5 = DataSignerCrc32(md5)
			}(wg1)
			wg1.Wait()

			out <- crc32 + "~" + crc32md5
		}(i, wg, mu)
	}
}

func MultiHash(in, out chan interface{}) {
	wg := &sync.WaitGroup{}
	for i := range in {
		wg.Add(1)
		go func(i interface{}, out chan interface{}, wg *sync.WaitGroup) {
			defer wg.Done()
			data, ok := i.(string)
			if !ok {
				fmt.Printf("Cant convert result data to string\n")
			}
			strings := make([]string, 6)

			wg1 := &sync.WaitGroup{}
			mu := &sync.Mutex{}
			for i := 0; i < 6; i++ {
				wg1.Add(1)
				go func(arr []string, ind int, wg1 *sync.WaitGroup, mu *sync.Mutex) {
					defer wg1.Done()
					crc32 := DataSignerCrc32(strconv.Itoa(ind) + data)
					mu.Lock()
					arr[ind] = crc32
					mu.Unlock()
				}(strings, i, wg1, mu)
			}
			wg1.Wait()

			var res string
			for i := 0; i < 6; i++ {
				res += strings[i]
			}
			out <- res
		}(i, out, wg)
	}
	wg.Wait()
}

func CombineResults(in, out chan interface{}) {
	result := make([]string, 0)
	for i := range in {
		data, ok := i.(string)
		if !ok {
			fmt.Printf("Can't convert result data to string\n")
		}
		result = append(result, data)
	}
	sort.Sort(MyString(result))

	var res string
	if len(result) > 0 {
		res = result[0]
	}
	for i := 1; i < len(result); i++ {
		res += "_" + result[i]
	}
	out <- res
}
