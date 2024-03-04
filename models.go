package main

import (
	"blog-aggregator/internal/database"
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name string `json:"name"`
	ApiKey string `json:"api_key"`
}

func databaseUserToUser(user database.User) User {
	return User{
		ID: user.ID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Name: user.Name,
		ApiKey: user.ApiKey,
	}
}


type Feed struct {
	ID uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name string `json:"name"`
	Url string `json:"url"`
	UserId uuid.UUID `json:"user_id"`
	LastFetchedAt *time.Time `json:"last_fetched_at"`
}

func databaseFeedToFeed(feed database.Feed) Feed {
	return Feed{
		ID: feed.ID,
		CreatedAt: feed.CreatedAt,
		UpdatedAt: feed.UpdatedAt,
		Name: feed.Name,
		Url: feed.Url,
		UserId: feed.ID,
		LastFetchedAt: &feed.LastFetchedAt.Time,
	}
}

func databaseFeedsToFeeds(feeds []database.Feed) []Feed {
	result := make([]Feed, len(feeds))
	for i, feed := range feeds {
		result[i] = databaseFeedToFeed(feed)
	}
	return result
}

type FeedFollows struct {
	ID uuid.UUID `json:"id"`
	FeedID uuid.UUID `json:"feed_id"`
	UserID uuid.UUID `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func databaseFeedFollowToFeedFollow(feed_follows database.FeedFollow) FeedFollows {
	return FeedFollows{
		ID: feed_follows.ID,
		FeedID: feed_follows.FeedID,
		UserID: feed_follows.UserID,
		CreatedAt: feed_follows.CreatedAt,
		UpdatedAt: feed_follows.UpdatedAt,
	}
}

func databaseFeedFollowsToFeedFollows(feed_follows []database.FeedFollow) []FeedFollows {
	result := make([]FeedFollows, len(feed_follows))
	for i, feed_follow := range feed_follows {
		result[i] = databaseFeedFollowToFeedFollow(feed_follow)
	}
	return result
}

type Post struct {
	ID          uuid.UUID `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Title       string `json:"title"`
	Url         string `json:"url"`
	Description string `json:"description"`
	PublishedAt time.Time `json:"publishedAt"`
	FeedID      uuid.UUID `json:"feed_id"`
}

func databasePostToPost(post database.Post) Post {
	return Post{
		ID: post.ID,
		CreatedAt: post.CreatedAt,
		UpdatedAt: post.UpdatedAt,
		Title: post.Title,
		Url: post.Url,
		Description: post.Description.String,
		PublishedAt: post.PublishedAt,
		FeedID: post.FeedID,
	}
}

func databasePostsToPosts(posts []database.Post) []Post {
	result := make([]Post, len(posts))
	for i, post := range posts {
		result[i] = databasePostToPost(post)
	}
	return result
}