Biodata 
Kode Peserta : 149368582100-372
Nama		 : Purwitasari
Link Github  : 


Paduan Running Aplikasi	: 
1. Jalankan Xampp Apcahe dan MySQL
2. Buat database dengan nama mygram pada http://localhost/phpmyadmin/
3. Buka VSCode di folder MyGram
4. Buka terminal dan jalankan go run main.go
5. Klik allow_access firewall
6. Buka Postman  http://localhost:8080/


Register and Login
	1. POST http://localhost:8080/users/register (formKey: user_name,user_password(6 character),user_age(>8),user_email)
	2. POST http://localhost:8080/users/login 	(formKey: user_email,user_password)

User (Bearer Token)
	1. GET      http://localhost:8080/users 
	2. PUT      http://localhost:8080/users (formKey: user_name,user_email,user_age)
	3. DELETE   http://localhost:8080/users 

Photo (Bearer Token)
	1. GET      http://localhost:8080/photos 
	2. POST     http://localhost:8080/photos     		(formKey: photo_url,photo_title,photo_caption)
	3. PUT      http://localhost:8080/photos/:photoId   (formKey: photo_url,photo_title,photo_caption)
	4. DELETE   http://localhost:8080/photos/:photoId 

Comment (Bearer Token)
	1. GET      http://localhost:8080/comments/:photoId 			
	2. POST     http://localhost:8080/comments/ 	        (formKey:message,photo_id)
	3. PUT      http://localhost:8080/comments/:commentId 	(formKey:message)
	4. DELETE   http://localhost:8080/comments/:commentId 	

Social Media (Bearer Token)
	1. GET      http://localhost:8080/socialmedias 			
	2. POST     http://localhost:8080/socialmedias 	 	            (formKey: name,sosmed_url)
	3. PUT      http://localhost:8080/socialmedias/:socialMediaId 	(formKey: name,sosmed_url)
	4. DELETE   http://localhost:8080/socialmedias/:socialMediaId 




Tahapan Pembuatan Final Project :

1. Membuat folder final-project
2. masuk ke vs code
3. membuka folder directory folder final project di vscode
4. Membuat folder MyGram
5. cd MyGram di terminal
6. go mod init MyGram di terminal
7. Install/go get package : 
        "github.com/dgrijalva/jwt-go"
		"github.com/gin-gonic/gin"
		"golang.org/x/crypto/bcrypt"
		"github.com/jinzhu/gorm"
		"github.com/asaskevich/govalidator"
		"github.com/go-sql-driver/mysql"
8. Persiapkan database mygram di mysql ( http://localhost/phpmyadmin/)
9. Muat folder controller, router, assets, database, middleware, helper, models
10. Membuat main.go
11. Membuat db.go untuk connect ke database mygram di mysql
12. Membuat struct pada folder models untuk auto migration
13. Membuat auth untuk register dan login 
14. Membuat authetikasi token
15. Membuat router