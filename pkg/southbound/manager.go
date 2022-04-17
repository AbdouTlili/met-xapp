package southbound

import (
	"context"

	"github.com/AbdouTlili/met-xapp/pkg/rnib"
	topoapi "github.com/onosproject/onos-api/go/onos/topo"
	"github.com/onosproject/onos-lib-go/pkg/logging"
	e2client "github.com/onosproject/onos-ric-sdk-go/pkg/e2/v1beta1"
)

type Manager struct {
	e2client   e2client.Client
	RnibClient rnib.Client
}

type SubManager interface {
	Start() error
}

var log = logging.GetLogger()

const metServiceModelOID = "1.3.6.1.4.1.53148.1.2.2.97"

// NewManager creates a new subscription manager
func NewManager() (Manager, error) {

	serviceModelName := e2client.ServiceModelName("e2sm_met")
	serviceModelVersion := e2client.ServiceModelVersion("v1")
	appID := e2client.AppID("10")
	e2Client := e2client.NewClient(
		e2client.WithServiceModel(serviceModelName, serviceModelVersion),
		e2client.WithAppID(appID),
		e2client.WithE2TAddress("onos-e2t", 5150))

	rnibClient, err := rnib.NewClient()
	if err != nil {
		return Manager{}, err
	}

	return Manager{
		e2client:   e2Client,
		RnibClient: rnibClient,
	}, nil

}

func (m *Manager) watchE2Connections(ctx context.Context) error {
	ch := make(chan topoapi.Event)
	err := m.RnibClient.WatchE2Connections(ctx, ch)
	if err != nil {
		log.Warn(err)
		return err
	}

	// creates a new subscription whenever there is a new E2 node connected and supports KPM service model
	for topoEvent := range ch {
		log.Debugf("Received topo event: %v", topoEvent)

		if topoEvent.Type == topoapi.EventType_ADDED || topoEvent.Type == topoapi.EventType_NONE {
			relation := topoEvent.Object.Obj.(*topoapi.Object_Relation)
			e2NodeID := relation.Relation.TgtEntityID
			if !m.RnibClient.HasMETRanFunction(ctx, e2NodeID, metServiceModelOID) {
				continue
			}

			log.Info("we are supposed to subscribe here to the node ", e2NodeID)

			// go func() {
			// 	err := m.newSubscription(ctx, e2NodeID)
			// 	if err != nil {
			// 		log.Warn(err)
			// 	}
			// }()

		}

	}
	return nil
}

// Start starts subscription manager
func (m *Manager) Start() error {
	ctx, _ := context.WithCancel(context.Background())
	err := m.watchE2Connections(ctx)
	if err != nil {
		return err
	}

	return nil
}

// e2node := client.Node(e2client.NodeID(e2nodeID))
//    subName := "met-sm-subscription" // A unique and constant subscription name
//    var eventTriggerData []byte     // Encode the service model specific event trigger
//    var actionDefinitionData []byte // Encode the service model specific Action Definitions
//    var actions []e2api.Action
//    action := e2api.Action{
//     ID:   100,
//     Type: e2api.ActionType_ACTION_TYPE_REPORT,
//     SubsequentAction: &e2api.SubsequentAction{
//         Type:       e2api.SubsequentActionType_SUBSEQUENT_ACTION_TYPE_CONTINUE,
//         TimeToWait: e2api.TimeToWait_TIME_TO_WAIT_ZERO,
//    },
//     Payload: actionDefinitionData,
//    }

//    subSpec := e2api.SubscriptionSpec{
//         Actions: actions,
//         EventTrigger: e2api.EventTrigger{
//              Payload: eventTriggerData,
//      },
// }

// ch := make(chan e2api.Indication)

// channelID, err := e2node.Subscribe(context.TODO(), subName, subSpec, ch)
// if err != nil {
//    return err
// }

// for ind := range ch {
//     fmt.Println(ind)
// }
