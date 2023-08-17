package hello

func ValidChar() []byte {
	var temp []byte
	for i := byte(32); i < 128; i++ {
		if isSpace(i) || isAlphaNumberic(i) || isPunctuation(i) {
			temp = append(temp, i)
		}
	}
	return temp
}

func isSpace(b byte) bool {
	return b == 32
}

func isAlphaNumberic(b byte) bool {
	return (b >= 65 && b <= 90) || (b >= 97 && b <= 122) || (b >= 48 && b <= 57)
}

func isPunctuation(b byte) bool {
	return (b >= 33 && b <= 47) || (b >= 58 && b <= 64) || (b >= 91 && b <= 96) || (b >= 123 && b <= 126)
}

func GetKey() (key string) {
	key = "Hello World!"

	return
}
