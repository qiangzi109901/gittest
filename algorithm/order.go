package main

import (
	"fmt"
	"math/rand"
	"time"
)


/**
 * 选择排序算法
 *
 * 每一轮i筛选一个最大值并放到len-1-i位置上
 */


func OrderByChoose(nums []int) {
	size := len(nums)
	for i:=0;i<size-1;i++ {
		max := nums[0]
		maxIndex := 0

		for j:=1;j<size-i;j++ {
			if max<nums[j] {
				max = nums[j]
				maxIndex = j
			}
		}
		//将最后一位置为max
		lastindex := size - 1 -i
		nums[maxIndex] = nums[lastindex]
		nums[lastindex] = max

	}
}


/**
 * 冒泡排序算法
 *
 * 两两比较，如果前者小于后者，则不操作，否则就更换两个元素的位置
 */
func OrderByBubble(nums []int) {
	size := len(nums)

	for i:=0; i<size-1; i++{

		//每轮在i到size-i-1进行比较
		for j:=0;j<size-2-i;j++ {
			var prev = nums[j]
			var next = nums[j+1]

			if prev > next {
				//更换两个元素位置,调整后应该是大数后调了
				nums[j] = next
				nums[j+1] = prev
			}
		}
	}
}


func Shuffle(nums []int) {
	size := len(nums)

	//打乱10次
	for i := 0; i<10; i++ {
		random := rand.New(rand.NewSource(time.Now().UnixNano()))
		r := random.Intn(size)
		r2 := random.Intn(size)
		fmt.Println("随机数:", r, r2)
		temp := nums[r]
		nums[r] = nums[r2]
		nums[r2]  = temp
	}
}

func main() {
	nums := []int{93,221,22,34,32,103,32,56,77,10,203,40}

	OrderByChoose(nums)
	fmt.Println(nums)

	Shuffle(nums)
	fmt.Println(nums)

	OrderByBubble(nums)
	fmt.Println(nums)

}