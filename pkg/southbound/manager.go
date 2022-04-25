package southbound

import (
	"context"

	"github.com/AbdouTlili/met-xapp/pkg/rnib"
	"github.com/AbdouTlili/onos-e2-sm/servicemodels/e2sm_met/pdubuilder"
	e2api "github.com/onosproject/onos-api/go/onos/e2t/e2/v1beta1"
	topoapi "github.com/onosproject/onos-api/go/onos/topo"
	"github.com/onosproject/onos-lib-go/pkg/logging"
	e2client "github.com/onosproject/onos-ric-sdk-go/pkg/e2/v1beta1"
	"google.golang.org/protobuf/proto"
)

type Manager struct {
	e2client   e2client.Client
	appID      string
	RnibClient rnib.Client
}

type SubManager interface {
	Start() error
}

var log = logging.GetLogger()

const metServiceModelOID = "1.3.6.1.4.1.53148.1.2.2.98"

// NewManager creates a new subscription manager
func NewManager() (Manager, error) {

	serviceModelName := e2client.ServiceModelName("e2sm_met")
	serviceModelVersion := e2client.ServiceModelVersion("v1")
	appID := e2client.AppID("onos-met")
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

			go func() {
				err := m.createSubscription(ctx, e2NodeID)
				if err != nil {
					log.Warn(err)
				}
			}()

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

func (m *Manager) createSubscription(ctx context.Context, e2nodeID topoapi.ID) error {
	log.Info("Creating subscription for E2 node ID with: ", e2nodeID)
	eventTriggerData, err := m.createEventTriggerData(64)
	if err != nil {
		log.Warn(err)
		return err
	}

	// aspects, err := m.rnibClient.GetE2NodeAspects(ctx, e2nodeID)
	// if err != nil {
	// 	log.Warn(err)
	// 	return err
	// }

	// _, err = m.getRanFunction(aspects.ServiceModels)
	// if err != nil {
	// 	log.Warn(err)
	// 	return err
	// }

	ch := make(chan e2api.Indication)
	node := m.e2client.Node(e2client.NodeID(e2nodeID))
	subName := "met-sub"
	subSpec := e2api.SubscriptionSpec{
		EventTrigger: e2api.EventTrigger{
			Payload: eventTriggerData,
		},
		Actions: m.createSubscriptionActions(),
	}
	log.Debugf("subSpec: %v", subSpec)

	_, err = node.Subscribe(ctx, subName, subSpec, ch)
	if err != nil {
		log.Warn(err)
		return err
	}

	// _, err = m.streams.OpenReader(ctx, node, subName, channelID, subSpec)
	// if err != nil {
	// 	log.Warn(err)
	// 	return err
	// }

	return nil
}

func (m *Manager) createSubscriptionActions() []e2api.Action {
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

// func (m *Manager) getRanFunction(serviceModelsInfo map[string]*topoapi.ServiceModelInfo) (*topoapi.RSMRanFunction, error) {
// 	for _, sm := range serviceModelsInfo {
// 		smName := strings.ToLower(sm.Name)
// 		if smName == string(m.serviceModel.Name) && sm.OID == oid {
// 			rsmRanFunction := &topoapi.RSMRanFunction{}
// 			for _, ranFunction := range sm.RanFunctions {
// 				if ranFunction.TypeUrl == ranFunction.GetTypeUrl() {
// 					err := prototypes.UnmarshalAny(ranFunction, rsmRanFunction)
// 					if err != nil {
// 						return nil, err
// 					}
// 					return rsmRanFunction, nil
// 				}
// 			}
// 		}
// 	}
// 	return nil, errors.New(errors.NotFound, "cannot retrieve ran functions")

// }

func (m *Manager) createEventTriggerData(rtPeriod int64) ([]byte, error) {
	e2SmMetEventTriggerDefinition, err := pdubuilder.CreateE2SmMetEventTriggerDefinition(rtPeriod)
	if err != nil {
		return []byte{}, err
	}

	err = e2SmMetEventTriggerDefinition.Validate()
	if err != nil {
		return []byte{}, err
	}

	protoBytes, err := proto.Marshal(e2SmMetEventTriggerDefinition)
	if err != nil {
		return []byte{}, err
	}

	return protoBytes, nil
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
