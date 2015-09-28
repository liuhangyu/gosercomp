package gosercomp

//go:generate msgp -o msgp_gen.go -io=false -tests=false
//go:generate protoc --go_out=. protobuf.proto
//go:generate  protoc --gofast_out=.  -I. -I$GOPATH/src  -I$GOPATH/src/github.com/gogo/protobuf/protobuf mygogo.proto
//go:generate flatc -g -o .. flatbuffers.fbs

type ColorGroup struct {
	Id     int      `json:"id" xml:"id,attr" msg:"id"`
	Name   string   `json:"name" xml:"name" msg:"name"`
	Colors []string `json:"colors" xml:"colors" msg:"colors"`
}
