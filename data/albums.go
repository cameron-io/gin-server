package data

import "cameron.io/albums/src/models"

func GetAlbums() []models.Album {
	return []models.Album{
		{ID: "1", Title: "Dream On", Artist: "Aerosmith", Price: 10.99},
		{ID: "2", Title: "Ain't No Worries", Artist: "One Republic", Price: 11.99},
		{ID: "3", Title: "Don't Forget Me", Artist: "Lewis Capaldi", Price: 9.99},
	}
}
