package client

import (
	"fmt"
	"net"
	"strconv"
	"strings"

	"github.com/ContinuumLLC/platform-common-lib/src/plugin/protocol"
	chttp "github.com/ContinuumLLC/platform-common-lib/src/plugin/protocol/http"
)

func WriteToPipeWithConn(conn net.Conn) {
	rsp, err := sendToConn("response from the server", conn, "/path", 0)
	if err != nil {
		errStr := fmt.Sprintf("Handshake failed. Err: %v", err)
		panic(errStr)
	}
	fmt.Println(rsp)
}

func sendToConn(message string, conn net.Conn, path string, i int64) (*protocol.Response, error) {

	//***  Use common-lib code to create HTTP client request-response ***
	//***  instead of directly using  golang http package as above  ***
	request := protocol.NewRequest()
	request.Path = path
	request.Body = strings.NewReader(message)
	request.Headers.SetKeyValue(protocol.HdrTransactionID, "ID"+strconv.Itoa(int(i)))
	fmt.Printf("Sending data:%s, to path= %s \n", message, request.Path)
	fmt.Printf("Sending data:%+v \n", request)

	client := chttp.ClientHTTPFactory{}.GetClient(conn, conn)
	err := client.SendRequest(request)
	if err != nil {
		fmt.Printf("Error while sending plugin data to AC %+v\n", err)
		return nil, err
	}

	return nil, err
}
