package wrapper

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
	OnTextMessageRecieved     TextMessageHandlerFunc
	OnImageMessageRecieved    ImageMessageHandlerFunc
	OnVideoMessageRecieved    VideoMessageHandlerFunc
	OnAudioMessageRecieved    AudioMessageHandlerFunc
	OnLocationMessageRecieved LocationMessageHandlerFunc
	OnStickerMessageRecieved  StickerMessageHandlerFunc

	// Follow/Unfollow Event
	OnFollowed   FollowEventHandlerFunc
	OnUnfollowed UnfollowEventHandlerFunc

	// Join/Leave Event
	OnJoinGroup     JoinGroupEventHandlerFunc
	OnLeaveGroup    LeaveGroupEventHandlerFunc
	OnJoinTalkRoom  JoinTalkRoomEventHandlerFunc
	OnLeaveTalkRoom LeaveTalkRoomEventHandlerFunc

	// Postback Event
	OnPostback PostbackEventHandlerFunc

	// Beacon Event
	OnBeaconEnter        BeaconEnterEventHandlerFunc
	OnBeaconLeave        BeaconLeaveEventHandlerFunc
	OnBeaconBannerTapped BeaconBannerEventHandlerFunc

	OnError func(http.ResponseWriter, *http.Request, error)
}

// New returns a Handler using *linebot.Client
func New(client *linebot.Client) (*Handler, error) {
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
			replyMessages = handler.OnTextMessageRecieved(event, message)
		case *linebot.ImageMessage:
			replyMessages = handler.OnImageMessageRecieved(event, message)
		case *linebot.VideoMessage:
			replyMessages = handler.OnVideoMessageRecieved(event, message)
		case *linebot.AudioMessage:
			replyMessages = handler.OnAudioMessageRecieved(event, message)
		case *linebot.LocationMessage:
			replyMessages = handler.OnLocationMessageRecieved(event, message)
		case *linebot.StickerMessage:
			replyMessages = handler.OnStickerMessageRecieved(event, message)
		}

	case linebot.EventTypeFollow:
		replyMessages = handler.OnFollowed(event, event.Source.UserID)

	case linebot.EventTypeUnfollow:
		replyMessages = handler.OnUnfollowed(event, event.Source.UserID)

	case linebot.EventTypeJoin:
		switch event.Source.Type {
		case linebot.EventSourceTypeGroup:
			replyMessages = handler.OnJoinGroup(event, event.Source.GroupID)
		case linebot.EventSourceTypeRoom:
			replyMessages = handler.OnJoinTalkRoom(event, event.Source.RoomID)
		}

	case linebot.EventTypeLeave:
		switch event.Source.Type {
		case linebot.EventSourceTypeGroup:
			replyMessages = handler.OnLeaveGroup(event, event.Source.GroupID)
		case linebot.EventSourceTypeRoom:
			replyMessages = handler.OnLeaveTalkRoom(event, event.Source.RoomID)
		}

	case linebot.EventTypePostback:
		replyMessages = handler.OnPostback(event, event.Postback.Data)

	case linebot.EventTypeBeacon:
		switch event.Beacon.Type {
		case linebot.BeaconEventTypeEnter:
			replyMessages = handler.OnBeaconEnter(event, event.Beacon.Hwid)
		case linebot.BeaconEventTypeLeave:
			replyMessages = handler.OnBeaconLeave(event, event.Beacon.Hwid)
		case linebot.BeaconEventTypeBanner:
			replyMessages = handler.OnBeaconBannerTapped(event, event.Beacon.Hwid)
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
