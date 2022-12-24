package array

func plusOne(digits []int) []int {
	carry := 1
	result := make([]int, 0)
	for i := len(digits) - 1; i >= 0; i-- {
		sum := digits[i] + carry
		digits[i] = sum % 10
		carry = sum / 10
	}
	if carry == 1 {
		result = append(result, carry)
	}
	result = append(result, digits...)
	return result
}
