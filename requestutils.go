package utils

import (
	"net/http"
	"strings"
)

func PostJson(url, data string) error {
	r := strings.NewReader(data)
	res, err := http.Post(url, "application/json", r)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	return nil
}
