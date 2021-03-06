package main

import (
	_ "github.com/go-chassis/go-chassis/bootstrap"
	_ "github.com/go-chassis/go-chassis/config-center"

	"github.com/go-chassis/go-chassis"
	"net/http"
	rf "github.com/go-chassis/go-chassis/server/restful"
	"github.com/go-chassis/go-chassis/client/rest"
	"github.com/go-chassis/go-chassis/core"
	"context"
	"github.com/go-chassis/go-chassis/core/lager"
)

func main() {
	chassis.Init()
	chassis.Run()
}


func init() {
	chassis.RegisterSchema("rest", &RestFulMessage{})
}

type RestFulMessage struct {
}

func (r *RestFulMessage) Saymessage(b *rf.Context) {
	id := b.ReadPathParameter("name")

	var req *rest.Request

	restinvoker := core.NewRestInvoker()
	req, _ = rest.NewRequest("GET", "cse://Server/saymessage/"+id)
	resp1, err := restinvoker.ContextDo(context.TODO(), req)
	if err != nil {
		b.WriteError(http.StatusInternalServerError, err)
		lager.Logger.Errorf("call request fail: %s",err)
		return
	}

	b.Write(resp1.ReadBody())
}

func (s *RestFulMessage) URLPatterns() []rf.Route {
	return []rf.Route{
		{http.MethodGet, "/saymessage/{name}", "Saymessage"},
	}
}
