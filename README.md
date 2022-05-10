# STL-tech-companies-REST-API
A REST API for accessing data on the 100 fastest growing tech companies in St. Louis

## Initial Setup
Before running, a PostgreSQL server must be setup locally, as there is not a remote server active at this time. The local server must contain a database called "tech_companies" (or you can call it whatever you want, as long as the DB_NAME field in the .env file is changed to reflect the database name).

Once the server is set up and active, you can run the main.go file as follows:

```
go run main.go
```

If everything is setup correctly, you should see no error messages, but a message from the Fiber framework that the API is active, along with a local host url for accessing the API.


## populating the database through the API
Upon initial setup, the database will currently be empty, and requires an additional step for populating it. There is a Python script, called webscraper.py that is inside the main project directory that will accomplish this. This script scrapes the website linked below, collecting all of the data for each company listed, and makes subsequent POST api calls for each company to populate the database. Once this script is ran, the database should be populated, and can then be interacted with through the API by the methods detailed in the next section.

**website that the company data was collected from**: https://growjo.com/city/St_Louis

## Interacting with the API through http requests
The base address for the API is as follows:
  http://*localhost IP*:8080/api **NOTE**: The IP address may be different depending on whatever your localhost port is.
  
**Add data to database(POST)**:
  http://*localhost IP*:8080/api/create_company

**Delete a company record (DELETE)**:
  http://*localhost IP*:8080/api/delete_company/:id

**Get a specific company by ID (GET)**:
  http://*localhost IP*:8080/api/get_company/:id

**Get a full list of every company record in the database (GET)**:
  http://*localhost IP*:8080/api/tech_companies
  
**NOTE**: "localhost IP" should be replaced by the localhost port IP address of your machine (it is likely 127.0.0.1, but could be different). The Fiber success message once the main.go file is ran will display the correct one.


