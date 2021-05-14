# icmpping from zabbix-agent

## An example of compile in clean installation of Ubuntu server 20.04
sudo apt install golang libz-dev libpcre3-dev -y  
wget https://raw.githubusercontent.com/klimenk0aa/zabbix_agent2-ping/master/icmpping.patch  
wget https://cdn.zabbix.com/zabbix/sources/stable/5.0/zabbix-5.0.11.tar.gz  
tar -xzf zabbix-5.0.11.tar.gz  
cd zabbix-5.0.11  
patch -s -p1 <../icmpping.patch  
./configure --enable-agent2 --enable-static  
make  
sudo setcap cap_net_raw=+ep src/go/bin/zabbix_agent2

### test
./src/go/bin/zabbix_agent2 -t icmpping[8.8.8.8]

### result
icmpping[8.8.8.8]                             [s|1]
