package util

import (
	"encoding/json"
	"log"
)

var Json JSON

type JSON struct {
}

func (*JSON) Marshal(v any) string {
	data, err := json.Marshal(v)
	if err != nil {
		log.Println(err)
	}
	return string(data)
}

func (*JSON) Unmarshal(data string, v any) {
	err := json.Unmarshal([]byte(data), &v)
	if err != nil {
		log.Println(err)
	}
}

func (j *JSON) DecodeMap(in any, out any) {
	input := j.Marshal(in)
	j.Unmarshal(input, &out)
}
