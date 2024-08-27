package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// type InteractionType string
// type TicketStatus string

const (
	ROLE_ADMIN = "ADMIN"
	ROLE_USER  = "USER"

	// InteractionTask     InteractionType = "task"
	// InteractionMeeting  InteractionType = "meeting"
	// InteractionFollowup InteractionType = "followup"

	TICKET_OPEN       = "open"
	TICKETIN_PROGRESS = "in_progress"
	TICKET_RESOLVED   = "resolved"
	TICKET_CLOSED     = "closed"
)

// User model
type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name     *string            `bson:"name" json:"name" validate:"required"`
	Password *string            `bson:"password" json:"password" validate:"required,min=2,max=100"`
	Email    *string            `bson:"email" json:"email" validate:"email,required"`
	Role     *string            `bson:"role" json:"role" validate:"required,eq=ADMIN|eq=USER"`
	// Company   *string            `bson:"company,omitempty" json:"company,omitempty"`
	// PhoneNo   *string            `bson:"phone_no,omitempty" json:"phone_no,omitempty"`
	Token     *string   `bson:"token,omitempty" json:"token,omitempty"`
	CreatedAt time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time `bson:"updated_at" json:"updated_at"`
	UserId    string    `bson:"user_id" json:"user_id"`
}

// Customer model
type Customer struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name       *string            `bson:"name" json:"name" validate:"required"`
	Email      *string            `bson:"email" json:"email" validate:"email,required"`
	Password   *string            `bson:"password" json:"password" validate:"required,min=2,max=100"`
	Company    *string            `bson:"company,omitempty" json:"company,omitempty"`
	Phone      *string            `bson:"phone,omitempty" json:"phone,omitempty"`
	Token      *string            `bson:"token,omitempty" json:"token,omitempty"`
	CreatedAt  time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt  time.Time          `bson:"updated_at" json:"updated_at"`
	CustomerId string             `bson:"customer_id" json:"customer_id"`
}

// Interaction model
type Interaction struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UserID     primitive.ObjectID `bson:"user_id" json:"user_id"`
	CustomerID primitive.ObjectID `bson:"customer_id" json:"customer_id"`
	// Type          *InteractionType   `bson:"type" json:"type" validate:"required,eq=task|eq=meeting|eq=followup"`
	Title         *string   `bson:"title,omitempty" json:"title,omitempty"`
	Description   *string   `bson:"description" json:"description"`
	StartTime     time.Time `bson:"start_time,omitempty" json:"start_time,omitempty"`
	CreatedAt     time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt     time.Time `bson:"updated_at" json:"updated_at"`
	InteractionId string    `bson:"interaction_id" json:"interaction_id"`
}

// Ticket model
type Ticket struct {
	ID            primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	InteractionID primitive.ObjectID `bson:"interaction_id" json:"interaction_id"`
	CustomerID    primitive.ObjectID `bson:"customer_id" json:"customer_id"`
	Status        *string            `bson:"status" json:"status" validate:"required,eq=open|eq=in_progress|eq=resolved|eq=closed"`
	Description   *string            `bson:"description" json:"description"`
	CreatedAt     time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt     time.Time          `bson:"updated_at" json:"updated_at"`
	TicketId      string             `bson:"ticket_id" json:"ticket_id"`
}
