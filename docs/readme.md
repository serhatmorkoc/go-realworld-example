## TODOs


- [x] Users and Authentication
  - [x] `POST /user/login`: Existing user login
  - [x] `POST /users`: Register a new user
  - [x] `GET /user`: Get current user
  - [x] `PUT /user`: Update current user
- [x] Profiles
  - [x] `GET /profiles/{username}`: Get a profile
  - [ ] `POST /profiles/{username}/follow`: Follow a user
  - [ ] `DELETE /profiles/{username}/follow`: Unfollow a user
- [x] Articles
  - [ ] `GET /articles/feed`: Get recent articles from users you follow
  - [ ] `GET /articles`: Get recent articles globally
  - [ ] `POST /articles `: Create an article
  - [ ] `GET /articles/{slug}`: Get an article
  - [ ] `PUT /articles/{slug}`: Update an article
  - [ ] `DELETE /articles/{slug}`: Delete an article
- [x] Comments
  - [ ] `GET /articles/{slug}/comments`: Get comments for an article
  - [ ] `POST /articles/{slug}/comments`: Create a comment for an article
  - [ ] `DELETE /articles/{slug}/comments/{id}`: Delete a comment for an article
- [x] Favorites
  - [ ] `POST /articles/{slug}/favorite`: Favorite an article
  - [ ] `DELETE /articles/{slug}/favorite`: Unfavorite an article
