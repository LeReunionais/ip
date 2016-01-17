package ip

import (
  "net"
)

// Ips will return not loopback ip slice of current host.
func Ips () ([]string) {
  addresses,_ := net.InterfaceAddrs()
  var ips []string
  for _, address := range addresses {
    ipnet, ok := address.(*net.IPNet)
    if ok && !ipnet.IP.IsLoopback() && ipnet.IP.To4() != nil {
      ips = append(ips, ipnet.IP.String())
    }
  }
  return ips
}
