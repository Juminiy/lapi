package storage

import "log"

func init(){

	// 如果连接失败,尝试重连,尽可能不让服务失败
	if err := RedisConnect() ; err != nil {
		log.Fatalln(err)
	}
	if err := OSSConnect() ; err != nil {
		log.Fatalln(err)
	}
}