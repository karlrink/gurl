package http

import (
    "fmt"
    "net/http"
)

func Request(method string, url string, header map[string]string, data map[string]string) (*http.Response, error) {

    // don't follow redirect
    client := &http.Client {
        CheckRedirect: func(req *http.Request, via []*http.Request) error {
            return http.ErrUseLastResponse
        },
    }

    req, _ := http.NewRequest(method, url, nil)

    for key, value := range header {
        req.Header.Set(key, value)
    }

    response, err := client.Do(req)
    if err != nil {
        return nil, fmt.Errorf("Client Error making HTTP request: %v", err)
    }

    return response, nil
}
