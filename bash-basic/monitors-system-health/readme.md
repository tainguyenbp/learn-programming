# System monitoring Bash script
Write a bash script that monitors system health and sends an alert if any of the following conditions are met:

The CPU usage exceeds 80% for more than 5 minutes.
The available disk space on the root partition is less than 10%.
If any condition is met, the script should output an appropriate message to the standard error (stderr) indicating the issue.
You can use the following template to create your script:
```
check_cpu() {
    local threshold=80
    #... other variables

    while true; do
        # Check CPU usage
        # Your CPU usage check logic here
    done
}

check_disk_space() {
    local threshold=10
    #... other variables

    # Check disk space
    # Your disk space check logic here
}

check_disk_space
check_cpu  

```

# Solution
Below is an example of a bash script that monitors CPU usage and disk space, sending an alert if the conditions are met.

Your solution may vary and be written differently, however if you are not familiar with bash scripting, we recommend you to learn the basics of bash scripting before attempting this question.
```
function _check_cpu() {
    local threshold=80
    local count5min=0
    local countTotal=0
    local iterations=3 # 3x10 sec
    local max_count=$(bc <<< 300/10) # 5 minutes / 10 seconds interval

    while true; do
        cpu_usage=$(top -bn2 | grep "Cpu(s)" | sed -n '2p' | awk '{print $2}' | cut -d'.' -f1)
        ((countTotal++))
        if [ "$cpu_usage" -gt "$threshold" ]; then
            ((count++))
            if [ "$count" -ge "$max_count" ]; then
                echo "CPU usage has been above $threshold% for more than 5 minutes" >&2
                break # Exit the loop after alerting once
            fi
        else
            count5min=0 # Reset counter because CPU usage dropped below threshold
        fi
        if [ "$countTotal" -eq "$iterations" ]; then
            break # Exit the loop after 5 minutes
        fi
        echo "CPU usage is $cpu_usage% - Threshold: $threshold% .. waiting.."
        sleep 10 # Check every 10 seconds
    done
}

function _check_disk_space() {
    local threshold=10 # Set threshold to 10%
    local disk_usage=$(df / | grep / | awk '{ print $5 }' | sed 's/%//g')

    if [ "$disk_usage" -gt $((100 - threshold)) ]; then
        echo "Disk space on root partition is less than $threshold% available." &>2
    else
        echo "Disk space on root partition is sufficient."
    fi
}    

_check_disk_space
_check_cpu

```
