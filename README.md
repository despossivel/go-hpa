## Go Lang em um container docker com CI/CD e Kubernetes (Gcloud Build) 

> Este repositório foi criado com o objetivo de praticar os conhecimentos adquiridos nos meus últimos estudos sobre DevOps, mas precisamente sobre CI/CD no Google cloud platform e Kubenetes.

![enter image description here](https://miro.medium.com/max/3284/1*dgWpWWIZPR_QASrzMRqF8w.png) 

**O que você vai encontrar aqui?**  Um servidor http, escrito em Go Lang, um contêiner docker para execução da aplicação, um simples teste unitário, uma esteira de CI & CD executado cinco passos no Cloud Build do GCP e os arquivos de configuração do K8s. 

**Passos da esteira de CI & CD**

 1. Rodando docker-compose
 2. Build image do contêiner
 3. Build application
 4. Rodando teste unitário 
 5. Push da imagem do contêiner para o registry


**K8s**

 1. Deployment e service do SERVER
 2. Deployment, service e configMap do NGINX 
 3. Deployment, service e persistent volume do Mysql





##### A imagem do container está disponível no registry do [Docker Hub clicando aqui](https://hub.docker.com/r/despossivel/server-go-lang)


**obs:** Utilize a tag beta

`
docker pull despossivel/server-go-lang:beta
`



@[despossivel](http://instagram.com/despossivel)