package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User model
type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Name      *string            `json:"name" validate:"required"`
	Password  *string            `json:"password" validate:"required,min=2,max=100"`
	Email     *string            `json:"email" validate:"email,required"`
	Role      *string            `json:"role" validate:"required,eq=ADMIN|eq=USER"`
	Token     *string            `json:"token,omitempty"`
	CreatedAt time.Time          `json:"created_at"`
	UpdatedAt time.Time          `json:"updated_at"`
	UserId    string             `json:"user_id"`
}

// Customer model
type Customer struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	Name       *string            `json:"name" validate:"required"`
	Email      *string            `json:"email" validate:"email,required"`
	Password   *string            `json:"password" validate:"required,min=2,max=100"`
	Status     *string            `json:"status" validate:"required,eq=lead|eq=prospect|eq=customer"` // lead, prospect, customer
	Notes      *string            `json:"notes,omitempty"`
	Token      *string            `json:"token,omitempty"`
	CreatedAt  time.Time          `json:"created_at"`
	UpdatedAt  time.Time          `json:"updated_at"`
	CustomerId string             `json:"customer_id"`
}

// Interaction model
type Interaction struct {
	ID            primitive.ObjectID `bson:"_id,omitempty"`
	UserID        primitive.ObjectID `json:"user_id" validate:"required"`
	CustomerID    primitive.ObjectID `json:"customer_id" validate:"required"`
	Type          *string            `json:"type" validate:"required,eq=task|eq=meeting|eq=followup"` // task, meeting, follow-ups
	Title         *string            `json:"title,omitempty"`
	Description   *string            `json:"description"`
	CreatedAt     time.Time          `json:"created_at"`
	UpdatedAt     time.Time          `json:"updated_at"`
	InteractionId string             `json:"interaction_id"`
}

// Ticket model
type Ticket struct {
	ID            primitive.ObjectID `bson:"_id,omitempty"`
	InteractionID primitive.ObjectID `json:"interaction_id" validate:"required"`
	UserID        primitive.ObjectID `json:"user_id" validate:"required"`
	Status        *string            `json:"status" validate:"required,eq=open|eq=in_progress|eq=resolved|eq=closed"`
	Description   *string            `json:"description"`
	CreatedAt     time.Time          `json:"created_at"`
	UpdatedAt     time.Time          `json:"updated_at"`
	TicketId      string             `json:"ticket_id"`
}
