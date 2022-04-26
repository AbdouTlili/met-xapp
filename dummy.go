package main

import (
	"context"
	"fmt"

	e2api "github.com/onosproject/onos-api/go/onos/e2t/e2/v1beta1"
	topoapi "github.com/onosproject/onos-api/go/onos/topo"
	rc "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc_pre/pdubuilder"
	e2client "github.com/onosproject/onos-ric-sdk-go/pkg/e2/v1beta1"
	toposdk "github.com/onosproject/onos-ric-sdk-go/pkg/topo"
	"google.golang.org/protobuf/proto"
)

const (
	serviceModelName    = "oran-e2sm-rc-pre"
	serviceModelVersion = "v2"
)

func main() {
	fmt.Println("Getting E2 node ids...")
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	sdkClient, err := toposdk.NewClient()
	if err != nil {
		fmt.Println("topoclient creation error: ", err.Error())
	}
	objects, err := sdkClient.List(ctx, toposdk.WithListFilters(getControlRelationFilter()))
	if err != nil {
		fmt.Println("topoclient list nodes error: ", err.Error())
	}

	e2NodeIDs := make([]topoapi.ID, len(objects))
	for _, object := range objects {
		relation := object.Obj.(*topoapi.Object_Relation)
		e2NodeID := relation.Relation.TgtEntityID
		if e2NodeID != "" {
			e2NodeIDs = append(e2NodeIDs, e2NodeID)
		}
	}
	fmt.Printf("E2 nodes: %+v", e2NodeIDs)

	fmt.Println("Setting E2 client...")
	appID := e2client.AppID("sample-xapp")

	client := e2client.NewClient(e2client.WithE2TAddress("onos-e2t", 5150),
		e2client.WithServiceModel(e2client.ServiceModelName(serviceModelName), e2client.ServiceModelVersion(serviceModelVersion)),
		e2client.WithAppID(appID),
		e2client.WithEncoding(e2client.ProtoEncoding))

	e2node := client.Node(e2client.NodeID(e2NodeIDs[len(e2NodeIDs)-1]))
	subName := "sample-subscription"
	eventTriggerData, err := CreatePciEventTrigger()
	if err != nil {
		fmt.Println("Error initiating event trigger ", err.Error())
		return
	}
	fmt.Printf("\n event trigger %+v", eventTriggerData)
	actions := CreateSubscriptionActions()
	fmt.Printf("\n action trigger %+v", actions)
	subSpec := e2api.SubscriptionSpec{
		Actions: actions,
		EventTrigger: e2api.EventTrigger{
			Payload: eventTriggerData,
		},
	}
	fmt.Println("Subscription Details configured:")
	fmt.Println(subSpec)
	ch := make(chan e2api.Indication)
	ctx1, cancel1 := context.WithCancel(context.Background())
	defer cancel1()
	channelID, err := e2node.Subscribe(ctx1, subName, subSpec, ch)
	if err != nil {
		fmt.Println("subscription error: ", err.Error())
	}

	fmt.Println("channel id", channelID)
}

func CreateSubscriptionActions() []e2api.Action {
	actions := make([]e2api.Action, 0)
	action := &e2api.Action{
		ID:   int32(0),
		Type: e2api.ActionType_ACTION_TYPE_REPORT,
		SubsequentAction: &e2api.SubsequentAction{
			Type:       e2api.SubsequentActionType_SUBSEQUENT_ACTION_TYPE_CONTINUE,
			TimeToWait: e2api.TimeToWait_TIME_TO_WAIT_ZERO,
		},
	}
	actions = append(actions, *action)
	return actions

}

func CreatePciEventTrigger() ([]byte, error) {
	e2smRcEventTriggerDefinition, err := rc.CreateE2SmRcPreEventTriggerDefinitionUponChange()
	if err != nil {
		return []byte{}, err
	}

	err = e2smRcEventTriggerDefinition.Validate()
	if err != nil {
		return []byte{}, err
	}

	protoBytes, err := proto.Marshal(e2smRcEventTriggerDefinition)
	if err != nil {
		return []byte{}, err
	}

	return protoBytes, nil
}
func getControlRelationFilter() *topoapi.Filters {
	controlRelationFilter := &topoapi.Filters{
		KindFilter: &topoapi.Filter{
			Filter: &topoapi.Filter_Equal_{
				Equal_: &topoapi.EqualFilter{
					Value: topoapi.CONTROLS,
				},
			},
		},
	}
	return controlRelationFilter
}
