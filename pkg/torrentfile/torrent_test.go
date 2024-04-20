package torrent

import (
	"encoding/json"
	"flag"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var update = flag.Bool("update", false, "update .golden.json files")

func TestOpen(t *testing.T) {
	torrent, err := Open("testdata/debian-12.5.0-amd64-DVD-1.iso.torrent")
	require.Nil(t, err)

	input := "testdata/debian-12.5.0-amd64-DVD-1.iso.torrent.golden.json"
	if *update {
		serialized, err := json.MarshalIndent(torrent, "", "  ")
		require.Nil(t, err)
		os.WriteFile(input, serialized, 0644)
	}

	expected := TorrentFile{}
	golden, err := os.ReadFile(input)
	require.Nil(t, err)
	err = json.Unmarshal(golden, &expected)
	require.Nil(t, err)

	assert.Equal(t, expected, torrent)
}
