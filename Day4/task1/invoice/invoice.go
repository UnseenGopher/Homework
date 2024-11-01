package invoice

import (
	"errors"
	"reflect"
	"strconv"
	"unicode/utf8"
)

func inc(ind *int) int {
	*ind++
	return *ind
}

type Product struct {
	name string
}

func NewProduct(input string) (Product, error) {
	l := utf8.RuneCountInString(input)
	if l < 1 || l > 100 {
		return Product{}, errors.New("error, product name must have from 1 to 100 symbols")
	}
	return Product{
		name: input,
	}, nil
}

type Address struct {
	index     [6]uint
	city      string
	street    string
	house     string
	apartment string
}

func NewAdress(input [5]string) (Address, error) {
	var err_msg string
	val := reflect.ValueOf(Address{})
	for i, inp := range input {
		if len(inp) == 0 {
			err_msg += val.Type().Field(i).Name + " is empty\n"
		}
	}
	if utf8.RuneCountInString(input[0]) != 6 {
		err_msg += "index must have 6 symbols\n"
	}
	var index [6]uint
	for i, char := range input[0] {
		digit, err := strconv.Atoi(string(char))
		if err != nil || digit < 0 {
			err_msg += "invalid  index input: " + string(char) + "\n"
		} else {
			if i < 6 {
				index[i] = uint(digit)
			}
		}
	}
	// var j int = 1
	// cant use ++ operator while subscripting, holy hell
	if len(err_msg) == 0 {
		return Address{
			index:     index,
			city:      input[1],
			street:    input[2],
			house:     input[3],
			apartment: input[4],
		}, nil
	}
	return Address{}, errors.New(err_msg)

}

type FullName struct {
	surname    string
	name       string
	patronymic string
}

func NewFullName(input [3]string) (FullName, error) {
	var err_msg string
	val := reflect.ValueOf(FullName{})
	for i, inp := range input {
		if len(inp) == 0 {
			err_msg += val.Type().Field(i).Name + " is empty\n"
		}
	}
	var j = 0
	if len(err_msg) == 0 {
		return FullName{
			surname:    input[j],
			name:       input[inc(&j)],
			patronymic: input[inc(&j)],
		}, nil
	}
	return FullName{}, errors.New(err_msg)
}

type PhoneNumber struct {
	data [10]uint
}

func NewPhoneNumber(input string) (PhoneNumber, error) {

	var err_msg string
	var data [10]uint
	l := utf8.RuneCountInString(input)
	if l != 10 {
		err_msg += "phone number must have 10 digits\n"
	}
	for i, c := range input {
		digit, err := strconv.Atoi(string(c))
		if err != nil || digit < 0 {

			err_msg += "invalid digit at position: " + strconv.Itoa(i) + " in number"
		} else {
			if i < 10 {
				data[i] = uint(digit)
			}
		}

	}
	if len(err_msg) != 0 {
		return PhoneNumber{}, errors.New(err_msg)
	}
	return PhoneNumber{
		data: data,
	}, nil

}

type Client struct {
	credentials FullName
	number      PhoneNumber
	address     Address
}

func NewClient(credentials FullName, number PhoneNumber, address Address) (Client, error) {

	var err_msg string
	if (credentials == FullName{}) {
		err_msg += "empty FullName\n"
	}
	if (number == PhoneNumber{}) {
		err_msg += "empty phone number\n"
	}

	if (address == Address{}) {
		err_msg += "empty adrress\n"
	}

	if len(err_msg) != 0 {
		return Client{}, errors.New(err_msg)
	}

	return Client{
		credentials: credentials,
		number:      number,
		address:     address,
	}, nil
}

type Invoice struct {
	client   *Client
	product  *Product
	quantity uint
}

func NewInvoice(client *Client, product *Product, q int) (*Invoice, error) {
	var err_msg string
	if client == nil {
		err_msg += "internal error: nil client ptr\n"
	}

	if product == nil {
		err_msg += "internal error: nil product ptr\n"
	}
	if q < 1 || q > 100 {
		err_msg += "product quantity must be within [1,100]\n"
	}
	if len(err_msg) != 0 {
		return nil, errors.New(err_msg)
	}
	return &Invoice{
		client:   client,
		product:  product,
		quantity: uint(q),
	}, nil
}
