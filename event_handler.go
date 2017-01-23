package wrapper

import (
	"net/http"

	"github.com/line/line-bot-sdk-go/linebot"
)

// Handler is a Handler
type Handler struct {
	handlerFuncMap map[linebot.EventType]func(*linebot.Event)

	// Message Event
	textMessageHandlerFunc     func(*linebot.Event, *linebot.TextMessage)
	imageMessageHandlerFunc    func(*linebot.Event, *linebot.ImageMessage)
	videoMessageHandlerFunc    func(*linebot.Event, *linebot.VideoMessage)
	audioMessageHandlerFunc    func(*linebot.Event, *linebot.AudioMessage)
	locationMessageHandlerFunc func(*linebot.Event, *linebot.LocationMessage)
	stickerMessageHandlerFunc  func(*linebot.Event, *linebot.StickerMessage)

	// Beacon Event
	beaconEnterHandlerFunc func(*linebot.Event)
	beaconLeaveHandlerFunc func(*linebot.Event)
}

// HandleEvents handles events
func (handler *Handler) HandleEvents(events []*linebot.Event, r *http.Request) {
	for _, event := range events {

		f, ok := handler.handlerFuncMap[event.Type]
		if ok {
			f(event)
			continue
		}

		switch event.Type {
		case linebot.EventTypeMessage:

			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				if handler.textMessageHandlerFunc != nil {
					handler.textMessageHandlerFunc(event, message)
				}
			case *linebot.ImageMessage:
				if handler.imageMessageHandlerFunc != nil {
					handler.imageMessageHandlerFunc(event, message)
				}
			case *linebot.VideoMessage:
				if handler.videoMessageHandlerFunc != nil {
					handler.videoMessageHandlerFunc(event, message)
				}
			case *linebot.AudioMessage:
				if handler.audioMessageHandlerFunc != nil {
					handler.audioMessageHandlerFunc(event, message)
				}
			case *linebot.LocationMessage:
				if handler.locationMessageHandlerFunc != nil {
					handler.locationMessageHandlerFunc(event, message)
				}
			case *linebot.StickerMessage:
				if handler.stickerMessageHandlerFunc != nil {
					handler.stickerMessageHandlerFunc(event, message)
				}
			}
		case linebot.EventTypeBeacon:
			if event.Beacon == nil {
				break
			}
			switch event.Beacon.Type {
			case linebot.BeaconEventTypeEnter:
				if handler.beaconEnterHandlerFunc != nil {
					handler.beaconEnterHandlerFunc(event)
				}
			case linebot.BeaconEventTypeLeave:
				if handler.beaconLeaveHandlerFunc != nil {
					handler.beaconLeaveHandlerFunc(event)
				}
			}
		}
	}
}
