module axiom

go 1.12

//require "botapi" v0.0.0
replace botapi => ./api

require (
	botapi v0.0.0-00010101000000-000000000000
	github.com/bwmarrin/discordgo v0.19.0
)
