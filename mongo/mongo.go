package mongo

import (
	"encoding/json"
	"fmt"
	"github.com/globalsign/mgo"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"wdp/conf"
)

var (
	DBS map[string]*mgo.Session = map[string]*mgo.Session{}
)

type DbInfo struct {
	Dsn string `json:"dsn"`
}

func walkFunc(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}

	if info.IsDir() {
		return nil
	}

	if strings.HasPrefix(filepath.Base(path), ".") {
		return nil
	}

	if filepath.Ext(path) != ".json" {
		return nil
	}

	fileContent, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	var mongoConf map[string]*DbInfo
	if err = json.Unmarshal(fileContent, &mongoConf); err != nil {
		return err
	}

	dbInfo, ok := mongoConf[conf.Env]
	if !ok {
		return fmt.Errorf("dbInfo of %s not found", conf.Env)
	}

	if dbInfo.Dsn == "" {
		return fmt.Errorf("%s dsn is nil", path)
	}

	session, err := mgo.Dial(dbInfo.Dsn)
	if err != nil {
		return err
	}

	key := strings.TrimSuffix(filepath.Base(path), filepath.Ext(path))
	DBS[key] = session
	return nil
}

func init() {
	dir := "mongo"

	err := filepath.Walk(dir, walkFunc)
	if err != nil {
		os.Exit(-1)
	}

	log.Printf("mongo conf: %v\n", DBS)
	return
}
