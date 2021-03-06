package photoprism

import (
	"fmt"

	"github.com/kris-nova/photoprism-client-go"
	"github.com/kris-nova/photoprism-client-go/api/v1"
)

// PhotoFinder is used to find photos by rules :)
type PhotoFinder struct {
	rules    map[int]FindPhotoRule
	client   photoprism.Client
	albumUID string
	photos   []api.Photo
}

// NewDefaultPhotoFinder will load the default (and probably most common)
// rules and give you a photo finder :)
func NewDefaultPhotoFinder(client photoprism.Client, albumUID string) *PhotoFinder {
	return &PhotoFinder{
		rules: map[int]FindPhotoRule{
			0: FindFavoritePhoto,
			1: FindUnusedPhoto,
			2: FindOldestPhotoCustom,
			3: FindRandomPhoto,
		},
		client:   client,
		albumUID: albumUID,
	}
}

// Find will iterate through the rules (in order!)
// and will match on the first rule that finds a
// photo. Otherwise will concat error messages
// and return combine error messages as a single
// error.
func (p *PhotoFinder) Find() (*api.Photo, error) {
	errMsg := ""
	err := p.LoadPhotos()
	if err != nil {
		return nil, err
	}
	for i := 0; i < len(p.rules); i++ {
		f := p.rules[i] // This is done to ensure we preserve rule order!
		p, err := f(p.photos)
		if err != nil {
			if errMsg == "" {
				errMsg = err.Error()
			} else {
				errMsg = fmt.Sprintf("%s %s", errMsg, err.Error())
			}
			continue // Continue here because error
		}
		return p, nil
	}
	return nil, fmt.Errorf("unable to find photo: %s", errMsg)
}

// LoadPhotos is the default photo lookup mechanism for a photo finder
func (p *PhotoFinder) LoadPhotos() error {
	// TODO What happens if we have a lot of lubbi photos?!
	// TODO We probably need to paginate
	client := p.client
	albumUID := p.albumUID
	photos, err := client.V1().GetPhotos(&api.PhotoOptions{
		AlbumUID: albumUID,
		Count:    DefaultAlbumCount,
	})
	if err != nil {
		return fmt.Errorf("unable to find photos in album [%s] check albumUID in config: %v", albumUID, err)
	}
	if len(photos) < 1 {
		return fmt.Errorf("unable to find >0 photos in album [%s] check albumUID in config", albumUID)

	}
	p.photos = photos
	return nil
}
