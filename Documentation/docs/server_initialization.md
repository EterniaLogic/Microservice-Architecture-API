# Server Initialization

## First Step: Connect to the server.

#### https://www.digitalocean.com/community/tutorials/how-to-connect-to-your-droplet-with-ssh

## Second step: download all required software.

#### Server admin utils
```Bash
apt-get install aptitude screen htop iptraf sudo ufw nano -y
```

#### Mysql stuff
```Bash
aptitude install mysql-server mysql-client -y
```

#### Go stuff
```Bash
aptitude install golang-go golang-doc -y
```

#### Git
```Bash
aptitude install git git-extras -y
```

#### Create database (Note, tables are automatically created by the server)

Tables are automatically created by the server executable.
```Bash
mysql --user="root" --password="PASSWORD" --execute="CREATE DATABASE `go_serv` /*!40100 DEFAULT CHARACTER SET latin1 */;"
```

#### Create user
```Bash
adduser goserv -d /home/goserv
mkdir /home/goserv
chown goserv /home/goserv
```


#### Firewall
```Bash
ufw allow in on eth1 to any port 6100:6107 proto tcp
ufw enable
```

#### Downloading code
```Bash
git clone https://github.com/EterniaLogic/Microservice-Architecture-API.git microservices
cd ~/microservices
bash buildall.sh
```


#### Service Creation

nano /etc/systemd/system/goserv.service
```Bash
[Unit]
Description=Server
After=syslog.target network.target

[Service]
User=goserv

[Service]
WorkingDirectory=/home/goserv
ExecStart=/home/goserv/server
ExecReload=/bin/kill -HUP $MAINPID
KillMode=process
Restart=always
RestartSec=1s
PrivateTmp=true

[Install]
WantedBy=multi-user.target
```

Starting service
```Bash
sudo systemctl dameon-reload
sudo systemctl start goserv.service
```