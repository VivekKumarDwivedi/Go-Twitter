# GoTwitter

A microservices-based Twitter clone built with Go, implementing clean architecture patterns with separate services for users and tweets.

## 🚀 Tech Stack

- **Go 1.25.4** - Backend programming language
- **Chi v5.2.5** - HTTP router and middleware framework
- **MySQL** - Database
- **JWT** - Authentication tokens
- **Godotenv** - Environment variable management
- **Go Playground Validator** - Input validation

## 📁 Project Structure

```
GoTwitter/
├── TweetService/          # Tweet management microservice
│   ├── app/
│   ├── config/
│   │   ├── db/
│   │   └── env/
│   ├── controllers/
│   ├── db/
│   │   └── repositories/
│   ├── dto/
│   ├── middlewares/
│   ├── models/
│   ├── router/
│   ├── services/
│   └── utils/
└── UserService/           # User management microservice
    ├── app/
    ├── config/
    │   ├── db/
    │   └── env/
    ├── controllers/
    ├── db/
    │   └── repositories/
    ├── dto/
    ├── middlewares/
    ├── models/
    ├── router/
    ├── services/
    └── utils/
```

## 🏗️ Architecture

### Microservices Design
- **TweetService**: Handles tweet creation, retrieval, hashtag management
- **UserService**: Manages user registration, authentication, profile management

### Clean Architecture Layers
- **Controllers**: HTTP request handlers
- **Services**: Business logic implementation
- **Repositories**: Database operations
- **Models**: Database entities
- **DTOs**: Data transfer objects
- **Middlewares**: Authentication, validation, context management
- **Utils**: Helper functions (pagination, response formatting, hashtag extraction)

## 🔧 Features Implemented

### TweetService
- Tweet creation and management
- Hashtag extraction and tracking
- Pagination support
- Tag-based tweet filtering

### UserService
- User registration and authentication
- JWT token-based authentication
- Input validation
- Password hashing with cryptographic functions

## 📋 Setup Instructions

### Prerequisites
- Go 1.25.4 or higher
- MySQL database
- Git

### Installation

1. **Clone the repository**
   ```bash
   git clone <repository-url>
   cd GoTwitter
   ```

2. **Setup TweetService**
   ```bash
   cd TweetService
   go mod download
   ```

3. **Setup UserService**
   ```bash
   cd ../UserService
   go mod download
   ```

4. **Environment Configuration**
   
   Create `.env` files in both service directories:
   
   **TweetService/.env**
   ```
   DB_HOST=localhost
   DB_PORT=3306
   DB_USER=your_username
   DB_PASSWORD=your_password
   DB_NAME=tweet_db
   PORT=8081
   ```
   
   **UserService/.env**
   ```
   DB_HOST=localhost
   DB_PORT=3306
   DB_USER=your_username
   DB_PASSWORD=your_password
   DB_NAME=user_db
   JWT_SECRET=your_jwt_secret_key
   PORT=8080
   ```

5. **Database Setup**
   - Create MySQL databases: `tweet_db` and `user_db`
   - Run migrations (if available)

### Running the Services

1. **Start UserService**
   ```bash
   cd UserService
   go run main.go
   ```

2. **Start TweetService**
   ```bash
   cd TweetService
   go run main.go
   ```

## 🔌 API Endpoints

### UserService (Port 8080)
- `POST /users/register` - User registration
- `POST /users/login` - User login
- `GET /users/profile` - Get user profile
- `PUT /users/profile` - Update user profile

### TweetService (Port 8081)
- `POST /tweets` - Create tweet
- `GET /tweets` - Get tweets with pagination
- `GET /tweets/hashtag/{tag}` - Get tweets by hashtag
- `GET /tags` - Get all hashtags
- `GET /ping` - Health check

## 🔐 Authentication

The UserService implements JWT-based authentication:
- Login endpoint returns JWT token
- Token must be included in Authorization header for protected routes
- Token format: `Bearer <jwt_token>`

## 📊 Database Schema

### Users Table
- id, username, email, password_hash, created_at, updated_at

### Tweets Table  
- id, user_id, content, created_at, updated_at

### Tags Table
- id, name, created_at

### Tweet_Tags Table (Many-to-Many)
- tweet_id, tag_id

## 🛠️ Development

### Running Tests
```bash
go test ./...
```

### Building for Production
```bash
# TweetService
cd TweetService
go build -o tweet-service main.go

# UserService  
cd UserService
go build -o user-service main.go
```

## 📝 TODO

- [ ] Add Docker configuration
- [ ] Implement rate limiting
- [ ] Add comprehensive logging
- [ ] Create database migrations
- [ ] Add API documentation (Swagger)
- [ ] Implement caching layer
- [ ] Add unit and integration tests
- [ ] Setup CI/CD pipeline

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch
3. Commit your changes
4. Push to the branch
5. Create a Pull Request

## 📄 License

This project is licensed under the MIT License.

## 📞 Contact

For any questions or suggestions, please reach out to the project maintainer.
