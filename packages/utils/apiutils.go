package utils
import (
	"github.com/go-resty/resty/v2"

);
func PerformGetRequestWIthQueryString(client *resty.Client) (*resty.Response, error) {
	return client.R(). EnableTrace().
    Get("https://reqres.in/api/users")


}