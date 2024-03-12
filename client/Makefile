HOST = http://localhost:8080

.PHONY: postAlbum getAlbums getAlbumById

postAlbum:
	curl $(HOST)/albums \
		--include \
		--header "Content-Type: application/json" \
		--request "POST" \
		--data '{"id": "4","title": "The Modern Sound of Betty Carter","artist": "Betty Carter","price": 49.99}'

getAlbums:
	curl $(HOST)/albums

getAlbumById:
	@if [ -n "$(albumId)" ]; then\
		curl $(HOST)/albums/"$(albumId)";\
	else\
		echo "Usage: make albumId=<ID> getAlbumById";\
	fi
