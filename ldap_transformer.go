package main

import (
	ldap "github.com/zonradkuse/go-ldap-authenticator"

	"log"
	"regexp"
	"strconv"
	"strings"
)

type LDAPTransformer struct{}

func (this LDAPTransformer) Transform(entry *ldap.Entry) interface{} {
	user := NewUserData()
	log.Printf("%+v", entry)
	for _, attr := range entry.Attributes {
		log.Printf("%+v", *attr)
		if attr.Name == "mail" {
			user.Email = attr.Values[0]
		}
		if attr.Name == "createTimestamp" {
			re := regexp.MustCompile("[0-9]+")
			numbers := re.FindAllString(attr.Values[0], -1)
			id, err := strconv.ParseInt(strings.Join(numbers, ""), 10, 64)
			user.Id = id

			if err != nil {
				panic(err)
			}
		}
		if attr.Name == "cn" {
			user.Name = attr.Values[0]
		}
		if attr.Name == "uid" {
			user.Username = "sog_" + attr.Values[0]
		}
	}

	return user
}
