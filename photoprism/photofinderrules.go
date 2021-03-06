package photoprism

import (
	"fmt"
	"time"

	"github.com/kris-nova/novaarchive/random"

	"github.com/kris-nova/photoprism-client-go/api/v1"
)

const (
	//TODO this is a lot of photos!
	DefaultAlbumCount int = 500
)

// FindPhotoRule will just define the function type signature
// for how we find our little photos with our little rules.
type FindPhotoRule func(photos []api.Photo) (*api.Photo, error)

// FindFavoritePhoto will look through all photos in an album
// and return the first favorited (hearted) photo in that album.
func FindFavoritePhoto(photos []api.Photo) (*api.Photo, error) {
	for _, photo := range photos {
		if photo.PhotoFavorite {
			return &photo, nil
		}
	}
	return nil, fmt.Errorf("unable to find favorite photo")
}

func FindUnusedPhoto(photos []api.Photo) (*api.Photo, error) {
	for _, photo := range photos {
		cd := GetCustomData(photo)
		if cd == nil {
			return &photo, nil
		}
	}
	return nil, fmt.Errorf("unable to find unused photo")
}

func FindOldestPhotoCustom(photos []api.Photo) (*api.Photo, error) {
	var highestDelta time.Duration
	var oldestPhoto *api.Photo
	now := time.Now()
	// 100
	for _, photo := range photos {
		u := photo.UpdatedAt
		delta := now.Sub(u)
		if delta > highestDelta {
			oldestPhoto = &photo
			highestDelta = delta
		}
	}
	if oldestPhoto != nil {
		return oldestPhoto, nil
	}
	return nil, fmt.Errorf("unable to find oldest photo")
}

func FindRandomPhoto(photos []api.Photo) (*api.Photo, error) {
	r := random.IBetween(0, len(photos))
	p := photos[r-1]
	return &p, nil
}

// func FindNewestPhotoCustom(photos []api.Photo) (*api.Photo, error) {
//
//}

// func FindNewestPhoto(photos []api.Photo) (*api.Photo, error) {
//
//}

// func FindOldestPhoto(photos []api.Photo) (*api.Photo, error) {
//
//}
