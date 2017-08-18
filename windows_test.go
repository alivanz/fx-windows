package windows_test

import (
  "testing"
  "."
  "fmt"
)

func Test(t *testing.T){
  PNP := windows.PNPDeviceID("PCI\\VEN_1002&DEV_666F&SUBSYS_380C17AA&REV_00\\4&3420519A&0&00E4")
  fmt.Println(PNP.DeviceID())
}
