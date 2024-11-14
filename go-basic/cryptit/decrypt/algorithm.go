package decrypt

func Decrypt(str string) string{
	decryptedString := ""
	for _, c := range str {
		asciiCode := int(c)
		character := string(rune(asciiCode - 3)) //todo check rune
		decryptedString += character
	}
	return decryptedString;
}