package utils

import (
	"bytes"
	"encoding/json"
	"log"
)

func PrettyJSON(headerJson []byte) {
	log.Printf("Headers: %s", headerJson)
	var prettyJSON bytes.Buffer
	_ = json.Indent(&prettyJSON, headerJson, "", " ")
	log.Printf("Headers: \n%s", prettyJSON.String())
}
