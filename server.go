package zcache

import (
	"fmt"
	"strings"

	"github.com/armnerd/zcache/server/common"

	"github.com/aceld/zinx/ziface"
	"github.com/aceld/zinx/znet"
)

// Router 路由
type Router struct {
	znet.BaseRouter
}

type Server struct {
}

// Handle 处理器
func (rt *Router) Handle(request ziface.IRequest) {
	var command = string(request.GetData())
	// 读取客户端的数据
	fmt.Println("recv from client : msgId=", request.GetMsgID(), ", data=", command)

	// 处理参数, 分发处理
	var args []string
	var temp = strings.Split(command, " ")
	for _, arg := range temp {
		if arg != "" {
			args = append(args, arg)
		}
	}
	var res = common.Handler(args)

	// 回写消息
	var echo = fmt.Sprint(res)
	err := request.GetConnection().SendBuffMsg(0, []byte(echo))
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	s := znet.NewServer()
	common.Init()
	s.AddRouter(0, &Router{})
	s.Serve()
}
