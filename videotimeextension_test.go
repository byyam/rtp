package rtp

import (
	"bytes"
	"testing"
)

func TestVideoTimeExtension(t *testing.T) {
	v1 := VideoTimeExtension{}
	rawData := []byte{
		0x01, 0x00, 0x14, 0x00, 0x18, 0x00, 0x18, 0x00, 0x2d, 0x00, 0x00, 0x00, 0x00,
	}

	if err := v1.Unmarshal(rawData); err != nil {
		t.Fatal("Unmarshal error on video time extension", err)
	}
	t.Logf("Unmarshal data:%s", v1.String())
	v2 := VideoTimeExtension{
		encodeStartDeltaMs:         256,
		encodeFinishDeltaMs:        5120,
		packetizationFinishDeltaMs: 6144,
		pacerExitDeltaMs:           6144,
		networkTimestampDeltaMs:    11520,
		network2TimestampDeltaMs:   0,
		flags:                      0,
	}
	dstData, err := v2.Marshal()
	if err != nil {
		t.Fatal("Marshal error on video time extension", err)
	}
	if !bytes.Equal(dstData, rawData) {
		t.Error("Marshal failed")
	}
	t.Logf("Marshal data:0x%0x", dstData)
}
