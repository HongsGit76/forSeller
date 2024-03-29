package server

import (
    "bytes"
    "net"
)

// func main() {
//     fmt.Printf("MAC: %16.16X\n", macUint64())
// }

func MacUint64() uint64 {
    interfaces, err := net.Interfaces()
    if err != nil {
        return uint64(0)
    }

    for _, i := range interfaces {
        if i.Flags&net.FlagUp != 0 && bytes.Compare(i.HardwareAddr, nil) != 0 {

            // Skip locally administered addresses
            if i.HardwareAddr[0]&2 == 2 {
                continue
            }

            var mac uint64
            for j, b := range i.HardwareAddr {
                if j >= 8 {
                    break
                }
                mac <<= 8
                mac += uint64(b)
            }

            return mac
        }
    }

    return uint64(0)
}