package gigasecond

// import path for the time package from the standard library
import "time"

// This returns the giga second added to given time
func AddGigasecond(t time.Time) time.Time {

	   next_date = t.Add(time.Hour * 1000000000 + time.Minute * 0+ time.Second * 0)
	return next_date
	}

