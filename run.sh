export LOGGER_FILENAME="./otto-logger.log"

#main.go
export OTTOSFA_API_APK_PORT=0.0.0.0:8044
export MAXPROCS="1"

#routes/route.go
export OTTOSFA_API_APK="OTTOSFA_API_APK"
export READ_TIMEOUT="120"
export WRITE_TIMEOUT="120"
export JAEGER_HOSTURL="13.250.21.165:5775"

#db/postgres.go
#export DB_POSTGRES_USER=ubuntu
#export DB_POSTGRES_PASS="Ubuntu!23"
#export DB_POSTGRES_NAME="otto-sfa-admin-api_production"
export DB_POSTGRES_ADDRESS="34.101.208.156"
#export DB_POSTGRES_PORT="5432"
export DB_POSTGRES_DEBUG="true"
export DB_POSTGRES_TYPE="postgres"
export DB_POSTGRES_SSL_MODE="disable"

export DB_POSTGRES_USER="sfa_web_usr"
export DB_POSTGRES_PASS="sfa32021"
export DB_POSTGRES_NAME="otto-sfa-admin-api"
export DB_POSTGRES_HOST="34.101.208.156"
export DB_POSTGRES_PORT="6432"

#database/postgresrose/postgresrose.go
export DB_ROSE_POSTGRES_USER="ottoagcfg"
export DB_ROSE_POSTGRES_PASS="dTj*&56$es"
export DB_ROSE_POSTGRES_NAME="rosedb"
export DB_ROSE_POSTGRES_ADDRESS="13.228.23.160"
export DB_ROSE_POSTGRES_PORT="8432"
export DB_ROSE_POSTGRES_DEBUG="true"
export DB_ROSE_POSTGRES_TYPE="postgres"
export DB_ROSE_POSTGRES_SSL_MODE="disable"
export DB_ROSE_POSTGRES_TIMEOUT="30"

#hosts/stringbuilder
export STRINGBUILDER_HOST="http://13.228.25.85:8995"
export STRINGBUILDER_ENDPOINT_REVERSE_QR="/merchant/reverseqr"
export STRINGBUILDER_HEALTH_CHECK_KEY"="OTTO-SFA-API_HEALTH_CHECK:STRINGBUILDER"
export STRINGBUILDER_NAME"="QRIS.STRINGBUILDER"

#ROSE-API-SERVICE
export ROSE_API_SERVICE_HOST=http://13.228.25.85:8899/rose-api-service/v0.0.1
export ROSE_API_SERVICE_ENDPOINT_INQUIRY_MERCHANT=/inquiry-merchant/find
export ROSE_API_SERVICE_ENDPOINT_LOOK_UP_GROUP="/lookup/lookupgroup"
export ROSE_API_SERVICE_ENDPOINT_USER_CATEGORY="/user-category/find"
export ROSE_API_SERVICE_ENDPOINT_KEY_APP_ID="3"
export ROSE_API_SERVICE_ENDPOINT_UPDATE_MERCHANT="/updated-data-merchant"
export ROSE_API_SERVICE_ENDPOINT_FIND_BY_MID="/owner/get-merchant"

#ROSE-OP-SERVICE
export ROSE_OP_SERVICE_HOST="http://13.228.25.85:8914/rose-op-service/v0.0.1"
export ROSE_OP_SERVICE_ENDPOINT_MERCHANT_PROFIL="/merchant/info/profile"

#REDIS
export REDIS_HOST_CLUSTER1=13.228.23.160:8079
export REDIS_HOST_CLUSTER2=13.228.23.160:8078
export REDIS_HOST_CLUSTER3=13.228.23.160:8077
export REDIS_KEY_AKUISISI="BACKUP:SFA:ACQUITISIONS:CHECK:PhoneNumber:"
export REDIS_KEY_EXP_TOKEN="2h"

#ROSE/LOOKUP
export LOOKUP_LOKASI_BISNIS="BUSINESS_LOCATION"
export LOOKUP_JENIS_LOKASI_BISNIS="MERCHANT_LOCATIONS"
export LOOKUP_TIPE_BISNIS="JENIS_USAHA"
export LOOKUP_KATEGORI_BISNIS=""
export LOOKUP_JAM_OPERASIONAL="OPERATION_HOURS"
export LOOKUP_JAM_KUNJUNGAN_TERBAIK="BEST_VISITS"

#ACQUITISION
export ACQUITISIONS_HOST="http://13.228.25.85:8871/v3"
export ACQUITISIONS_ENDPOINT="/merchants/new"

#router/routers.go
#swagger
export APPS_ENV="local"
export SWAGGER_HOST_LOCAL="localhost:8044"
export SWAGGER_HOST_DEV="13.228.25.85:8044"

#FDS
export FDS_HOST="https://admin.ottopay.id/merchant/rest/sfa" #"https://admin.ottopay.id/merchant/rest/sfa"
export FDS_ENDPOINT_GET_LONGITUDE_LATITUDE="/getGeotag"

#REDISCLUSTER
export REDIS_MASTER="13.228.23.160:8079;13.228.23.160:8078;13.228.23.160:8077"
export REDIS_SLAVE="13.228.23.160:6479;13.228.23.160:6478;13.228.23.160:6477"

export REDIS_MASTER_1="34.101.208.23:8177"
export REDIS_MASTER_2="34.101.208.23:8178"
export REDIS_MASTER_3="34.101.208.23:8179"
export REDIS_SLAVE_1="34.101.208.23:8174"
export REDIS_SLAVE_2="34.101.208.23:8175"
export REDIS_SLAVE_3="34.101.208.23:8176"

export URL_IMAGE="http://13.228.25.85:9000/rose/"

#IDM
export IDM_HOST="http://ottopayqavm.southeastasia.cloudapp.azure.com:8885/indomarco/v0.1.0"
export IDM_ENDPOINT_CHECK_ACCOUNT="/users/check_account"
export IDM_ENDPOINT_UPDATE_MERCHANT="/users/update_merchant"
export IDM_HEALTH_CHECK_KEY="OTTOSFA-API-APK_HEALTH_CHECK:INDOMARCO"

#HOSTS/NOTIF
export NOTIF_KEY="AAAAU4tck2Q:APA91bFXkMQOmQQxBiTq6VUrB52af9Q2oyaIdKCNtbRqVF8gJ-_P7XR8luW0c9FwtCDc8lgHwTrgBt4iLNK8zlvifQLhAL-sS3M0Wq3gTvBMqLXoItP5M_weYfsKksDxzXZrrLt4ZHuC"
export NOTIF_GRPC_ADDRESS="18.138.101.49:9191"

export ATTENDANCE_EVENT_TIME_1=" 13:00:00"
export ATTENDANCE_EVENT_TIME_2=" 13:59:59"
export REDIS_KEY_AKUISISI="SFA:ACQUITISIONS:CHECK:PhoneNumber:"

#message
export OTTOSFA_MESSAGE_DESC_SALES_PENDING="Akun kamu sedang dalam proses verifikasi"

#version
export VERSION_RELEASE="V.0.1"
export VERSION_GIT_BRANCH_NAME="sfa_3.0.0"
export VERSION_GIT_COMMIT="51ee8631f320b63bdad1e635c9eb882441cd4086"
export VERSION_DATE_TIME="13 Oct 2021"
export VERSION_TEAM_CREATED="Wayang Squad - Srikandi Team"
export VERSION_NAME_SERVICES="ottosfa-api-apk"
export VERSION_TYPE="SIT"

go run main.go