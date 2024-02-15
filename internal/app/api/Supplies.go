package api

import (
	"Apifbs/internal/app/models"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (api *API) GetSupplies(writer http.ResponseWriter, req *http.Request) {
	initHeaders(writer)
	var (
		filter models.PageRequest
	)

	url := "https://suppliers-api.wildberries.ru/api/v3/supplies"

	err := json.NewDecoder(req.Body).Decode(&filter)
	if err != nil {
		api.logger.Info("Invalid json recieved from products")
		msg := Message{
			StatusCode: 400,
			Message:    "Provided json is invalid",
			IsError:    true,
		}
		writer.WriteHeader(400)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	body, error := ioutil.ReadAll(req.Body)
	if error != nil {
		return
	}
	fmt.Println(body)
	var jsonByte, _ = json.Marshal(filter)

	fmt.Println(string(jsonByte))
	url = url + "?limit=" + strconv.Itoa(filter.Limit) + "&next=" + strconv.Itoa(filter.Next)
	fmt.Println(url)
	r, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		fmt.Println(err)
	}
	r.Header.Add("Authorization", "eyJhbGciOiJFUzI1NiIsImtpZCI6IjIwMjMxMDI1djEiLCJ0eXAiOiJKV1QifQ.eyJlbnQiOjEsImV4cCI6MTcxOTAwMTU5MywiaWQiOiJjZDkzMDhhYS1kYWE1LTRkOWItYTYyMC0zMDU4NDQzNDJkODkiLCJpaWQiOjQ3MTk2MzE5LCJvaWQiOjQzNDU3OCwicyI6NTEwLCJzaWQiOiI1ZWM2N2Q4Zi02MTdiLTQ2YmQtOTJhYi1iMmQ3Yzc0NGIwYmUiLCJ0IjpmYWxzZSwidWlkIjo0NzE5NjMxOX0.kWqbfxpl_1U1DHn_Yn93GVFYAPRWeU2DDg3zndcGhq93jSlChCgvU1EZNNHBf4Pp7Umi4XzQMKhcFwv7Bh9kvA")
	client := &http.Client{}
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
func (api *API) DeleteSupplies(writer http.ResponseWriter, req *http.Request) {
	initHeaders(writer)

	url := "https://suppliers-api.wildberries.ru/api/v3/supplies/"

	vars := mux.Vars(req)
	id, ok := vars["id"]
	if !ok {
		fmt.Println("id is missing in parameters")
	}

	url = url + id
	fmt.Println(url)
	r, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		fmt.Println(err)
	}
	r.Header.Add("Authorization", "eyJhbGciOiJFUzI1NiIsImtpZCI6IjIwMjMxMDI1djEiLCJ0eXAiOiJKV1QifQ.eyJlbnQiOjEsImV4cCI6MTcxOTAwMTU5MywiaWQiOiJjZDkzMDhhYS1kYWE1LTRkOWItYTYyMC0zMDU4NDQzNDJkODkiLCJpaWQiOjQ3MTk2MzE5LCJvaWQiOjQzNDU3OCwicyI6NTEwLCJzaWQiOiI1ZWM2N2Q4Zi02MTdiLTQ2YmQtOTJhYi1iMmQ3Yzc0NGIwYmUiLCJ0IjpmYWxzZSwidWlkIjo0NzE5NjMxOX0.kWqbfxpl_1U1DHn_Yn93GVFYAPRWeU2DDg3zndcGhq93jSlChCgvU1EZNNHBf4Pp7Umi4XzQMKhcFwv7Bh9kvA")
	client := &http.Client{}
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
