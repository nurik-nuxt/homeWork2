package itoa

func Itoa(n int) string  {
	result :=""
	if n == 0 {
		return string(0)
	}
	var nums []byte
	if n<0 {
		result +="-"
		n*= -1
	}
	if n>0{
		nums=append(nums,byte(n%10 + '0'))
		n/=10
	}
	for i:=0; i<len(nums)/2; i++ {
		nums[i], nums[len(nums) -1 -i] = nums[len(nums) -1-i],nums[i]
	}
	return result + string(nums)
}
