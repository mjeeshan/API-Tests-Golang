package main
import (
	"log"
	"io/ioutil"
	"fmt"
	"github.com/go-resty/resty/v2"
	"golearnings/packages/utils"
	"encoding/json"
);

type NewUser struct {
	Id string `json:"id"`
	CreatedAt string `json:"createdAt"`
}
func main()  {
	client := resty.New()
	res, _:= utils.PerformGetRequest(client)
	if res.StatusCode() != 200 {
		 fmt.Println("Response status  is not ok")
   }
   defer res.RawBody().Close();
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

   // creating header map
   headersMap := map[string][]string{
	"Accept": []string{
		"application/json"},
   };
   resp, _ := utils.PerformGetRequestWithQueyParams(client, "page=2", headersMap);

	if resp.StatusCode() != 200 {
		 fmt.Println("Response status  is not ok")
   }
   defer resp.RawBody().Close();
   fmt.Println(resp.StatusCode())

   json.Unmarshal(resp.Body(), &users)
   fmt.Printf("Page is %v, entries per page are %v, total pages are %v\n", users.Page, users.PerPage, users.TotalPages)

   for _, user := range users.Data {
	fmt.Printf("First name is %v. last name is %v, email is %v, and ID is %v\n", user.FirstName, user.LastName, user.Email, user.ID)
   }

  userResp, err := utils.PerformPostRequest(client, ReadFileAsString(), headersMap);
  if err != nil {
	  log.Panic(err)
  }
  defer userResp.RawBody().Close();

  newUser := NewUser{}
  json.Unmarshal(userResp.Body(), &newUser)
  fmt.Println(newUser.Id, newUser.CreatedAt)

}

func ReadFileAsString() string {
	data, err := ioutil.ReadFile("./resources/request.json")
	if err != nil {
		log.Fatal(err)
	}
	return string(data)
}






