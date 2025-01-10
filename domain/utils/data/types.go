package data

import "github.com/google/uuid"

type Obj map[string]any

func ConvToUuid(str string) uuid.UUID {
	var uuid [16]byte
	copy(uuid[:], str)
	return uuid
}
