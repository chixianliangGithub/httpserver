# httpserver
part1
kubectl apply -f webserver.yaml


part2

按照ingress.MD生成证书

kubectl apply -f nginx-ingress-deployment.yaml 
kubectl apply -f secret.yaml 
kubectl apply -f ingress.yaml
