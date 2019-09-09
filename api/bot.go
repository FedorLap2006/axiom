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
)

const (
	EVENTH_MESSAGE_CREATE  = EVENT_TYPE_MESSAGE_CREATE
	EVENTH_MESSAGE_DELETE  = EVENT_TYPE_MESSAGE_DELETE
	EVENTH_MESSAGE_EDIT    = EVENT_TYPE_MESSAGE_EDIT
	EVENTH_MESSAGE_BULKDEL = EVENT_TYPE_MESSAGE_BULKDEL
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
	}
}
