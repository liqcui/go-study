package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func EncryptAES(key []byte, plaintext string) string {

	c, err := aes.NewCipher(key)
	CheckError(err)

	out := make([]byte, len(plaintext))

	c.Encrypt(out, []byte(plaintext))

	return hex.EncodeToString(out)
}

func DecryptAES(key []byte, ct string) {
	ciphertext, _ := hex.DecodeString(ct)

	c, err := aes.NewCipher(key)
	CheckError(err)

	pt := make([]byte, len(ciphertext))
	c.Decrypt(pt, ciphertext)

	s := string(pt[:])
	fmt.Println("DECRYPTED:", s)
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func createHash(key string) string {
	hasher := md5.New()
	hasher.Write([]byte(key))
	return hex.EncodeToString(hasher.Sum(nil))
}

func encryptStr(data []byte, passphrase string) []byte {
	block, _ := aes.NewCipher([]byte(createHash(passphrase)))
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}
	ciphertext := gcm.Seal(nonce, nonce, data, nil)
	return ciphertext
}

func decryptStr(data []byte, passphrase string) []byte {
	key := []byte(createHash(passphrase))
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	nonceSize := gcm.NonceSize()
	nonce, ciphertext := data[:nonceSize], data[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		panic(err.Error())
	}
	return plaintext
}

func encryptFile(filename string, data []byte, passphrase string) {
	f, _ := os.Create(filename)
	defer f.Close()
	f.Write(encryptStr(data, passphrase))
}

func decryptFile(filename string, passphrase string) []byte {
	data, _ := ioutil.ReadFile(filename)
	return decryptStr(data, passphrase)
}

// base解码
func BASE64DecodeStr(src string) string {
	a, err := base64.StdEncoding.DecodeString(src)
	if err != nil {
		return ""
	}
	return string(a)
}

func main() {

	// plaintext
	dockerConfigPT := "ewoJImF1dGhzIjogewoJCSJkb2NrZXIuaW8iOiB7CgkJCSJhdXRoIjogIlpHRm5jbUY1T2pRMlltSmpOekkxTFROaE1HSXRORGs1WlMxaU1qUmtMVFJoTjJZME5XWmtZakJsTlE9PSIKCQl9LAoJCSJodHRwczovL2luZGV4LmRvY2tlci5pby92MS8iOiB7CgkJCSJhdXRoIjogIlpHRm5jbUY1T2pRMlltSmpOekkxTFROaE1HSXRORGs1WlMxaU1qUmtMVFJoTjJZME5XWmtZakJsTlE9PSIKCQl9Cgl9Cn0K"
	// cipher key
	// key := "redhatopenshiftpsapqesrophrase32"

	// c := EncryptAES([]byte(key), dockerConfigPT)

	// // plaintext
	// fmt.Println(dockerConfigPT)

	// // ciphertext
	// fmt.Println(c)

	// // decrypt
	// DecryptAES([]byte(key), c)
	//dockerConfigPT := "ZGFncmF5OjQ2YmJjNzI1LTNhMGItNDk5ZS1iMjRkLTRhN2Y0NWZkYjBlNQ=="
	fmt.Println("Starting the application...")
	ciphertext := encryptStr([]byte(dockerConfigPT), "openshift")
	fmt.Printf("Encrypted: %x\n", ciphertext)
	plaintext := decryptStr(ciphertext, "openshift")
	fmt.Printf("Decrypted: %s\n", plaintext)
	encryptFile("/tmp/multibuild-pull-secret.crypt", []byte("ewoJImF1dGhzIjogewoJCSJkb2NrZXIuaW8iOiB7CgkJCSJhdXRoIjogIlpHRm5jbUY1T2pRMlltSmpOekkxTFROaE1HSXRORGs1WlMxaU1qUmtMVFJoTjJZME5XWmtZakJsTlE9PSIKCQl9LAoJCSJodHRwczovL2luZGV4LmRvY2tlci5pby92MS8iOiB7CgkJCSJhdXRoIjogIlpHRm5jbUY1T2pRMlltSmpOekkxTFROaE1HSXRORGs1WlMxaU1qUmtMVFJoTjJZME5XWmtZakJsTlE9PSIKCQl9Cgl9Cn0K"), "controller-manager")
	fmt.Println(string(decryptFile("/tmp/multibuild-pull-secret.crypt", "controller-manager")))

	decodeBase64 := BASE64DecodeStr(string(plaintext))
	fmt.Print(decodeBase64)
}
