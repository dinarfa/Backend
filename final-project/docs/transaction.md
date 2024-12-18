# Transaction API Specification

## Create Transaction API -->

**Endpoint** : POST `http:/localhost:8080/transactions/`

**Request Body** :

```json
{
  "user_id": 1,
  "transaction_type_id": 1,
  "description": "Pembayaran listrik bulan Desember",
  "amount": 150000
}
```

**Response Body Success** :

```json
{
    "message": "Transaksi berhasil ditambahkan"
}
```

**Response Body Error** :

```json
{
  "message": "Mohon input dengan benar"
}
```

## Update Transaction API

**Endpoint** : PUT `http:/localhost:8080/transactions/id`

**Request Body** :

```json
{
  "user_id": 1,
  "transaction_type_id": 2,
  "description": "Pembayaran listrik bulan Desember",
  "amount": 150000
}

```

**Response Body Success** :

```json
{
    "message": "Transaksi berhasil di-update"
}
```

**Response Body Error** :

```json
{
  "message": "Mohon Input dengan benar"
}
```

## Get Transaction By Id API

**Endpoint** : GET `http:/localhost:8080/transactions/{id}`

**Headers** :

- Authorization : Bearer token

**Response Body Success** :

```json
{

  "id": 1,
  "user_id": 1,
  "transaction_type_id": 1,
  "description": "Pembayaran listrik bulan Desember",
  "amount": 150000,
  "date": "0001-01-01T00:00:00Z"
    
}
```
**Response Body Error** :

```json
{
  "message": "Maaf, data tidak ditemukan."
}
```
## Delete Transaction API

**Endpoint** : DELETE `http:/localhost:8080/transactions/{id}`

**Headers** :

- Authorization : Bearer token

**Response Body Success** :

```json
{
    "message": "Transaksi berhasil dihapus",
    "transaction_id": 3
}
```

**Response Body Error** :

```json
{
    "error": "Transaksi dengan ID tersebut tidak ditemukan"
}
```

## Get All Transaction API

**Endpoint** : GET `http:/localhost:8080/transactions/`

**Headers** :

- Authorization : Bearer token

**Response Body Success** :

```json
[{
  "id": 1,
  "user_id": 1,
  "transaction_type_id": 1,
  "description": "Pembayaran listrik bulan Desember",
  "amount": 150000,
  "date": "0001-01-01T00:00:00Z"
},
{
  "id": 1,
  "user_id": 2,
  "transaction_type_id": 1,
  "description": "Pembayaran pulsa",
  "amount": 20000,
  "date": "0001-01-01T00:00:00Z"
}
]
```

**Response Body Error** :

```json
{
  "message": "[]"
}
```

## Get All Transaction Types API

**Endpoint** : GET `http:/localhost:8080/transactions_types/`

**Headers** :

- Authorization : Bearer token

**Response Body Success** :

```json
[{
  "transaction_type_id": 1,
  "name": "Pemasukan"
    
},
{
  "transaction_type_id": 2,
  "name": "Pengeluaran"
    
}
]
```
**Response Body Error** :

```json
{
  "message": "pesan eror"
}
```

## Create Transaction API -->

**Endpoint** : POST `http:/localhost:8080/transactions_types/`

**Request Body** :

```json
{
  "name": "Pemasukan"
}
```

**Response Body Success** :

```json
{
  "transaction_type_id": 1,
  "name": "Pemasukan"
}
```

**Response Body Error** :

```json
{
  "message": "pesan eror"
}
```
