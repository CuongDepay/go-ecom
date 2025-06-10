# Go E-commerce REST API

A robust and scalable e-commerce REST API built with Go, featuring user authentication, product management, shopping cart functionality, and order processing.

## 🚀 Features

- **User Authentication & Authorization**

  - User registration and login
  - JWT-based authentication
  - Password hashing with bcrypt
  - Protected routes for authenticated users

- **Product Management**

  - Browse all products
  - Get product details by ID
  - Create new products (authenticated users only)
  - Product inventory management

- **Shopping Cart & Orders**

  - Add items to cart
  - Checkout functionality
  - Order creation and management
  - Order item tracking
  - Inventory updates on purchase

- **Database Integration**
  - MySQL database with migration support
  - Clean architecture with repository pattern
  - Database connection pooling

## 🛠 Tech Stack

- **Backend**: Go 1.24.3
- **Database**: MySQL
- **Authentication**: JWT (JSON Web Tokens)
- **HTTP Router**: Gorilla Mux
- **Database Migrations**: golang-migrate
- **Validation**: go-playground/validator
- **Password Hashing**: bcrypt
- **Environment Management**: godotenv

## 📁 Project Structure

```
go-ecom/
├── cmd/
│   ├── api/                # API server setup
│   ├── migrate/            # Database migrations
│   └── main.go            # Application entry point
├── config/                # Configuration management
├── db/                    # Database connection
├── service/               # Business logic
│   ├── auth/             # Authentication middleware
│   ├── user/             # User management
│   ├── product/          # Product management
│   ├── cart/             # Shopping cart
│   └── order/            # Order processing
├── types/                # Data models and interfaces
├── utils/                # Utility functions
├── go.mod               # Go module dependencies
├── Makefile            # Build and run commands
└── README.md           # Project documentation
```

## 📋 Prerequisites

Before running this project, make sure you have the following installed:

- [Go 1.24.3+](https://golang.org/dl/)
- [MySQL 8.0+](https://dev.mysql.com/downloads/)
- [golang-migrate](https://github.com/golang-migrate/migrate/tree/v4.17.0/cmd/migrate) (for database migrations)

## 🔧 Installation & Setup

### 1. Clone the repository

```bash
git clone https://github.com/CuongDepay/go-ecom.git
cd go-ecom
```

### 2. Install dependencies

```bash
go mod download
```

### 3. Environment Configuration

Create a `.env` file in the root directory with the following variables:

```bash
# Server Configuration
PUBLIC_HOST=http://localhost
PORT=8080

# Database Configuration
DB_HOST=127.0.0.1
DB_PORT=3306
DB_USER=root
DB_PASSWORD=mypassword
DB_NAME=ecom

# JWT Configuration
JWT_SECRET=your-super-secret-jwt-key
JWT_EXPIRATION_IN_SECONDS=604800  # 7 days
```

### 4. Database Setup

Create a MySQL database:

```sql
CREATE DATABASE ecom;
```

Run database migrations:

```bash
make migrate-up
```

### 5. Run the application

```bash
make run
```

The server will start on `http://localhost:8080`

## 🔄 Available Commands

```bash
# Build the application
make build

# Run the application
make run

# Run tests
make test

# Run database migrations up
make migrate-up

# Run database migrations down
make migrate-down

# Create a new migration
make migration <migration_name>
```

## 📚 API Endpoints

### Authentication

| Method | Endpoint                 | Description         | Auth Required |
| ------ | ------------------------ | ------------------- | ------------- |
| POST   | `/api/v1/register`       | Register a new user | No            |
| POST   | `/api/v1/login`          | Login user          | No            |
| GET    | `/api/v1/users/{userID}` | Get user by ID      | Yes           |

### Products

| Method | Endpoint                       | Description        | Auth Required |
| ------ | ------------------------------ | ------------------ | ------------- |
| GET    | `/api/v1/products`             | Get all products   | No            |
| GET    | `/api/v1/products/{productID}` | Get product by ID  | No            |
| POST   | `/api/v1/products`             | Create new product | Yes           |

### Cart & Orders

| Method | Endpoint                | Description         | Auth Required |
| ------ | ----------------------- | ------------------- | ------------- |
| POST   | `/api/v1/cart/checkout` | Checkout cart items | Yes           |

## 🧪 API Usage Examples

### User Registration

```bash
curl -X POST http://localhost:8080/api/v1/register \
  -H "Content-Type: application/json" \
  -d '{
    "firstName": "John",
    "lastName": "Doe",
    "email": "john@example.com",
    "password": "password123"
  }'
```

### User Login

```bash
curl -X POST http://localhost:8080/api/v1/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "john@example.com",
    "password": "password123"
  }'
```

Response:

```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
}
```

### Get All Products

```bash
curl -X GET http://localhost:8080/api/v1/products
```

### Create Product (Authenticated)

```bash
curl -X POST http://localhost:8080/api/v1/products \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN" \
  -d '{
    "name": "Laptop",
    "description": "High-performance laptop",
    "image": "laptop.jpg",
    "price": 999.99,
    "quantity": 10
  }'
```

### Cart Checkout

```bash
curl -X POST http://localhost:8080/api/v1/cart/checkout \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN" \
  -d '{
    "items": [
      {
        "productID": 1,
        "quantity": 2
      },
      {
        "productID": 2,
        "quantity": 1
      }
    ]
  }'
```

## 🗄 Database Schema

The application uses the following main tables:

- **users**: User account information
- **products**: Product catalog
- **orders**: Order records
- **order_items**: Individual items within orders

## 🔒 Authentication

This API uses JWT (JSON Web Tokens) for authentication. Include the token in the `Authorization` header:

```
Authorization: Bearer <your-jwt-token>
```

JWT tokens expire after 7 days by default (configurable via `JWT_EXPIRATION_IN_SECONDS`).

## 🧪 Testing

Run the test suite:

```bash
make test
```

The project includes unit tests for handlers and services.

## 🚧 Development

### Adding New Migrations

```bash
make migration add_new_table
```

This creates new migration files in `cmd/migrate/migrations/`.

### Project Architecture

The project follows a clean architecture pattern:

- **cmd/**: Application entry points
- **service/**: Business logic and HTTP handlers
- **types/**: Data models and interfaces
- **db/**: Database connection and queries
- **config/**: Configuration management
- **utils/**: Utility functions

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📝 License

This project is licensed under the MIT License.

## 🐛 Troubleshooting

### Common Issues

1. **Database connection failed**: Ensure MySQL is running and credentials are correct
2. **Migration errors**: Check if database exists and user has proper permissions
3. **JWT token invalid**: Ensure token is not expired and JWT_SECRET matches

### Database Connection

Make sure your MySQL server is running and accessible with the credentials specified in your `.env` file.

## 📞 Support

If you encounter any issues or have questions, please open an issue on GitHub.

