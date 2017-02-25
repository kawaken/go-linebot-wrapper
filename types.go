package rod

import "github.com/line/line-bot-sdk-go/linebot"

// TextMessageHandlerFunc is a function to handle a linebot.TextMessage.
// If It requires reply messages it returns them.
type TextMessageHandlerFunc func(*linebot.Event, *linebot.TextMessage) []linebot.Message

// ImageMessageHandlerFunc is a function to handle a linebot.ImageMessage.
// If It requires reply messages it returns them.
type ImageMessageHandlerFunc func(*linebot.Event, *linebot.ImageMessage) []linebot.Message

// VideoMessageHandlerFunc is a function to handle a linebot.VideoMessage.
// If It requires reply messages it returns them.
type VideoMessageHandlerFunc func(*linebot.Event, *linebot.VideoMessage) []linebot.Message

// AudioMessageHandlerFunc is a function to handle a linebot.AudioMessage.
// If It requires reply messages it returns them.
type AudioMessageHandlerFunc func(*linebot.Event, *linebot.AudioMessage) []linebot.Message

// LocationMessageHandlerFunc is a function to handle a linebot.LocationMessage.
// If It requires reply messages it returns them.
type LocationMessageHandlerFunc func(*linebot.Event, *linebot.LocationMessage) []linebot.Message

// StickerMessageHandlerFunc is a function to handle a linebot.StickerMessage.
// If It requires reply messages it returns them.
type StickerMessageHandlerFunc func(*linebot.Event, *linebot.StickerMessage) []linebot.Message

// FollowEventHandlerFunc is a function to handle follow event.
// Second argument is userId.
type FollowEventHandlerFunc func(*linebot.Event, string) []linebot.Message

// UnfollowEventHandlerFunc is a function to handle unfollow event.
// Second argument is userId.
type UnfollowEventHandlerFunc func(*linebot.Event, string) []linebot.Message

// JoinGroupEventHandlerFunc is a function to handle event when join the group.
// Second argument is the groupId.
type JoinGroupEventHandlerFunc func(*linebot.Event, string) []linebot.Message

// LeaveGroupEventHandlerFunc is a function to handle event when leave the group.
// Second argument is the groupId.
type LeaveGroupEventHandlerFunc func(*linebot.Event, string) []linebot.Message

// JoinTalkRoomEventHandlerFunc is a function to handle event when join the talk room.
// Second argument is the roomId.
type JoinTalkRoomEventHandlerFunc func(*linebot.Event, string) []linebot.Message

// LeaveTalkRoomEventHandlerFunc is a function to handle event when leave the talk room.
// Second argument is the roomId.
type LeaveTalkRoomEventHandlerFunc func(*linebot.Event, string) []linebot.Message

// PostbackEventHandlerFunc is a function to handle postback event.
// Second argument is the postback.data.
type PostbackEventHandlerFunc func(*linebot.Event, string) []linebot.Message

// BeaconEnterEventHandlerFunc is a function to handle event when enter the range of a beacon.
// Second argument is the beacon.hwid.
type BeaconEnterEventHandlerFunc func(*linebot.Event, string) []linebot.Message

// BeaconLeaveEventHandlerFunc is a function to handle event when leave the range of a beacon.
// Second argument is the beacon.hwid.
type BeaconLeaveEventHandlerFunc func(*linebot.Event, string) []linebot.Message

// BeaconBannerEventHandlerFunc is a function to handle event when user tapped beacon banner.
// Second argument is the beacon.hwid.
type BeaconBannerEventHandlerFunc func(*linebot.Event, string) []linebot.Message
