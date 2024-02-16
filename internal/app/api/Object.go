package api

import (
	"fmt"
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

func (api *API) GetCategory(writer http.ResponseWriter, req *http.Request) {
	initHeaders(writer)

	vars := mux.Vars(req)
	lang, ok := vars["lang"]
	if !ok {
		fmt.Println("id is missing in parameters")
	}

	url := "https://suppliers-api.wildberries.ru/content/v2/object/parent/all"

	url = url + "?locale=" + lang
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
func (api *API) rrr(writer http.ResponseWriter, req *http.Request) {
	initHeaders(writer)

	url := "https://suppliers-api.wildberries.ru/content/v2/object/all"

	r, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
	}

	// add authorization header to the req
	r.Header.Add("Authorization", "eyJhbGciOiJFUzI1NiIsImtpZCI6IjIwMjMxMDI1djEiLCJ0eXAiOiJKV1QifQ.eyJlbnQiOjEsImV4cCI6MTcxOTAwMTU5MywiaWQiOiJjZDkzMDhhYS1kYWE1LTRkOWItYTYyMC0zMDU4NDQzNDJkODkiLCJpaWQiOjQ3MTk2MzE5LCJvaWQiOjQzNDU3OCwicyI6NTEwLCJzaWQiOiI1ZWM2N2Q4Zi02MTdiLTQ2YmQtOTJhYi1iMmQ3Yzc0NGIwYmUiLCJ0IjpmYWxzZSwidWlkIjo0NzE5NjMxOX0.kWqbfxpl_1U1DHn_Yn93GVFYAPRWeU2DDg3zndcGhq93jSlChCgvU1EZNNHBf4Pp7Umi4XzQMKhcFwv7Bh9kvA")
	if err != nil {
		fmt.Println(err)

	}

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
