package rod

import (
	"errors"

	"github.com/line/line-bot-sdk-go/linebot"
)

// ErrNoHandleFunc is error
var ErrNoHandleFunc = errors.New("no handler")

// TextMessageHandleFunc is a function to handle a linebot.TextMessage.
// If It requires reply messages it returns them.
func (h *Handler) TextMessageHandleFunc(f func(*linebot.Event, *linebot.TextMessage) []linebot.Message) {
	h.textMessageHandlerFunc = f
}

// ImageMessageHandleFunc is a function to handle a linebot.ImageMessage.
// If It requires reply messages it returns them.
func (h *Handler) ImageMessageHandleFunc(f func(*linebot.Event, *linebot.ImageMessage) []linebot.Message) {
	h.imageMessageHandlerFunc = f
}

// VideoMessageHandleFunc is a function to handle a linebot.VideoMessage.
// If It requires reply messages it returns them.
func (h *Handler) VideoMessageHandleFunc(f func(*linebot.Event, *linebot.VideoMessage) []linebot.Message) {
	h.videoMessageHandlerFunc = f
}

// AudioMessageHandleFunc is a function to handle a linebot.AudioMessage.
// If It requires reply messages it returns them.
func (h *Handler) AudioMessageHandleFunc(f func(*linebot.Event, *linebot.AudioMessage) []linebot.Message) {
	h.audioMessageHandlerFunc = f
}

// LocationMessageHandleFunc is a function to handle a linebot.LocationMessage.
// If It requires reply messages it returns them.
func (h *Handler) LocationMessageHandleFunc(f func(*linebot.Event, *linebot.LocationMessage) []linebot.Message) {
	h.locationMessageHandlerFunc = f
}

// StickerMessageHandleFunc is a function to handle a linebot.StickerMessage.
// If It requires reply messages it returns them.
func (h *Handler) StickerMessageHandleFunc(f func(*linebot.Event, *linebot.StickerMessage) []linebot.Message) {
	h.stickerMessageHandlerFunc = f
}

// FollowEventHandleFunc is a function to handle follow event.
// Second argument is userId.
func (h *Handler) FollowEventHandleFunc(f func(*linebot.Event, string) []linebot.Message) {
	h.followEventHandlerFunc = f
}

// UnfollowEventHandleFunc is a function to handle unfollow event.
// Second argument is userId.
func (h *Handler) UnfollowEventHandleFunc(f func(*linebot.Event, string) []linebot.Message) {
	h.unfollowEventHandlerFunc = f
}

// JoinGroupEventHandleFunc is a function to handle event when join the group.
// Second argument is the groupId.
func (h *Handler) JoinGroupEventHandleFunc(f func(*linebot.Event, string) []linebot.Message) {
	h.joinGroupEventHandlerFunc = f
}

// LeaveGroupEventHandleFunc is a function to handle event when leave the group.
// Second argument is the groupId.
func (h *Handler) LeaveGroupEventHandleFunc(f func(*linebot.Event, string) []linebot.Message) {
	h.leaveGroupEventHandlerFunc = f
}

// JoinTalkRoomEventHandleFunc is a function to handle event when join the talk room.
// Second argument is the roomId.
func (h *Handler) JoinTalkRoomEventHandleFunc(f func(*linebot.Event, string) []linebot.Message) {
	h.joinTalkRoomEventHandlerFunc = f
}

// LeaveTalkRoomEventHandleFunc is a function to handle event when leave the talk room.
// Second argument is the roomId.
func (h *Handler) LeaveTalkRoomEventHandleFunc(f func(*linebot.Event, string) []linebot.Message) {
	h.leaveTalkRoomEventHandlerFunc = f
}

// PostbackEventHandleFunc is a function to handle postback event.
// Second argument is the postback.data.
func (h *Handler) PostbackEventHandleFunc(f func(*linebot.Event, string) []linebot.Message) {
	h.postbackEventHandlerFunc = f
}

// BeaconEnterEventHandleFunc is a function to handle event when enter the range of a beacon.
// Second argument is the beacon.hwid.
func (h *Handler) BeaconEnterEventHandleFunc(f func(*linebot.Event, string) []linebot.Message) {
	h.beaconEnterEventHandlerFunc = f
}

// BeaconLeaveEventHandleFunc is a function to handle event when leave the range of a beacon.
// Second argument is the beacon.hwid.
func (h *Handler) BeaconLeaveEventHandleFunc(f func(*linebot.Event, string) []linebot.Message) {
	h.beaconLeaveEventHandlerFunc = f
}

// BeaconBannerEventHandleFunc is a function to handle event when user tapped beacon banner.
// Second argument is the beacon.hwid.
func (h *Handler) BeaconBannerEventHandleFunc(f func(*linebot.Event, string) []linebot.Message) {
	h.beaconBannerEventHandlerFunc = f
}

// VerifyMessageHandleFunc is a function to handle verify event.
func (h *Handler) VerifyMessageHandleFunc(f func(*linebot.Event)) {
	h.verifyMessageHandlerFunc = f
}
