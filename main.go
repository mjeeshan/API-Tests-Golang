package main
import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"golearnings/packages/utils"
	"encoding/json"
);
func main()  {
	client := resty.New()
	res, _:= utils.PerformGetRequestWIthQueryString(client)
	if res.StatusCode() != 200 {
		 fmt.Println("Response status  is not ok")
   }
   type Data struct {
	ID int `json:"id"`
	Email string `json:"email"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Avatar string `json:"avatar"`
}
   type UserResponse struct {
		Page int `json:"page"`
		PerPage int `json:"per_page"`
		TotalPages int `json:"total_pages"`
		Data []Data
   }
   users := UserResponse{}
   json.Unmarshal(res.Body(), &users)
   fmt.Printf("Page is %v, entries per page are %v, total pages are %v\n", users.Page, users.PerPage, users.TotalPages)

   for _, user := range users.Data {
	fmt.Printf("First name is %v. last name is %v, email is %v, and ID is %v\n", user.FirstName, user.LastName, user.Email, user.ID)
   }
}



