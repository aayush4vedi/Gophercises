package  main

import (
	"flag"
	"fmt"
	"github.com/stock"
	"os"
)

func main(){
	filename := flag.String("file", "data.json", "Enter input file")
	flag.Parse()
	fmt.Printf("Using inputs in file %v\n", *filename)

	f, err := os.Open(*filename)
	if err != nil{
		panic(err)
	}
	stocks, err :=  stock.JsonParser(f)
	if err != nil{
		panic(err)
	}
	stock.StockExchange(stocks)
	os.Exit(1)
}