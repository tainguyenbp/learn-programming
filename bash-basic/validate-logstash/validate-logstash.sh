```
#!/bin/bash

NAME_LOGSTASH='logstash.conf'

PATH_LOGSTASH='/etc/logstash/conf.d'

echo "The check command verifies the syntax of Logstash configuration files:"
mustache check "${PATH_LOGSTASH}/${NAME_LOGSTASH}"

echo "The lint command checks for problems in Logstash configuration files: "
mustache lint --auto-fix-id "${PATH_LOGSTASH}/${NAME_LOGSTASH}"

echo "The format command, mustache returns the provided configuration files in a standardized format: "
#mustache format --write-to-source "${PATH_LOGSTASH}/${NAME_LOGSTASH}"
```

```
/usr/share/logstash/bin/logstash --config.test_and_exit -f /etc/logstash/conf.d/logstash.conf
```
