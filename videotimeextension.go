package rtp

import (
	"encoding/binary"
	"fmt"
)

const (
	videoTimeExtensionSize = 13
)

// VideoTimeExtension is a extension payload format in
// http://www.webrtc.org/experiments/rtp-hdrext/video-timing
type VideoTimeExtension struct {
	encodeStartDeltaMs         uint16
	encodeFinishDeltaMs        uint16
	packetizationFinishDeltaMs uint16
	pacerExitDeltaMs           uint16
	networkTimestampDeltaMs    uint16
	network2TimestampDeltaMs   uint16
	flags                      uint8
}

func (v *VideoTimeExtension) Marshal() ([]byte, error) {
	buf := make([]byte, videoTimeExtensionSize)
	binary.BigEndian.PutUint16(buf[0:], v.encodeStartDeltaMs)
	binary.BigEndian.PutUint16(buf[2:], v.encodeFinishDeltaMs)
	binary.BigEndian.PutUint16(buf[4:], v.packetizationFinishDeltaMs)
	binary.BigEndian.PutUint16(buf[6:], v.pacerExitDeltaMs)
	binary.BigEndian.PutUint16(buf[8:], v.networkTimestampDeltaMs)
	binary.BigEndian.PutUint16(buf[10:], v.network2TimestampDeltaMs)
	buf[12] = v.flags
	return buf, nil
}

func (v *VideoTimeExtension) Unmarshal(rawData []byte) error {
	if len(rawData) < videoTimeExtensionSize {
		return errTooSmall
	}
	v.encodeStartDeltaMs = binary.BigEndian.Uint16(rawData[0:])
	v.encodeFinishDeltaMs = binary.BigEndian.Uint16(rawData[2:])
	v.packetizationFinishDeltaMs = binary.BigEndian.Uint16(rawData[4:])
	v.pacerExitDeltaMs = binary.BigEndian.Uint16(rawData[6:])
	v.networkTimestampDeltaMs = binary.BigEndian.Uint16(rawData[8:])
	v.network2TimestampDeltaMs = binary.BigEndian.Uint16(rawData[10:])
	v.flags = rawData[12]

	return nil
}

func (v VideoTimeExtension) String() string {
	var flags string
	switch v.flags {
	case 0:
		flags = "NotTriggered"
	case 1 << 0:
		flags = "TriggeredByTimer"
	case 1 << 1:
		flags = "TriggeredBySize"
	default:
		flags = "UnknownTriggered"
	}

	out := "VideoTimeExtension HEADER:\n"

	out += fmt.Sprintf("\tencodeStartDeltaMs: %d\n", v.encodeStartDeltaMs)
	out += fmt.Sprintf("\tencodeFinishDeltaMs: %d\n", v.encodeFinishDeltaMs)
	out += fmt.Sprintf("\tpacketizationFinishDeltaMs: %d\n", v.packetizationFinishDeltaMs)
	out += fmt.Sprintf("\tpacerExitDeltaMs: %d\n", v.pacerExitDeltaMs)
	out += fmt.Sprintf("\tnetworkTimestampDeltaMs: %d\n", v.networkTimestampDeltaMs)
	out += fmt.Sprintf("\tnetwork2TimestampDeltaMs: %d\n", v.network2TimestampDeltaMs)
	out += fmt.Sprintf("\tflags: %s\n", flags)

	return out
}
