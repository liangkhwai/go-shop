package rbac


func IntToBinary(n ,length int) []int{
	binary := make([]int, length)
	for i := 0; n > 0; i++ {
		binary[i] = n % 2
		n = n / 2
	}
	return binary
}