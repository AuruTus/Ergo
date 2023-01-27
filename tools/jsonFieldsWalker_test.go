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
		},
	}

	for _, a := range args {
		val, err := GetJsonField([]byte(a.j), a.path)
		if err != nil {
			if pathErr := (&PathError{}); errors.As(err, &pathErr) {
				t.Logf("patherr: %v\n\n", pathErr)
			} else if panicErr := (&WrappedPanic{}); errors.As(err, &panicErr) {
				t.Logf("panicErr: %v\n", panicErr)
			} else {
				t.Logf("%v\n", err)
			}
		}
		t.Logf("%+v %+v\n", val.Kind(), val)
	}
}
