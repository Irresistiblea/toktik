package dao

import (
	"strconv"
	"strings"
)

func FindVideoById(cid int) string {
	sql := "select content from video where cid = "
	sql += strconv.Itoa(cid)
	// Do the query...
	return strings.Join([]string{"Hello World ", strconv.Itoa(cid)}, "")
}
