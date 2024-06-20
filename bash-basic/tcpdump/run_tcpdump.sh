#!/bin/bash

CURRENT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

year=$(date +"%Y")
month=$(date +"%-m")
day=$(date +"%-d")

hour=$(date +"%H")
munites=$(date +"%M")
second=$(date +"%S")

name=$(hostname)

hms=$hour$munites$second
ymd=$year$month$day

mkdir -p "$CURRENT_DIR"/"$ymd"

DIRECTORY="$CURRENT_DIR"/"$ymd"

        if [ -d "$DIRECTORY" ];
                then
                        echo "Folder $DIRECTORY exists"
                                sleep 2
                                tcpdump -i eth1 -n -s 0 port 5060 -vvv -w "$CURRENT_DIR"/"$ymd"/log-eth1-"$name"-"$ymd"-"$hms".pcap &
                                tcpdump -i eth0 -n -s 0 port 5060 -vvv -w "$CURRENT_DIR"/"$ymd"/log-eth0-"$name"-"$ymd"-"$hms".pcap &

        elif [ ! -d "$DIRECTORY" ];
                then
                        echo "Folder $DIRECTORY doesnt exists"
                              mkdir -p "$CURRENT_DIR"/"$ymd"
                        echo "Creating Foleder $DIRECTORY Done !!!"
                                sleep 2
                                tcpdump -i eth1 -n -s 0 port 5060 -vvv -w "$CURRENT_DIR"/"$ymd"/log-eth1-"$name"-"$ymd"-"$hms".pcap &
                                tcpdump -i eth0 -n -s 0 port 5060 -vvv -w "$CURRENT_DIR"/"$ymd"/log-eth0-"$name"-"$ymd"-"$hms".pcap &
        fi
