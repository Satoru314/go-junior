package models

import "time"

var (
	Comment1 = Comment{
		CommentID: 1,
		ArticleID: 1,
		Message:   "first death",
		CreatedAt: time.Now(),
	}
	Comment2 = Comment{
		CommentID: 2,
		ArticleID: 1,
		Message:   "second death",
		CreatedAt: time.Now(),
	}
)
var (
	Article1 = Article{
		ID:          1,
		Title:       "first article",
		Contents:    "orehatsmonogadaisukida!!!",
		UserName:    "saki",
		NiceNum:     1,
		CommentList: []Comment{Comment1, Comment2},
		CreatedAt:   time.Now(),
	}
	Article2 = Article{
		ID:        2,
		Title:     "second article",
		UserName:  "saki",
		NiceNum:   2,
		CreatedAt: time.Now(),
	}
)
