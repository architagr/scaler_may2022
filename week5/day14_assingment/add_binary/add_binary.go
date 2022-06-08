package add_binary

func AddBinary(A, B string) string {
	first, second := A, B
	result := ""
	if len(B) > len(A) {
		first = B
		second = A
	}
	countPrefix := (len(first) - len(second))
	for i := 0; i < countPrefix; i++ {
		second = "0" + second
	}
	c := "0"
	for i := len(second) - 1; i > -1; i-- {

		cf := string(first[i])
		cs := string(second[i])

		value, carry := addBit(cf, cs)
		value, c = addBit(value, c)
		c, _ = addBit(c, carry)

		result = value + result
	}
	if c == "1" {
		result = c + result
	}
	return result
}
func addBit(bit1, bit2 string) (value string, carry string) {
	if bit1 == bit2 {
		value = "0"
		carry = bit1
	} else {
		value = "1"
		carry = "0"
	}
	return
}
