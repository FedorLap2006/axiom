package botapi

import (
	"log"
	//"fmt"
	dsgo "github.com/bwmarrin/discordgo"
)

const (
	_EVENT_CLASS_ZERO = iota
	EVENT_CLASS_GUILD
	EVENT_CLASS_MESSAGE
	_EVENT_TYPE_ZERO = iota
	EVENT_TYPE_MESSAGE_CREATE
	EVENT_TYPE_MESSAGE_DELETE
	EVENT_TYPE_MESSAGE_EDIT
	EVENT_TYPE_MESSAGE_BULKDEL

	EVENT_TYPE_GUILD_BAN_ADD
	EVENT_TYPE_GUILD_BAN_DEL
	EVENT_TYPE_GUILD_MEMBER_ADD
	EVENT_TYPE_GUILD_MEMBER_DEL
	EVENT_TYPE_GUILD_MEMBER_UPD

	EVENT_TYPE_GUILD_ROLE_CREATE
	EVENT_TYPE_GUILD_ROLE_DEL
	EVENT_TYPE_GUILD_ROLE_UPD

	EVENT_TYPE_GUILD_CHAN_CREATE
	EVENT_TYPE_GUILD_CHAN_DEL
	EVENT_TYPE_GUILD_CHAN_UPD
)

const (
	EVENTH_MESSAGE_CREATE  = EVENT_TYPE_MESSAGE_CREATE
	EVENTH_MESSAGE_DELETE  = EVENT_TYPE_MESSAGE_DELETE
	EVENTH_MESSAGE_EDIT    = EVENT_TYPE_MESSAGE_EDIT
	EVENTH_MESSAGE_BULKDEL = EVENT_TYPE_MESSAGE_BULKDEL
	EVENTH_GUILD_BAN_ADD   = EVENT_TYPE_GUILD_BAN_ADD
	EVENTH_GUILD_BAN_DEL   = EVENT_TYPE_GUILD_BAN_DEL
	EVENTH_TYPE_MEMBER_ADD = EVENT_TYPE_GUILD_MEMBER_ADD
	EVENTH_TYPE_MEMBER_DEL = EVENT_TYPE_GUILD_MEMBER_DEL
	EVENTH_TYPE_MEMBER_UPD = EVENT_TYPE_GUILD_MEMBER_UPD

	EVENTH_GUILD_ROLE_CREATE = EVENT_TYPE_GUILD_ROLE_CREATE
	EVENTH_GUILD_ROLE_DEL    = EVENT_TYPE_GUILD_ROLE_DEL
	EVENTH_GUILD_ROLE_UPD    = EVENT_TYPE_GUILD_ROLE_UPD

	EVENTH_GUILD_CHAN_CREATE = EVENT_TYPE_GUILD_CHAN_CREATE
	EVENTH_GUILD_CHAN_DEL    = EVENT_TYPE_GUILD_CHAN_DEL
	EVENTH_GUILD_CHAN_UPD    = EVENT_TYPE_GUILD_CHAN_UPD
)

type Event struct {
	Class   int
	Type    int
	Data    interface{}
	Session *dsgo.Session
}

type EventHandler func(*Event)

type HandlersList map[int]EventHandler

type Bot struct {
	Ds       *dsgo.Session
	Botu     *dsgo.User
	Handlers HandlersList
}

func (self *Bot) Init(token string) {
	self.Ds, _ = dsgo.New("Bot " + token)

	botu, err := self.Ds.User("@me")
	if err != nil {
		log.Fatal(err)
	}
	self.Botu = botu
	self.Handlers = make(HandlersList)
	self.Ds.AddHandler(self.BaseHandler)
}

func (self *Bot) Run() error {
	err := self.Ds.Open()
	return err
}

func (self *Bot) Stop() error {
	err := self.Ds.Close()
	return err
}

func (self *Bot) BaseHandler(s *dsgo.Session, event interface{}) {
	ec := Event{}
	switch event.(type) {
	case *dsgo.MessageCreate:
		ec.Class = EVENT_CLASS_MESSAGE
		ec.Type = EVENT_TYPE_MESSAGE_CREATE
		ec.Data = event.(*dsgo.MessageCreate).Message
		self.Handlers[ec.Type](&ec)
	case *dsgo.MessageDelete:
		ec.Class = EVENT_CLASS_MESSAGE
		ec.Type = EVENT_TYPE_MESSAGE_CREATE
		ec.Data = event.(*dsgo.MessageDelete).Message
		self.Handlers[ec.Type](&ec)
	case *dsgo.MessageUpdate:
		ec.Class = EVENT_CLASS_MESSAGE
		ec.Type = EVENT_TYPE_MESSAGE_EDIT
		ec.Data = event.(*dsgo.MessageUpdate).Message
		self.Handlers[ec.Type](&ec)
	case *dsgo.GuildBanAdd:
		ec.Class = EVENT_CLASS_GUILD
		ec.Type = EVENT_TYPE_GUILD_BAN_ADD
		ec.Data = event.(*dsgo.GuildBanAdd)
		self.Handlers[ec.Type](&ec)
	case *dsgo.GuildBanRemove:
		ec.Class = EVENT_CLASS_GUILD
		ec.Type = EVENT_TYPE_GUILD_BAN_DEL
		ec.Data = event.(*dsgo.GuildBanRemove)
		self.Handlers[ec.Type](&ec)
	case *dsgo.GuildMemberAdd:
		ec.Class = EVENT_CLASS_GUILD
		ec.Type = EVENT_TYPE_GUILD_MEMBER_ADD
		ec.Data = event.(*dsgo.GuildMemberAdd)
		self.Handlers[ec.Type](&ec)
	case *dsgo.GuildMemberRemove:
		ec.Class = EVENT_CLASS_GUILD
		ec.Type = EVENT_TYPE_GUILD_MEMBER_DEL
		ec.Data = event.(*dsgo.GuildMemberRemove)
		self.Handlers[ec.Type](&ec)
	case *dsgo.GuildMemberUpdate:
		ec.Class = EVENT_CLASS_GUILD
		ec.Type = EVENT_TYPE_GUILD_MEMBER_UPD
		ec.Data = event.(*dsgo.GuildMemberUpdate)
		self.Handlers[ec.Type](&ec)
	case *dsgo.GuildRoleCreate:
		ec.Class = EVENT_CLASS_GUILD
		ec.Type = EVENT_TYPE_GUILD_ROLE_CREATE
		ec.Data = event.(*dsgo.GuildRoleCreate)
		self.Handlers[ec.Type](&ec)
	case *dsgo.GuildRoleDelete:
		ec.Class = EVENT_CLASS_GUILD
		ec.Type = EVENT_TYPE_GUILD_ROLE_DEL
		ec.Data = event.(*dsgo.GuildRoleDelete)
		self.Handlers[ec.Type](&ec)
	case *dsgo.GuildRoleUpdate:
		ec.Class = EVENT_CLASS_GUILD
		ec.Type = EVENT_TYPE_GUILD_ROLE_UPD
		ec.Data = event.(*dsgo.GuildRoleUpdate)
		self.Handlers[ec.Type](&ec)
	case *dsgo.ChannelCreate:
		ec.Class = EVENT_CLASS_GUILD
		ec.Type = EVENT_TYPE_GUILD_CHAN_CREATE
		ec.Data = event.(*dsgo.ChannelCreate).Channel
		self.Handlers[ec.Type](&ec)
	case *dsgo.ChannelDelete:
		ec.Class = EVENT_CLASS_GUILD
		ec.Type = EVENT_TYPE_GUILD_CHAN_DEL
		ec.Data = event.(*dsgo.ChannelDelete).Channel
		self.Handlers[ec.Type](&ec)
	case *dsgo.ChannelUpdate:
		ec.Class = EVENT_CLASS_GUILD
		ec.Type = EVENT_TYPE_GUILD_CHAN_UPD
		ec.Data = event.(*dsgo.ChannelUpdate).Channel
		self.Handlers[ec.Type](&ec)

	}
}

// bot api methods
func (self *Bot) SendMsg(chid string, msg string, tss bool) {
	if tss {
		self.Ds.ChannelMessageSend(chid, msg)
	} else {
		self.Ds.ChannelMessageSendTTS(chid, msg)
	}
}
