package main

import "fmt"

func main() {
	obj := Constructor()
	obj.Inc("hello")
	obj.Inc("hello")
	fmt.Println(obj.GetMaxKey())
	fmt.Println(obj.GetMinKey())
	obj.Inc("leet")
	obj.Dec("hello")
	fmt.Println(obj.GetMaxKey())
	fmt.Println(obj.GetMinKey())
}

type AllOne struct {
	counter map[string]int
	words   map[int]map[string]struct{}
}

func Constructor() AllOne {
	return AllOne{
		counter: map[string]int{},
		words:   map[int]map[string]struct{}{},
	}
}

func (this *AllOne) Inc(key string) {
	oldCounter, ok := this.counter[key]
	if !ok {
		this.counter[key] = 0
		oldCounter = 0
	}

	delete(this.words[oldCounter], key)
	this.counter[key] += 1
	if this.words[this.counter[key]] == nil {
		this.words[this.counter[key]] = map[string]struct{}{}
	}
	this.words[this.counter[key]][key] = struct{}{}
}

func (this *AllOne) Dec(key string) {
	oldCounter := this.counter[key]

	delete(this.words[oldCounter], key)
	this.counter[key] -= 1

	if this.counter[key] == 0 {
		return
	}
	if this.words[this.counter[key]] == nil {
		this.words[this.counter[key]] = map[string]struct{}{}
	}
	this.words[this.counter[key]][key] = struct{}{}
}

func (this *AllOne) GetMaxKey() string {
	if len(this.words) == 0 {
		return ""
	}

	for i := len(this.words); i > 0; i-- {
		for j := range this.words[i] {
			return j
		}
	}

	return ""
}

func (this *AllOne) GetMinKey() string {
	if len(this.words) == 0 {
		return ""
	}

	for i := 1; i <= len(this.words); i++ {
		for j := range this.words[i] {
			return j
		}
	}

	return ""
}
