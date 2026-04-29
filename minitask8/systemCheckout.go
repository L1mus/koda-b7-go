package minitask8

import (
	"errors"
	"fmt"
	"strconv"
)

func Reduce[T, M any](s []T, initValue M, f func(M, T) M) M {
	acc := initValue
	for _, v := range s {
		acc = f(acc, v)
	}
	return acc
}

func SystemCheckout(total, pay int, payment MethodPayment) {
	_, _, err := payment.pembayaran(total, pay)
	if err != nil {
		fmt.Println(err)
	}
}

type Bank struct {
	name string
	code string
}

func (b Bank) NewBank(name, code string) *Bank {
	return &Bank{
		name: name,
		code: code,
	}
}

func (b *Bank) pembayaran(price int, pay int) (int, string, error) {
	if pay < price {
		err := errors.New("Pembayaran tidak dapat dilakukan karena dana tidak mencukupi")
		return price, "", err
	}
	calculate := pay - price
	massage, _ := fmt.Printf("Pembayaran berhasil dengan total harga yang harus dibayar sebesar %d \n Terima kasih telah menggunakan kami %s sebagai metode pembayaran Anda ", price, b.code)
	return calculate, strconv.Itoa(massage), nil

}

type Online struct {
	name string
	code string
}

func (o Online) NewOnline(name, code string) *Online {
	return &Online{
		name: name,
		code: code,
	}
}

func (o *Online) pembayaran(price int, pay int) (int, string, error) {
	if pay < price {
		err := errors.New("Pembayaran tidak dapat dilakukan karena dana tidak mencukupi")
		return price, "", err
	}
	calculate := pay - price
	massage, _ := fmt.Printf("Pembayaran berhasil dengan total harga yang harus dibayar sebesar %d \n Terima kasih telah menggunakan kami %s sebagai metode pembayaran Anda", price, o.code)
	return calculate, strconv.Itoa(massage), nil
}

type MethodPayment interface {
	pembayaran(price int, pay int) (int, string, error)
}
