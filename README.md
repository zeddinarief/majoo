# majoo using Go

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
