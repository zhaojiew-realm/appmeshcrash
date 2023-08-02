#!/bin/bash
#

POD_NAME=`kubectl get pods -o wide -n istio-system | grep istio-ingressgateway | awk -F" " '{print $1}'`
GW_IP=`kubectl get pods -o wide -n istio-system | grep istio-ingressgateway | awk -F" " '{print $6}'`
FQDN="front.liarlee.site"

echo -e "Pod Name:\t $POD_NAME"
echo -e "Pod IP:\t   $GW_IP"
echo -e "Doamin Name:\t   $FQDN"

sleep 1s;

curl -X GET http://"$FQDN":8090/forward;

# while true
# do 
#     curl -X GET http://"$FQDN":8090/forward;
#     sleep 0.1s;
# done

