package solid

import (
	"encoding/json"
	"testing"
)

func TestMarshal(t *testing.T) {
	data := Marshal()
	raw, err := json.Marshal(data)
	t.Log(err)
	t.Log(string(raw))

	var tmp Notify
	err = json.Unmarshal(raw, &tmp)
	t.Log(err)
	t.Logf("%+v", tmp)
}
