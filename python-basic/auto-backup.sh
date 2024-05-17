#!/bin/sh
# Delete cookies.txt file if exists and start fresh.
if [ -f cookies.txt ]; then
rm cookies.txt
fi<
# Get cookie values.
/usr/bin/curl -k -b cookies.txt -c cookies.txt --data 'login=Login&usernamefld=backup&passwordfld=password' http://192.168.0.1:8080/diag_backup.php
# Download the configuration.
/usr/bin/curl -k -b cookies.txt -o /home/jon/pfsense-backups/config-router-`date +%Y%m%d%H%M%S`.xml --data 'Submit=download&donotbackuprrd=no' http://192.168.0.1:8080/diag_backup.php
