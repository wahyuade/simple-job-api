package helpers

import "net/url"

func GetMappedQuery(query string) map[string]interface{} {
	map_query := make(map[string]interface{})
	params, err := url.ParseQuery(query)
	if err == nil {
		for key, value := range params {
			if key == "" {
				continue
			}
			if len(value) == 0 {
				continue
			}
			map_query[key] = value[0]
		}
	}
	return map_query
}
