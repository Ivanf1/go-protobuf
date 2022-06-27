package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	pb "github.com/Ivanf1/go-protobuf/protos"

	"github.com/julienschmidt/httprouter"
	"google.golang.org/protobuf/proto"
)

func DeviceMessageHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	device := pb.Device{}

	data, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Println(err)
	}

	if err := proto.Unmarshal(data, &device); err != nil {
		fmt.Println(err)
	}

	println(device.Id, ":", device.Name)
}

func main() {
	router := httprouter.New()
	router.POST("/", DeviceMessageHandler)

	log.Fatal(http.ListenAndServe(":3000", router))
}
