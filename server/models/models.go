package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SignupStruct struct {
	ID            primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Email         string             `json:"email" bson:"email,omitempty"`
	Password      string             `json:"password" bson:"password,omitempty"`
	FullName      string             `json:"fullName" bson:"fullName,omitempty"`
	Phone         string             `json:"phone" bson:"phone,omitempty"`
	Address       string             `json:"address" bson:"address,omitempty"`
	Token         string             `json:"token" bson:"token,omitempty"`
	Last_login    string             `json:"last_login" bson:"last_login ,omitempty"`
	Last_login_IP string             `json:"last_login_ip" bson:"last_login_ip ,omitempty"`
}

type EditProfile struct {
	Email    string `json:"email"  bson:"email,omitempty"`
	FullName string `json:"fullName" bson:"fullName,omitempty"`
	Phone    string `json:"phone" bson:"phone,omitempty"`
	Address  string `json:"address"  bson:"address,omitempty"`
}

type UpdatePassword struct {
	Password    string `json:"password" bson:"password,omitempty"`
	NewPassword string `json:"newPassword"`
	Phone       string `json:"phone" bson:"phone,omitempty"`
}

type SigninStruct struct {
	Password string `json:"password"`
	Phone    string `json:"phone"`
}
type OTP struct {
	UserId primitive.ObjectID `json:"userId" bson:"userId,omitempty"`
	Otp    string             `json:"otp" bson:"otp,omitempty"`
}
type ResponseStruct struct {
	Status string      `json:"status"`
	Token  string      `json:"token"`
	Body   interface{} `json:"body"`
}

type RequestBody struct {
	ID             primitive.ObjectID  `json:"_id,omitempty" bson:"_id,omitempty"`
	RequestType    string              `json:"requestType"  bson:"requestType,omitempty"`
	FullName       string              `json:"fullName" bson:"fullName,omitempty"`
	Phone          string              `json:"phone" bson:"phone,omitempty"`
	Language       string              `json:"language" bson:"language,omitempty"`
	Address        string              `json:"address" bson:"address,omitempty"`
	State          string              `json:"state" bson:"state,omitempty"`
	City           string              `json:"city" bson:"city,omitempty"`
	Amount         string              `json:"amount" bson:"amount,omitempty"`
	Purpose        string              `json:"purpose" bson:"purpose,omitempty"`
	Files          []map[string]string `json:"files" bson:"files,omitempty"`
	AmountSaved    string              `json:"amountSaved" bson:"amountSaved,omitempty"`
	PaybackMethod  string              `json:"paybackMethod" bson:"paybackMethod,omitempty"`
	AddIntrest     string              `json:"addInterest" bson:"addInterest,omitempty"`
	Urgency        string              `json:"urgency" bson:"urgency,omitempty"`
	MeetAgent      string              `json:"meetAgent" bson:"meetAgent,omitempty"`
	BankName       string              `json:"bankName" bson:"bankName,omitempty"`
	AccountNumber  string              `json:"accountNumber" bson:"accountNumber,omitempty"`
	AccountName    string              `json:"accountName" bson:"AccountName,omitempty"`
	ShareTestimony string              `json:"shareTestimony" bson:"shareTestimony,omitempty"`
	Bvn            string              `json:"bvn" bson:"bvn,omitempty"`
	DateCreated    string              `json:"date" bson:"date,omitempty"`
	UserId         primitive.ObjectID  `json:"userId" bson:"userId,omitempty"`
}

type TalkStruct struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Date        string             `json:"date" bson:"date,omitempty"`
	Medium      string             `json:"medium" bson:"medium,omitempty"`
	Note        string             `json:"note" bson:"note,omitempty"`
	Phone       string             `json:"phone" bson:"phone,omitempty"`
	Email       string             `json:"email" bson:"email,omitempty"`
	Talktype    string             `json:"talktype" bson:"talkType,omitempty"`
	DateCreated string             `json:"date_created" bson:"date_created,omitempty"`
	UserId      primitive.ObjectID `json:"userId" bson:"userId,omitempty"`
}

type MediaItem struct {
	ID       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Imageurl string             `json:"imageurl"  bson:"imageurl,omitempty"`
	Title    string             `json:"title"  bson:"title,omitempty"`
	Desc     string             `json:"desc"  bson:"desc,omitempty"`
	Url      string             `json:"url"  bson:"url,omitempty"`
	Type     string             `json:"type"  bson:"type,omitempty"`
}
type BlogItem struct {
	Files    []map[string]string `json:"files" bson:"files,omitempty"`
	Title    string              `json:"title"  bson:"title,omitempty"`
	Subtitle string              `json:"subtitle"  bson:"subtitle,omitempty"`
	Tag      string              `json:"tag"  bson:"tag,omitempty"`
	Publish  string              `json:"publish"  bson:"publish,omitempty"`
	Desc     string              `json:"desc"  bson:"desc,omitempty"`
	Category string              `json:"category"  bson:"category,omitempty"`
	Author   string              `json:"author"  bson:"author,omitempty"`
	DateCreated    string              `json:"date" bson:"date,omitempty"`
	UserId         primitive.ObjectID  `json:"userId" bson:"userId,omitempty"`
}

type ContentsStruct struct {
	Audios   []MediaItem `json:"audios"  bson:"audios,omitempty"`
	Videos   []MediaItem `json:"videos"  bson:"videos,omitempty"`
	Articles []MediaItem `json:"articles"  bson:"articles,omitempty"`
}
