#!/bin/bash

CURRENT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

year_global=$(date +"%Y")
month_global=$(date +"%m")
day_global=$(date +"%d")
ymd_global=$year_global$month_global$day_global

name_file_log="file-large-size-"$ymd_global".log"

path_file_log=""$CURRENT_DIR"/"$name_file_log""
path_find_file='/'
file_size='+10M'
head='-500'
string_remove_grep='grep -v \"/var/lib/mysql|/srv/minio|postgresql\"'

function _remove_log_find_file() {

        for i in `seq 10 20`; do
                year_ago=$(date -d "$i days ago" +"%Y")
                month_ago=$(date -d "$i days ago" +"%m")
                day_ago=$(date -d "$i days ago" +"%d")
                ymd_ago=$year_ago$month_ago$day_ago
                rm -rf "$path_file_log"/_"$ymd_ago"
        done
}

function _report_find_file_large() {

        find $path_find_file -xdev -type f -size $file_size -exec du -sh {} ';'| sort -rh | head $head | $string_remove_grep >> $path_file_log &

}

_report_find_file_large
