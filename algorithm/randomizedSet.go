package main

import "math/rand"

type RandomizedSet struct {
	nums    []int
	indices map[int]int
}

func RandomizedSetConstructor() RandomizedSet {
	return RandomizedSet{
		nums:    []int{},
		indices: map[int]int{},
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.indices[val]; ok {
		return false
	}
	this.indices[val] = len(this.nums)
	this.nums = append(this.nums, val)
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	if _, ok := this.indices[val]; !ok {
		return false
	}

	index := this.indices[val]
	last := this.nums[len(this.nums)-1]
	this.nums[index] = last
	this.indices[last] = index
	this.nums = this.nums[:len(this.nums)-1]
	delete(this.indices, val)
	return true
}

func (this *RandomizedSet) GetRandom() int {
	return this.nums[rand.Intn(len(this.nums))]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
