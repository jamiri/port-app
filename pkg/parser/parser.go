package parser

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)


func ReadFile(path string) (io.Reader, error) {
	f, _ := os.Open(path)
	defer func() {
		err := f.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	jsonBeginMarker := []byte("{")
	jsonEndMarker := []byte("}")

	scanner := bufio.NewScanner(f)
	firstLine := true
	var buffer bytes.Buffer
	data := make(map[string]interface{})
	// Skip first line as it is { and continue
	for scanner.Scan(); scanner.Scan(); {
		line := scanner.Bytes()
		var portCode []byte

		if firstLine {
			portCode = line[3:8]
			line = jsonBeginMarker
			firstLine = false
		}

		if string(line) == "  }," {
			line = jsonEndMarker
			buffer.Write(line)
			err := json.Unmarshal(buffer.Bytes(), &data)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(string(portCode))
			fmt.Println(data)

			fmt.Println("\n\n")

			buffer.Reset()
			firstLine = true
			continue
		}
		buffer.Write(line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return nil, nil
}
