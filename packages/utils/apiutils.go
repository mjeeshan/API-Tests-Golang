package utils
import (
	"github.com/go-resty/resty/v2"

);
func PerformGetRequest(client *resty.Client) (*resty.Response, error) {
	return client.R().EnableTrace().
    Get("https://reqres.in/api/users")
}


func PerformGetRequestWithQueyParams(client *resty.Client, params string, headers map[string][]string) (*resty.Response, error) {
	return client.R().EnableTrace().
	SetQueryString(params).SetHeaderMultiValues(headers).
    Get("https://reqres.in/api/users")
}

func PerformPostRequest(client *resty.Client, body string, headers map[string][]string) (*resty.Response, error) {
	return client.R().EnableTrace().
	SetHeaderMultiValues(headers).
	SetBody(body).
    Post("https://reqres.in/api/users")
}