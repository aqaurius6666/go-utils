package utils

import (
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

func ParseUUID(id interface{}) (uuid.UUID, error) {
	switch t := id.(type) {
	case string:
		return uuid.Parse(t)
	case uuid.UUID:
		return t, nil
	case [16]byte:
		return uuid.UUID(t), nil
	case []byte:
		if len(t) != 16 {
			return uuid.Nil, errors.New("invalid uuid")
		}
		tmp := [16]byte{}
		for i, b := range t {
			tmp[i] = b
		}
		return uuid.UUID(tmp), nil
	default:
		return uuid.Nil, errors.New("invalid uuid")
	}
}
