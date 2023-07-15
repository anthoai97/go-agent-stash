package serializer

import (
	"io/ioutil"
	"os"

	"google.golang.org/grpc/grpclog"
)

func CustomLogger() grpclog.LoggerV2 {
	return grpclog.NewLoggerV2(os.Stdout, ioutil.Discard, ioutil.Discard)
}
