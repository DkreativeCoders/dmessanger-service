# Api Guide Lines Rules
- An API is a user interface for developers. Put the necessary effort in it to ensure it’s not just functional but pleasant to use. Arsene Tochemey Gandote

### Key Requirements
An API is a developer’s UI-just like any UI. It is important to ensure the user’s experience is thought out carefully. So an API should:

- use web standards where they make sense.
- be friendly to the developer and be explorable.
- be simple, intuitive and consistent to make adoption not only easy but pleasant.
- provide enough flexibility to power the majority of the enchant UI.
- be efficient, scalable, while maintaining balance with the other requirements

### URLs and Actions
The key principles of REST involve separating your API into logical resources.
These resources are manipulated using HTTP requests where the method (GET, POST, PUT, PATCH, DELETE) has a specific meaning.

A resource should be a noun, not a verb. Examples:

- GET /tickets — Retrieves a list of tickets
- GET /tickets/12 — Retrieves a specific ticket #12
- POST /tickets — Creates a new ticket
- PUT /tickets/12 — Updates a given ticket #12
- PATCH /tickets/12 — Partially updates a given ticket #12
- DELETE /tickets/12 — Deletes a ticket #12

	```go
	router.HandleFunc("/api/users", handler.create).Methods("POST")

	```

### URLs and Actions — Relations
If a relationship can only exist within another resource, RESTful principles provide useful guidance. Let us look at few examples:

A ticket in enchants consists of a number of messages. These messages can be logically mapped to the /tickets endpoint as follows:

- GET /tickets/12/messages — Retrieves a list of messages for ticket #12
- GET /tickets/12/messages/5 — Retrieves message #5 for ticket #12
- POST /tickets/12/messages — Creates a new message in ticket #12
- PUT /tickets/12/messages/5 — Updates message #5 for ticket #12
- PATCH /tickets/12/messages/5 — Partially updates message #5 for ticket #12
- DELETE /tickets/12/messages/5 — Deletes message #5 for ticket #12



	```go

	router.HandleFunc("/api/couriers/2/deliveries", handler.create).Methods("GET")

	```
### Versioning

Always version your API. Versioning helps you iterate faster and prevents invalid requests from hitting updated endpoints. It also helps smooth over any major API version transitions as you can continue to offer old API versions for a period of time.

Use the version number in the URL. Example: 

- /v1/tickets/
    ```go
      router.HandleFunc("/api/v1/users", handler.create).Methods("POST")
	```

- /V2/tickets/

    ```go
      router.HandleFunc("/api/v2/users", handler.create).Methods("POST")
	```
 
#### JSON and XML

The API should either return JSON or XML data. Those two data formats are easily parsable. The recommended one is the JSON data type. One can also return other datatypes if needed.

If you are using JSON the "right" thing to do is to follow JavaScript naming conventions - and that means camelCase for field names.

By default the following naming conventions are adopted:


Variable name is generally using Camel Case style, but when you have unique nouns, should apply following rules:

- If the variable is private, and the unique noun is the first word, then use lower cases, e.g. `apiClient`.
- Otherwise, use the original cases for the unique noun, e.g. `APIClient`, `repoID`, `UserID`.

Here is a list of words which are commonly identified as unique nouns:

```go
// A GonicMapper that contains a list of common initialisms taken from golang/lint
type User struct {
	gorm.Model
	FirstName  string `json:"firstName"`
	LastName  string `json:"lastName"`
	Age string `json:"age"`
	Email string `json:"email"`
	PhoneNumber string `json:"phoneNumber"`
	Password string `json:"-"`
	Address string `json:"address"`
}


```
### link to docs
- https://docs.google.com/document/d/1txDyiV84GxpxBJw61p4RBsn3ekBXqGaNMZS3Pb5RcS0/edit#