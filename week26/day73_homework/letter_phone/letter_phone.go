package letter_phone

var result []string
var keyMap map[byte][]byte

func letterCombinations(A string) []string {
	keyMap = make(map[byte][]byte)
	keyMap['0'] = []byte{'0'}
	keyMap['1'] = []byte{'1'}
	keyMap['2'] = []byte{'a', 'b', 'c'}
	keyMap['3'] = []byte{'d', 'e', 'f'}
	keyMap['4'] = []byte{'g', 'h', 'i'}
	keyMap['5'] = []byte{'j', 'k', 'l'}
	keyMap['6'] = []byte{'m', 'n', 'o'}
	keyMap['7'] = []byte{'p', 'q', 'r', 's'}
	keyMap['8'] = []byte{'t', 'u', 'v'}
	keyMap['9'] = []byte{'w', 'x', 'y', 'z'}

	result = make([]string, 0)
	buffer := make([]byte, len(A))
	get(A, 0, buffer)
	return result
}
func get(str string, i int, buffer []byte) {
	if i == len(str) {
		d := make([]byte, 0)
		d = append(d, buffer...)
		result = append(result, string(d))
		return
	}
	letterList := keyMap[byte(str[i])]
	for _, letter := range letterList {
		buffer[i] = letter
		get(str, i+1, buffer)
	}
}
