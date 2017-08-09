package rod

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/line/line-bot-sdk-go/linebot"
)

// Handler is a Handler
type Handler struct {
	client *linebot.Client

	// Message Event
	textMessageHandlerFunc     TextMessageHandlerFunc
	imageMessageHandlerFunc    ImageMessageHandlerFunc
	videoMessageHandlerFunc    VideoMessageHandlerFunc
	audioMessageHandlerFunc    AudioMessageHandlerFunc
	locationMessageHandlerFunc LocationMessageHandlerFunc
	stickerMessageHandlerFunc  StickerMessageHandlerFunc

	// Follow/Unfollow Event
	followEventHandlerFunc   FollowEventHandlerFunc
	unfollowEventHandlerFunc UnfollowEventHandlerFunc

	// Join/Leave Event
	joinGroupEventHandlerFunc     JoinGroupEventHandlerFunc
	leaveGroupEventHandlerFunc    LeaveGroupEventHandlerFunc
	joinTalkRoomEventHandlerFunc  JoinTalkRoomEventHandlerFunc
	leaveTalkRoomEventHandlerFunc LeaveTalkRoomEventHandlerFunc

	// Postback Event
	postbackEventHandlerFunc PostbackEventHandlerFunc

	// Beacon Event
	beaconEnterEventHandlerFunc  BeaconEnterEventHandlerFunc
	beaconLeaveEventHandlerFunc  BeaconLeaveEventHandlerFunc
	beaconBannerEventHandlerFunc BeaconBannerEventHandlerFunc

	OnError func(http.ResponseWriter, *http.Request, error)
}

// New returns a Handler
func New(channelSecret string, channelToken string) (*Handler, error) {
	client, err := linebot.New(channelSecret, channelToken)
	if err != nil {
		return nil, err
	}

	return NewWithClient(client)
}

// NewWithClient returns a Handler using *linebot.Client
func NewWithClient(client *linebot.Client) (*Handler, error) {
	if client == nil {
		return nil, fmt.Errorf("linebot.Client is nil")
	}

	return &Handler{
		client: client,
	}, nil
}

func (handler *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	events, err := handler.client.ParseRequest(r)
	if err == linebot.ErrInvalidSignature {
		w.WriteHeader(400)
		return
	} else if err != nil {
		w.WriteHeader(500)
		return
	}

	handler.HandleEvents(events)
}

// HandleEvents handles events
func (handler *Handler) HandleEvents(events []*linebot.Event) {
	var wg sync.WaitGroup

	for _, event := range events {
		// check verification
		if event.ReplyToken == "00000000000000000000000000000000" {
			log.Print("verification")
			return
		}
		wg.Add(1)
		go func(event *linebot.Event) {
			if err := handler.handleEvent(event); err != nil {
				log.Print(err)
			}
			wg.Done()
		}(event)
	}
	wg.Wait()
}

func (handler *Handler) handleEvent(event *linebot.Event) error {

	var replyMessages []linebot.Message

	switch event.Type {
	case linebot.EventTypeMessage:

		switch message := event.Message.(type) {
		case *linebot.TextMessage:
			if handler.textMessageHandlerFunc == nil {
				return ErrNoHandleFunc
			}
			replyMessages = handler.textMessageHandlerFunc(event, message)

		case *linebot.ImageMessage:
			if handler.imageMessageHandlerFunc == nil {
				return ErrNoHandleFunc
			}
			replyMessages = handler.imageMessageHandlerFunc(event, message)

		case *linebot.VideoMessage:
			if handler.videoMessageHandlerFunc == nil {
				return ErrNoHandleFunc
			}
			replyMessages = handler.videoMessageHandlerFunc(event, message)

		case *linebot.AudioMessage:
			if handler.audioMessageHandlerFunc == nil {
				return ErrNoHandleFunc
			}
			replyMessages = handler.audioMessageHandlerFunc(event, message)

		case *linebot.LocationMessage:
			if handler.locationMessageHandlerFunc == nil {
				return ErrNoHandleFunc
			}
			replyMessages = handler.locationMessageHandlerFunc(event, message)

		case *linebot.StickerMessage:
			if handler.stickerMessageHandlerFunc == nil {
				return ErrNoHandleFunc
			}
			replyMessages = handler.stickerMessageHandlerFunc(event, message)

		}

	case linebot.EventTypeFollow:
		if handler.followEventHandlerFunc == nil {
			return ErrNoHandleFunc
		}
		replyMessages = handler.followEventHandlerFunc(event, event.Source.UserID)

	case linebot.EventTypeUnfollow:
		if handler.unfollowEventHandlerFunc == nil {
			return ErrNoHandleFunc
		}
		replyMessages = handler.unfollowEventHandlerFunc(event, event.Source.UserID)

	case linebot.EventTypeJoin:
		switch event.Source.Type {
		case linebot.EventSourceTypeGroup:
			if handler.joinGroupEventHandlerFunc == nil {
				return ErrNoHandleFunc
			}
			replyMessages = handler.joinGroupEventHandlerFunc(event, event.Source.GroupID)

		case linebot.EventSourceTypeRoom:
			if handler.joinTalkRoomEventHandlerFunc == nil {
				return ErrNoHandleFunc
			}
			replyMessages = handler.joinTalkRoomEventHandlerFunc(event, event.Source.RoomID)

		}

	case linebot.EventTypeLeave:
		switch event.Source.Type {
		case linebot.EventSourceTypeGroup:
			if handler.leaveGroupEventHandlerFunc == nil {
				return ErrNoHandleFunc
			}
			replyMessages = handler.leaveGroupEventHandlerFunc(event, event.Source.GroupID)

		case linebot.EventSourceTypeRoom:
			if handler.leaveTalkRoomEventHandlerFunc == nil {
				return ErrNoHandleFunc
			}
			replyMessages = handler.leaveTalkRoomEventHandlerFunc(event, event.Source.RoomID)

		}

	case linebot.EventTypePostback:
		if handler.postbackEventHandlerFunc == nil {
			return ErrNoHandleFunc
		}
		replyMessages = handler.postbackEventHandlerFunc(event, event.Postback.Data)

	case linebot.EventTypeBeacon:
		switch event.Beacon.Type {
		case linebot.BeaconEventTypeEnter:
			if handler.beaconEnterEventHandlerFunc == nil {
				return ErrNoHandleFunc
			}
			replyMessages = handler.beaconEnterEventHandlerFunc(event, event.Beacon.Hwid)

		case linebot.BeaconEventTypeLeave:
			if handler.beaconLeaveEventHandlerFunc == nil {
				return ErrNoHandleFunc
			}
			replyMessages = handler.beaconLeaveEventHandlerFunc(event, event.Beacon.Hwid)

		case linebot.BeaconEventTypeBanner:
			if handler.beaconBannerEventHandlerFunc == nil {
				return ErrNoHandleFunc
			}
			replyMessages = handler.beaconBannerEventHandlerFunc(event, event.Beacon.Hwid)

		}
	}

	if len(replyMessages) > 0 {
		_, err := handler.client.ReplyMessage(event.ReplyToken, replyMessages...).Do()
		if err != nil {
			return err
		}
	}

	return nil
}
