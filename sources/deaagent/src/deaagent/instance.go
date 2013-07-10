package deaagent

import (
	"code.google.com/p/gogoprotobuf/proto"
	"deaagent/loggregatorclient"
	"github.com/cloudfoundry/gosteno"
	"logMessage"
	"net"
	"path/filepath"
	"runtime"
	"strconv"
	"time"
)

type instance struct {
	applicationId       string
	wardenJobId         uint64
	wardenContainerPath string
	index               uint64
	logger              *gosteno.Logger
}

func (instance *instance) identifier() string {
	return filepath.Join(instance.wardenContainerPath, "jobs", strconv.FormatUint(instance.wardenJobId, 10))
}

func (inst *instance) startListening(loggregatorClient loggregatorclient.LoggregatorClient) {

	listen := func(messageType logMessage.LogMessage_MessageType) {

		newLogMessage := func(message []byte) *logMessage.LogMessage {
			currentTime := time.Now()
			sourceType := logMessage.LogMessage_DEA
			return &logMessage.LogMessage{
				Message:     message,
				AppId:       proto.String(inst.applicationId),
				MessageType: &messageType,
				SourceType:  &sourceType,
				Timestamp:   proto.Int64(currentTime.UnixNano()),
			}
		}

		socket := func(messageType logMessage.LogMessage_MessageType) (net.Conn, error) {
			var socketName string

			if messageType == logMessage.LogMessage_OUT {
				socketName = "stdout.sock"
			} else {
				socketName = "stderr.sock"
			}
			return net.Dial("unix", filepath.Join(inst.identifier(), socketName))
		}

		connection, err := socket(messageType)
		if err != nil {
			inst.logger.Errorf("Error while dialing into socket %s, %s", messageType, err)
			return
		}
		defer func() {
			connection.Close()
			inst.logger.Infof("Stopped reading from socket %s", messageType)
		}()

		buffer := make([]byte, bufferSize)

		for {
			readCount, err := connection.Read(buffer)
			if readCount == 0 && err != nil {
				inst.logger.Infof("Error while reading from socket %s, %s", messageType, err)
				break
			}
			inst.logger.Debugf("Read %d bytes from instance socket", readCount)

			data, err := proto.Marshal(newLogMessage(buffer[0:readCount]))
			if err != nil {
				inst.logger.Errorf("Error marshalling log message: %s", err)
			}

			loggregatorClient.Send(data)
			inst.logger.Debugf("Sent %d bytes to loggregator client", readCount)
			runtime.Gosched()
		}
	}

	go listen(logMessage.LogMessage_OUT)
	go listen(logMessage.LogMessage_ERR)
}
