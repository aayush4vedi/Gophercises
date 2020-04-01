package splitwise

// User Model
type User struct {
	ID int `json:"ID"`
	Name string `json:"Name"`
}

//Balance Model

type Balance struct {
	UserId int
	Amount int
}

/*json
[
	user1: [
			{Rs 100 to user 2},
			{Rs 10 to user 3}
			],
	{ user 2 ows:
		{Rs 70 to user 3},
		~~{Rs -100 to user 1}~~ => shouldn't come. Auto settle
	},
]
 */

// Group
type Group struct {
	Users []User
	Balances map[int][]Balance
}

// User Actions

type Action int

const(
	_ Action = iota
	Exit
	Entry
	Balances
)

// expense settlements
type Settlement int

const(
	_ Settlement = iota
	Equal
	Exact
	Percent
)

//Entry models
// Equal: PaidBy: U1, AmountPaid: x
// Exact: PaidBy: U1, AmountU2: x, AmountU3: y, ...
// Percent: PaidBy: U1, PercentU2: x, PercentU3: y, ...