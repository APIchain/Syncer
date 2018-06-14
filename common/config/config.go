package config

import (
	"bytes"
	"encoding/json"
	"github.com/Syncer/metadata"
	"io/ioutil"
	"log"
	"os"
)

var Config = "./config.json"

var Version string

const (
	DEFAULTGENLOGTIME = 60
)

type Configuration struct {
	Magic            int64              `json:"Magic"`
	Version          int                `json:"Version"`
	SeedList         []string           `json:"SeedList"`
	HttpRestPort     int                `json:"HttpRestPort"`
	HttpWsPort       int                `json:"HttpWsPort"`
	PrintLevel       int                `json:"PrintLevel"`
	TokenList        []*metadata.ScInfo `json:"TokenList"`
	SyncServer       string             `json:"SyncServer"`
	SyncServerPort   string             `json:"SyncServerPort"`
	LockExpireTime   int64              `json:"LockExpireTime"`
	MinAmount        int64              `json:"MinAmount"`
	MaxLogSize       int64              `json:"MaxLogSize"`
	SyncStart        int64              `json:"SyncStart"`
	DefaultPageCount int64              `json:"DefaultPageCount"`
}

type ConfigFile struct {
	ConfigFile Configuration `json:"Configuration"`
}

var Parameters *Configuration

func init() {
	file, e := ioutil.ReadFile(Config)
	if e != nil {
		log.Fatalf("File error: %v", e)
		os.Exit(1)
	}
	// Remove the UTF-8 Byte Order Mark
	file = bytes.TrimPrefix(file, []byte("\xef\xbb\xbf"))

	config := ConfigFile{}
	e = json.Unmarshal(file, &config)
	if e != nil {
		log.Fatalf("Unmarshal json file erro %v", e)
		os.Exit(1)
	}
	Parameters = &(config.ConfigFile)
}
