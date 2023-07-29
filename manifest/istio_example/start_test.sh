#!/bin/bash
#

POD_NAME=`kubectl get pods -o wide -n appmeshload | grep front | awk -F" " '{print $1}'`
FRONT_IP=`kubectl get pods -o wide -n appmeshload | grep front | awk -F" " '{print $6}'`

echo -e "Pod Name:\t $POD_NAME"
echo -e "Pod IP:\t   $FRONT_IP"

sleep 2s;

while true
do 
    curl -X GET -s http://$FRONT_IP:8090/forward;
    sleep 0.5s;
done

