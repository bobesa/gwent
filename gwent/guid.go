package gwent

// GUID defines size of the guid
type GUID uint64

var guidChannel = make(chan GUID, 1)

func init() {
	go func() {
		guid := GUID(0)
		for {
			guid++
			guidChannel <- guid
		}
	}()
}

// GetNextGUID reports new guid that is generated by our worker from init() call
func GetNextGUID() GUID {
	return <-guidChannel
}
