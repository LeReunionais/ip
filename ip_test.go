package ip

import (
  "testing"
  "fmt"
)

func TestIps(t *testing.T) {
  ips := Ips()
  for _, ip := range ips {
    fmt.Println(ip)
  }
}
