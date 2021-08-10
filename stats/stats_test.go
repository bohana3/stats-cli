package stats

import (
	"encoding/json"
	"errors"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

func TestAddThreeFilesAndGetStats(t *testing.T) {
	sv := NewStatsVisualizator()

	file1:=FileMetadata{Path:"C:\\Users\\benjamin\\Downloads\\010EditorWin64Installer801.exe",Size:16594864,IsBinary:true}
	file2:=FileMetadata{Path:"C:\\Users\\benjamin\\Downloads\\1513735129.pdf",Size:1552691,IsBinary:false}
	file3:=FileMetadata{Path:"C:\\Users\\benjamin\\Downloads\\1513739300.pdf",Size:19484,IsBinary:false}

	sv.AddFile(file1)
	sv.AddFile(file2)
	sv.AddFile(file3)

	stats := sv.GetStats()

	assert.Equal(t,int64(3), stats.NumFiles)
	assert.Equal(t,"C:\\Users\\benjamin\\Downloads\\010EditorWin64Installer801.exe", stats.LargestFile.Path)
	assert.Equal(t,int64(16594864), stats.LargestFile.Size)

	assert.Equal(t,(16594864.00+1552691.00+19484.00)/3.00, stats.AverageFileSize)
	assert.Equal(t,".pdf", stats.MostFrequentExt.Extension)
	assert.Equal(t,int64(2), stats.MostFrequentExt.NumOccurrences)
	assert.Equal(t,float32(2)/float32(3)*float32(100), stats.TextPercentage)

	assert.Equal(t,"C:\\Users\\benjamin\\Downloads\\010EditorWin64Installer801.exe", stats.MostRecentPaths[0])
	assert.Equal(t,"C:\\Users\\benjamin\\Downloads\\1513735129.pdf", stats.MostRecentPaths[1])
	assert.Equal(t,"C:\\Users\\benjamin\\Downloads\\1513739300.pdf", stats.MostRecentPaths[2])

	assert.Equal(t,3, len(stats.MostRecentPaths))
}

func TestAddManyFilesAndGetStats(t *testing.T) {
	var sv = NewStatsVisualizator()

	pathJson, _ := ioutil.ReadFile("testresources\\paths.json")
	var files []FileMetadata
	json.Unmarshal([]byte(pathJson), &files)

	for _, file := range files {
		sv.AddFile(file)
	}

	stats := sv.GetStats()

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


func TestNullFileAndGetStats(t *testing.T) {
	file:=FileMetadata{}
	var sv = NewStatsVisualizator()
	err := sv.AddFile(file)

	if assert.Error(t, err) {
		assert.Equal(t, errors.New("Path is empty"), err)
	}
}

func TestAddNoneFileAndGetStats(t *testing.T) {
	var sv = NewStatsVisualizator()
	stats := sv.GetStats()

	assert.Equal(t,int64(0), stats.NumFiles)
	assert.Equal(t,"", stats.LargestFile.Path)
	assert.Equal(t,int64(0), stats.LargestFile.Size)

	assert.Equal(t,0.00, stats.AverageFileSize)
	assert.Equal(t,"", stats.MostFrequentExt.Extension)
	assert.Equal(t,int64(0), stats.MostFrequentExt.NumOccurrences)
	assert.Equal(t,float32(0), stats.TextPercentage)

	assert.Equal(t,0, len(stats.MostRecentPaths))
}
