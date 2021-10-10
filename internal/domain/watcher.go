package domain

import "context"

type WatcherType int64

const (
	SubredditWatcher WatcherType = iota
	UserWatcher
	TrendingWatcher
)

func (wt WatcherType) String() string {
	switch wt {
	case SubredditWatcher:
		return "subreddit"
	case UserWatcher:
		return "user"
	case TrendingWatcher:
		return "trending"
	}

	return "unknown"
}

type Watcher struct {
	ID             int64
	CreatedAt      float64
	LastNotifiedAt float64
	Label          string

	DeviceID  int64
	AccountID int64
	Type      WatcherType
	WatcheeID int64

	Author  string
	Upvotes int64
	Keyword string
	Flair   string
	Domain  string
	Hits    int64

	// Related models
	Device  Device
	Account Account
}

type WatcherRepository interface {
	GetByID(ctx context.Context, id int64) (Watcher, error)
	GetBySubredditID(ctx context.Context, id int64) ([]Watcher, error)
	GetByUserID(ctx context.Context, id int64) ([]Watcher, error)
	GetByTrendingSubredditID(ctx context.Context, id int64) ([]Watcher, error)
	GetByDeviceAPNSTokenAndAccountRedditID(ctx context.Context, apns string, rid string) ([]Watcher, error)

	Create(ctx context.Context, watcher *Watcher) error
	Update(ctx context.Context, watcher *Watcher) error
	IncrementHits(ctx context.Context, id int64) error
	Delete(ctx context.Context, id int64) error
	DeleteByTypeAndWatcheeID(context.Context, WatcherType, int64) error
}