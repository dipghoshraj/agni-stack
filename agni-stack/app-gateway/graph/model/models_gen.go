// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type App struct {
	ID          string        `json:"id"`
	Name        string        `json:"name"`
	Description string        `json:"description"`
	Image       *string       `json:"image,omitempty"`
	Owner       *BasicUser    `json:"owner,omitempty"`
	Project     *BasicProject `json:"project,omitempty"`
}

type AppInput struct {
	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"`
	Image       *string `json:"image,omitempty"`
	ProjectID   *int32  `json:"project_id,omitempty"`
}

type AuthResponse struct {
	Token string `json:"token"`
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type BasicApp struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"`
	Image       *string `json:"image,omitempty"`
}

type BasicProject struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type BasicUser struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Mutation struct {
}

type Project struct {
	ID          string      `json:"id"`
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Owner       *BasicUser  `json:"owner,omitempty"`
	Apps        []*BasicApp `json:"apps"`
}

type ProjectInput struct {
	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"`
}

type Query struct {
}

type User struct {
	ID       string     `json:"id"`
	Name     string     `json:"name"`
	Email    string     `json:"email"`
	Projects []*Project `json:"projects"`
}

type UserInput struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
