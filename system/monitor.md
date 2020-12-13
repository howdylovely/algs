<!--
 * @Descripttion: 
 * @version: 
 * @Author: WangShuaibing
 * @Date: 2020-11-17 21:20:52
 * @LastEditors: WangShuaibing
 * @LastEditTime: 2020-11-30 10:01:10
-->
# 监控
14 

13

8
9

1qaz@WSX





10 0 * * * /data0/allinone/scripts/auto-del-7-days-ago-log.sh >/dev/null 2>&1


cat >> /etc/crontab <<EOF
* * * * * docker exec openfalcon sh ctrl.sh start graph hbs judge transfer nodata aggregator agent gateway api alarm
* * * * * docker exec php php /var/www/html/jd-api/yii light/light > /dev/null 2>&1
* * * * * docker exec php php /var/www/html/jd-api/yii fan/fancontrol > /dev/null 2>&1
EOF




./docker-compose.yml
./openresty/default.conf
./openresty/nginx/logs/jd-api-access_log
./openresty/www/config/dev_db.php
./php/applogs/jd_api_gateway.warning.log.20201111
./php/applogs/jd_api_gateway.error.log.20201111
./php/applogs/jd_api_gateway.info.log.20201111
./php/applogs/jd_api_gateway.notice.log.20201111
./php/applogs/jd_api_gateway.warning.log.20201112
./php/applogs/jd_api_gateway.notice.log.20201112
./php/applogs/jd_api_gateway.info.log.20201112
./vue/endpoint.json


rsync -avlK ./rack_release.tar.gz root@12.1.1.9:/data0/rack_release.tar.gz
grep -Ilr 192.168.15.44 ./ | xargs -n1 -- sed -i 's/192.168.56.6/12.1.1.9/g'

rack_release.tar.gz

chmod +x /etc/rc.d/rc.local

cat >> /etc/rc.d/rc.local <<EOF
/data0/erroralarm/control start > /dev/null 2>&1
/data0/jdsystem-agent-client/open-falcon start agent > /dev/null 2>&1
EOF


### ipmi
戴尔
联想
浪潮
Redfish



### 系统监控工具
free内存监控

vmstat

iostat

uptime

W

mpstat

pmap

sar


### metric

> 监控指标解释
https://satori-monitoring.readthedocs.io/zh/latest/builtin-metrics/net.html