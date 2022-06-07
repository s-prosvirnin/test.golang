package benchmarks

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"strings"
	"testing"

	"github.com/ugorji/go/codec"
)

func makeList() []string {
	var list []string
	for i := 0; i <= 100; i++ {
		list = append(list, "38467")
	}

	return list
}

func list() {
	list := makeList()
	listStr := strings.Trim(strings.Replace(fmt.Sprint(list), " ", ",", -1), "[]")
	_ = strings.Split(listStr, ",")
}

func msgpack() {
	var (
		r  io.Reader
		w  io.Writer
		b  []byte
		mh codec.MsgpackHandle
	)
	list := makeList()

	enc := codec.NewEncoder(w, &mh)
	enc = codec.NewEncoderBytes(&b, &mh)
	_ = enc.Encode(list)
	// fmt.Println(b)

	var listRes []string
	dec := codec.NewDecoder(r, &mh)
	dec = codec.NewDecoderBytes(b, &mh)
	_ = dec.Decode(&listRes)
	// fmt.Println(listRes)
}

func jsonEncoder() {
	list := makeList()

	buf := new(bytes.Buffer)
	jsonEncoder := json.NewEncoder(buf)
	_ = jsonEncoder.Encode(list)
	// fmt.Println(buf)

	var listRes []string
	jsonDecoder := json.NewDecoder(buf)
	_ = jsonDecoder.Decode(&listRes)
	// fmt.Println(listRes)
}

func jsonMarshal() {
	list := makeList()

	buf, _ := json.Marshal(list)
	// fmt.Println(buf)

	var listRes []string
	_ = json.Unmarshal(buf, &listRes)
	// fmt.Println(listRes)
}

func BenchmarkList(b *testing.B) {
	for i := 0; i < b.N; i++ {
		list()
	}
}

func BenchmarkMsgPack(b *testing.B) {
	for i := 0; i < b.N; i++ {
		msgpack()
	}
}

func BenchmarkJsonEncoder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		jsonEncoder()
	}
}

func BenchmarkJsonMarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		jsonMarshal()
	}
}
