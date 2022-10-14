package util

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"math/rand"
	"time"
)

var Generator generator

type generator struct {
}

func (*generator) MD5(str string) string {
	m := md5.New()
	m.Write([]byte(str))
	return hex.EncodeToString(m.Sum(nil))
}

func (*generator) UUID() string {
	u := uuid.NewV4()
	return u.String()
}

func (*generator) ValidateCode() string {
	code := fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))
	return code
}
