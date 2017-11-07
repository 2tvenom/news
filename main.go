package main

import (
	"bytes"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/golang/protobuf/proto"
)

const (
	searchUrl = "https://news.google.com/news/exec/fetchSectionedStories"
)

func main() {

	deviceId := genDeviceId()

	locale := "en"
	country := "US"
	//donk know
	someId := "31502265"
	//search keyword
	keyword := "singapore"
	var one int64 = 1

	reqData := &RequestRoot{
		Params: &RequestParams{
			Keyword: &keyword,
		},
		Device: &Device{
			Lang:         &locale,
			Country:      &country,
			SetOne:       &one,
			SetOne2:      &one,
			SetOne3:      &one,
			DeviceId:     &deviceId,
			SomeStringId: &someId,
		},
	}

	reqDataBytes, _ := proto.Marshal(reqData)

	buff := bytes.NewBuffer(reqDataBytes)

	req, _ := http.NewRequest("POST", searchUrl, buff)

	resp, err := http.DefaultClient.Do(req)
	defer resp.Body.Close()

	if err != nil {
		log.Fatal(err.Error())
	}

	p := &ResponseRoot{}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err.Error())
	}

	err = proto.Unmarshal(data, p)
	if err != nil {
		log.Fatal(err.Error())
	}

	out, _ := json.MarshalIndent(p, "", "	")
	fmt.Println(string(out))

}

func genDeviceId() string {
	lenRand := []int{4, 2, 2, 2, 6}

	out := ""

	for _, i := range lenRand {
		b := make([]byte, i)
		rand.Read(b)

		if len(out) > 0 {
			out += "-"
		}
		out += fmt.Sprintf("%x", b)
	}

	return out
}
