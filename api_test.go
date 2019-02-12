package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"testing"

	"github.com/berlingoqc/fxbe/api"
	"github.com/berlingoqc/fxbe/auth"
	"github.com/berlingoqc/fxbe/files"
)

var addr = "localhost:24323"
var wdir string
var ws *api.WebServer

var ChuckSize = 1024

var FileSize = []int{
	1024,
	1048576,
	1073741824,
	//1073741824 * 5,
	//1073741824 * 10,
}

var FileName = []string{
	"1KB", "1MB", "1GB", //"5GB", //"10GB",
}

func CreateTestFile(filepath string, size int) {
	f, er := os.Create(filepath)
	if er != nil {
		panic(er)
	}
	defer f.Close()
	b := make([]byte, ChuckSize)
	nbrBlock := size / ChuckSize
	for i := 0; i < nbrBlock; i++ {
		f.Seek(int64(i*ChuckSize), 0)
		_, err := f.Write(b)
		if err != nil {
			panic(err)
		}
	}
}

func SetupTest() {
	auth.GetIAuth = func() auth.IAuth {
		d := &auth.DumpAuth{
			Account: make([]*auth.User, 0),
		}
		p, _ := auth.GetSaltedHash("admin")

		d.Account = append(d.Account, &auth.User{
			ID:       0,
			Username: "admin",
			SaltedPW: p,
		})

		return d
	}

	wdir, _ = os.Getwd()
	wdir = filepath.Join(wdir, "test")

	config := &api.Config{
		Bind: addr,
		Root: wdir,
	}
	if err := config.Save("config.json"); err != nil {
		panic(err)
	}

	folders := []string{
		filepath.Join(wdir, "upload"), filepath.Join(wdir, "download"),
	}

	if err := files.EnsureFolderExists(folders); err != nil {
		panic(err)
	}

	for i := 0; i < len(FileName); i++ {
		filetest := filepath.Join(wdir, "upload", FileName[i])
		CreateTestFile(filetest, FileSize[i])
	}

	ws, _ = api.GetWebServer(config)
}

func CleanupTest() {
	os.RemoveAll(wdir)
}

func TestApi(t *testing.T) {
	SetupTest()
	ws.StartAsync()

	defer func() {
		CleanupTest()
		ws.Stop()
	}()

	client := &http.Client{}

	addr = "http://" + addr

	if resp, err := client.Get(addr + "/auth/?username=admin&password=admin"); err != nil {
		t.Fatal(err)
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		println(string(data))
	}

	resp, err := client.Get(addr + "/files/")
	if err != nil {
		t.Fatal(err)
	}
	var f []files.FileInfo
	data, _ := ioutil.ReadAll(resp.Body)
	print(string(data))
	if err = json.Unmarshal(data, &f); err != nil {
		t.Fatal(err)
	}

}
