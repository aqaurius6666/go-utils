package utils

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Src struct {
	Offset  string
	Limit   string
	Another string
}
type Dest struct {
	Offset  int
	Limit   int
	Another string
}

func TestMarshalOffsetLimit(t *testing.T) {

	testcase := []map[string]interface{}{
		{
			"s":  NewSrc("1", "321"),
			"d":  NewDest(1, 321),
			"ok": true,
			"e":  NewDest(1, 321),
		},
		{
			"s":  NewSrc("1a", "321"),
			"d":  NewDest(0, 2),
			"ok": false,
			"e":  NewDest(1, 321),
		},
		{
			"s":  NewSrc("1", "321a"),
			"d":  NewDest(1, 321),
			"ok": false,
			"e":  NewDest(1, 321),
		},
		{
			"s":  NewSrc("-1", "321"),
			"d":  NewDest(1, 321),
			"ok": false,
			"e":  NewDest(1, 321),
		},
		{
			"s":  NewSrc("1", "321.3"),
			"d":  NewDest(1, 321),
			"ok": false,
			"e":  NewDest(1, 321),
		},
		{
			"s":  NewSrc("1", "-321"),
			"d":  NewDest(1, 321),
			"ok": false,
			"e":  NewDest(1, 321),
		},
	}
	for _, tcase := range testcase {
		err := MarshalOffsetLimit(tcase["d"], tcase["s"])
		if tcase["ok"].(bool) {
			assert.Nil(t, err)
			assert.Equal(t, tcase["e"].(*Dest).Limit, tcase["d"].(*Dest).Limit)
			assert.Equal(t, tcase["e"].(*Dest).Offset, tcase["d"].(*Dest).Offset)

		} else {
			fmt.Printf("err: %v\n", err)
			assert.NotNil(t, err)
		}

	}
}

func NewSrc(o, l string) interface{} {
	if rand.Int()%2 == 0 {
		return &Src{
			Offset:  o,
			Limit:   l,
			Another: "132",
		}
	}
	return Src{
		Offset:  o,
		Limit:   l,
		Another: "132",
	}
}

func NewDest(o, l int) interface{} {
	return &Dest{
		Offset:  o,
		Limit:   l,
		Another: "132",
	}
}
