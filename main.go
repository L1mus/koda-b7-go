package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/L1mus/koda-b7-go/minitask1"
	"github.com/L1mus/koda-b7-go/minitask2"
	minitask3 "github.com/L1mus/koda-b7-go/minitask3"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Masukan Fungsi yang anda ingin lakukan : ")
	fmt.Println("A.Calculate Circle\n B.Create Triangle\n C.Insert new Number To Array")
	scanner.Scan()
	whichSelected := scanner.Text()

	if whichSelected == "A" {
		fmt.Print("Masukan Jari Jari Lingkaran : ")
		scanner.Scan()
		r := scanner.Text()
		f64, err := strconv.ParseFloat(r, 32)
		if err != nil {
			fmt.Println("error", err)
			return
		}
		luas, keliling := minitask1.CalculateCircle(float32(f64))
		fmt.Printf("Keliling Lingkaran %.2f , Luas Lingkaran %.2f\n", keliling, luas)
	} else if whichSelected == "B" {
		fmt.Print("Masukan panjang segitiga : ")
		scanner.Scan()
		n := scanner.Text()
		num, err := strconv.Atoi(n)
		if err != nil {
			fmt.Println("error", err)
			return
		}
		minitask2.CreateTriangel(num)
	} else if whichSelected == "C" {
		fmt.Print("Masukan Array yang ingin di insert dengan format angka integer di pisahkan dengan spasi: ")
		scanner.Scan()
		input := scanner.Text()
		parts := strings.Fields(input)
		var arr []int
		for _, part := range parts {
			num, err := strconv.Atoi(part)
			if err == nil {
				arr = append(arr, num)
			}
		}
		fmt.Print("Masukan dari index berapa yang ingin di insert : ")
		scanner.Scan()
		fromIndexStr := scanner.Text()
		fromIndex, _ := strconv.Atoi(fromIndexStr)
		fmt.Print("Masukan Angka yang ingin dimasukan bisa satu atau lebih angka dengan format yang sama seperti sebelunya : ")
		scanner.Scan()
		inputElement := scanner.Text()
		elements := strings.Fields(inputElement)
		var arrElements []int
		for _, element := range elements {
			num, err := strconv.Atoi(element)
			if err == nil {
				arrElements = append(arrElements, num)
			}
		}
		minitask3.InsertNumber(arr, fromIndex, arrElements)
	}

}
