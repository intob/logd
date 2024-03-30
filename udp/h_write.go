package udp

import (
	"fmt"
	"strings"

	"github.com/inneslabs/logd/cmd"
	"github.com/inneslabs/logd/sign"
	"google.golang.org/protobuf/proto"
)

func (svc *UdpSvc) handleWrite(c *cmd.Cmd, pkg *sign.Pkg) {
	valid, err := svc.signer.Verify(svc.writeSecret, pkg)
	if !valid || err != nil {
		return
	}
	if !svc.guard.Good(pkg.Sum) {
		return
	}
	msgBytes, err := proto.Marshal(c.Msg)
	if err != nil {
		return
	}
	msgKey := c.Msg.GetKey()
	segments := strings.Split(msgKey, "/")
	if len(segments) < 3 {
		return
	}
	storeKey := fmt.Sprintf("/%s/%s", segments[1], segments[2])
	svc.logStore.Write(storeKey, msgBytes)
	svc.forSubs <- &ProtoPair{
		Msg:   c.Msg,
		Bytes: msgBytes,
	}
}