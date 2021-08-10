package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"github.com/bohana3/stats-cli/stats"
	"log"
	"os"
)

func main() {
	log.Println("Start stats-cli")
	sv := stats.NewStatsVisualizator()

	//scan console standard input
	scn := bufio.NewScanner(os.Stdin)
	for {
		log.Println("Enter Lines: (to stop the capture, press twice Enter!)")
		for scn.Scan() {
			line := scn.Text()
			if line == "" { //stop the scan if a empty line is detected
				break
			}
			if metadata, error := ParseMetadata(line); error != nil {
				log.Fatalf("skip the record [%s] since the deserialization failed with error [%s]", line, error.Error())
			} else {
				sv.AddFile(metadata)
			}
		}

		if err := scn.Err(); err != nil {
			log.Fatalf("error during scanning standard input [%ws]", err.Error())
		}

		break
	}

	if res, err := json.MarshalIndent(sv.GetStats(), "", "  ");err != nil {
		log.Println(err)
		log.Println(sv.GetStats())
	} else {
		log.Println(string(res))
	}

	log.Println("Finish stats-cli")
}

func ParseMetadata(line string) (stats.FileMetadata, error) {
	var file stats.FileMetadata
	if error := json.Unmarshal([]byte(line), &file);error != nil {
		return file, error
	}

	if len(file.Path) == 0 {
		return file, errors.New("Path is empty")
	}

	return file,nil
}