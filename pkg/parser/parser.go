package parser

import (
	"bufio"
	"bytes"
	"encoding/json"
	"log"
	"os"
)

var (
	jsonBeginMarker = []byte("{")
	jsonEndMarker   = []byte("}")
	jsonEndLine     = []byte("  },")
	lastJsonEndLine = []byte("  }")
)

func ReadPorts(path string) <-chan Port {
	f, _ := os.Open(path)
	portsChan := make(chan Port)

	go func() {
		scanner := bufio.NewScanner(f)
		firstLine := true
		var buffer bytes.Buffer
		data := Port{}

		// Skip first line as it is { and continue
		for scanner.Scan(); scanner.Scan(); {
			line := scanner.Bytes()

			if firstLine {
				line = jsonBeginMarker
				firstLine = false
			}

			if bytes.Equal(line, jsonEndLine) || bytes.Equal(line, lastJsonEndLine) {
				line = jsonEndMarker
				buffer.Write(line)
				err := json.Unmarshal(buffer.Bytes(), &data)
				portsChan <- data
				if err != nil {
					log.Fatal(err)
				}
				buffer.Reset()
				firstLine = true
				continue
			}
			buffer.Write(line)
		}

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}

		err := f.Close()
		if err != nil {
			log.Fatal(err)
		}
		close(portsChan)
	}()

	return portsChan
}
