#!/bin/bash

# root user, Generates a full server
# bash MicroserverGen.sh

########## Prequisites ##########
# server admin utils
apt-get install aptitude screen htop iptraf sudo ufw nano -y

# mysql stuff
aptitude install mysql-server mysql-client -y

# go stuff
aptitude install golang-go golang-doc -y

# git
aptitude install git git-extras -y

# Prompt user/pass for mysql
echo "Save these later for the config for the private user server"
muser="root"
echo "Mysql root password (From server install):"
read mpass
echo "Mysql New Database:"
read mdatabase
echo "Mysql New Table:"
read mtable


########## Set up users ##########
echo "Adding user auth"
useradd auth -d /home/auth

########## Mysql configuration ##########
echo "Adding Mysql generation"
mysql --user="$muser" --password="$mpassword" --execute="CREATE DATABASE $mdatabase /*!40100 DEFAULT CHARACTER SET latin1 */;"
mysql --user="$muser" --password="$mpassword" --execute="CREATE TABLE `$mtable` (`UUID` binary(16) NOT NULL,`username` varchar(50) NOT NULL,`password` blob NOT NULL,`oauth2_token` varchar(45) NOT NULL,`oauth2_website` varchar(45) NOT NULL,`oauth2_id` varchar(45) NOT NULL,`auth_token` blob NOT NULL,`auth_level` varchar(25) NOT NULL DEFAULT '',`auth_level_gen` blob NOT NULL,`join_date` datetime NOT NULL,`last_login` datetime NOT NULL DEFAULT '0000-00-00 00:00:00',PRIMARY KEY (`UUID`)) ENGINE=InnoDB DEFAULT CHARSET=latin1 COMMENT='Store direct user accounts here.\nThis includes accounts that use OAuth keys generated from google, facebook.\n- Brent Clancy (EterniaLogic)';"

########## DNS Configuration ##########
echo "TODO: Add DNS server settings in /etc/resolv.conf and /etc/hosts"

########## SSH security ##########
echo "TODO: Add SSH server settings"

########## VPN ##########
echo "TODO: Add VPN server settings"
# TODO: remove sleep
sleep 6

########## UFW Firewall ##########
ufw allow ssh
ufw allow 6100
# TODO: add admin RESTful port
ufw enable

########## Git Download ##########
git clone https://github.com/EterniaLogic/Microservice-Architecture-API.git microservices
chmod 0700 microservices -R

########## Build Auth ##########
cd microservices/AuthServer
go build

echo "Copying over AuthServer and config"
cp AuthServer /home/auth/AuthServer
cp conf.json /home/auth/conf.json

# set perms for auth
echo "Setting basic perms for AuthServer and config"
chmod 100 /home/auth/AuthServer
chmod 400 /home/auth/conf.json
chown auth /home/auth/AuthServer

echo "Please edit config file..."
sleep 1
nano /home/auth/conf.json