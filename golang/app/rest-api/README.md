## Event Booking REST API (Go)

Built a RESTful Event Booking API in Go using the Gin web framework that enables users to register, authenticate, create and manage events, and register or cancel event bookings. The application uses JWT-based authentication and authorization to secure protected endpoints, ensuring only authenticated users can perform restricted actions and only event owners can update or delete their events.

The project follows a modular architecture with separate packages for routing, middleware, models, database access, and utility functions. User passwords are securely hashed before storage, and data is persisted in SQLite using parameterized SQL queries. The API exposes standard REST endpoints for user management, event CRUD operations, and event registration, with proper HTTP status codes and JSON responses.

**Key features:**

* User registration and login with JWT authentication
* Secure password hashing using bcrypt
* CRUD operations for events
* Event registration and cancellation
* Authorization to restrict event modifications to owners
* SQLite database integration
* Modular project structure using Gin middleware and route handlers
* RESTful API design with JSON request/response handling
* HTTP endpoint testing using `.http` request files

---