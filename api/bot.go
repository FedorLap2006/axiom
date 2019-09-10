package botapi

import (
	"log"
	//"fmt"
	dsgo "github.com/bwmarrin/discordgo"
	"plugin"
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

	EVENT_TYPE_GUILD_CREATE
	EVENT_TYPE_GUILD_DEL
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
	EVENTH_GUILD_CREATE      = EVENT_TYPE_GUILD_CREATE
	EVENTH_GUILD_DEL         = EVENT_TYPE_GUILD_DEL
)

type Event struct {
	Class   int
	Type    int
	Data    interface{}
	Session *dsgo.Session
}

type EventHandler func(*Event)

type HandlersList map[int]EventHandler

type IBotPlugin interface {
	Invoke(*Bot) error
}

type BotPluginFile struct {
	Name         string `json:"name"`
	Filename     string `json:"filename"`
	FunctionName string `json:"funcname"`
}

type BotPluginFunc struct {
	Name     string `json:"name"`
	Function func(*Bot) error
}

func (self BotPluginFile) Invoke(bot *Bot) error {
	plug, err := plugin.Open(self.Filename)
	if err != nil {
		return err
	}
	inv, err := plug.Lookup(self.FunctionName)
	if err != nil {
		return err
	}

	return inv.(func(bot *Bot) error)(bot)
}

func (self BotPluginFunc) Invoke(bot *Bot) error {
	return self.Function(bot)
}

type Bot struct {
	Discord  *dsgo.Session
	Botu     *dsgo.User
	Handlers HandlersList
	Plugins  map[string]IBotPlugin
}

func (self *Bot) RegPluginFunc(name string, f func(*Bot) error) {
	self.Plugins[name] = BotPluginFunc{
		Name:     name,
		Function: f,
	}
}

func (self *Bot) RegPluginFile(name string, filename string, function string) {
	self.Plugins[name] = BotPluginFile{
		Name:         name,
		Filename:     filename,
		FunctionName: function,
	}
}

func (self *Bot) Init(token string) {
	self.Discord, _ = dsgo.New("Bot " + token)

	botu, err := self.Discord.User("@me")
	if err != nil {
		log.Fatal(err)
	}
	self.Botu = botu
	self.Handlers = make(HandlersList)
	self.Plugins = make(map[string]IBotPlugin)
	self.Discord.AddHandler(self.BaseHandler)
}

func (self *Bot) Run() error {
	for key, val := range self.Plugins {
		log.Printf("[Plugins: %s]: Trying Run...\n", key)
		err := val.Invoke(self)
		if err != nil {
			log.Printf("[Plugins: %s]: OOPS => Plugin Ended With Error ( %s )\n", key, err)
		} else {
			log.Printf("[Plugins: %s]: DONE => Plugin Is Running", key)
		}
	}
	err := self.Discord.Open()
	return err
}

func (self *Bot) Stop() error {
	err := self.Discord.Close()
	return err
}

func (self *Bot) BaseHandler(s *dsgo.Session, event interface{}) {
	ec := Event{}
	switch event.(type) {
	case *dsgo.MessageCreate:
		ec.Class = EVENT_CLASS_MESSAGE
		ec.Type = EVENT_TYPE_MESSAGE_CREATE
		ec.Data = event.(*dsgo.MessageCreate).Message
		if s, ok := self.Handlers[ec.Type]; ok {
			s(&ec)
		}
	case *dsgo.MessageDelete:
		ec.Class = EVENT_CLASS_MESSAGE
		ec.Type = EVENT_TYPE_MESSAGE_CREATE
		ec.Data = event.(*dsgo.MessageDelete).Message
		if s, ok := self.Handlers[ec.Type]; ok {
			s(&ec)
		}
	case *dsgo.MessageUpdate:
		ec.Class = EVENT_CLASS_MESSAGE
		ec.Type = EVENT_TYPE_MESSAGE_EDIT
		ec.Data = event.(*dsgo.MessageUpdate).Message
		if s, ok := self.Handlers[ec.Type]; ok {
			s(&ec)
		}
	case *dsgo.GuildBanAdd:
		ec.Class = EVENT_CLASS_GUILD
		ec.Type = EVENT_TYPE_GUILD_BAN_ADD
		ec.Data = event.(*dsgo.GuildBanAdd)
		if s, ok := self.Handlers[ec.Type]; ok {
			s(&ec)
		}
	case *dsgo.GuildBanRemove:
		ec.Class = EVENT_CLASS_GUILD
		ec.Type = EVENT_TYPE_GUILD_BAN_DEL
		ec.Data = event.(*dsgo.GuildBanRemove)
		if s, ok := self.Handlers[ec.Type]; ok {
			s(&ec)
		}
	case *dsgo.GuildMemberAdd:
		ec.Class = EVENT_CLASS_GUILD
		ec.Type = EVENT_TYPE_GUILD_MEMBER_ADD
		ec.Data = event.(*dsgo.GuildMemberAdd)
		if s, ok := self.Handlers[ec.Type]; ok {
			s(&ec)
		}
	case *dsgo.GuildMemberRemove:
		ec.Class = EVENT_CLASS_GUILD
		ec.Type = EVENT_TYPE_GUILD_MEMBER_DEL
		ec.Data = event.(*dsgo.GuildMemberRemove)
		if s, ok := self.Handlers[ec.Type]; ok {
			s(&ec)
		}
	case *dsgo.GuildMemberUpdate:
		ec.Class = EVENT_CLASS_GUILD
		ec.Type = EVENT_TYPE_GUILD_MEMBER_UPD
		ec.Data = event.(*dsgo.GuildMemberUpdate)
		if s, ok := self.Handlers[ec.Type]; ok {
			s(&ec)
		}
	case *dsgo.GuildRoleCreate:
		ec.Class = EVENT_CLASS_GUILD
		ec.Type = EVENT_TYPE_GUILD_ROLE_CREATE
		ec.Data = event.(*dsgo.GuildRoleCreate)
		if s, ok := self.Handlers[ec.Type]; ok {
			s(&ec)
		}
	case *dsgo.GuildRoleDelete:
		ec.Class = EVENT_CLASS_GUILD
		ec.Type = EVENT_TYPE_GUILD_ROLE_DEL
		ec.Data = event.(*dsgo.GuildRoleDelete)
		if s, ok := self.Handlers[ec.Type]; ok {
			s(&ec)
		}
	case *dsgo.GuildRoleUpdate:
		ec.Class = EVENT_CLASS_GUILD
		ec.Type = EVENT_TYPE_GUILD_ROLE_UPD
		ec.Data = event.(*dsgo.GuildRoleUpdate)
		if s, ok := self.Handlers[ec.Type]; ok {
			s(&ec)
		}
	case *dsgo.ChannelCreate:
		ec.Class = EVENT_CLASS_GUILD
		ec.Type = EVENT_TYPE_GUILD_CHAN_CREATE
		ec.Data = event.(*dsgo.ChannelCreate).Channel
		if s, ok := self.Handlers[ec.Type]; ok {
			s(&ec)
		}
	case *dsgo.ChannelDelete:
		ec.Class = EVENT_CLASS_GUILD
		ec.Type = EVENT_TYPE_GUILD_CHAN_DEL
		ec.Data = event.(*dsgo.ChannelDelete).Channel
		if s, ok := self.Handlers[ec.Type]; ok {
			s(&ec)
		}
	case *dsgo.ChannelUpdate:
		ec.Class = EVENT_CLASS_GUILD
		ec.Type = EVENT_TYPE_GUILD_CHAN_UPD
		ec.Data = event.(*dsgo.ChannelUpdate).Channel
		if s, ok := self.Handlers[ec.Type]; ok {
			s(&ec)
		}
	case *dsgo.GuildCreate:
		ec.Class = EVENT_CLASS_GUILD
		ec.Type = EVENT_TYPE_GUILD_CREATE
		ec.Data = event.(*dsgo.GuildCreate).Guild
		if s, ok := self.Handlers[ec.Type]; ok {
			s(&ec)
		}
	case *dsgo.GuildDelete:
		ec.Class = EVENT_CLASS_GUILD
		ec.Type = EVENT_TYPE_GUILD_DEL
		ec.Data = event.(*dsgo.GuildDelete).Guild
		if s, ok := self.Handlers[ec.Type]; ok {
			s(&ec)
		}

	}
}

// bot api methods
func (self *Bot) SendMsg(chid string, msg string, tss bool) {
	if tss {
		self.Discord.ChannelMessageSend(chid, msg)
	} else {
		self.Discord.ChannelMessageSendTTS(chid, msg)
	}
}
func (self *Bot) GetMessage(chid string, msgid string) (*dsgo.Message, error) {
	return self.Discord.ChannelMessage(chid, msgid)
}
func (self *Bot) DeleteMessage(chid string, msgid string) error {
	return self.Discord.ChannelMessageDelete(chid, msgid)
}
