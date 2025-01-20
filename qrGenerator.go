package main


import(
	"fmt"
)

func findVersion(inputBytes []byte){
	fmt.Println(cap(inputBytes))
}

func main(){
	var qrInput string
	fmt.Println("Input QR code string: ")
	fmt.Scanln(&qrInput)
	findVersion([]byte(qrInput))
}
