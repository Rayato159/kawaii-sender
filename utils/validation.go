package utils

import (
	"fmt"
	"regexp"
)

func BodyValidate(method string, body any) error {
	msg := fmt.Sprintf("body is invalid for method: %v", method)
	switch method {
	case "GET":
		if body != nil {
			return fmt.Errorf(msg)
		}
	case "DELETE":
		if body != nil {
			return fmt.Errorf(msg)
		}
	}
	return nil
}

func UrlValidate(url string) error {
	match, err := regexp.MatchString(
		"https?:\\/\\/(www\\.)?[-a-zA-Z0-9@:%._\\+~#=]{1,256}\\.[a-zA-Z0-9()]{1,6}\b([-a-zA-Z0-9()@:%_\\+.~#?&//=]*)",
		url,
	)
	if err != nil {
		return fmt.Errorf("regexp error: %v", err)
	}
	if !match {
		return fmt.Errorf("http url pattern error: %s", url)
	}
	return nil
}
