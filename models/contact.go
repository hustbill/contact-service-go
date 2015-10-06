package models

import (
    "time"
)

type Review struct {
    Id          int64         `json:"id"`
    Code        string      `json:"code"`
    CompanyCode string      `json:"company-code"`
    Type        string      `json:"review-type"`
    BodyHtml    string      `json:"body_html"`
    Body        string      `json:"body"`
    UserId      int64         `json:"user-id"`
    StarRating  int64         `json:"star-rating"`
    LikeCount   int64         `json:"like-count"`
    Active      bool        `json:"active"`
    DeletedAt   time.Time   `json:"deleted-at"` 
    CreatedAt   time.Time   `json:"created-at"` 
    UpdatedAt   time.Time   `json:"updated-at"` 
    ClientId    int64         `json:"client-id"`
    TypeId      int64         `json:"review-type-id"`
}

type Reviews []Review
