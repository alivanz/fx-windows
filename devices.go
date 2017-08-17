package windows

import (
  "golang.org/x/sys/windows/registry"
  "strconv"
  "github.com/StackExchange/wmi"
  "strings"
  "log"
)

type VideoController struct{
  Name string
  AdapterRAM uint32
  PNPDeviceID PNPDeviceID
}

type PNPDeviceID string
type LocationInformation struct{
  Bus int
  Device int
  Function int
}

/* Sort */
type VideoControllerByBusID []VideoController
func (l VideoControllerByBusID) Len() int { return len(l) }
func (l VideoControllerByBusID) Swap(i, j int) { l[i], l[j] = l[j], l[i] }
func (l VideoControllerByBusID) Less(i, j int) bool {
  return l[i].PNPDeviceID.LocationInformation().Bus < l[j].PNPDeviceID.LocationInformation().Bus
}

func (PNPDeviceID PNPDeviceID) LocationInformation() LocationInformation {
  k, err := registry.OpenKey(registry.LOCAL_MACHINE, `SYSTEM\CurrentControlSet\Enum\`+string(PNPDeviceID), registry.QUERY_VALUE)
  if err != nil {
    log.Panic(err)
  }
  defer k.Close()
  s, _, err := k.GetStringValue("locationinformation")
  if err != nil {
  	log.Panic(err)
  }
  locinfo := strings.Split( s[strings.Index(s, "(")+1:len(s)-1], "," )
  bus,_ := strconv.Atoi(locinfo[0])
  dev,_ := strconv.Atoi(locinfo[1])
  fun,_ := strconv.Atoi(locinfo[2])
  return LocationInformation{ bus,dev,fun }
}

func ListVideoController() []VideoController {
  var dst []VideoController
  if err := wmi.Query("SELECT Name,AdapterRAM,PNPDeviceID FROM Win32_VideoController", &dst); err!= nil{
		log.Panic(err)
	}
  return dst
}
