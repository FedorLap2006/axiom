package main

import (
	api "botapi"
	dsgo "github.com/bwmarrin/discordgo"
	//"log"
	strs "strings"
)

var botutils_CMDS map[string]api.EventHandler
var botutils_PREFIX string
var botutils_SUFFIX string

func botutils_CMDSHandler(e *api.Event) {
	for k, v := range botutils_CMDS {
		// log.Println(botutils_PREFIX + e.Data.(*dsgo.Message).Content + botutils_SUFFIX)
		// log.Println(strs.HasPrefix(e.Data.(*dsgo.Message).Content, botutils_PREFIX))
		// log.Println(strs.HasSuffix(e.Data.(*dsgo.Message).Content, botutils_SUFFIX))
		if strs.HasPrefix(e.Data.(*dsgo.Message).Content, botutils_PREFIX) && strs.HasSuffix(e.Data.(*dsgo.Message).Content, botutils_SUFFIX) && botutils_PREFIX+k+botutils_SUFFIX == e.Data.(*dsgo.Message).Content {
			v(e)
		}
	}
}
