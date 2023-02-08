package mqttHandler

type Places int

const (
	PlaceHall Places = iota
	PlaceTrees
	PlaceTV
	PlaceCron // Virtual place used for a cron trigger app
)

func (s Places) String() string {
	places := [...]string{"hall", "trees", "tv", "cron"}
	if s < PlaceHall || s > PlaceCron {
		panic("Wrong place given")
	}

	return places[s-1]
}
