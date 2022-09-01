## GO WEB SERVER
# Go Web Sunucusunu Docker Üzerinden Yayınlamak için:
* sudo docker build -t go_web_server .
	
* sudo docker run -p 8080:8080 go_web_server

# Kubernets Deployment:
* kubectl apply -f deployment.yaml

* kubectl get deployments(Oluşan deploymentı görüntelemek için)
