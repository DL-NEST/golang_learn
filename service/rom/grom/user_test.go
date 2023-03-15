package main

import (
	"context"
	"encoding/json"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
	"log"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

type Runtime struct {
	startTime int64
}

func start() Runtime {
	r := Runtime{
		startTime: time.Now().UnixNano(),
	}
	return r
}

func (r Runtime) end(index uint) string {
	end := float64((time.Now().UnixNano()-r.startTime)/(time.Second.Nanoseconds()/100)) / 100
	return fmt.Sprintf("运行时长%0.3f\n\t\t\t每秒次数:%0.f", end, float64(index)/end)
}

func BenchmarkName(b *testing.B) {
	db := LinkMySql()
	b.Log()
	for i := 0; i < b.N; i++ {
		err := db.Create(&DeviceMsg{
			BaseModel:     BaseModel{},
			DeviceSubject: "deviceName",
			DeviceMsg:     "bool",
		}).Error
		if err != nil {
			log.Printf("%v", err)
		}
	}
}

func BenchmarkRedis(b *testing.B) {
	RdbAuth := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	deviceMsg := DeviceMsg{
		BaseModel:     BaseModel{},
		DeviceSubject: "deviceName",
		DeviceMsg:     "bool",
	}
	res, _ := json.Marshal(deviceMsg)
	for i := 0; i < b.N; i++ {
		RdbAuth.LPush(context.Background(), "cache", res)
	}
}

func TestMySqlInsert(t *testing.T) {
	t.Run("MySqlInsert", func(t *testing.T) {
		db := LinkMySql()
		r := start()
		for i := 0; i < 1000; i++ {
			err := db.Create(&DeviceMsg{
				BaseModel:     BaseModel{},
				DeviceSubject: "deviceName",
				DeviceMsg:     "bool",
			}).Error
			if err != nil {
				log.Printf("%v", err)
			}
		}
		t.Logf(" %s", r.end(1000))
	})
	t.Run("MySqlInsertPipe", func(t *testing.T) {
		db := LinkMySql()
		r := start()
		err := db.Transaction(func(tx *gorm.DB) error {
			for i := 0; i < 1000; i++ {
				tx.Create(&DeviceMsg{
					BaseModel:     BaseModel{},
					DeviceSubject: "deviceName",
					DeviceMsg:     "bool",
				})
			}
			return nil
		})
		if err != nil {
			log.Panic("提交失败")
		}
		t.Logf(" %s", r.end(1000))
	})
	t.Run("MySqlInsertMultiplePipe", func(t *testing.T) {
		db := LinkMySql()
		r := start()
		// mysql的max_allowed_packet
		maxAllowed := 1024
		err := db.Transaction(func(tx *gorm.DB) error {
			for i := 0; i < 10000; {
				var outStr string
				outStr = "INSERT INTO `device_msgs` (`created_at`,`updated_at`,`deleted_at`,`device_subject`,`device_msg`) VALUES "
				for a := 0; a < maxAllowed/177; a++ {
					times := time.Now().Format("2006-01-02 15:04:05.000")
					outStr += "(" + "'" + times + "','" + times + "',NULL,'" + "topic" + "','" + "msg" + "')"
					if a < maxAllowed/177-1 {
						outStr += ","
					} else {
						outStr += ";"
					}
					i++
				}
				tx.Exec(outStr)
			}
			return nil
		})
		if err != nil {
			log.Panic("提交失败")
		}
		t.Logf(" %s", r.end(10000))
	})
}

type msgType struct {
	Topic string
	Msg   string
}

func BenchmarkJson(b *testing.B) {
	aa := msgType{
		Topic: "sas",
		Msg:   "sca",
	}
	var cc msgType
	for i := 0; i < b.N; i++ {
		res, _ := json.Marshal(aa)
		_ = json.Unmarshal(res, &cc)
	}
}

type MsgType struct {
	Topic string
	Msg   string
}

func TestRedis(t *testing.T) {
	RdbAuth := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	deviceMsg := MsgType{
		Topic: "test",
		Msg:   "{fwsfvwv}",
	}
	res, _ := json.Marshal(deviceMsg)
	r := start()
	_, err := RdbAuth.Pipelined(context.Background(), func(pipe redis.Pipeliner) error {
		for i := 0; i < 1000; i++ {
			pipe.RPush(context.Background(), "logCache", res)
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
	t.Logf(" %s", r.end(100000))
}

func Worker() {
	//RdbAuth := redis.NewClient(&redis.Options{
	//	Addr:     "127.0.0.1:6379",
	//	Password: "", // no password set
	//	DB:       0,  // use default DB
	//})
	//lens := RdbAuth.LLen(context.Background(), "C").Val()
	//if lens > 1 {
	//	val := RdbAuth.LRange(context.Background(), "C", 0, lens).Val()
	//	for _, s := range val {
	//		pj("ss", s)
	//	}
	//}
}

func BenchmarkSQLGenerate(b *testing.B) {
	db := LinkMySql()
	b.Run("ToSQL", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = db.ToSQL(func(tx *gorm.DB) *gorm.DB {
				return tx.Create(&DeviceMsg{
					BaseModel:     BaseModel{},
					DeviceSubject: "deviceName",
					DeviceMsg:     "bool",
				})
			})
		}
	})
	b.Run("字符串拼接", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = fmt.Sprintf("INSERT INTO `device_msgs` (`created_at`,`updated_at`,`deleted_at`,`device_subject`,`device_msg`) "+
				"VALUES (%s,%s,NULL,%s,%s)",
				time.Now(), time.Now(), "deviceName", "bool")
		}
	})
	b.Run("字符串拼接+", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			topic := "ss/ca"
			msg := "{}"
			times := time.Now().Format("2006-01-02 15:04:05.000")
			_ = "INSERT INTO `device_msgs` (`created_at`,`updated_at`,`deleted_at`,`device_subject`,`device_msg`) VALUES " +
				"(" + "'" + times + "','" + times + "',NULL,'" + topic + "','" + msg + "');"
		}
	})
}

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

func TestMqttTest(t *testing.T) {
	//mqtt.DEBUG = log.New(os.Stdout, "", 0)
	//mqtt.ERROR = log.New(os.Stdout, "", 0)

	var cList []mqtt.Client

	// 创建1000个客户端
	for i := 0; i < 10000; i++ {
		cList = append(cList, mqttCreate(fmt.Sprintf("test/%d", i)))
	}

	time.Sleep(20 * time.Second)

	var wg sync.WaitGroup

	wg.Add(len(cList))

	for index, _ := range cList {
		go func() {
			for i := 0; i < 10; i++ {
				cList[index].Publish("test/state", 0, false, fmt.Sprintf("ss:%d", 3)).Wait()
				cList[index].Publish("test/ctrl", 0, false, fmt.Sprintf("ss:%d", 3)).Wait()
			}
			wg.Done()
		}()
	}
	wg.Wait()
}

func TestLong(t *testing.T) {
	var wg sync.WaitGroup
	var num int64 = 0
	var atomicNum atomic.Int64
	atomicNum.Store(0)
	wg.Add(10000)

	for i := 0; i < 10000; i++ {
		go func() {
			num++
			atomicNum.Add(1)
			wg.Done()
		}()
	}

	wg.Wait()

	t.Logf("num:%d", num)
	t.Logf("num:%d", atomicNum.Load())
}
