<pre style="font-size: 1.4vw;">
<p align="center">
                                          
 _              _                         
| |_  ____  ___| |_  ____ ____  ___  ____ 
|  _)/ _  )/___)  _)/ ___) _  |/___)/ _  )
| |_( (/ /|___ | |_( (__( ( | |___ ( (/ / 
 \___)____|___/ \___)____)_||_(___/ \____)
                                          
</p>
</pre>
<p align="center">
<a href="https://golang.org/">
    <img src="https://img.shields.io/badge/Made%20with-Go-1f425f.svg">
</a>
<a href="/LICENSE">
    <img src="https://img.shields.io/badge/License-MIT-green.svg">
</a>
</p>
<p align="center">
<b>Go - Testcase</b> is a test preparation </b>
</p>

# Testcase API Guide

## 🔀 How to Run :

```js
- clone the project using, git clone https://github.com/Yoga-Saputra/testcase.git
- open the directory project and run go mod tidy then go mod vendor
- create file .config.yaml like the example and set the configuration
- this project can be run using air or using command go run main.go --run
- if you want to run this project using air, make sure air already install  on your laptop
```

## 🔀 Compatible Route Endpoint

| NO  | Use                    | Endpoint                   | Example                                      | Action |
| --- | ---------------------- | -------------------------- | -------------------------------------------- | ------ |
| 1   | List Product datatable | api/v1/product/list        | http://localhost:4040/v1/product/list        | GET    |
| 2   | Create Product         | api/v1/product/create      | http://localhost:4040/v1/product/create      | POST   |
| 3   | Update Product         | api/v1/product/update/{id} | http://localhost:4040/v1/product/update/{id} | PUT    |
| 4   | Delete Product         | api/v1/product/delete/{id} | http://localhost:4040/v1/product/delete/{id} | DELETE |
| 5   | Create Brand           | api/v1/brand/create        | http://localhost:4040/v1/brand/create        | POST   |
| 6   | Delete Brand           | api/v1/brand/delete/{id}   | http://localhost:4040/v1/brand/delete/{id}   | DELETE |

---

## 📖 Compatible JSON Payload Testcase API

This is the JSON payload that's sended to Testcase API

### 💸 List Product Datatable Request

```js
curl --location --request GET 'localhost:4040/v1/product/list' \
--header 'Content-Type: application/json' \
--data '{
    "page": 0,
    "length": 2,
    "offset": 0
}'
```

### 💸 List Product Datatable Response

```js
{
    "draw": 1,
    "recordsTotal": 3,
    "data": [
        {
            "id": 3,
            "brand_id": 1,
            "nama_product": "GGCOK",
            "harga": 2000,
            "quantity": 4,
            "brand_name": "ABC",
            "created_at": "2024-12-28T11:12:12+07:00",
            "updated_at": "2024-12-28T11:20:33+07:00"
        },
        {
            "id": 4,
            "brand_id": 1,
            "nama_product": "GGC",
            "harga": 1000,
            "quantity": 1,
            "brand_name": "ABC",
            "created_at": "2024-12-28T14:34:50+07:00",
            "updated_at": "0001-01-01T00:00:00Z"
        }
    ]
}
```

### 💸 Create Product Request

```js
curl --location 'localhost:4040/v1/product/create' \
--header 'Content-Type: application/json' \
--data '{
	"nama_product": "test",
    "brand_id": 1,
    "harga": 1000,
    "qty": 1
}'
```

### 💸 Create Product Response

````js
{
    "success": true,
    "code": 2400,
    "data": "Product [test] successfully created",
    "error": null
}
```

### 💸 Update Product Request

```js
curl --location --request PUT 'localhost:4040/v1/product/update/1' \
--header 'Content-Type: application/json' \
--data '{
	"nama_product": "Test Update",
    "brand_id": 1,
    "harga": 12000,
    "qty": 5
}'
````

### 💸 Update Product Response

```js
{
    "success": true,
    "code": 2400,
    "data": "Data product successfully updated",
    "error": null
}
```

### 💸 Delete Product Request

```js
curl --location --request DELETE 'localhost:4040/v1/product/delete/10'
```

### 💸 Delete Product Response

```js
{
    "success": true,
    "code": 2400,
    "data": "Brand ID [10] successfully deleted",
    "error": null
}
```

### 💸 Create Brand Request

```js
curl --location 'localhost:4040/v1/brand/create' \
--header 'Content-Type: application/json' \
--data '{
	"brand_name": "Test Brand"
}'
```

### 💸 Create Brand Response

```js
{
    "success": true,
    "code": 2400,
    "data": "Brand name [Test Brand] successfully created",
    "error": null
}
```

### 💸 Delete Brand Request

```js
curl --location --request DELETE 'localhost:4040/v1/brand/delete/3'
```

### 💸 Delete Brand Response

```js
{
    "success": true,
    "code": 2400,
    "data": "Brand ID [3] successfully deleted",
    "error": null
}
```

## 📖 Excersise Essay

### 💸 Project review dan Project planning

```js
## Project review adalah sebuah proses evaluasi dan tinjauan terhadap kemajuan , kualitas dari proyek yang sedang kita kembangkan.  Tujuannya untuk memastikan bahwa proyek yang sedang dikembangkan berjalan sesuai timeline yang sudah direncanakan
## Project planning adalah sebuah tahapan perencanaan atau analisa yang dilakukan sebelum memulai pengembangan sistem atau aplikasi.
```

### 💸 Load balance dan Security group di AWS EC2

```js
## Load Balance adalah layanan yang mendistribusikan trafik masuk (incoming traffic) ke beberapa instances untuk memastikan aplikasi tetap berjalan secara optimal, ketika ada lonjakan trafik. Load balancer bertujuan untuk meningkatkan ketersediaan (availability) dan skalabilitas aplikasi dengan mendistribusikan beban secara merata, sehingga menghindari beban berlebih pada satu instance atau server.

## dikarenakan saya belum pernah mempelajari security group di aws EC2 atau devops, saya kurang begitu paham, tapi saya sayang antusian untuk mempelajari lebih lanjut untuk aws dan dari hasil saya membaca beberapa artikel berfungsi sebagai firewall virtual yang mengontrol trafik masuk dan keluar ke EC2 instances. Security group memungkinkan untuk menetapkan aturan (rules) berdasarkan IP address, protokol, dan port untuk menentukan jenis trafik yang diperbolehkan atau ditolak ke atau dari instance EC2.
```

### 💸 Menangani issue memory leak di golang

```js
## Contoh :
    func memoryLeak() {
        leakMap := make(map[int]*[]int) // Map yang menyimpan slice
        for i := 0; i < 1000; i++ {
            // Membuat slice baru untuk setiap iterasi
            nSlice := make([]int, 10)
            leakMap[i] = &nSlice  // Menyimpan referensi ke slice dalam map
        }

            // Setelah selesai digunakan, kita hapus referensinya
        for i := 0; i < 1000; i++ {
            delete(leakMap, i)  // Menghapus entri map setelah selesai digunakan
        }
    }
Berdasarkan contoh diatas, dapat disimpulkan bahwa kita perlu memastikan bahwa objek yang tidak digunakan lagi bisa dibebaskan dari memori dengan cara menghapus referensinya.
```
