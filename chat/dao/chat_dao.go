package dao

import (
	"strconv"
	"strings"
)

func FindChatById(cid int) string {
	sql := "select content from chat where cid = "
	sql += strconv.Itoa(cid)
	// Do the query...
	return strings.Join([]string{"Hello World ", strconv.Itoa(cid)}, "")
}
