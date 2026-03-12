package redis

import (
	"fmt"
	"gateway/core/app"
	"gateway/core/enums"
	"gateway/infrastructure/queue/redis/controllers"
	"time"
)

type ControllerContract interface {
	Handle(queueName string, payload string)
	HandleError(queueName string, error error)
}
type Controller struct {
	AppService app.ApplicationService
}

func (n Controller) Handle(queueName string, payload string) {
	r := getRouter()
	route, found := r[queueName]
	if found {
		route.Action(n.AppService, queueName, []byte(payload), time.Now().UnixMilli())
	}
	//n.AppService.Context.Logger.Log("tm_ctrl_handle_message", fmt.Sprintf("Received Message: %s %s\n", queueName, payload))
	//fmt.Printf("TM Received Message: %s %s\n", queueName, payload)
}
func (n Controller) HandleError(streamName string, error error) {
	//n.AppService.Context.Logger.Error("tm_ctrl_handle_error", error, fmt.Sprintf("Queue name: %s", streamName))
	//fmt.Println("Error during Fetch(): ", error)
}

var routes []ListenerRoute

func init() {
	/*for i := 0; i < 100; i++ {
		h := fmt.Sprintf("%s:%02d", enums.QSearchCardsByPinflTask, i)
		routes = append(routes, ListenerRoute{
			QueueName:    h,
			ThreadsCount: 1, // important
			Action:       controllers.CardController{}.SearchCardsByPinfl,
		})
	}*/
	routes = append(routes, ListenerRoute{
		QueueName:    enums.QSearchCardsByPhoneTask,
		ThreadsCount: 5,
		Action:       controllers.CardController{}.SearchCardsByPhone,
	})
	routes = append(routes, ListenerRoute{
		QueueName:    enums.QSearchCardsByCustomerTask,
		ThreadsCount: 10,
		Action:       controllers.CardController{}.SearchCardsByCustomer,
	})
	routes = append(routes, ListenerRoute{
		QueueName:    enums.QSearchFullCard,
		ThreadsCount: 10,
		Action:       controllers.CardController{}.SearchFullCard,
	})
	routes = append(routes, ListenerRoute{
		QueueName:    enums.QCheckPayment,
		ThreadsCount: 10,
		Action:       controllers.PaymentController{}.CheckPayment,
	})
	routes = append(routes, ListenerRoute{
		QueueName:    enums.QCancelPayment,
		ThreadsCount: 10,
		Action:       controllers.PaymentController{}.CancelPayment,
	})
	routes = append(routes, ListenerRoute{
		QueueName:    enums.QReturnPayment,
		ThreadsCount: 10,
		Action:       controllers.PaymentController{}.ReturnPayment,
	})
	routes = append(routes, ListenerRoute{
		QueueName:    enums.QReconciliationAuth,
		ThreadsCount: 1,
		Action:       controllers.PaymentController{}.ReconciliationAuth,
	})
	routes = append(routes, ListenerRoute{
		QueueName:    enums.QReconciliationApprove,
		ThreadsCount: 1,
		Action:       controllers.PaymentController{}.ReconciliationApprove,
	})
	routes = append(routes, ListenerRoute{
		QueueName:    enums.QPhoneFoundTask,
		ThreadsCount: 10,
		Action:       controllers.UserController{}.PhoneFound,
	})
	routes = append(routes, ListenerRoute{
		QueueName:    enums.QAddPhoneTask,
		ThreadsCount: 5,
		Action:       controllers.UserController{}.AddPhone,
	})
	/*routes = append(routes, ListenerRoute{
		QueueName:    enums.QCheckIntendUsersTask,
		ThreadsCount: 5,
		Action:       controllers.BgController{}.CheckIntendUsers,
	})*/
	routes = append(routes, ListenerRoute{
		QueueName:    enums.QCheckKnownUsersTask,
		ThreadsCount: 5,
		Action:       controllers.BgController{}.CheckKnownUsers,
	})
	routes = append(routes, ListenerRoute{
		QueueName:    enums.QCheckCardPinflTask,
		ThreadsCount: 10,
		Action:       controllers.BgController{}.CheckCardByPinfl,
	})
}
func GetRoutes() []ListenerRoute {
	return routes
}

func getRouter() map[string]ListenerRoute {
	res := map[string]ListenerRoute{}
	for _, route := range routes {
		res[route.QueueName] = route
	}
	return res
}

type ControllerAction func(appSrv app.ApplicationService, queueName string, payload []byte, eventAt int64)
