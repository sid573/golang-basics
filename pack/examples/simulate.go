package pack

import (
	"fmt"
	"math/rand"
)

// Bot for creating each bot
type Bot struct {
	ID        int
	Happiness bool
}

// Region for creating different region
type Region struct {
	ID        int
	Available float64
	BotCount  int
}

// SDSController to be used as reciever for functions implemented
type SDSController struct {
	Bots    []*Bot
	Regions []*Region
}

// SDSEnv to create environment for simulation
func SDSEnv(botCount int, regionCount int) *SDSController {
	bots := make([]*Bot, 0)
	regions := make([]*Region, 0)
	for i := 0; i < botCount; i++ {
		bot := &Bot{
			ID: i,
		}
		r := rand.New(rand.NewSource(10))
		available := r.Float64()
		region := &Region{
			ID:        i,
			Available: available,
		}
		bots = append(bots, bot)
		regions = append(regions, region)
	}
	controller := &SDSController{
		Bots:    bots,
		Regions: regions,
	}
	return controller
}

// DisplayEnv will show all details of bots
// and region of the environment
func (env *SDSController) DisplayEnv() {
	fmt.Println("Details of all bots in the environment")
	for _, bot := range env.Bots {
		fmt.Println(*bot)
	}
	fmt.Println("Details of all regions in the environment")
	for _, region := range env.Regions {
		fmt.Println(*region)
	}
}
