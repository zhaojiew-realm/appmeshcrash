#!/bin/bash
#

POD_NAME=`kubectl get pods -o wide -n appmeshload | grep front | awk -F" " '{print $1}'`
FRONT_IP=`kubectl get pods -o wide -n appmeshload | grep front | awk -F" " '{print $6}'`

echo -e "Pod Name:\t $POD_NAME"
echo -e "Pod IP:\t   $FRONT_IP"

curl -X GET -s -H "X-Request-Id: 66666666" -H "X-B3-Traceid: 123123213" http://$FRONT_IP:8090/forward;

# while true
# do 
#     curl -X GET -s http://$FRONT_IP:8090/forward;
#     sleep 0.5s;
# done

