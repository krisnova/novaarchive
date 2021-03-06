package photoprism

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/kris-nova/photoprism-client-go/api/v1"
)

type CustomData struct {
	// Deprecated
	LastTweet *time.Time

	Updated     time.Time
	NoteStrings []string
	KeyValue    map[string]string
	Description string
}

// SetCustomData will always set the updated time!
func SetCustomData(d *CustomData, photo *api.Photo) error {
	d.Updated = time.Now()
	jBytes, err := json.Marshal(&d)
	if err != nil {
		return fmt.Errorf("unable to set custom photo data: %v", err)
	}
	photo.PhotoDescription = string(jBytes)
	return nil
}

func GetCustomData(photo api.Photo) *CustomData {
	d := &CustomData{}
	if photo.PhotoDescription == "" {
		return nil
	}
	noteStr := photo.PhotoDescription
	err := json.Unmarshal([]byte(noteStr), &d)
	if err != nil {
		// Libraries houldnt log but wtf
		//logger.Warning("INVALID JSON in Notes: %v", err)
		return nil
	}
	return d
}
