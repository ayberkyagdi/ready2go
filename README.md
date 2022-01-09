# Ready2Go
Ready2Go is an online retail site demo. Customer and store accounts can be created then shopping can be done via shopping carts. 
The output of order information is displayed in json format. It can also be managed on the admin page.

# Installation & Requirements 

**Go** and **MySQL** 
<br />
<br />
Clone the project, set **.env** file according to the DB connection. 
<br />
DB_USER=**example_root**
<br />
DB_PASSWORD=**example_pass**
<br />
DB_NAME=**example_name**

# Usage

Simply run **go run main.go** and can visit **localhost:8080**

UI will guide you. You can simply create customer and store accounts. Also add products when you log in with your store account. Then, customer accounts will be able to view the shops in **their city** and place an order. When the customer request an order and approve the shopping cart, an order summary will be created in the **Orders** folder. All customers and stores can be seen on the admin page and has been authorized to **edit** and **delete**. <br /><br />**ADMIN PAGE**<br />Username: **admin** <br />Password: **admin**<br /><br /> must be used on the user login page to log in to the admin account.
