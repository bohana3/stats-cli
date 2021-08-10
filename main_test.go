package main

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/bohana3/stats-cli/stats"
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"testing"
)

func TestParseMetadata_ValidJson(t *testing.T) {
	correctFile :=`{"path": "C:\\Users\\benjamin\\Downloads\\010EditorWin64Installer801.exe","size": 16594864,"is_binary": true}`
	file,err := ParseMetadata(correctFile)

	assert.Equal(t,nil,err)
	assert.Equal(t,"C:\\Users\\benjamin\\Downloads\\010EditorWin64Installer801.exe",file.Path)
	assert.Equal(t,int64(16594864),file.Size)
	assert.Equal(t,true,file.IsBinary)
}

func TestParseMetadata_NoPath(t *testing.T) {
	correctFile :="{\"size\": 16594864,\"is_binary\": true}"
	_,err := ParseMetadata(correctFile)

	if assert.Error(t, err) {
		assert.Equal(t, errors.New("Path is empty"), err)
	}
}

func TestParseMetadata_PathEmpty(t *testing.T) {
	correctFile :="{\"path\": \"\": 16594864,\"is_binary\": true}"
	_,err := ParseMetadata(correctFile)

	assert.Error(t, err)
}

func TestParseMetadata_InvalidJson(t *testing.T) {
	correctFile :="\"path\": \"\",\"size\": 16594864,\"is_binary\": true}"
	_,err := ParseMetadata(correctFile)
	assert.Error(t, err)
}

func TestParseManyFilesAndGetStats(t *testing.T) {
	var sv = stats.NewStatsVisualizator()

	file, err := os.Open("testresources\\standardinput.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if metadata, error := ParseMetadata(scanner.Text()); error != nil {
			assert.Error(t, fmt.Errorf("unable to deserialization the record [%ws]  with error [%s]",scanner.Text(),error.Error()))
		} else {
			sv.AddFile(metadata)
		}
	}

	stats := sv.GetStats()

	//same assert than stats_test.TestAddManyFiles()
	assert.Equal(t,int64(467), stats.NumFiles)
	assert.Equal(t,"C:\\Users\\benjamin\\Downloads\\ideaIC-2020.3.exe", stats.LargestFile.Path)
	assert.Equal(t,int64(645770200), stats.LargestFile.Size)

	assert.Equal(t,1.290615485653105e+07, stats.AverageFileSize)
	assert.Equal(t,".zip", stats.MostFrequentExt.Extension)
	assert.Equal(t,int64(94), stats.MostFrequentExt.NumOccurrences)
	assert.Equal(t,float32(82.22698), stats.TextPercentage)

	assert.Equal(t,"C:\\Users\\benjamin\\Downloads\\~$CV_v1.docx", stats.MostRecentPaths[0])
	assert.Equal(t,"C:\\Users\\benjamin\\Downloads\\בינימין אוחנה 39742.pdf", stats.MostRecentPaths[1])
	assert.Equal(t,"C:\\Users\\benjamin\\Downloads\\תרגיל גרפים 1.xlsx", stats.MostRecentPaths[9])
	assert.Equal(t,10, len(stats.MostRecentPaths))
}


