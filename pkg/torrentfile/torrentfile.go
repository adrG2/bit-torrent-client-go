package torrent

import (
	"bytes"
	"crypto/sha1"
	"fmt"
	"os"

	"github.com/jackpal/bencode-go"
)

// TorrentFile encodes the metadata from a .torrent file
type TorrentFile struct {
	Announce    string
	BenInfoHash [20]byte
	PieceHashes [][20]byte
	PieceLength int
	Length      int
	Name        string
}

type bencodeInfo struct {
	Pieces      string `bencode:"pieces"`
	PieceLength int    `bencode:"piece length"`
	Length      int    `bencode:"length"`
	Name        string `bencode:"name"`
}

type bencodeTorrent struct {
	Announce string      `bencode:"announce"`
	Info     bencodeInfo `bencode:"info"`
}

// Parse a torrent file
func Open(path string) (TorrentFile, error) {
	file, err := os.Open(path)
	if err != nil {
		return TorrentFile{}, err
	}
	defer file.Close()

	bto := bencodeTorrent{}
	err = bencode.Unmarshal(file, &bto)
	if err != nil {
		return TorrentFile{}, err
	}
	return bto.toTorrentFile()
}

func (bto *bencodeTorrent) toTorrentFile() (TorrentFile, error) {
	infoHash, err := bto.Info.hash()
	if err != nil {
		return TorrentFile{}, err
	}
	pieceHashes, err := bto.Info.splitPieceHashes()
	if err != nil {
		return TorrentFile{}, err
	}
	t := TorrentFile{
		Announce:    bto.Announce,
		BenInfoHash: infoHash,
		PieceHashes: pieceHashes,
		PieceLength: bto.Info.PieceLength,
		Length:      bto.Info.Length,
		Name:        bto.Info.Name,
	}
	return t, nil
}

// hashSize specifies the length of each torrent piece SHA-1 hash in bytes.
const hashSize = 20

// Compute SHA-1 hash for bencodeInfo(name, size and piece hashes). The SHA-1 hash is 20 bytes long
func (i *bencodeInfo) hash() ([hashSize]byte, error) {
	var buf bytes.Buffer
	err := bencode.Marshal(&buf, *i)
	if err != nil {
		return [hashSize]byte{}, err
	}
	h := sha1.Sum(buf.Bytes())
	return h, nil
}

// splitPieceHashes Splits the hashes of the parts into a [20]byte slice.
// It takes 20-byte ranges from the pieces buffer, this represents each hash of each piece, and copies it into an array of hashes.
func (i *bencodeInfo) splitPieceHashes() ([][hashSize]byte, error) {
	buf := []byte(i.Pieces)
	if isValidPieces(buf) {
		err := fmt.Errorf("Received malformed pieces of length %d", len(buf))
		return nil, err
	}

	numHashes := calculateNumHashes(buf)

	// Initialize a slice to store the piece hashes.
	// Each element of the slice is an array of 20 bytes representing a hash.
	hashes := make([][hashSize]byte, numHashes)

	// Copy piece hashes, 20 bytes at a time, from the buffer.
	// These segments are of length 20 bytes, which is the length of each piece hash.
	// It increments by 20 by multiplying the loop index by the hash length in bytes.
	// example: buf[0 -> byte,..,20 -> byte] -> hashes[0 -> [0 -> byte,..,20 -> byte]]
	for i := 0; i < numHashes; i++ {
		dst := hashes[i][:]
		src := buf[i*hashSize : (i+1)*hashSize]
		copy(dst, src)
	}
	return hashes, nil
}

// isValidPieces validates the pieces by ensuring each piece has a length of 20 bytes.
// If the total number of bytes is not a multiple of 20, it indicates malformed pieces.
func isValidPieces(buf []byte) bool {
	return len(buf)%hashSize != 0
}

// calculateNumHashes Compares length of the pieces buffer with the size of each hash, which is 20 bytes.
func calculateNumHashes(buf []byte) int {
	return len(buf) / hashSize
}
