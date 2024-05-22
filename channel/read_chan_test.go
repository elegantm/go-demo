package channel

import "testing"

func TestReadCloseChannelNodata(t *testing.T) {
	ReadCloseChannelWithData()

	t.Log("-------------")

	defer func() {
		m := recover()
		t.Log(m)
	}()
	ReadCloseChannel()

}

func TestChanAndArray(t *testing.T) {
	ChanAndArray()
}
