#!/bin/bash

screen -dmS servers htop

sleep 0.2
screen -S servers -X screen 0
sleep 0.2
screen -S servers -p 0 -X exec sh -c "clear;echo Admin; cd Servers/AdminServer; ./server"

sleep 0.2
screen -S servers -X screen 1
sleep 0.2
screen -S servers -p 1 -X exec sh -c "clear;echo Auth; cd Servers/AuthServer; ./server"

sleep 0.2
screen -S servers -X screen 2
sleep 0.2
screen -S servers -p 2 -X exec sh -c "clear;echo Comment; cd Servers/CommentServer; ./server"

sleep 0.2
screen -S servers -X screen 3
sleep 0.2
screen -S servers -p 3 -X exec sh -c "clear;echo Feedback; cd Servers/FeedbackServer; ./server"

sleep 0.2
screen -S servers -X screen 4
sleep 0.2
screen -S servers -p 4 -X exec sh -c "clear;echo Feeds; cd Servers/FeedsServer; ./server"

sleep 0.2
screen -S servers -X screen 5
sleep 0.2
screen -S servers -p 5 -X exec sh -c "clear;echo Profile; cd Servers/ProfileServer; ./server"

sleep 0.2
screen -S servers -X screen 6
sleep 0.2
screen -S servers -p 6 -X exec sh -c "clear;echo Search; cd Servers/SearchServer; ./server"

sleep 0.2
screen -S servers -X screen 7
sleep 0.2
screen -S servers -p 7 -X exec sh -c "clear;echo Video; cd Servers/VideoServer; ./server"