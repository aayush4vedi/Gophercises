package stock

import (
	"encoding/json"
	"fmt"
	"io"
	"math"
	"sort"
)

func JsonParser(r io.Reader) ([]Stock, error){
	d := json.NewDecoder(r)
	stocks := make([]Stock, 0)
	if err := d.Decode(&stocks); err != nil{
		return nil, err
	}
	return stocks,nil
}

var Portfolio = make(map[string][]float32)
var StocksPortfolio = make(map[string]float32)

func StockExchange(stocks []Stock){
	availableStocks := make([]Stock,0)
	for _, s := range stocks{
		if s.Action == "SELL"{
			availableStocks = append(availableStocks, s)
			sort.Slice(availableStocks, func(i,j int )bool{
				if availableStocks[i].Price == availableStocks[j].Price{
					return availableStocks[i].Time < availableStocks[j].Time
				}
				return availableStocks[i].Price < availableStocks[j].Price
			})
		}else{  //"BUY"ing case
			s0 := s
			qmax := s0.Quantity  // qumax = qmax/k
			//Hanles the case where time & price both are equal.So buy all of those in equal proportion
			//k := 0
			//for i := 0; i< len(availableStocks)-1; i++{
			//	for availableStocks[i].Price == availableStocks[i+1].Price && availableStocks[i].Time == availableStocks[i+1].Time{
			//		k++
			//	}
			//}
			//if k > 0{
			//	qmax = qmax/k
			//}

			for i, a := range availableStocks{
				if s0.Price >= a.Price && qmax >0 && a.Quantity > 0{
					p0 := a.Price
					q0 := math.Min(float64(qmax), float64(a.Quantity))
					qmax -= int(q0)
					availableStocks[i].Quantity -= int(q0)
					fmt.Printf("%d %d %v %v \n", a.OrderID, s0.OrderID, q0, p0)

					//Update the Portfolio
					Portfolio[a.Name] = append(Portfolio[a.Name], p0)
					StocksPortfolio[a.Name] += float32(q0)
				}
			}
		}
	}

	//Display Portfolios
	fmt.Println("*********** Displaying Portfolios ***************\n")
	for k, v := range Portfolio{
		fmt.Printf("Stock Name: %v \n", k)
		pmin := v[0]
		pmax := float32(0)
		psum := float32(0)
		for _, p := range v{
			//fmt.Printf("%v ",p)
			pmin = float32(math.Min(float64(p), float64(pmin)))
			pmax = float32(math.Max(float64(p), float64(pmax)))
			psum += p
		}
		pavg := psum/float32(len(v))
		fmt.Printf("Total Stocks Purchased: %v\n", StocksPortfolio[k])
		fmt.Printf("Average Buying Price: %v\n", pavg)
		fmt.Printf("Min Buying Price: %v\n", pmin)
		fmt.Printf("Max Buying Price: %v\n", pmax)
	}
}
