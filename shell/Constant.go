package shell

var Reboot = New("reboot")
var NetworkStatus = New("networkctl status eth0")
var NetworkRestart = New("systemctl restart systemd-networkd.service")
var Log = New("journalctl ")
var IFCONFIG = New("ifconfig")
var OS_RELEASE = New("cat /etc/os-release")
var MemAvailable = New("grep MemAvailable /proc/meminfo")
var MemTotal = New("grep MemTotal /proc/meminfo")
var SYSTEMTIME = New("date -d \"$(awk -F. '{print $1}' /proc/uptime) second ago\" +\"%Y-%m-%d %H:%M:%S\" ")
var SYSCPU = New("top -n 1 | grep CPU: |grep idle |awk '{print $8}'")
var SYSROM = New("df -h |grep /userdata |awk '{print $2,$3}'")
