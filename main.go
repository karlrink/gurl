
package main

import (
    "fmt"
    "io/ioutil"
    "os"
    "github.com/karlrink/gurl/http"
    "github.com/karlrink/gurl/cmd"
)

func main() {

    fmt.Printf("Random string: %s\n", cmd.RandStr(20))

    url := os.Args[1]

    header := map[string]string{
        "Content-Type": "text/html;charset=utf-8",
	}

    data := map[string]string{}

    method := "GET"

    response, err := http.Request(method, url, header, data)
    if err != nil {
        fmt.Printf("Failed GET: %v \n", err)
        os.Exit(1)
    }

    for name, values := range response.Header {
    // Loop over all values for the name.
        for _, value := range values {
            fmt.Println(name, value)
        }
    }


    content, err := ioutil.ReadAll(response.Body)
	if err != nil {
	    fmt.Printf("Read body: %v", err)
    }

    responseString := string(content)

    fmt.Printf("%s", responseString)

    fmt.Println(response.StatusCode)


}
