package external_service

import (
	"io/ioutil"
	"net/http"
)

/// ExternalService

type ExternalService interface {
	ReadPage(url string) string
}

/// YaSearch

type YaSearch struct {
	Service ExternalService
}

func (this *YaSearch) Get() string {
	return this.Service.ReadPage("http://www.ya555.com")
}

/// HttpRequest

type HttpRequest struct {
}

func (*HttpRequest) ReadPage(url string) string {
	var resp, err = http.Get(url)

	if err != nil {
		return "PageNotFound"
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return "PageNotFound"
	}

	body, _ := ioutil.ReadAll(resp.Body)

	return string(body)
}

