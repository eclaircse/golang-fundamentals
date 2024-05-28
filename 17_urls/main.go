package main

import (
	"fmt"
	"net/url"
)

const testUrl = "https://api.github.com:3000/users/eclairjit?username=eclairjit&repoCount=14"

func main() {

	result, _ := url.Parse(testUrl)

	fmt.Println(result.Scheme) // https
	fmt.Println(result.Host) // api.github.com:3000
	fmt.Println(result.Path) // /users/eclairjit
	fmt.Println(result.RawQuery) // username=eclairjit&repoCount=14
	fmt.Println(result.Port()) // 3000

	params := result.Query()

	fmt.Println(params) // map[repoCount:[14] username:[eclairjit]]
	fmt.Println(params["username"]) // [eclairjit]

	// creating an url

	partsOfUrl := &url.URL{
		Scheme: "https",
		Host: "localhost:3000",
		Path: "/user",
		RawQuery: "username=eclair",
	}

	anotherUrl := partsOfUrl.String()
	fmt.Println(anotherUrl)

}