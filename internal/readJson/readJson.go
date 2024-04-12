package readjson

import (
	"encoding/json"
	"fmt"
)

type Shape struct {
	Records []map[string]interface{} `json:"records"`
}

func Process(content []byte) ([][]string, error) {
	d := Shape{}

	err := json.Unmarshal(content, &d)
	if err != nil {
		return nil, err
	}
	atributes := map[string]struct{}{}
	for _, rec := range d.Records {

		for atr := range rec {
			atributes[atr] = struct{}{}
		}
	}
	atr := []string{}
	for a := range atributes {
		atr = append(atr, a)
	}

	data := [][]string{}
	data = append(data, atr)
	for _, rec := range d.Records {
		line := make([]string, len(atr))
		for i, v := range atr {
			if rec[v] == nil {
				line[i] = "null"
				continue
			}
			tem := fmt.Sprintf("%v", rec[v])

			line[i] = tem
		}
		data = append(data, line)

	}
	return data, nil
}
