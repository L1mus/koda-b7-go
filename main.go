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
	"github.com/L1mus/koda-b7-go/minitask5"
	"github.com/L1mus/koda-b7-go/minitask7"
	"github.com/L1mus/koda-b7-go/minitask8"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Masukan Fungsi yang anda ingin lakukan : ")
	fmt.Println("A.Calculate Circle\n B.Create Triangle\n C.Insert new Number To Array\n D.Read File")
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
	} else if whichSelected == "D" {
		fmt.Print("Masukan path file yang ingin di baca isinya di console: ")
		scanner.Scan()
		input := scanner.Text()
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recover panic", r)
			}
		}()

		data, err := minitask5.ReadFile(input)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println(data)
	} else if whichSelected == "E" {
		fmt.Print("Masukan nama anda: ")
		scanner.Scan()
		inputNama := scanner.Text()
		fmt.Print("Masukan Alamat anda: ")
		scanner.Scan()
		inputAlamat := scanner.Text()
		fmt.Print("Masukan nomor telepon anda: ")
		scanner.Scan()
		inputTelepon := scanner.Text()

		minitask7.CreatePersonData(inputNama, inputAlamat, inputTelepon)

	} else if whichSelected == "F" {
		fmt.Print("Masukan Harga Barang yang di beli dengan format angka integer di pisahkan dengan spasi antar harga barang lainnya : ")
		scanner.Scan()
		input := scanner.Text()
		parts := strings.Fields(input)
		var data []int
		for _, part := range parts {
			num, err := strconv.Atoi(part)
			if err == nil {
				data = append(data, num)
			}
		}

		for idx, v := range data {
			fmt.Printf("Harga barang ke-%d: %d\n", idx+1, v)
		}

		total := minitask8.Reduce(data, 0, func(acc, initValue int) int {
			return acc + initValue
		})

		fmt.Print("Pilih Metode Pembayaran Bank atau Online:")
		scanner.Scan()
		methodPayment := scanner.Text()
		if strings.ToLower(methodPayment) == "bank" {
			b := minitask8.Bank{}.NewBank("Bank Rakyat Indonesia", "BRI")
			fmt.Print("Masukan Uang anda : ")
			scanner.Scan()
			value := scanner.Text()
			pay, _ := strconv.Atoi(value)
			if pay < total {
				println("Uang anda Kurang!")
				return
			}
			minitask8.SystemCheckout(total, pay, b)

		} else if strings.ToLower(methodPayment) == "online" {
			o := minitask8.Online{}.NewOnline("OVO", "OVO")
			fmt.Print("Masukan Uang anda : ")
			scanner.Scan()
			value := scanner.Text()
			pay, _ := strconv.Atoi(value)
			if pay < total {
				println("Uang anda Kurang!")
				return
			}
			minitask8.SystemCheckout(total, pay, o)
		}
	}
}
