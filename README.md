# Greeting-App
Simple greeting app using Go

![image](https://github.com/MiguelAngelHorta/Greeting-App/assets/106134627/0107808f-c77d-4c37-bcae-720a1ed83ef3)

## Summary

This Go application sets up a basic HTTP server that serves an HTML file and handles requests to a `/greet` endpoint.

### `handleGreet`
This function handles requests to the `/greet` endpoint. It retrieves the `name` parameter from the request URL, generates a greeting message using the retrieved name, and writes the greeting as the HTTP response.

### `main`
This is the main function of the program. It sets up the HTTP server by:
- Serving the `index.html` file at the root endpoint `/`.
- Registering the `handleGreet` function to handle requests to the `/greet` endpoint.
- Starting the server on port `8080` and prints a message to indicate that the server is running.

![image](https://github.com/MiguelAngelHorta/Greeting-App/assets/106134627/b4876a0c-5d19-4dfc-be1b-7fcda1942219)
