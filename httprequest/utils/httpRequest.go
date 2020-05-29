package httpRequest
import (
    "encoding/json"
    "net/http"
    "io/ioutil"
)


req := http.NewRequest("GET", "http://test.com", nil)
req.Header.Set("Content-type", "application/json")
client := &http.Client{}
response,err := client.Do(req)

if err == nil {
    // 解析Response
    returnMap,err := util.ParseResponse(response)
}


func ParseResponse(response *http.Response) (map[string]interface{}, error){
	var result map[string]interface{}
	body,err := ioutil.ReadAll(response.Body)
	if err == nil {
		err = json.Unmarshal(body, &result)
	}

	return result,err
}
