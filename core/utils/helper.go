package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"io"
	"regexp"
	"strings"
)

var nonAlphanumericRegex = regexp.MustCompile(`[^a-zA-Z0-9 ]+`)

var nonDigitsRegex = regexp.MustCompile(`[^0-9]+`)

func StrLeaveOnlyDigits(str string) string {
	return nonDigitsRegex.ReplaceAllString(str, "")
}

func Encrypt(key []byte, message string) (encoded string, err error) {
	//Create byte array from the input string
	plainText := []byte(message)

	//Create a new AES cipher using the key
	block, err := aes.NewCipher(key)

	//IF NewCipher failed, exit:
	if err != nil {
		return
	}

	//Make the cipher text a byte array of size BlockSize + the length of the message
	cipherText := make([]byte, aes.BlockSize+len(plainText))

	//iv is the ciphertext up to the blocksize (16)
	iv := cipherText[:aes.BlockSize]
	if _, err = io.ReadFull(rand.Reader, iv); err != nil {
		return
	}

	//Encrypt the data:
	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(cipherText[aes.BlockSize:], plainText)

	//Return string encoded in base64
	return hex.EncodeToString(cipherText), err
}

func Decrypt(key []byte, secure string) (decoded string, err error) {
	cipherText, err := hex.DecodeString(secure)
	if err != nil {
		return
	}

	//Create a new AES cipher with the key and encrypted message
	block, err := aes.NewCipher(key)

	//IF NewCipher failed, exit:
	if err != nil {
		return
	}

	//IF the length of the cipherText is less than 16 Bytes:
	if len(cipherText) < aes.BlockSize {
		err = errors.New("Ciphertext block size is too short!")
		return
	}

	iv := cipherText[:aes.BlockSize]
	cipherText = cipherText[aes.BlockSize:]

	//Decrypt the message
	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(cipherText, cipherText)

	return string(cipherText), err
}

var fioTrims = []string{
	"`", "'", "'", "`", "‘", "ʼ", "’", "`", "ʻ",
}

var fioMods = map[int]map[string]string{
	0:  {"SH": "6"},
	1:  {"CH": "4"},
	2:  {"KH": "X"},
	3:  {"H": "X"},
	4:  {"DJ": "J"},
	5:  {"##": "#"},
	6:  {"G#": "G"},
	7:  {"Q": "K"},
	8:  {"LL": "L"},
	9:  {"BB": "B"},
	10: {"DD": "D"},
	11: {"NN": "N"},
	12: {"AYEV": "AEV"},
	13: {"AYEVA": "AEVA"},
	14: {"IYEV": "IEV"},
	15: {"IYEVA": "IEVA"},
	16: {"IN#OM": "INOM"},
	17: {"E#ZOZ": "EZOZ"},
	18: {"MM": "M"},
	19: {"BOYEVA": "BOEVA"},
	20: {"BOYEV": "BOEV"},
	21: {"LOYEVA": "LOEVA"},
	22: {"LOYEV": "LOEV"},
	23: {"RA#NO": "RANO"},
}

var fioMods1 = map[int]map[string]string{
	0: {"O#": "U"},
}

var fioMods2 = map[int]map[string]string{
	0: {"O#": "O"},
}
var fioMods3 = map[int]map[string]string{
	0: {"#": ""},
}

func CalcTrustLevel(fullName string, firstName string, lastName string) int {
	_fullName := strings.TrimSpace(strings.ToUpper(fullName))
	_firstName := strings.TrimSpace(strings.ToUpper(firstName))
	_lastName := strings.TrimSpace(strings.ToUpper(lastName))

	_fullNames := strings.Fields(_fullName)

	_firstName = normalizeFioPart(_firstName)
	_lastName = normalizeFioPart(_lastName)

	for i, name := range _fullNames {
		_fullNames[i] = normalizeFioPart(name)
		//print(name, " ")
	}

	//println(_firstName, " ", _lastName)

	scoreModRaw := calcTrustScoreModRaw(_firstName, _lastName, _fullNames)
	scoreMod0, _firstNameMod0, _lastNameMod0, _fullNamesMod0 := calcTrustScoreMod0(_firstName, _lastName, _fullNames)
	scoreMod1 := calcTrustScoreMod1(_firstNameMod0, _lastNameMod0, _fullNamesMod0)
	scoreMod2 := calcTrustScoreMod2(_firstNameMod0, _lastNameMod0, _fullNamesMod0)
	scoreMod3 := calcTrustScoreMod3(_firstNameMod0, _lastNameMod0, _fullNamesMod0)

	return max(scoreModRaw, scoreMod0, scoreMod1, scoreMod2, scoreMod3)
}

func calcTrustScore(_firstName string, _lastName string, _fullNames []string, mod int) int {

	maxScore := 50
	switch mod {
	case -1:
		maxScore = 50
		break
	case 0:
		maxScore = 48
		break
	case 1:
		maxScore = 46
		break
	case 2:
		maxScore = 46
		break
	case 3:
		maxScore = 47
		break
	default:
		maxScore = 45
		break
	}

	fNameEquals := false
	lNameEquals := false
	score := 0

	for _, name := range _fullNames {
		//println(name, _firstName)
		//println(name, _lastName)
		if name == _firstName {
			fNameEquals = true
		} else if name == _lastName {
			lNameEquals = true
		}
	}

	if fNameEquals {
		score += maxScore
	} else {
		levenshteinDistance := 100
		for _, name := range _fullNames {
			d := LevenshteinDistance(name, _firstName)
			levenshteinDistance = min(levenshteinDistance, d)
		}
		if levenshteinDistance <= 2 {
			score += 45 - levenshteinDistance*10
		}
	}

	if lNameEquals {
		score += maxScore
	} else {
		levenshteinDistance := 100
		for _, name := range _fullNames {
			d := LevenshteinDistance(name, _lastName)
			levenshteinDistance = min(levenshteinDistance, d)
		}
		if levenshteinDistance <= 2 {
			score += 45 - levenshteinDistance*10
		}
	}

	return score
}
func calcTrustScoreModRaw(_firstName string, _lastName string, _fullNames []string) int {

	return calcTrustScore(_firstName, _lastName, _fullNames, -1)
}
func calcTrustScoreMod0(_firstName string, _lastName string, _fullNames []string) (int, string, string, []string) {

	_firstName = normalizeFioPartMod0(_firstName)
	_lastName = normalizeFioPartMod0(_lastName)

	for i, name := range _fullNames {
		_fullNames[i] = normalizeFioPartMod0(name)
	}

	return calcTrustScore(_firstName, _lastName, _fullNames, 0), _firstName, _lastName, _fullNames
}

func calcTrustScoreMod1(_firstName string, _lastName string, _fullNames []string) int {

	_firstName = normalizeFioPartMod1(_firstName)
	_lastName = normalizeFioPartMod1(_lastName)

	for i, name := range _fullNames {
		_fullNames[i] = normalizeFioPartMod1(name)
	}

	return calcTrustScore(_firstName, _lastName, _fullNames, 1)
}
func calcTrustScoreMod2(_firstName string, _lastName string, _fullNames []string) int {

	_firstName = normalizeFioPartMod2(_firstName)
	_lastName = normalizeFioPartMod2(_lastName)

	for i, name := range _fullNames {
		_fullNames[i] = normalizeFioPartMod2(name)
	}

	return calcTrustScore(_firstName, _lastName, _fullNames, 2)
}
func calcTrustScoreMod3(_firstName string, _lastName string, _fullNames []string) int {

	_firstName = normalizeFioPartMod3(_firstName)
	_lastName = normalizeFioPartMod3(_lastName)

	for i, name := range _fullNames {
		_fullNames[i] = normalizeFioPartMod3(name)
	}

	return calcTrustScore(_firstName, _lastName, _fullNames, 3)
}

func normalizeFioPart(input string) string {
	for _, trim := range fioTrims {
		input = strings.Replace(input, trim, "#", -1)
	}
	return input
}

func normalizeFioPartMod0(input string) string {
	for i := 0; i < len(fioMods); i++ {
		for find, replace := range fioMods[i] {
			input = strings.Replace(input, find, replace, -1)
		}
	}
	return input
}
func normalizeFioPartMod1(input string) string {
	for i := 0; i < len(fioMods1); i++ {
		for find, replace := range fioMods1[i] {
			input = strings.Replace(input, find, replace, -1)
		}
	}
	return input
}
func normalizeFioPartMod2(input string) string {
	for i := 0; i < len(fioMods2); i++ {
		for find, replace := range fioMods2[i] {
			input = strings.Replace(input, find, replace, -1)
		}
	}
	return input
}
func normalizeFioPartMod3(input string) string {
	for i := 0; i < len(fioMods3); i++ {
		for find, replace := range fioMods3[i] {
			input = strings.Replace(input, find, replace, -1)
		}
	}
	return input
}
