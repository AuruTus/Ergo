package tools

import (
	"errors"
	"testing"
)

func TestJsonFieldsGet(t *testing.T) {
	type tArg struct {
		path     string
		j        string
		expected any
	}

	args := []tArg{
		{
			path: "name",
			j: `
			{
				"name": "text",
				"data": {
					"text": " plz "
				}
			}`,
			expected: "text",
		}, {
			path: "vis.h",
			j: `
			{
				"data": {
					"text": "plz"
				},
				"name": "text",
				"vis": {
					"h": "123"
				}
			}`,
			expected: "123",
		}, {
			path: "vis.H",
			j: `
			{
				"data": {
					"text": "plz"
				},
				"name": "text",
				"vis": {
					"H": 4589
				}
			}`,
			expected: 4589.0,
		}, {
			path: "vis.1.3",
			j: `
			{
				"data": {
					"text": "plz"
				},
				"name": "text",
				"vis": {
					"1": "MTM0YQ=="
				}
			}`,
			expected: uint8('0'),
		}, {
			path: "more.0.happy.高兴",
			j: `
			{
				"name": "text",
				"data": {
					"text": " plz "
				},
				"1": 2,
				"more": [
					{
						"happy": {
							"高兴": 3.3213
						}
					}
				]
			}`,
			expected: 3.3213,
		},
	}

	for _, a := range args {
		val, err := GetJsonField([]byte(a.j), a.path)
		if err != nil {
			if pathErr := (&PathError{}); errors.As(err, &pathErr) {
				t.Errorf("patherr: %v\n\n", pathErr)
			} else if panicErr := (&WrappedPanic{}); errors.As(err, &panicErr) {
				t.Errorf("panicErr: %v\n", panicErr)
			} else {
				t.Errorf("%v\n", err)
			}
		}
		t.Logf("%+v %+v\n", val.Kind(), val)
		if val.Interface() != a.expected {
			t.Errorf("get wrong value %v for expected %v\n", val.Interface(), a.expected)
		}
	}
}
