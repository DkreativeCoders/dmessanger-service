# golang-crud-app-db
A service for real time delivery in golang

# Getting Started
1. Clone the project, go to the project root folder

2. Run <code> go get ./... </code> to install all dependencies

3. Install Mysql. The setup should contain mysql workbench 

4. Create a new file in the project root. The file name is ".env" (Note: Just ".env" no name), then copy the default 
parameters in the "example.env" file into the ."env" 

5. Open mysql workbench, open the default connection and create a new schema. 
(Note. The name you give it should match the one in the ".env"fle)

6. Open goland IDE, open the main.go file, right click and Run the file.

7. You should see the port displayed in the console

# Setting up swagger 
[Go Swagger installation guide](https://goswagger.io/install.html)

Use the [guide](https://goswagger.io/use/spec.html) to comment your code

Run <code>swagger generate spec -o ./swagger.json</code> to generate swagger specs

paste the generated swagger.json code into this [editor](https://editor.swagger.io/)

Updated the ReadMe.md file partially

@mr-m0nkey Still waiting on your updated version for easy installation, you can follow the apiguidlines.md file