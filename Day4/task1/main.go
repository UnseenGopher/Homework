package main

import (
	"bufio"
	"flag"
	"fmt"
	_ "fmt"
	"os"
	"strconv"
	inv "task1/invoice"
)

func PrintErrExit(err error, topic string) {
	if err != nil {
		fmt.Println("Error forming " + topic + ": " + err.Error())
		os.Exit(1)
	} else {
		fmt.Println("Succefully formed: " + topic)
	}
}

func usage() {
	fmt.Fprintf(os.Stderr, "usage: progname [inputfile]\n")
	flag.PrintDefaults()
	os.Exit(2)
}
func main() {
	flag.Usage = usage
	flag.Parse()
	args := flag.Args()
	if len(args) < 1 {
		fmt.Println("Input file is missing.")
		os.Exit(1)
	}
	file, err := os.Open(args[0])
	if err != nil {
		fmt.Println(err.Error())
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var input []string
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	if len(input) < 11 {
		fmt.Println("Error, not enough inputs in file")
	}
	var in1 [5]string
	copy(in1[:], input[0:5])
	in2 := input[5]
	var in3 [3]string
	copy(in3[:], input[6:9])
	in4 := input[9]

	quantity, err := strconv.Atoi(input[10])

	if err != nil {
		fmt.Println("Error creating int :", err.Error())
	}

	addr, err := inv.NewAdress(in1)
	PrintErrExit(err, "address")

	phone, err := inv.NewPhoneNumber(in2)
	PrintErrExit(err, "phone number")

	fullname, err := inv.NewFullName(in3)
	PrintErrExit(err, "full name")

	client, err := inv.NewClient(fullname, phone, addr)
	PrintErrExit(err, "client")

	product, err := inv.NewProduct(in4)
	PrintErrExit(err, "product")

	invoice, err := inv.NewInvoice(&client, &product, quantity)
	PrintErrExit(err, "invoice")

	_ = invoice
}
