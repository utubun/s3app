package models

import (
	"fmt"
	"time"
)

type Bucket struct {
	CreationDate *time.Time `type:"timestamp"`
	Name         string
}

type Buckets struct {
	Bucket []*Bucket `locationNameList:"Bucket" type:"list"`
	Owner  *Owner    `type:"structure"`
}

func NewBucket(Name string) Bucket {
	CreationDate := time.Now()
	return Bucket{&CreationDate, Name}
}

func (b Bucket) String() string {
	return fmt.Sprintf("%s created %v", b.Name, b.CreationDate)
}
