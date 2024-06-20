#!/bin/bash

CURRENT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

year_global=$(date +"%Y")
month_global=$(date +"%m")
day_global=$(date +"%d")
ymd_global=$year_global$month_global$day_global

name_file_log="log-inode-system-"$ymd_global".log"
path_file_log=""$CURRENT_DIR"/"$name_file_log""

echo "" > $path_file_log
path_find_inode='/home/*'

function _remove_log_find_file() {

        for i in `seq 10 20`; do
                year_ago=$(date -d "$i days ago" +"%Y")
                month_ago=$(date -d "$i days ago" +"%m")
                day_ago=$(date -d "$i days ago" +"%d")
                ymd_ago=$year_ago$month_ago$day_ago
                rm -rf "$path_file_log"/_"$ymd_ago"
        done
}

function _report_find_inode_system(){

        echo "Path Directory Find Inode System: $path_find_inode"
        echo "Path Directory Find Inode System: $path_find_inode" >> $path_file_log
        for directory_path in $path_find_inode;
                do result_find_inode=$(find $directory_path | wc -l);
                printf "$result_find_inode\t\t- $directory_path\n";
                printf "$result_find_inode\t\t- $directory_path\n" >> $path_file_log
        done
}

_report_find_inode_system
