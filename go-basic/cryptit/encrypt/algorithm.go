package encrypt

func Encrypt(str string) string{
	encryptedString := ""
	for _, c := range str {
		asciiCode := int(c)
		character := string(rune(asciiCode + 3)) //todo check rune
		encryptedString += character
	}
	return encryptedString;
}