package shuffle_array

import "math/rand"

type Solution struct {
	original []int
}


func Constructor(nums []int) Solution {
	res := Solution{}
	res.original = nums
	return res
}


/** Resets the array to its original configuration and return it. */
func (this *Solution) Reset() []int {
	res := []int{}
	res = append(res, this.original...)
	return res
}


/** Returns a random shuffling of the array. */
func (this *Solution) Shuffle() []int {
	tmp := []int{}
	tmp = append(tmp,this.original...)
	for i:=0; i< len(tmp)-1;i++ {
		randomIndex := rand.Intn(len(tmp)-i-1) + (i+1)
		tmp[randomIndex],tmp[i] = tmp[i],tmp[randomIndex]
	}
	return tmp
}

