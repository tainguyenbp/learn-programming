#!/bin/bash

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
logPath="$DIR/log_kill_tcpdump"

log_message() {
        echo -e $1 >> $logPath
}

kill_tcpdump() {

        killall tcpdump
}
remove_log_tcpdump() {

        for i in `seq 2 10`; do
                year=$(date -d "$i days ago" +"%Y")
                month=$(date -d "$i days ago" +"%-m")
                day=$(date -d "$i days ago" +"%-d")
                ymd=$year$month$day
                rm -rf "$DIR"/"$ymd"
        done
}
start_tcpdump() {
        "$DIR"/check_tcpdump.sh
}

log_message "\n\n\n\n\n----------- Checking System Info -----------"
log_message "Checking Time: $(date)"
kill_tcpdump
remove_log_tcpdump
start_tcpdump
