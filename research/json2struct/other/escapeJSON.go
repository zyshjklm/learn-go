package main

import (
	"encoding/json"
	"fmt"
)

// LogPayload for payload
type LogPayload struct {
	Message string `json:"message"`
}

type fakeLogPayload LogPayload

// LogEntry for json
type LogEntry struct {
	ID      int        `json:"id"`
	Name    string     `json:"name"`
	Payload LogPayload `json:"payload"`
}

// UnmarshalJSON for UnmarshalJSON
func (lp *LogPayload) UnmarshalJSON(b []byte) (err error) {
	var s string
	if err = json.Unmarshal(b, &s); err != nil {
		return err
	}
	// fmt.Printf("str:%s\n", s)

	var fake fakeLogPayload
	if err = json.Unmarshal([]byte(s), &fake); err != nil {
		return err
	}
	*lp = LogPayload(fake)

	return nil
}

func main() {
	doc := []byte(`{
       "id": 12345,
       "name": "Test Document",
       "payload": "{\"message\":\"test\"}"
	}`)
	//"payload": "{\"message\":\"test\"}"
	var entry LogEntry
	if err := json.Unmarshal(doc, &entry); err != nil {
		fmt.Println("Error!", err)
	}

	fmt.Printf("%#v\n", entry)
	fmt.Printf("%#v\n", entry.Payload)
}
