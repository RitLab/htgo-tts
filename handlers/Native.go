package handlers

import (
	"bytes"
	"os"
	"time"

	"github.com/ebitengine/oto/v3"
	"github.com/hajimehoshi/go-mp3"
)

type Native struct {
}

func (n *Native) Play(fileName string) error {
	// Read the mp3 file into memory
	fileBytes, err := os.ReadFile(fileName)
	if err != nil {
		return err
	}

	fileBytesReader := bytes.NewReader(fileBytes)

	// Decode file
	decodedMp3, err := mp3.NewDecoder(fileBytesReader)
	if err != nil {
		return err
	}

	numOfChannels := 2

	otoCtx, readyChan, err := oto.NewContext(&oto.NewContextOptions{
		SampleRate:   decodedMp3.SampleRate(),
		ChannelCount: numOfChannels,
		Format:       oto.FormatSignedInt16LE,
	})
	if err != nil {
		return err
	}
	<-readyChan

	player := otoCtx.NewPlayer(decodedMp3)

	player.Play()

	for player.IsPlaying() {
		time.Sleep(time.Millisecond)
	}

	return player.Close()
}
