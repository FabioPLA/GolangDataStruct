package main

import "sort"

type Num struct {
	num  int
	freq int
}

func frequencySort(nums []int) []int {
	//1. 统计每个数字出现的频数
	numsFreqMap := map[int]int{}
	for _, num := range nums {
		if _, ok := numsFreqMap[num]; ok {
			numsFreqMap[num]++
		} else {
			numsFreqMap[num] = 1
		}
	}

	numsFreq := make([]Num, 0)
	for key, val := range numsFreqMap {
		numsFreq = append(numsFreq, Num{
			num:  key,
			freq: val,
		})
	}

	//2. 按频数排序
	sort.Slice(numsFreq, func(i, j int) bool {
		if numsFreq[i].freq == numsFreq[j].freq {
			return numsFreq[i].num > numsFreq[j].num
		}
		return numsFreq[i].freq < numsFreq[j].freq
	})

	res := make([]int, 0)

	for _, val := range numsFreq {
		for i := 0; i < val.freq; i++ {
			res = append(res, val.num)
		}
	}

	return res
}
func main() {

}
