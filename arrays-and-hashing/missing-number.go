// get the sum of a range from 1 to n and then subtract 
// all numbers we have to get the number we're missing

func missingNumber(nums []int) int {
	sum := (len(nums) * (len(nums) + 1)) / 2

	for _, num := range nums {
		sum -= num
	}

	return sum
}

// x ^ x ^ y ^ y ^ z = z because XORing the 
// same number with itself = 0

func missingNumber(nums []int) int {
	ans := 0

	for idx, num := range nums {
		ans = ans ^ (idx + 1) ^ num
	}

	return ans
}

