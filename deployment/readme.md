# Deployment

## GCP
* Buat Instance (Compute Engine): https://cloud.google.com/compute/docs/instances/create-start-instance
* Generate SSH untuk connect ke Instance GCP: https://cloud.google.com/compute/docs/connect/create-ssh-keys#windows-10-or-later
* Menambahkan SSH ke Instance: https://cloud.google.com/compute/docs/connect/add-ssh-keys
* Setup Firewall Rule: https://cloud.google.com/vpc/docs/using-firewalls

Note saat melakukan generate SSH di windows, jika nama user windows nya ada spasi, maka tambahkan tanda "" di url tempat penyimpanan file ssh nya.
contoh "c:/Users/Fakhry Firdaus/.ssh/sshgcpalta"

## Connect to Compute Engine / EC2
```bash
ssh -i </directory/namafilessh> <username-server>@<public-ipv4>

# example:
[GCP] ssh -i ~/.ssh/sshgcp fakhry@18.0.1.2

[AWS] ssh -i ~/.ssh/file.pem ubuntu@18.0.1.2
```
Note:
* Untuk connect ke GCP, kita perlu generate ssh key terlebih dahulu, dan menambahkan ssh key public nya ke setup instance GCP.
* Untuk connect ke AWS, kita dapat menggunakan file .pem yang didapat saat awal pembuatan Instance untuk connect ke Instance AWS. 
* untuk nama username di GCP, biasanya sesuai dengan nama alamat email kita. misal fakhry@gmail.com
maka username instance kita adalah fakhry.
* Untuk AWS, jika anda menggunakan ubuntu. maka username nya adalah ubuntu.

* Jika ada error saat akses server via ssh (permission denied). coba ubah hak akses file ssh nya
    ```
    chmod 400 namafile
    ```


## Update & Upgrade OS Ubuntu Instance
```bash
sudo apt-get update
sudo apt-get upgrade
```

## Transfer File/Folfer to Server using SCP
Tambahkan `-r` jika ingin transfer folder

```bash
scp -i </direktori/ssh-key-private> </direktori/nama-file-transfer> <username>@<public-ip-server>:/home/fakhry

# example: folder
 scp -i ~/.ssh/sshgcpalta -r 14-deployment fakhry@10.170.120.170:/home/fakhry

 # file
 scp -i ~/.ssh/sshgcpalta 14-deployment/main.go fakhry@10.170.120.170:/home/fakhry
```


# DOCKER

## Install Docker on Ubuntu Server
```bash
sudo apt install docker.io
```
Note: pastikan install berhasil dan kita sudah bisa menjalankan perintah `docker -v`

## Jika ada error (Permission Denied) ketika run docker di Ubuntu
```bash
sudo usermod -a -G docker ubuntu

or

sudo chmod 777 /var/run/docker.sock
```

## Build Docker Image
```bash
docker build -t <nama-image>:<tag> .

# example
docker build -t be15-images:latest .
```

## Show Image List
```bash
docker images

docker images list
```

## Delete Docker Image
```bash
docker rmi <image-id>
#or
docker rmi <image-name>

# example:
docker rmi be15-images
```

## Create Docker Container
Note:
* https://docs.docker.com/engine/reference/commandline/run/
* -d digunakan agar app berjalan di background
* host-port : isi dengan port yang akan digunakan di dockernya
* container-port / app port: isi dengan port yang digunakan di app golang (di bagian e.Start())

```bash
docker run -d
-p <host-port>:<container-port>
-e <env-name>=<env-value>
-e <env-name>=<env-value>
-v <host-volume>:<container-volume>
--name <container-name> <image-name>:<tag>

# example:
docker run -p 80:80 --name apiContainer api-images:latest

# example with env
docker run -d -p 80:80 -e JWT_KEY=blabla -e DBUSER=root -e DBPASS=abcdef -e DBHOST=10.10.20.30 -e DBPORT=3306 -e DBNAME=dbapi --name be15Clean be15clean-images:latest

#or
docker run -d -p 80:80 --env-file=path/.env --name be15Clean be15clean-images:latest
```

## Show Container
```bash
# melihat container yang sedang running
docker ps

# melihat seluruh container, termasuk yang sedang stop
docker ps -a
```

## Start/Stop Container
```bash
docker stop <container-name>

docker start <container-name>
```

## Remove Docker Container
```bash
docker rm <container-name>

docker rm <container-id>

# example
docker rm apiContainer
```

## Docker Logs Container
```bash
docker logs <container-name>
```

## Push Image to Docker Hub
```bash
docker login -u <username-dockerhub>

docker build -t <username-dockerhub>/<image-name>:<tag> .

docker push <username-dockerhub>/<image-name>
```

## Pull image dari docker hub
```bash
docker pull <image-name>
```

## Menjalankan MySQL diatas Docker Container
if you are using docker-desktop, the containers can access host os by using host.docker.internal name.

Otherwise, you can use default host IP address: 172.17.0.1
```
# How to create and run mysql DB on container docker
docker run -p 3307:3306 \
-e MYSQL_ROOT_PASSWORD=rahasia \
-e MYSQL_DATABASE=altagormdb \
-v mysql_data:/var/lib/mysql \
--name mysqldocker \
-d \
mysql

# How to create and run our app on container docker
docker run -d -p 5000:8080 -e APP_PORT="8080" -e DB_PORT="3306" -e DB_DRIVER="mysql" -e DB_NAME="altagormdb" -e DB_ADDRESS="172.17.0.1" -e DB_USERNAME="root" -e DB_PASSWORD="rahasia" --name altacontainer alta-images:latest

```

## Check IP docker container
```bash
docker inspect -f '{{range.NetworkSettings.Networks}}{{.IPAddress}}{{end}}' container_name_or_id
```

## Create docker compose for MySQL
```yaml
version: '3.8'

services:
  mysql:
    image: mysql:8
    container_name: mysql_container
    environment:
      MYSQL_ROOT_PASSWORD: your_root_password
      MYSQL_DATABASE: your_database_name
      MYSQL_USER: your_mysql_user
      MYSQL_PASSWORD: your_mysql_password
    ports:
      - "3306:3306"
    volumes:
      - mysql_data:/var/lib/mysql

volumes:
  mysql_data:
```