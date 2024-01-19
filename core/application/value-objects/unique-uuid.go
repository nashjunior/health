package valueobjects

import (
	"encoding/json"
	"log"
	"reflect"

	"github.com/google/uuid"
)

type UniqueEntityUUID struct {
	Id uuid.UUID `json:"id"`
}

func (v *UniqueEntityUUID) Equals(obj *UniqueEntityUUID) bool {
	if obj == nil {
		return false
	}

	if obj.Id.String() != v.Id.String() {
		return false
	}

	if reflect.TypeOf(obj) != reflect.TypeOf(v) {
		return false
	}

	objType := reflect.TypeOf(obj.Id)
	thisType := reflect.TypeOf(v.Id)

	if objType.Kind() == reflect.Map && thisType.Kind() == reflect.Map {
		keysThis := reflect.ValueOf(v.Id).MapKeys()
		keysObj := reflect.ValueOf(obj.Id).MapKeys()

		if len(keysThis) != len(keysObj) {
			return false
		}

		for _, key := range keysThis {
			found := false
			for _, objKey := range keysObj {
				if reflect.DeepEqual(key.Interface(), objKey.Interface()) {
					found = true
					break
				}
			}
			if !found {
				return false
			}
		}
	}

	return true
}

func (v *UniqueEntityUUID) ToString() string {

	jsonBytes, err := json.Marshal(v)
	if err != nil {
		log.Fatalf("Error marshaling struct: %v", err)
	}

	// Convert the JSON bytes to a string
	return string(jsonBytes)

}

func NewUniqueUUID(prop uuid.NullUUID) UniqueEntityUUID {
	var instance uuid.UUID

	if prop.Valid {
		instance = uuid.New()

	} else {
		instance = prop.UUID
	}

	return UniqueEntityUUID{Id: instance}
}
