#### for testing only
1. wget -q -O - https://raw.githubusercontent.com/k3d-io/k3d/main/install.sh | bash
1. k3d cluster create -p "80:80@loadbalancer" -p "443:443@loadbalancer"
2. kubectl apply -f tools/swagger/swagger.yml
3. kubectl create token new-admin-sa
4. replace token in tools/swagger/swagger.yml
5. kubectl apply -f tools/swagger/swagger.yml
6. echo "127.0.0.1 kubernetes" >> /etc/hosts
7. echo "127.0.0.1 swagger-ui" >> /etc/hosts
8. open https://kubernetes and accept self-signed SSL - IMPORTANT
9. open https://swagger-ui