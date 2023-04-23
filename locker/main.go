package main

import (
	"fmt"
	"locker/utils"
	"log"
)

func main() {

	filePath := "C:/Users/georg/Dendron-1/dependencies/localhost/redis/notes/crypto/wallet.md"
	password := "ThaDiTaIn"

	buffer, err := utils.ReadFromFile(filePath)
	if err != nil {
		return
	}

	fileInfo, err := utils.GetFileStat(filePath)
	if err != nil {
		return
	}

	newFileName := utils.ChangeFileExtension(fileInfo.Name(), ".aed")
	file, err := utils.CreateFile(newFileName)
	if err != nil {
		return
	}

	utils.GenerateTextToHash("ThaDiTaIn@1110")

	log.Println("New File Name:", utils.ChangeFileExtension(fileInfo.Name(), ".aed"))

	log.Printf("%s\n", buffer)

	log.Printf("file : %v\n", file)

	hash := utils.GenerateTextToHash(password)

	encryptedText := utils.Encrypt(fmt.Sprintf("%s", &buffer), hash)

	log.Println("encryptedText :", encryptedText)

	err = utils.WriteToFile(file, encryptedText)
	if err != nil {
		return
	}

	encryptedBuffer, err := utils.ReadFromFile(newFileName)
	if err != nil {
		return
	}

	log.Printf("encryptedBuffer :%s\n", encryptedBuffer)

	decryptedText := utils.Decrypt(encryptedBuffer, hash)

	log.Println("decryptedText :", decryptedText)
}
