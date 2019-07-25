package relay

import (
	"github.com/chris-pikul/go-wormhole-server/config"
	"github.com/chris-pikul/go-wormhole/msg"
)

//Service encompases the actual relay service for use by
//clients. Essentially most of this package handles
//the actual network connections and message handling.
//This object is the actual implementation (or at least
//the start of it)
type Service struct {
	Welcome msg.WelcomeInfo
}

//NewService initializes the relay service object
//and returns it as a pointer, if we can not start
//one then nil, error is returned instead
func NewService() (*Service, error) {
	srv := &Service{}

	//Setup the welcome message stuff
	if config.Opts.Relay.WelcomeMOTD != "" {
		//Pointers should bind this so we can change it later?
		srv.Welcome.MOTD = &config.Opts.Relay.WelcomeMOTD
	}

	if config.Opts.Relay.WelcomeError != "" {
		//Pointers should bind this so we can change it later?
		srv.Welcome.Error = &config.Opts.Relay.WelcomeError
	}

	if config.Opts.Relay.AdvertisedVersion != "" {

		//Pointers should bind this so we can change it later?
		srv.Welcome.Version = &config.Opts.Relay.AdvertisedVersion
	}

	return srv, nil
}
