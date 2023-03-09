package main

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"sync"
	"time"
)

func mqttCreate(id string) mqtt.Client {
	opts := mqtt.NewClientOptions().AddBroker("mqtt://127.0.0.1:1883").SetClientID(id)

	opts.SetKeepAlive(60 * time.Second)
	// 设置消息回调处理函数
	opts.SetDefaultPublishHandler(func(client mqtt.Client, msg mqtt.Message) {
		fmt.Printf("TOPIC: %s\n", msg.Topic())
		fmt.Printf("MSG: %s\n", msg.Payload())
	})
	opts.SetPingTimeout(1 * time.Second)

	c := mqtt.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	return c
}

func main() {
	var cList []mqtt.Client
	// 创建1000个客户端
	for i := 0; i < 10000; i++ {
		cList = append(cList, mqttCreate(fmt.Sprintf("test/%d", i)))
	}

	time.Sleep(2 * time.Second)

	var wg sync.WaitGroup
	wg.Add(len(cList))

	for index, _ := range cList {
		go func() {
			for i := 0; i < 2; i++ {
				cList[index].Publish("test/state", 0, false, fmt.Sprintf("ss:%d", 3)).Wait()
				//cList[index].Publish("test/ctrl", 0, false, fmt.Sprintf("ss:%d", 3)).Wait()
			}
			wg.Done()
		}()
	}
	wg.Wait()
}

func LinkMySql() *gorm.DB {
	// 打开数据库
	db, err := gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@(%s)/%s?charset=%s&parseTime=True&loc=Local",
		"root",
		"example",
		"127.0.0.1:3306",
		"test",
		"utf8")), &gorm.Config{
		CreateBatchSize: 1000,
	})
	if err != nil {
		log.Fatalln("连接数据库失败")
	}
	return db
}
