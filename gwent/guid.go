package gwent

type GUID int64

var guidChannel chan GUID = make(chan GUID, 1)

func init() {
	go func() {
		guid := GUID(0)
		for {
			guid++
			guidChannel <- guid
		}
	}()
}

func GetNextGUID() GUID {
	return <-guidChannel
}
