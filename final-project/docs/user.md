# User API Specification

## Register User API -->

**Endpoint** : POST `http:/localhost:8080/users/register`

**Request Body** :

```json
{
  "name": "Estes",
  "email": "estes@gmail.com",
  "password": "rahasia",
  "confirm-password": "rahasia",
  "balance": 20000
}
```

**Response Body Success** :

```json
{
  "message": "Registrasi berhasil",
}
```

**Response Body Error** :

```json
{
  "message": "Email sudah terpakai"
}
```

## Login User API

**Endpoint** : POST `http:/localhost:8080/users/login`

**Request Body** :

```json
{
  "email": "dinar@gmail.com",
  "password": "rahasia"
}
```

**Response Body Success** :

```json
{
  "message": "Login Berhasil",
  "access_token": "unique-token"
  
}
```

**Response Body Error** :

```json
{
  "message": "Password salah"
}
```

## Get User By Id API

**Endpoint** : GET `http:/localhost:8080/users/{id}`

**Headers** :

- Authorization : Bearer token

**Response Body Success** :

```json
{
  "message": "Data ditemukan",
  "name": "Dinar",
  "email": "dinar@gmail.com" 
}
```

**Response Body Error** :

```json
{
  "message": "Maaf, data tidak ditemukan."
}
```
## Get balance By Id API

**Endpoint** : GET `http:/localhost:8080/users/{id}/balance`

**Headers** :

- Authorization : Bearer token

**Response Body Success** :

```json
{
    "current_balance": 200000
}
```

**Response Body Error** :

```json
{
    "error": "user dengan id 2 tidak ditemukan"
}
```
## Delete user API

**Endpoint** : DELETE `http:/localhost:8080/users/{id}`

**Headers** :

- Authorization : Bearer token

**Response Body Success** :

```json
{
    "message": "User berhasil di-hapus"
}
```

**Response Body Error** :

```json
{
    "message": "Gagal menghapus user"
}
```

## Update User API -->

**Endpoint** : PUT `http:/localhost:8080/users/{id}`

**Request Body** :

```json
{
  "name": "Estes",
  "email": "estes@gmail.com",
  "password": "rahasia",
  "balance": 20000
}
```

**Response Body Success** :

```json
{
    "message": "User berhasil di-update"
}
```

**Response Body Error** :

```json
{
  "message": "Mohon input dengan benar"
}
```