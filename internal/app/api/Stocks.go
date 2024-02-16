package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (api *API) TestPut(writer http.ResponseWriter, req *http.Request) {
	initHeaders(writer)

	id, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		api.logger.Info("Unable to parse request target")
		msg := Message{
			StatusCode: 400,
			Message:    "Bad request id",
			IsError:    true,
		}
		writer.WriteHeader(400)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	url := "https://suppliers-api.wildberries.ru/api/v3/stocks/" + strconv.Itoa(id)
	fmt.Println(url)
	body, error := ioutil.ReadAll(req.Body)
	if error != nil {
		return
	}
	fmt.Println(string(body))

	client := &http.Client{}

	r, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(body))
	if err != nil {
		fmt.Println(err)
	}
	r.Header.Add("Authorization", "eyJhbGciOiJFUzI1NiIsImtpZCI6IjIwMjMxMDI1djEiLCJ0eXAiOiJKV1QifQ.eyJlbnQiOjEsImV4cCI6MTcxOTAwMTU5MywiaWQiOiJjZDkzMDhhYS1kYWE1LTRkOWItYTYyMC0zMDU4NDQzNDJkODkiLCJpaWQiOjQ3MTk2MzE5LCJvaWQiOjQzNDU3OCwicyI6NTEwLCJzaWQiOiI1ZWM2N2Q4Zi02MTdiLTQ2YmQtOTJhYi1iMmQ3Yzc0NGIwYmUiLCJ0IjpmYWxzZSwidWlkIjo0NzE5NjMxOX0.kWqbfxpl_1U1DHn_Yn93GVFYAPRWeU2DDg3zndcGhq93jSlChCgvU1EZNNHBf4Pp7Umi4XzQMKhcFwv7Bh9kvA")
	r.Header.Set("Content-Type", "application/json")

	response, err := client.Do(r)
	if err != nil {
		fmt.Println("error sending request: ", err)
		http.Error(writer, "internal error", http.StatusInternalServerError)
		return
	}

	defer response.Body.Close()
	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("error reading response data: ", err)
		http.Error(writer, "internal error", http.StatusInternalServerError)
		return
	}
	fmt.Println("writing response")
	writer.WriteHeader(response.StatusCode)
	writer.Write(responseData)
}
