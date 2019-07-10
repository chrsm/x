package main

import (
	"log"
	"os"
	"time"

	"github.com/bwmarrin/discordgo"
)

const (
	email    = "thisisnt@myrealemail.org"
	password = "hunter2"
)

var myID = ""

func main() {
	disc, err := discordgo.New(email, password)

	if err != nil {
		log.Fatalf("couldn't auth discord: %s", err)
	}

	me, err := disc.User("@me")
	if err != nil {
		log.Fatalf("failed to get @me: %s", err)
	}
	myID = me.ID

	myg, err := disc.UserGuilds(0, "", "")
	if err != nil {
		log.Fatalf("failed to get guilds: %s", err)
	}

	for i := range myg {
		g := myg[i]

		gch, err := disc.GuildChannels(g.ID)
		if err != nil {
			log.Fatalf("failed to get guild channels for %s: %s", g.ID, err)
		}

		for j := range gch {
			ch := gch[j]

			if ch.Type != discordgo.ChannelTypeGuildText {
				continue
			}

			log.Printf("getting message IDs in %s(%s)", ch.Name, ch.ID)
			ids := getMIDs(disc, ch.ID)
			for k := range ids {
				if err := disc.ChannelMessageDelete(ch.ID, ids[k]); err != nil {
					log.Printf("failed to delete %s in %s: %s", ids[k], ch.ID, err)
				}
			}
		}
	}

	os.Exit(1)

	myc, err := disc.UserChannels()
	if err != nil {
		log.Fatalf("couldn't get userchannels: %s", err)
	}

	log.Printf("got dms, len: %d", len(myc))
	for i := range myc {
		dm := myc[i]

		ids := getMIDs(disc, dm.ID)

		for j := range ids {
			if err := disc.ChannelMessageDelete(dm.ID, ids[j]); err != nil {
				log.Printf("couldn't delete msg %s in %s: %s", ids[j], dm.ID, err)
			}
		}
	}
}

func getMIDs(disc *discordgo.Session, dmID string) []string {
	var ids []string

	msgs, err := disc.ChannelMessages(dmID, 100, "", "", "")
	if err != nil {
		log.Printf("failed to get messages in c(%s): %s", dmID, err)
		return nil
	}

	if len(msgs) == 0 {
		return nil
	}

	var last = msgs[len(msgs)-1].ID

	s := time.Now()
	for {
		for i := range msgs {
			if msgs[i].Author.ID != myID {
				continue
			}

			ids = append(ids, msgs[i].ID)
		}

		msgs, err = disc.ChannelMessages(dmID, 100, last, "", "")
		if err != nil {
			log.Fatalf("failed to get messages in ch(%s): %s", dmID, err)
		}
		if len(msgs) == 0 || msgs[len(msgs)-1].ID == last || time.Now().Sub(s).Seconds() > 60 {
			break
		}

		last = msgs[len(msgs)-1].ID
		time.Sleep(time.Millisecond * 500)
	}

	return ids
}
