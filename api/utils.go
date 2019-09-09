package botapi

import (
	dsgo "github.com/bwmarrin/discordgo"
	//"log"
	strs "strings"
)

var Utils_CMDS map[string]EventHandler
var Utils_PREFIX string
var Utils_SUFFIX string

func Utils_CMDSHandler(e *Event) {
	for k, v := range Utils_CMDS {
		// log.Println(Utils_PREFIX + e.Data.(*dsgo.Message).Content + Utils_SUFFIX)
		// log.Println(strs.HasPrefix(e.Data.(*dsgo.Message).Content, Utils_PREFIX))
		// log.Println(strs.HasSuffix(e.Data.(*dsgo.Message).Content, Utils_SUFFIX))
		if strs.HasPrefix(e.Data.(*dsgo.Message).Content, Utils_PREFIX) && strs.HasSuffix(e.Data.(*dsgo.Message).Content, Utils_SUFFIX) && Utils_PREFIX+k+Utils_SUFFIX == e.Data.(*dsgo.Message).Content {
			v(e)
		}
	}
}
