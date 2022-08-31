package helpers

import (
	"fmt"
	"strconv"
	"strings"
)

func Validate(fields map[string]interface{}, rules map[string][]string) (bool, map[string]string, map[string]interface{}) {
	err := make(map[string]string)
	clean_fields := make(map[string]interface{})
	for field, rule := range rules {
		value := fields[field]
		for _, r := range rule {
			if r == "required" && value == nil {
				err[field] = r
				return false, err, nil
			}
			if value == nil {
				continue
			}
			if r == "string" {
				value = strings.Trim(fmt.Sprintf("%v", value.(string)), " ")
			}
			if r == "boolean" {
				if value != "true" && value != "false" {
					err[field] = fmt.Sprintf("Value must be boolean [true, false]")
					return false, err, nil
				}
			}
			if r == "number" {
				if _, errCheckNumber := strconv.Atoi(value.(string)); errCheckNumber != nil {
					err[field] = "Must be a number"
					return false, err, nil
				}
			}
		}
		clean_fields[field] = value
	}
	return true, nil, clean_fields
}
