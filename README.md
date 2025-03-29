# Project FileShare

This is a fileshare and management application like filetransfer. Here users can register and login to upload their files and share the file url to whoever they want. files will be expired after 2 days.

## Installation

To install and run this application, follow these steps:

1. Make sure you have Docker and Docker Compose installed on your machine.

2. Clone the repository to your local machine.

3. Edit the following environment variables in the docker-compose.yaml file to match your desired configuration.

```
      - JWT_SECRET=
      - AWS_S3_REGION=
      - AWS_S3_BUCKET=
      - AWS_ACCESS_KEY_ID=
      - AWS_SECRET_ACCESS_KEY=

```

4. Open a terminal and navigate to directory where docker-compose.yaml available

5. Run the following command to start the application using Docker Compose:

    ```
    docker-compose up
    ```

6. The application should now be running. You can access it by opening your web browser and navigating to `http://127.0.0.1:8080`.

7. You can check the server status by `curl -X GET http://127.0.0.1:8080/status` 

## Usage

1. This application is also hosted in aws. You can also use those API endpoints to directly access the application without running in your local machine.


## API Documentation

- Cloud hosted application URL
    ```
    http://fileshare.kavirajk.me/<endpoint>
    ```

## API Documentation

1. To check server status

    - **Endpoint:** `/status`
    - **Method:** GET
    - **Request Body:** N/A
    - **Output:** Returns the status of the server.

2. To register a new user

    - **Endpoint:** `/signup`
    - **Method:** POST
    - **Request Body:**
    ```
    {
      "username": "user",
      "email":"user@email.com",
      "password": "password"
    }
    ```
    - **Desired Output:** Returns the user registration success message.

3. To login

    - **Endpoint:** `/login`
    - **Method:** POST
    - **Request Body:**
    ```
    {
      "username": "user",
      "password": "password"
    }
    ```
    - **Desired Output:** Returns a JWT token for authentication. 

4. To upload a file

    - **Endpoint:** `/upload`
    - **Method:** POST
    - **Request Body:** Form data with the following fields:
      - `file`: The file to be uploaded
      - key:`file`
      - value: path of file 
    - **Desired Output:** Returns the metadata of the uploaded file.

5. To get a list of uploaded files

    - **Endpoint:** `/files`
    - **Method:** GET
    - **Request Body:** N/A
    - **Desired Output:** Returns a list of uploaded files with its metadata of the logged in user.

6. To share a file based on file id 
    - **Endpoint:** `/share`
    - **Method:** GET
    - **Request Body:** 
    ```
    {
      "file_id": 1
    }
    ```
    - **Desired Output:** Returns metadata of the file of logged in user.

7. To delete a file of a user
    - **Endpoint:** `/delete`
    - **Method:** DELETE
    - **Request Body:** 
    ```
    {
      "file_id": 1
    }
    ```
    - **Desired Output:** Return delete success message of user's file.

8. To Search a file based on filename
    - **Endpoint:** `/search`
    - **Method:** GET
    - **Request Body:** 
    ```
    {
      "filename": "text.txt"
    }
    ```
    - **Desired Output:** Returns metadata of the file of logged in user.


