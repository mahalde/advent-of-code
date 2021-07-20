package req

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/mahalde/advent-of-code/utils/secrets"
)

func MakeRequest(year, day int) string {
	url := fmt.Sprintf("https://adventofcode.com/%v/day/%v/input", year, day)
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		panic(err)
	}

	// Add the session cookie
	req.AddCookie(&http.Cookie{
		Name:  "session",
		Value: secrets.SessionID,
	})

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}

	return string(body)
}
