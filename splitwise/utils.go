package splitwise

import (
	"encoding/json"
	"fmt"
	"io"
)

func UserParser(r io.Reader)([]User, error){
	d := json.NewDecoder(r)
	users := make([]User, 0)
	if err := d.Decode(&users); err != nil{
		return nil, err
	}
	return users, nil
}

func PrintUsers(users []User){
	for _, u := range users{
		fmt.Printf("%+v\n", u)
	}
}





































