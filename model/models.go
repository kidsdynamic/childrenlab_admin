package model

import "time"

type AdminLogin struct {
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type AdminSignUpRequest struct {
	Name      string `json:"name" binding:"required"`
	Password  string `json:"password" binding:"required"`
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
}

type User struct {
	ID        int64  `json:"id"`
	Email     string `json:"email"`
	FirstName string `json:"firstName" db:"first_name"`
	LastName  string `json:"lastName" db:"last_name"`
}

type AccessToken struct {
	ID          int64
	Email       string    `db:"email"`
	Token       string    `db:"token"`
	LastUpdated time.Time `db:"last_updated"`
	AccessCount int       `db:"access_count"`
}

type TokenRequest struct {
	Email string `query:"email" binding:"required"`
	Token string `query:"token" binding:"required"`
}

// Dashboard
type Dashboard struct {
	TotalUserCount      int64                 `json:"totalUserCount"`
	Signup              []SignupCountByDate   `json:"signup"`
	TotalActivityCount  int64                 `json:"totalActivityCount"`
	Activity            []ActivityCountByDate `json:"activity"`
	ActivityByEventDate []ActivityCountByDate `json:"activityByEventDate"`
}

type SignupCountByDate struct {
	SignupCount int64  `db:"signup",json:"signupCount"`
	Date        string `db:"date",json:"date"`
}

type ActivityCountByDate struct {
	ActivityCount int64  `db:"count",json:"activityCount"`
	UserCount     int64  `db:"userCount",json:"userCount"`
	IndoorSteps   string `db:"indoorSteps",json:"indoorSteps"`
	OutdoorSteps  string `db:"outdoorSteps",json:"outdoorSteps"`
	Date          string `db:"date",json:"date"`
}
