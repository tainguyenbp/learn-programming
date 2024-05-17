#!/bin/bash
export d=`date +%d-%m-%Y-%H-%M-%S`

# Copy it!
/usr/bin/scp -q micro@trogdor.muppetz.com:/cf/conf/config.xml /external/tim/syncconfig/Sync/Trogdor/trogdor-backup-$d.xml

# Compress it!
/usr/bin/xz -9e -T0 /external/tim/syncconfig/Sync/Trogdor/trogdor-backup-$d.xml

# Email it!
/usr/bin/mutt -e "set record=/dev/null" -s "Daily Trogdor Backup" remote@email.address.com -a /external/tim/syncconfig/Sync/Trogdor/trogdor-backup-$d.xml.xz < /dev/null
