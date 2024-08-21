# Test Backend - Sharing Vision

## CDatabase (Bobot 20%)

1.  Create Table Article Manually

```
CREATE DATABASE IF NOT EXISTS `article`;
USE `article`;

CREATE TABLE `postsposts` (
    `id` INT AUTO_INCREMENT PRIMARY KEY,
    `title` VARCHAR(200),
    `content` TEXT ,
    `category` INT(100),
    `created_date` TIMESTAMP,
    `updated_date` TIMESTAMP,
    `status` VARCHAR(100)
);
```

2.  Create Table Article with migration

````
- Install gorm
    go get -u gorm.io/gorm
    go get -u gorm.io/driver/mysql

- Create Model
    type Posts struct {
        Id          int       `json:"id" gorm:"primaryKey;autoIncrement"`
        Title       string    `json:"title" gorm:"type:varchar(200)"`
        Content     string    `json:"content" gorm:"type:text"`
        Category    string    `json:"category" gorm:"type:varchar(100)"`
        CreatedDate time.Time `json:"created_date" gorm:"autoCreateTime"`
        UpdatedDate time.Time `json:"updated_date" gorm:"autoUpdateTime"`
        Status      string    `json:"status" gorm:"type:varchar(100)"`
    }
    ```

- Create Migration and call in file main.go
   func DBMigrate() {
	err := DB.AutoMigrate(&models.Posts{})

	if err != nil {
		log.Println(err)
	}
	fmt.Println("Migration Success")
}

````

## Microservice (Bobot 80%)

3. Create Endpoint

- Create Article

```
    Endpoint : /article
    Method : POST
```

- Get All Article

```
    Endpoint : /article/:limit/:offset
    Method : GET
```

- Get Article By Id

```
    Endpoint : /article/:id
    Method : GET
```

- Update Article

```
    Endpoint : /article/:id
    Method : PUT
```

- Delete Article

```
    Endpoint : /article/:id
    Method : DELETE
```

4. Create Postman Collection and this request

```
    attached
```

# Installation

1. Git Clone https://github.com/komporbakar/backend-technical-test-sinar-vision.git

```
git clone https://
```

2. Open Project

```
cd backend-technical-test-sinar-vision
```

3. Running Project

```
air
```

Thank You,
