"# golang-crud" 

**1. database**

CREATE TABLE `employee` (
  `id` int(11) NOT NULL AUTO_INCREMENT PRIMARY KEY,
  `name` varchar(40) NOT NULL,
  `city` varchar(40) NOT NULL
) 

**2. Install driver for Go’s MySQL**

go get -u github.com/go-sql-driver/mysql


**3. Install the mux driver for routing the url**

go get github.com/gorilla/mux


