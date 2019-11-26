package util

import (
	uuid1 "github.com/google/uuid"
	"github.com/rs/xid"
	uuid2 "github.com/satori/go.uuid"
	"github.com/segmentio/ksuid"
	"strings"
)

func NewXid() string {
	return xid.New().String()
}
func NewKsuid() string {
	return ksuid.New().String()
}

func NewUuid() string {
	u1, err := uuid1.NewUUID()
	if err != nil {
		u2 := uuid2.NewV4();
		return strings.ReplaceAll(u2.String(), "-", "")
	} else {
		return strings.ReplaceAll(u1.String(), "-", "")
	}
}

