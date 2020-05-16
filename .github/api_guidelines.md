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

### Constants

- Constant should use all capital letters and use underscore `_` to separate words.

	```go
	const APP_VER = "0.7.0.1110 Beta"
	```

- If you need enumerated type, you should define the corresponding type first:

	```go
	type Scheme string

	const (
		HTTP  Scheme = "http"
		HTTPS Scheme = "https"
	)
	```

- If functionality of the module is relatively complicated and easy to mixed up with constant name, you can add prefix to every constant:

	```go
	type PullRequestStatus int

	const (
		PULL_REQUEST_STATUS_CONFLICT PullRequestStatus = iota
		PULL_REQUEST_STATUS_CHECKING
		PULL_REQUEST_STATUS_MERGEABLE
	)
	```

### Variables

- A variable name should follow general English expression or shorthand.
- In relatively simple (less objects and more specific) context, variable name can use simplified form as follows:
    - `user` to `u`
    - `userID` to `uid`
- If variable type is `bool`, its name should start with `Has`, `Is`, `Can` or `Allow`, etc.

	```go
	var isExist bool
	var hasConflict bool
	var canManage bool
	var allowGitHook bool
	```

- The last rule also applies for defining structs:

	```go
	// Webhook represents a web hook object.
	type Webhook struct {
		ID           int64 `xorm:"pk autoincr"`
		RepoID       int64
		OrgID        int64
		URL          string `xorm:"url TEXT"`
		ContentType  HookContentType
		Secret       string `xorm:"TEXT"`
		Events       string `xorm:"TEXT"`
		*HookEvent   `xorm:"-"`
		IsSSL        bool `xorm:"is_ssl"`
		IsActive     bool
		HookTaskType HookTaskType
		Meta         string     `xorm:"TEXT"` // store hook-specific attributes
		LastStatus   HookStatus // Last delivery status
		Created      time.Time  `xorm:"CREATED"`
		Updated      time.Time  `xorm:"UPDATED"`
	}
	```

#### Variable Naming Convention

Variable name is generally using Camel Case style, but when you have unique nouns, should apply following rules:

- If the variable is private, and the unique noun is the first word, then use lower cases, e.g. `apiClient`.
- Otherwise, use the original cases for the unique noun, e.g. `APIClient`, `repoID`, `UserID`.

Here is a list of words which are commonly identified as unique nouns:

```go
// A GonicMapper that contains a list of common initialisms taken from golang/lint
var LintGonicMapper = GonicMapper{
	"API":   true,
	"ASCII": true,
	"CPU":   true,
	"CSS":   true,
	"DNS":   true,
	"EOF":   true,
	"GUID":  true,
	"HTML":  true,
	"HTTP":  true,
	"HTTPS": true,
	"ID":    true,
	"IP":    true,
	"JSON":  true,
	"LHS":   true,
	"QPS":   true,
	"RAM":   true,
	"RHS":   true,
	"RPC":   true,
	"SLA":   true,
	"SMTP":  true,
	"SSH":   true,
	"TLS":   true,
	"TTL":   true,
	"UI":    true,
	"UID":   true,
	"UUID":  true,
	"URI":   true,
	"URL":   true,
	"UTF8":  true,
	"VM":    true,
	"XML":   true,
	"XSRF":  true,
	"XSS":   true,
}
```
