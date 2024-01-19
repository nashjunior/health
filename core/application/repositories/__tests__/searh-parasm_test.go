package tests

import (
	"health/core/application/repositories"
	"math"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldHandleParamPage(t *testing.T) {
	arrange := []interface{}{-1, 1.1, '1', true}
	var params repositories.SearchParams[any]

	for _, test := range arrange {
		var testInt int

		switch v := test.(type) {
		case int:
			testInt = test.(int)

		case float64:
			testInt = int(v)
		case bool:
			testInt = 0
		case string:
			_, err := strconv.Atoi(v)
			if err != nil {
				testInt = int(math.NaN())
			}
		}
		params = repositories.NewSearchParams[any](&testInt, nil, nil, nil)

		assert.Equal(t, 1, params.GetPage(), "Should dbe 1 when invalid page")

	}

	val := 10
	params = repositories.NewSearchParams[any](&val, nil, nil, nil)

	assert.Equal(t, 10, params.GetPage(), "Should be the assigned value")
}

func TestShouldHandleParamPerPage(t *testing.T) {
	arrange := []interface{}{-1, 'a', true}
	var params repositories.SearchParams[any]

	for _, test := range arrange {
		var testInt int

		switch v := test.(type) {
		case int:
			testInt = test.(int)

		case float64:
			testInt = int(v)
		case bool:
			testInt = 0
		case string:
			_, err := strconv.Atoi(v)
			if err != nil {
				testInt = int(math.NaN())
			}
		}
		params = repositories.NewSearchParams[any](nil, &testInt, nil, nil)

		assert.Equal(t, 10, params.GetPerPage(), "Should dbe 1 when invalid page")
	}

	val := 10
	params = repositories.NewSearchParams[any](&val, nil, nil, nil)

	assert.Equal(t, 10, params.GetPage(), "Should be the assigned value")
}
