package splitwise

import "fmt"

func Setup(group *Group){
	for _, u := range group.Users{
		for _, v := range group.Users{
			if u.ID != v.ID{
				newB := Balance{
					UserId:v.ID,
					Amount:0,
				}
				group.Balances[u.ID] = append(group.Balances[u.ID], newB)
			}
		}
	}
}

func ShowBalance(group *Group){
	for u, bal := range group.Balances{
		fmt.Printf("User: %s",u," owns :\n")
		for _, b := range bal{
			fmt.Printf(" Rs. %d from %v\n", b.Amount ,b.UserId)
		}
	}
}

func MakeEntry(group *Group, settlement Settlement){
	switch settlement {
	case Equal:
		EqualSettlement(group)
	//case Exact:
	//	ExactSettlement(group)
	//case Percent:
	//	PercentSettlement(group)
	default:
		fmt.Println("Invalid Input.")
	}
}

func EqualSettlement(group *Group)  {
	fmt.Printf("Paid By?\n")
	var u int
	fmt.Scanf("%d\n", u)
	fmt.Printf("Amount?\n")
	var a int
	fmt.Scanf("%d\n", a)
	for _, bal := range group.Balances{
		for _, b := range bal{
			b.Amount += a/ len(group.Users)
		}
	}
}
