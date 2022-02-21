package utils

import (
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
)

func Get(url string) ([]byte, error) {
	var bytes = []byte{}
	resp, err := http.Get(url)
	if err != nil {
		return bytes, err
	}
	defer resp.Body.Close()
	return checkStatus(resp)
}

func checkStatus(h *http.Response) ([]byte, error) {
	if h.StatusCode == http.StatusOK {
		dataByte, err := ioutil.ReadAll(h.Body)
		if err != nil {
			return []byte{}, err
		}
		return dataByte, nil
	} else {
		return []byte{}, errors.New(strconv.Itoa(http.StatusNotFound))
	}
}
