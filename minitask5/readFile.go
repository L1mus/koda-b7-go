package minitask5

import (
	"log"
	"os"
)

func ReadFile(filepath string) (string, error) {
	//scenner := bufio.NewScanner(os.Stdin)
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	content := make([]byte, 1024)
	data, err := file.Read(content)
	if err != nil {
		log.Println("error to Read file")
	}
	return string(content[:data]), nil
}
