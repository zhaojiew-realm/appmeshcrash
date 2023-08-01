## File list:
- destinationrules-all-with-outlierdetection.yaml # 开启所有服务的outlierdetection
- destinationrules-default.yaml # 设置destionationrule， 启用Version控制。
- direct2frontapplication.sh # 测试工具。直接发送请求到front容器， 不使用Ingressgateway， **必须**设置front deployment replicas： 1 
- egressgateway-to-center.yaml # 使用egressgateway 约束指定域名的出向流量走 egressgateway
- ingressgateway-tofrontservie.yaml # 使用 ingressgateway 管理入向流量， 使用域名SNI区分流量是否可以正常的进入网格。
- kubernetes-application-default.yaml # kubernetes 资源。
- kubernetes-application_egressgateway.yaml # kubernetes 资源。
- sendto_ingressgateway.sh # 测试工具。直接发送请求到Ingressgateway。 replicas： 1 。
- vs-all-shiftingtov1.yaml # 默认virtualservice， 所有流量指向v1版本。
- vs-all-shiftingtov2.yaml # 默认所有指向v2。
- vs-all-splitv1v2-5050.yaml # 默认所有 v1 v2 各50%。
- vs-color-v1-faultinjection-abort.yaml # abort 5xx error.
- vs-color-v1-faultinjection-fixdelay.yaml # color 服务FaultInjection.
- vs-fruit-v1-faultinjection-fixdelay.yaml # fruit 服务FaultInjection.

-# Example:

### 最小

默认没有ingressgateway， 没有egressgateway， 所有版本指向v1 

```bash 
kubectl apply -f kubernetes-application-default.yaml
kubectl apply -f destinationrules-default.yaml
kubectl apply -f vs-all-shiftingtov1.yaml

kubectl scale -n appmeshload deployment/deploy-front --replicas 1
direct2frontapplication.sh 
```

### 启用Ingressgateway

开启ingressgateway， 没有egressgateway, 所有版本指向v1.

```bash
kubectl apply -f kubernetes-application-default.yaml
kubectl apply -f destinationrules-default.yaml
kubectl apply -f vs-all-shiftingtov1.yaml

kubectl scale -n istio-system deployment/istio-ingressgateway --replicas 1
sendto_ingressgateway.sh
```

### 启用outlier detection

对所有envoy配置严格的 outlierdetection 策略： 

```bash
kubectl apply -f destinationrules-all-with-outlierdetection.yaml 

sendto_ingressgateway.sh
```

### 启用egressgateway

使用egressgateway 管理流量， 前提是istio安装的时候已经开启了这个gateway。

```bash
kubectl apply -f kubernetes-application_egressgateway.yaml 
# 指向两个不同的出向目的， 一个是center， 这个是私有的地址， 可以随便改。 另一个是baidu， baidu没有配置serviceentry ，所以会显示成 passthrouth cluster的透传流量。 如果istio严格控制网格的流量不许未定义的流量出去， 那么baidu会直接访问不到。
kubectl apply -f egressgateway-to-center.yaml

sendto_ingressgateway.sh
```

### 启用故障注入

注入abort 故障， 自定义响应错误代码， 故障比率。

```bash
kubectl apply -f vs-all-splitv1v2-5050.yaml # 初始化vs， 将流量分发到v1v2, 取消故障可以直接apply这个文件。  
kubectl apply -f vs-color-v1-faultinjection-abort.yaml
# 注入 abort 故障
kubectl apply -f vs-fruit-v1-faultinjection-fixdelay.yaml
# 注入 延迟故障， 请求 delay 一定的时间。 单位； s
```



