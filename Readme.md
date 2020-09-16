## Returns Logistics Operations

The project has implemented a rest api for a returns management system. it can be used by any logistics company, the api can store data in PostGreSQl database and packaged as a Docker container.

The following tools are used for this project

       Go SDK 1.10.3
       Gorilla Mux 
       go-pg           // Go Postgress Library 
       Postgres
       Docker 
       Ubuntu 20.04
       
 #### Setting up Envoirment
 
##### Installing postgres in Ubuntu 20.04
    
    sudo apt install postgresql
    // Verify if the Postgresql Server is running 
    ss -nlt
 By default the Server is running at port 5432. But the configuration 
 can be changed by editing the following file
    
    sudo nano /etc/postgresql/12/main/postgresql.conf

##### Installing Postgres client 
 
     sudo apt install postgresql-client

   
   Creating user, database and adding access on PostgreSQL

    Connecting to the psql CLI 

            sudo -u postgres psql
 
     Create Database, User and assign the role to user.
    
            create database returnsdb;
            create user test with encrypted password 'test123';
            grant all privileges on database returnsdb to test;





### Go API SERVER

The following api endpoints would be helpful in our Senario. A Note For Frontend Developer
 
    1. content-type  on your request must be x-www-form-urlencoded
       Because according to the docs on r.ParseForm() the body won't be parsed 
       
       unless it's x-www-form-urlencoded
    
 
CRUD FOR CONSIGNMENTS
 
 GET
     
     /api/consignment/{id}
     /api/consignment/all
 
 POST
    
     
     /api/consignment/new
 
 PUT 
 
    /api/consignment/update/{id}
 
 DELETE 
    
     /api/consignment/delete/{id}
     
CRUD FOR SUPPLIERS
 
 GET
     
      /api/supplier/{id}
 
 POST
    
    /api/supplier/new
 
 PUT 
 
    /api/supplier/update/{id}
 
 DELETE 
    
     /api/supplier/delete/{id}

CRUD FOR CUSTOMERS
 
 GET
     
     /api/customer/{id}
 
 POST
    
    /api/customer/new
 
 PUT 
 
    /api/customer/update/{id}
 
 DELETE 
    
     /api/customer/delete/{id}

CRUD FOR DeliveryVan
 
 GET
     
     /api/customer/{id}
 
 POST
    
    /api/customer/new
 
 PUT 
 
    /api/van/update/{id}
 
 DELETE 
    
    /api/van/delete/{van_id}
