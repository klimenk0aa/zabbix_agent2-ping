package a_ping
import "zabbix.com/pkg/plugin"
import "github.com/sparrc/go-ping"

import "errors" 
import "time"

type Plugin struct {
        plugin.Base
}
var impl Plugin


func (p *Plugin) Export(key string, params []string, ctx plugin.ContextProvider) (result interface{}, err error){
        if len(params) != 1 {
                return nil, errors.New("Wrong parameters.")
        }

        pinger, err := ping.NewPinger(params[0])
        if err != nil {
                panic(err)
        }
        pinger.SetPrivileged(true)
        pinger.Count = 4
        pinger.Timeout = 1 * time.Second
        pinger.Run()
        stats := pinger.Statistics()
        res := stats.PacketLoss

        return (100-res)/100, nil
}


func init() {
        plugin.RegisterMetrics(&impl, "a_ping", "a_ping", "Ping address.")
}

