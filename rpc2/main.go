package main

import (
	"log"
	calculator "rpctest2/kitex_gen/calculator/calculator"
)

func main() {
	svr := calculator.NewServer(new(CalculatorImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
