#!/bin/bash

export GOPATH=/home/eternia/go

sudo systemctl stop goserv.service

cd common/main
go build

sudo cp main /home/goserv/server

sudo systemctl start goserv.service

cp main ../../Servers/AdminServer/server
cp main ../../Servers/AuthServer/server
cp main ../../Servers/CommentServer/server
cp main ../../Servers/FeedsServer/server
cp main ../../Servers/FeedbackServer/server
cp main ../../Servers/ProfileServer/server
cp main ../../Servers/SearchServer/server
cp main ../../Servers/VideoServer/server
