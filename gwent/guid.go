package gwent

type GUID int64

var guidChannel chan GUID

func init() {
	guidChannel = make(chan GUID)
	go func(){
		guid := GUID(0)
		for {
			guid++
			guidChannel <- guid
		}
	}()
}

func GetNextGUID() GUID {
	return <- guidChannel
}