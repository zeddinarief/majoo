# MAJOO login test using Go
This is using Gin for routing framework and Gorm for Go ORM Mysql
with service-object pattern 

endpoint:
POST /api/login
raw
body:
{
    "username" : "ZEZEZE",
    "password" : "cobapass"
}

POST /api/register
raw
body:
{
    "username" : "ZEZEZE",
    "password" : "cobapass",
    "nama_lengkap" : "ze ze"
}

POST /api/foto
form-data type file
key: filefoto
value: image.jpg
