package stats

import (
	"errors"
	"path/filepath"
	"sort"
	"sync"
)

type StatsVisualization interface {
	AddFile(metadata FileMetadata) error
	GetStats() FileStats
}

type statsVisualizator struct {
	sync.RWMutex
	NumFiles int64
	LargestFile LargestFileInfo
	TotalFilesSize int64
	ExtensionInfos map[string]int64
	NumFilesText int64
	MostRecentPaths *Queue
}

func NewStatsVisualizator() StatsVisualization {
	var sv statsVisualizator
	sv.ExtensionInfos = map[string]int64{}
	sv.MostRecentPaths = CreateQueue(10)
	return &sv
}

func (sv *statsVisualizator) AddFile(metadata FileMetadata) error {
	sv.Lock()
	defer sv.Unlock()

	if error := sv.isValidFile(metadata);error != nil {
		return error
	}

	sv.NumFiles += 1
	if metadata.Size > sv.LargestFile.Size {
		var largestFile LargestFileInfo
		largestFile.Size = metadata.Size
		largestFile.Path = metadata.Path
		sv.LargestFile = largestFile
	}

	sv.TotalFilesSize += metadata.Size

	extension := filepath.Ext(metadata.Path)
	if _, ok := sv.ExtensionInfos[extension]; ok {
		sv.ExtensionInfos[extension] +=1
	} else {
		sv.ExtensionInfos[extension] =1
	}

	if !metadata.IsBinary {
		sv.NumFilesText +=1
	}

	error := sv.MostRecentPaths.Insert(metadata.Path)
	if error != nil {
		return error
	}

	return nil
}

func (sv *statsVisualizator) GetStats() FileStats {
	if sv.NumFiles == 0 {
		return NewFileStats()
	}

	var fileStat FileStats
	fileStat.NumFiles = sv.NumFiles
	fileStat.LargestFile = sv.LargestFile
	fileStat.AverageFileSize = float64(sv.TotalFilesSize) / float64(sv.NumFiles)

	if len(sv.ExtensionInfos) != 0 {
		fileStat.MostFrequentExt = sv.getMostFrequentExtension()
	}

	fileStat.TextPercentage =  float32(sv.NumFilesText) / float32(sv.NumFiles) * float32(100)
	fileStat.MostRecentPaths = sv.MostRecentPaths.q

	return fileStat
}

func (sv *statsVisualizator) getMostFrequentExtension() ExtInfo{
	var ss []ExtInfo
	for k, v := range sv.ExtensionInfos {
		ss = append(ss, ExtInfo{k, v})
	}

	sort.Slice(ss, func(i, j int) bool {
		return ss[i].NumOccurrences > ss[j].NumOccurrences
	})

	return ss[0]
}

func (sv *statsVisualizator) isValidFile(metadata FileMetadata) error {
	if len(metadata.Path) == 0 {
		return errors.New("Path is empty")
	}
	return nil
}


type FileMetadata struct {
	Path string `json:"path"` // the file's absolute path
	Size int64 `json:"size"` // the file size in bytes
	IsBinary bool `json:"is_binary"` // whether the file is a binary file or a simple text file
}

type FileStats struct {
	NumFiles int64 `json:"num_files"` //number of files received
	LargestFile LargestFileInfo `json:"largest_file"` //largest file received
	AverageFileSize float64 `json:"avg_file_size"`  //average file size
	MostFrequentExt ExtInfo `json:"most_frequent_ext"`  //most frequent file extension (including number of occurences)
	TextPercentage float32 `json:"text_percentage"` //percentage of text file in all files received
	MostRecentPaths []string `json:"most_recent_paths"` //list of latest 10 file paths received
}

type LargestFileInfo struct {
	Path string `json:"path"`
	Size int64 `json:"size"`
}

type ExtInfo struct {
	Extension string `json:"extension"`
	NumOccurrences int64 `json:"num_occurrences"`
}

func NewFileStats() FileStats {
	stats := FileStats{}
	stats.LargestFile = LargestFileInfo{Path:"",Size:0}
	stats.AverageFileSize = 0.00
	stats.MostFrequentExt = ExtInfo{Extension:"",NumOccurrences:0}
	stats.TextPercentage = float32(0)
	stats.MostRecentPaths = []string{}
	return stats
}