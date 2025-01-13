package main


import(
	"fmt"
)

func main(){
	var qrInput string
	fmt.Println("Input QR code string: ")
	fmt.Scanln(&qrInput)

	fmt.Println(qrInput)
}
