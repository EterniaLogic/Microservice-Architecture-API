package common

// User: Brent Clancy (EterniaLogic)
// Date: 12/11/2015

import (
    "crypto/aes"
    "crypto/cipher"
    "crypto/rand"
    "encoding/base64"
	//"encoding/hex"
	"errors"
	"io"
	//"log"
)

var cipheraes = "IEPNlw78I5KsW64xFzswTBaXcYyftZEa";
var iv = []byte{34, 62, 35, 57, 10, 23, 31, 36, 65, 56, 45, 23, 32, 86, 31, 15};

//http://stackoverflow.com/questions/18817336/golang-encrypting-a-string-with-aes-and-base64
func encrypt(key, text []byte) ([]byte, error) {
    block, err := aes.NewCipher(key)
    if err != nil {
        return nil, err
    }
    b := base64.StdEncoding.EncodeToString(text)
    ciphertext := make([]byte, aes.BlockSize+len(b))
    iv := ciphertext[:aes.BlockSize]
    if _, err := io.ReadFull(rand.Reader, iv); err != nil {
        return nil, err
    }
    cfb := cipher.NewCFBEncrypter(block, iv)
    cfb.XORKeyStream(ciphertext[aes.BlockSize:], []byte(b))
    return ciphertext, nil
}

//http://stackoverflow.com/questions/18817336/golang-encrypting-a-string-with-aes-and-base64
func decrypt(key, text []byte) ([]byte, error) {
    block, err := aes.NewCipher(key)
    if err != nil {
        return nil, err
    }
    if len(text) < aes.BlockSize {
        return nil, errors.New("ciphertext too short")
    }
    iv := text[:aes.BlockSize]
    text = text[aes.BlockSize:]
    cfb := cipher.NewCFBDecrypter(block, iv)
    cfb.XORKeyStream(text, text)
    data, err := base64.StdEncoding.DecodeString(string(text))
    if err != nil {
        return nil, err
    }
    return data, nil
}


// Encrypt and ID with AES-256
func EncryptID(id string)(idenc string){
	/*cc,err := encrypt([]byte(cipheraes),[]byte(id));
	if(err != nil){
		common.AsyncPrintln("EncryptID error:",err);
	}
	return hex.EncodeToString(cc);*/
	return id;
}

// Encrypt and ID with AES-256
func DecryptID(id string)(idenc string){
	/*cc,err := decrypt([]byte(cipheraes),[]byte(id));
	if(err != nil){
		common.AsyncPrintln("DecryptID error:",err);
	}
	return hex.EncodeToString(cc);*/
	return id;
}