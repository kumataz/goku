package crypto

import (
    "crypto/aes"
    "crypto/cipher"
    "encoding/base64"
    "fmt"
    "os"
)

var commonIV = []byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f}

func Encrypt(key string, password []byte) []byte {
    // create aes
    c, err := aes.NewCipher([]byte(key))
    if err != nil {
        fmt.Printf("Error: NewCipher(%d bytes) = %s", len(key), err)
        os.Exit(-1)
    }

    // encrypt string
    cfb := cipher.NewCFBEncrypter(c, commonIV)
    ciphertext := make([]byte, len(password))
    cfb.XORKeyStream(ciphertext, password)
    return ciphertext
}

func Decrypt(key string, ciphercode []byte) []byte {
    // create aes
    c, err := aes.NewCipher([]byte(key))
    if err != nil {
        fmt.Printf("Error: NewCipher(%d bytes) = %s", len(key), err)
        os.Exit(-1)
    }

    // decrypt string
    cfbdec := cipher.NewCFBDecrypter(c, commonIV)
    password := make([]byte, 30)
    cfbdec.XORKeyStream(password, ciphercode)
    return password
}

func Base64(ciphercode []byte) string{
	return base64.StdEncoding.EncodeToString(ciphercode)
}


/*

func main() {
    // text length must be under 30
    text := []byte("i need crypto")
    // key string
    key := "s8*WQ0@KO#CN*raoua8ofCTx*oxqCk46"

    ciphercode := crypto.Encrypt(key, text)

	fmt.Printf("明文：      %s\n", text)
    fmt.Printf("hex密文   ：%02x\n", ciphercode)
    fmt.Printf("base64密文：%s\n", crypto.Base64(ciphercode))
	fmt.Printf("解密：%02x\n", crypto.Decrypt(key, ciphercode))
}

*/