package main

import (
	"math/rand"
)

type RandomizedSet struct {
	s        []int
	indexMap map[int]int
}

func Constructor() RandomizedSet {
	s := make([]int, 0)
	return RandomizedSet{
		s:        s,
		indexMap: make(map[int]int),
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.indexMap[val]; !ok {
		this.s = append(this.s, val)
		this.indexMap[val] = len(this.s) - 1
		return true
	} else {
		return false
	}

}

func (this *RandomizedSet) Remove(val int) bool {
	if _, ok := this.indexMap[val]; ok {
		i := this.indexMap[val]
		// 把要刪除的數，放入slice最後的數
		this.s[i] = this.s[len(this.s)-1]
		this.indexMap[this.s[len(this.s)-1]] = this.indexMap[val]
		// 刪除最後一個數
		this.s = this.s[:len(this.s)-1]
		delete(this.indexMap, val)
		return true
	} else {
		return false
	}
}

func (this *RandomizedSet) GetRandom() int {
	return this.s[rand.Intn(len(this.s))]
}
