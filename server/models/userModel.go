package models

import (
    "time"
    "go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct{
    ID              primative.ObjectID      `bson: "_id"`
    first_name      *string                 `json:"first_name" validate: "required, min=2, max=100"`
    last_name       *string                 `json:"last_name" validate: "required, min=2, max=100"`
    password        *string                 `json:"Password" validate: "required, min=8`
    email           *string                 `json:"email" validate: "email, required"`
    token           *string                 `json:"token"`
    refresh_token   *string                 `json:"refresh_token"`
    user_id         *string                 `json:"user_id"`
    user_type       *string                 `json:"user_type" validate: "required, eq=ADMIN|eq=USER"`
    created_at      time.Time               `json:"create_at"`
    updated_at      time.Time               `json:"updated_at"`
}
