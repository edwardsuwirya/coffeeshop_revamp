package models

type Coffee interface {
	ProductInfo() string
	Brew(status chan string)
}
