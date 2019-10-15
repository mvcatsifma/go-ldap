package main

import (
	"gopkg.in/korylprince/go-ad-auth.v2"
	"log"
	"os"
)

// username
// password
// server
// port
// groups
func main() {
	// without tls
	config := &auth.Config{
		Server:   "localhost",
		Port:     1389,
		BaseDN:   "DC=testdomain,DC=internal",
		Security: auth.SecurityNone,
	}

	username := "tdelnoij"
	password := "Xs4u2019"

	err := os.Setenv("", "")

	groups := []string{"Lyla-Users", "Another-Group"}
	status, entry, userGroups, err := auth.AuthenticateExtended(config, username, password, []string{}, groups)

	if err != nil {
		log.Fatal(err)
	}

	if !status {
		log.Fatal("not authenticated")
	}

	entry.Print()
	for _, attr := range entry.Attributes {
		log.Printf("Name: %v, Values: %v", attr.Name, attr.Values)
	}
	for _, gr := range userGroups {
		log.Println(gr)
	}
}
