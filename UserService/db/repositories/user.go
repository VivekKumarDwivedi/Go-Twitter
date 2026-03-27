package repositories

import (
	"database/sql"
	"fmt"
	"time"
	"userservice/models"
)

type UserRepository interface {
	CreateUser(username string, email string, password string) (*models.User, error)
	GetUserByID(id string) (*models.User, error)
	GetUserByEmail(email string) (*models.User, error)
	DeleteUser(id string) error
	UpdateUser(id string, username string, email string) (*models.User, error)
}

type UserRepositoryImpl struct {
	db *sql.DB
}

// constructor
func NewUserRepository(_db *sql.DB) UserRepository {
	return &UserRepositoryImpl{
		db: _db,
	}
}

// Create User
func (r *UserRepositoryImpl) CreateUser(username string, email string, password string) (*models.User, error) {
	query := `INSERT INTO users (username, email,password)
			  VALUES (?,?,?)`

	result, err := r.db.Exec(query, username, email, password)

	if err != nil {
		fmt.Println("error during creating user:", err)
		return nil, err
	}

	lastInsertID, rowErr := result.LastInsertId()

	if rowErr != nil {
		fmt.Println("error getting last insert Id:", rowErr)
		return nil, rowErr
	}

	user := &models.User{
		Id:        lastInsertID,
		Username:  username,
		Email:     email,
		Password:  password,
		CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
		UpdatedAt: time.Now().Format("2006-01-02 15:04:05"),
	}

	fmt.Println("user created successfully:", user)

	return user, nil
}

func (r *UserRepositoryImpl) GetUserByID(id string) (*models.User, error) {
	fmt.Println("Fetching user in UserRepository")

	// prepare the query

	query := `SELECT id, username,email,password,createdAt,updatedAt FROM users WHERE id=?`

	// execute the query
	row := r.db.QueryRow(query, id)

	// process the row data
	user := &models.User{}

	err := row.Scan(&user.Id, &user.Username, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("No user found with the given Id")
			return nil, err
		} else {
			fmt.Println("error scaning users:", err)
			return nil, err
		}
	}

	// return the user details
	fmt.Println("user fetched successfully:", user)

	return user, nil
}

func (r *UserRepositoryImpl) GetUserByEmail(email string) (*models.User, error) {
	query := `SELECT id, email, password FROM users WHERE email=?`

	row := r.db.QueryRow(query, email)

	user := &models.User{}

	err := row.Scan(&user.Id, &user.Email, &user.Password)

	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("No user found with the given email")
			return nil, err
		} else {
			fmt.Println("error scaning users:", err)
			return nil, err
		}
	}
	return user, nil
}

func (r *UserRepositoryImpl) DeleteUser(id string) error {
	query := `DELETE FROM users WHERE id=?`

	result, err := r.db.Exec(query, id)

	if err != nil {
		fmt.Println("error deleting user:", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		fmt.Println("error getting rows affected:", err)
		return err
	}

	if rowsAffected == 0 {
		fmt.Println("No user found with the given id")
		return fmt.Errorf("no user found with the given id: %s", id)
	}

	fmt.Println("user deleted successfully")
	return nil
}

func (r *UserRepositoryImpl) UpdateUser(id string, username string, email string)(*models.User, error){

		query := `UPDATE users SET username=?, email=?, updatedAt=? WHERE id=?`
		
		currentTime := time.Now().Format("2006-01-02 15:04:05")

		result, err := r.db.Exec(query, username, email, currentTime, id)
		
		if err != nil {
			fmt.Println("error updating user:", err)
			return nil, err
		}
		
		rowsAffected, err := result.RowsAffected()
		
		if err != nil {
			fmt.Println("error getting rows affected:", err)
			return nil, err
		}
		
		if rowsAffected == 0 {
			fmt.Println("No user found with the given id")
			return nil, fmt.Errorf("no user found with the given id: %s", id)
		}
		
		fmt.Println("user updated successfully")
		return r.GetUserByID(id)
}