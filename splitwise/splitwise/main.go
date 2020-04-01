package  main

import (
	"flag"
	"fmt"
	"github.com/splitwise"
	"os"
)


func main(){
	filename := flag.String("file", "data.json", "The json file you want to pass")
	flag.Parse()

	//read input
	f, err := os.Open(*filename)
	if err != nil{
		panic(err)
	}

	users, err := splitwise.UserParser(f)
	if err != nil{
		panic(err)
	}

	group := splitwise.Group{
		Users:    users,
		Balances: make(map[int][]splitwise.Balance),
	}
	//fmt.Printf("%+v\n", group)

	splitwise.Setup(&group)
	var input splitwise.Action
	fmt.Println("Press 2 to begin")
	fmt.Scanf("%d\n", &input)

	for input != splitwise.Exit{
		switch input {
		case splitwise.Exit:
			fmt.Println("You pressed 1. Exiting now")
		case splitwise.Entry:
			fmt.Println("You pressed 2. Make Entry:: \n>> Choose Settlements-  Equal-1: Exact-2, Percentage-3")
			var settlement splitwise.Settlement
			fmt.Scanf("%d\n", &settlement)
			if settlement == 1{
				fmt.Printf("Paid By?\n")
				var u int
				fmt.Scanf("%d\n", &u)
				fmt.Printf("Amount?\n")
				var a int
				fmt.Scanf("%d\n", &a)
				for v, bal := range group.Balances{
					if v == u{
						for i, _ := range bal{
							bal[i].Amount += a/ len(group.Users)
						}
					}
				}
				for u, bal := range group.Balances{
					fmt.Printf("User: %s",u," ows :\n")
					for _, b := range bal{
						fmt.Printf(" Rs. %d to %v\n", b.Amount ,b.UserId)
					}
				}
			}else if settlement == 2{

			}
		case splitwise.Balances:
			fmt.Println("Pressed 3. Showing Balances")
			splitwise.ShowBalance(&group)
		default:
			fmt.Println("Didn't get you.Press 0,1,2 only")
		}
		fmt.Println("Options at any point: 1 to quit, 2 to Make an entry, 3 : show BalanceSheet")
		fmt.Scanf("%d\n", &input)
	}
	fmt.Println("Bye!")
	os.Exit(1)
}


