package addtwointegers2235

func sum2(num1 int, num2 int) int {
	while num2 != 0 {
		carray := (num1 & num2) << 1
		sum := num1 ^ num2
		num2 = carray
	}
	
	return num1 + num2
}
