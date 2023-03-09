package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/redis/go-redis/v9"
	"google.golang.org/grpc"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net"
	"strconv"
	"strings"
	"sync"
	"time"

	pb "golang_learn/amqp/emqx/proto"
	utils "golang_learn/amqp/emqx/utils"
)

const (
	port = ":9981"
)

var Rdb *redis.Client
var index uint
var cnter *utils.Counter = utils.NewCounter(0, 100)
var lock sync.Mutex

// server is used to implement emqx_exhook_v1.s *server
type server struct {
	pb.UnimplementedHookProviderServer
}

// HookProviderServer callbacks

func (s *server) OnProviderLoaded(ctx context.Context, in *pb.ProviderLoadedRequest) (*pb.LoadedResponse, error) {
	cnter.Count(1)
	fmt.Println("hook 创建")
	hooks := []*pb.HookSpec{
		//&pb.HookSpec{Name: "client.connect"},
		//&pb.HookSpec{Name: "client.connack"},
		//&pb.HookSpec{Name: "client.connected"},
		//&pb.HookSpec{Name: "client.disconnected"},
		//&pb.HookSpec{Name: "client.authenticate"},
		//&pb.HookSpec{Name: "client.authorize"},
		//&pb.HookSpec{Name: "client.subscribe"},
		//&pb.HookSpec{Name: "client.unsubscribe"},
		//&pb.HookSpec{Name: "session.created"},
		//&pb.HookSpec{Name: "session.subscribed"},
		//&pb.HookSpec{Name: "session.unsubscribed"},
		//&pb.HookSpec{Name: "session.resumed"},
		//&pb.HookSpec{Name: "session.discarded"},
		//&pb.HookSpec{Name: "session.takenover"},
		//&pb.HookSpec{Name: "session.terminated"},
		&pb.HookSpec{Name: "message.publish"},
		//&pb.HookSpec{Name: "message.delivered"},
		//&pb.HookSpec{Name: "message.acked"},
		//&pb.HookSpec{Name: "message.dropped"},
	}
	return &pb.LoadedResponse{Hooks: hooks}, nil
}

func out(title string, a ...any) {
	fmt.Println(title)
	fmt.Println("==============================================================")
	fmt.Printf("%v\n", a)
	fmt.Printf("==============================================================\n\n")
}

func (s *server) OnProviderUnloaded(ctx context.Context, in *pb.ProviderUnloadedRequest) (*pb.EmptySuccess, error) {
	cnter.Count(1)
	fmt.Println("hook 注销")
	return &pb.EmptySuccess{}, nil
}

func (s *server) OnClientConnect(ctx context.Context, in *pb.ClientConnectRequest) (*pb.EmptySuccess, error) {
	cnter.Count(1)
	return &pb.EmptySuccess{}, nil
}

func (s *server) OnClientConnack(ctx context.Context, in *pb.ClientConnackRequest) (*pb.EmptySuccess, error) {
	cnter.Count(1)
	return &pb.EmptySuccess{}, nil
}

func (s *server) OnClientConnected(ctx context.Context, in *pb.ClientConnectedRequest) (*pb.EmptySuccess, error) {
	cnter.Count(1)
	//out("客户端连接 hook", in)
	return &pb.EmptySuccess{}, nil
}

func (s *server) OnClientDisconnected(ctx context.Context, in *pb.ClientDisconnectedRequest) (*pb.EmptySuccess, error) {
	cnter.Count(1)
	//out("客户端关闭连接 hook", in)
	return &pb.EmptySuccess{}, nil
}

func (s *server) OnClientAuthenticate(ctx context.Context, in *pb.ClientAuthenticateRequest) (*pb.ValuedResponse, error) {
	cnter.Count(1)
	reply := &pb.ValuedResponse{}
	reply.Type = pb.ValuedResponse_STOP_AND_RETURN
	reply.Value = &pb.ValuedResponse_BoolResult{BoolResult: true}
	return reply, nil
}

func (s *server) OnClientAuthorize(ctx context.Context, in *pb.ClientAuthorizeRequest) (*pb.ValuedResponse, error) {
	cnter.Count(1)
	reply := &pb.ValuedResponse{}
	reply.Type = pb.ValuedResponse_STOP_AND_RETURN
	reply.Value = &pb.ValuedResponse_BoolResult{BoolResult: true}
	return reply, nil
}

func (s *server) OnClientSubscribe(ctx context.Context, in *pb.ClientSubscribeRequest) (*pb.EmptySuccess, error) {
	cnter.Count(1)
	//out("客户端订阅 hook", in)
	return &pb.EmptySuccess{}, nil
}

func (s *server) OnClientUnsubscribe(ctx context.Context, in *pb.ClientUnsubscribeRequest) (*pb.EmptySuccess, error) {
	cnter.Count(1)
	return &pb.EmptySuccess{}, nil
}

func (s *server) OnSessionCreated(ctx context.Context, in *pb.SessionCreatedRequest) (*pb.EmptySuccess, error) {
	cnter.Count(1)
	return &pb.EmptySuccess{}, nil
}
func (s *server) OnSessionSubscribed(ctx context.Context, in *pb.SessionSubscribedRequest) (*pb.EmptySuccess, error) {
	cnter.Count(1)
	//out("客户端会话订阅 hook", in)
	return &pb.EmptySuccess{}, nil
}

func (s *server) OnSessionUnsubscribed(ctx context.Context, in *pb.SessionUnsubscribedRequest) (*pb.EmptySuccess, error) {
	cnter.Count(1)
	return &pb.EmptySuccess{}, nil
}

func (s *server) OnSessionResumed(ctx context.Context, in *pb.SessionResumedRequest) (*pb.EmptySuccess, error) {
	cnter.Count(1)
	return &pb.EmptySuccess{}, nil
}

func (s *server) OnSessionDiscarded(ctx context.Context, in *pb.SessionDiscardedRequest) (*pb.EmptySuccess, error) {
	cnter.Count(1)
	return &pb.EmptySuccess{}, nil
}

func (s *server) OnSessionTakenover(ctx context.Context, in *pb.SessionTakenoverRequest) (*pb.EmptySuccess, error) {
	cnter.Count(1)
	return &pb.EmptySuccess{}, nil
}

func (s *server) OnSessionTerminated(ctx context.Context, in *pb.SessionTerminatedRequest) (*pb.EmptySuccess, error) {
	cnter.Count(1)
	return &pb.EmptySuccess{}, nil
}

type msgType struct {
	Topic string
	Msg   string
}

func (s *server) OnMessagePublish(ctx context.Context, in *pb.MessagePublishRequest) (*pb.ValuedResponse, error) {
	cnter.Count(1)
	//out("消息 hook", in)
	//fmt.Printf("%s\n", in.Message.Topic)
	//fmt.Printf("%s\n", in.Message.Payload)
	msg := msgType{
		Topic: in.Message.Topic,
		Msg:   strconv.Itoa(int(index)),
	}
	res, _ := json.Marshal(msg)
	index++
	Rdb.RPush(ctx, "logCache", res)
	//	//fmt.Printf("%d\n", index)
	//	//in.Message.Payload = []byte("hardcode payload by exhook-svr-go :)")
	//	//reply := &pb.ValuedResponse{}
	//	//reply.Type = pb.ValuedResponse_STOP_AND_RETURN
	//	//reply.Value = &pb.ValuedResponse_Message{Message: in.Message}
	return &pb.ValuedResponse{}, nil
}

//case2: stop publish the `t/d` messages
//func (s *server) OnMessagePublish(ctx context.Context, in *pb.MessagePublishRequest) (*pb.ValuedResponse, error) {
//	cnter.Count(1)
//    if in.Message.Topic == "t/d" {
//        in.Message.Headers["allow_publish"] = "false"
//        in.Message.Payload = []byte("")
//    }
//	reply := &pb.ValuedResponse{}
//	reply.Type = pb.ValuedResponse_STOP_AND_RETURN
//	reply.Value = &pb.ValuedResponse_Message{Message: in.Message}
//	return reply, nil
//}

func (s *server) OnMessageDelivered(ctx context.Context, in *pb.MessageDeliveredRequest) (*pb.EmptySuccess, error) {
	cnter.Count(1)
	return &pb.EmptySuccess{}, nil
}

func (s *server) OnMessageDropped(ctx context.Context, in *pb.MessageDroppedRequest) (*pb.EmptySuccess, error) {
	cnter.Count(1)
	return &pb.EmptySuccess{}, nil
}

func (s *server) OnMessageAcked(ctx context.Context, in *pb.MessageAckedRequest) (*pb.EmptySuccess, error) {
	cnter.Count(1)
	return &pb.EmptySuccess{}, nil
}

func timeWork() {
	if Rdb == nil {
		return
	}
	db := LinkMySql()
	for {
		time.Sleep(2 * time.Second)
		lens := Rdb.LLen(context.Background(), "logCache").Val()
		if lens > 1 {
			fmt.Println(lens)
			dump(db, lens)
		}
	}
}

func dump(db *gorm.DB, list int64) {
	// mysql的max_allowed_packet
	maxAllowed := 1024
	outStrTemplate := "INSERT INTO `device_msgs` (`created_at`,`updated_at`,`deleted_at`,`device_subject`,`device_msg`) VALUES "

	dumpList := Rdb.LRange(context.Background(), "logCache", 0, list).Val()
	var sqlList = make([]string, 1)
	var i uint = 0
	// 可以提取出来统一转换
	for ii, s := range dumpList {
		// 反射取到的数据
		var msgList msgType
		err := json.Unmarshal([]byte(s), &msgList)
		if err != nil {
			fmt.Printf("erro\n")
			return
		}
		// 添加的插入
		times := time.Now().Format("2006-01-02 15:04:05.000")
		newStr := "(" + "'" + times + "','" + times + "',NULL,'" + msgList.Topic + "','" + msgList.Msg + "')"
		newLen := len(sqlList[i]) + len(outStrTemplate) + len(newStr)
		// 拼接sql语句
		if newLen < maxAllowed {
			sqlList[i] += newStr + ","
		} else {
			sqlList[i] = strings.TrimRight(sqlList[i], ",")
			sqlList = append(sqlList, newStr+",")
			i++
		}
		if ii == len(dumpList)-1 {
			sqlList[i] = strings.TrimRight(sqlList[i], ",")
		}
	}
	// 提交事务
	transaction := db.Begin()
	for _, ss := range sqlList {
		transaction.Exec(outStrTemplate + ss + ";")
	}
	err := transaction.Commit().Error
	if err != nil {
		// 回滚事件
		fmt.Println(transaction.Rollback().Error)
		return
	}
	for i2 := 0; i2 < len(dumpList); i2++ {
		Rdb.RPop(context.Background(), "logCache")
	}
	//Rdb.RPop(context.Background(), "logCache")
	//// 清除redis数据
	//Rdb.LTrim(context.Background(), "logCache", int64(len(dumpList)), 0)

}

func main() {
	Rdb = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	go timeWork()
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterHookProviderServer(s, &server{})
	log.Println("Started gRPC server on ::9981")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
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
