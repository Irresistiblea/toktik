package service

import "chat/dao"

func SayHelloService(cid int) string {
	return dao.FindVideoById(cid)
}
