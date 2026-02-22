package converter

import (
	"errors"
	"fmt"
	"io"
	"os"
	"time"

	"dav2mp4/pkg/dhplay"
)

type VideoFormat int

const (
	RAW VideoFormat = dhplay.DATA_RECORD_ORIGINAL
	AVI VideoFormat = dhplay.DATA_RECORD_AVI
	ASF VideoFormat = dhplay.DATA_RECORD_ASF
	MP4 VideoFormat = dhplay.DATA_RECORD_MP4
)

const (
	PORT           = 0
	SOURCEBUF_SIZE = 500 * 1024
	READ_LEN       = 8 * 1024
)

func initStream(outputPath string, outputFormat VideoFormat) (int, error) {
	port := PORT
	if !dhplay.SetStreamOpenMode(port, dhplay.STREAME_FILE) {
		return 0, errors.New("unable to set stream mode")
	}
	if !dhplay.OpenStream(port, SOURCEBUF_SIZE) {
		return 0, errors.New("unable to open stream")
	}
	if !dhplay.Play(port) {
		return 0, errors.New("unable to play on stream")
	}
	if !dhplay.StartDataRecordEx(port, outputPath, int(outputFormat)) {
		return 0, errors.New("unable to start data record")
	}
	return port, nil
}

func stopAndCloseStream(port int) {
	dhplay.StopDataRecord(port)
	dhplay.Stop(port)
	dhplay.CloseStream(port)
}

func Convert(inputFilePath, outputFilePath string, outputFormat VideoFormat) error {
	if _, err := os.Stat(inputFilePath); os.IsNotExist(err) {
		return fmt.Errorf("the input file '%s' does not exists", inputFilePath)
	}
	if _, err := os.Stat(outputFilePath); err == nil {
		return fmt.Errorf("the output file '%s' already exists", outputFilePath)
	}

	port, err := initStream(outputFilePath, outputFormat)
	if err != nil {
		return err
	}

	inputFile, err := os.Open(inputFilePath)
	if err != nil {
		stopAndCloseStream(port)
		return err
	}
	defer inputFile.Close()

	buf := make([]byte, READ_LEN)
	for {
		n, err := inputFile.Read(buf)
		if n > 0 {
			for !dhplay.InputData(port, buf[:n]) {
				time.Sleep(10 * time.Millisecond)
			}
		}
		if err != nil {
			if err == io.EOF {
				break
			}
			stopAndCloseStream(port)
			return err
		}
	}
	stopAndCloseStream(port)

	return nil
}
