Book Management CRUD API built using Gorilla Mux.

ROUTES:

1. GET -> /book : Returns all the books in the database in JSON format. 
2. GET -> /book/id : Returns the book with the given id in JSON format.
3. POST -> /book : Creates the book from the JSON passed in the body, and returns the same book in JSON format.
4. PUT -> /book/id : Updates the book with the given id, with the JSON passed in the body and returns the same.
5. DELETE -> /book/id : Deletes the book with the given id.
