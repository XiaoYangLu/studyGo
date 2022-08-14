package main

import (
	"fmt"
	"studyGo/packages/hash"
	"studyGo/packages/symmetricEncryptionAlgorithm"
)

func main() {
	text := "一篇诗，一斗酒，一曲长歌，一件天涯"
	hashStr := hash.StudyHASH(text, "sha256")
	fmt.Println(hashStr)
	symmetricEncryptionAlgorithm.StudyAES(text, "des", []byte("00000000"))
}
