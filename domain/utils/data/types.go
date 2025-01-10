package data

type Obj map[string]any

func StrToUuid(str string) [12]byte {
	var uuid [12]byte
	copy(uuid[:], str)
	return uuid
}
