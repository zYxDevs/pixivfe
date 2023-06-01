# PixivFE

A privacy-respecting alternative front-end for Pixiv that doesn't suck

## What is this?

This is an alternative front-end for Pixiv, also my second Go project, so the code may not look good for those who are experienced :(

**Why?** Because nobody made a front-end for Pixiv that I like.

Note that this project is still under its preparation stage.

## Previews

Homepage: currently contains recommended works, daily rankings, pixivision's articles and newest works on Pixiv. ~~Route: `/`~~ Currently borked due to recent API changes.
![Homepage](https://files.catbox.moe/053fzh.png)
![Homepage](https://files.catbox.moe/m64h7s.png)

Artwork's page: informations about the artwork, recent artworks by the artist and related artworks. Route: `/artworks/<id>` (id for the preview: 107664519)
![Artwork](https://files.catbox.moe/x3k85p.png)
![Artwork](https://files.catbox.moe/ocy4fq.png)

User's page: Route: `/user/<id>` (id for the preview: 11065404)
![User](https://files.catbox.moe/f055gn.png)

## Installation

This project only depends on one Go library: [Gin](https://github.com/gin-gonic/gin)
Run these commands below, then access the site on [localhost:8080](https://localhost:8080)

```
git clone https://codeberg.org/VnPower/pixivfe.git
cd pixivfe
go get -u github.com/gin-gonic/gin
cp config.example.yml config.yml # config file
```

You will have to add your own `PHPSESSID` in order to run. Go to `config.yml` for more details.

## To-do

- [ ] Base
  - [ ] Navigation bar
  - [ ] Searching
  - [ ] Pagination
  - [x] Configuration file
  - [x] Write a real independent API
- [ ] Index page
  - [x] Recommended artworks
  - [x] Daily rankings
  - [x] Spotlight (pixivision)
  - [ ] Discovery
    - [ ] Artworks
    - [ ] Users
  - [x] Newest by all
  - [ ] Trending tags
  - [ ] Switcher (illusts/mangas)
- [ ] Single pages
  - [x] User
  - [x] Artwork
  - [ ] Spotlight
- [ ] List pages
  - [ ] Recommended artworks
  - [ ] Daily rankings
  - [ ] Discovery
    - [ ] Artworks
    - [ ] Users
  - [x] Newest by all
  - [ ] Search results
  - [ ] Switcher
- [ ] Settings
  - [ ] Login
  - [ ] Local history
  - [ ] Toggling R-18, R-18G, AI
  - [ ] Custom `pximg` proxy
- [ ] Optimization
  - [x] Split web components into smaller templates
  - [x] Clean the models + JSON
  - [x] Navigation between pages
  - [x] Lazy load images
  - [ ] Better error handling
  - [ ] Fully proxy images from Pixiv

## License

[~~MIT~~](https://mit-license.org/) [AGPL3](https://www.gnu.org/licenses/agpl-3.0.txt)
