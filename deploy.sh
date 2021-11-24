#!/usr/bin/env bash

env GOOS=linux GOARCH=amd64 go build -o ottosfa-api-apk

 scp -i ~/.ssh/LightsailDefaultKey-ap-southeast-1-new.pem -P 22 ottosfa-api-apk ubuntu@13.228.25.85:/home/ubuntu

#scp -i /users/abdulah/OttopayAwsLite.pem ottosfa-api-apk ubuntu@13.228.25.85:/home/ubuntu/ottosfa-api-apk-1

#env GOOS=linux GOARCH=amd64 go build -o ottosfa-api-apk

#scp -i ~/.ssh/devsfa.priv -P 22 ottosfa-api-apk devsfa@34.101.141.240:/home/devsfa

#ssh devsfa@34.101.141.240 -i ~/.ssh/devsfa.priv

#sudo su - devsfa /opt/apps/wayang/SRIKANDI

#ssh devsfa@34.101.141.240 -i ~/.ssh/devsfa.pem

#env GOOS=linux GOARCH=amd64 go build -o ottosfa-api-apk

#scp -i ~/.ssh/LightsailDefaultKey-ap-southeast-1-new.pem -P 22 ottosfa-api-apk nc_ketut@10.10.43.22:/home/nc_ketut