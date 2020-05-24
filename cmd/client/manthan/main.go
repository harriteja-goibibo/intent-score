package main

import (
	"context"
	"fmt"
	"github.com/goibibo/intent-score/pkg/api/common"
	proto_files "github.com/goibibo/intent-score/pkg/api/grpc"
	"github.com/goibibo/intent-score/pkg/api/grpc/manthan"
	"google.golang.org/grpc"
)

var (
	hostname = "localhost:8066"
)

func main() {
	conn, err := grpc.Dial(hostname, grpc.WithInsecure())
	if err != nil {
		panic("Not able to make client connection")
	}
	defer conn.Close()

	client := proto_files.NewIntentScoreClient(conn)

	obj := &manthan.ManthanRealTimeData{
		Vertical:        common.Vertical_HOTELS,
		Org:             common.Org_GI,
		PageType:        common.PageHitType_SEARCH_PAGE,
		RequestDate:     "2020-05-27T15:01:01",
		UserId:          "test",
		EntityId:        "123",
		TravelStartDate: "20200601",
		TravelEndDate:   "20200603",
	}

	resp, err := client.ManthanRealTimeData(context.Background(), obj)
	if err != nil {
		fmt.Println("error", err.Error())
	}
	fmt.Println("resp", resp)

}
