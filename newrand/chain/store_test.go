package chain

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRoundToBytes(t *testing.T) {
	require.Equal(t, []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, RoundToBytes(0))
	require.Equal(t, []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01}, RoundToBytes(1))
	require.Equal(t, []byte{0x00, 0x00, 0x00, 0x2a, 0xec, 0x04, 0x83, 0xff}, RoundToBytes(184348345343))
	require.Equal(t, []byte{0xA1, 0xB2, 0xC3, 0xD4, 0xE5, 0xF6, 0xA7, 0xB8}, RoundToBytes(0xA1B2C3D4E5F6A7B8))
}
