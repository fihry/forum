# forum

## Table of Contents

- [Description](#description)
- [features](#features)
- [Api Documentation](#api-documentation)
- [Installation](#installation)
- [Usage](#usage)
- [Support](#support)


## Description

Forum is a simple forum application that allows users to create posts and comment on them. It is built using pure go lang template and javascript for the frontend and
sqlite3 for the database.
and go for the backend.
that hosted on gitlab.com


## Features

- [ ] login and register
- [ ] create post
- [ ] comment on post
- [ ] like and dislike post
- [ ] delete post
<!-- - [ ] edit post -->
- [ ] delete comment
<!-- - [ ] like and dislike comment -->


## Usage
    
```go
    go run main.go -p=<Port>
```
and open your browser and go to localhost:<Port>

## Api Documentation
you can use postman to test the api endpoints the link to the postman collection is [here](https://app.getpostman.com/join-team?invite_code=2eaa5c9bf99431776e8430984cfcd5b6&target_code=6fed2d2f1cab993025089b51d433f338)
### login
```http
  POST /api/login
```
| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `email`   | `string` | **Required**. user email   |
| `password`| `string` | **Required**. user password|

### register
```http
  POST api/register
```
| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `email`   | `string` | **Required**. user email   |
| `password`| `string` | **Required**. user password|
| `username`| `string` | **Required**. user username|

### create post
```http
  POST api/post
```
| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `title`   | `string` | **Required**. post title   |
| `content` | `string` | **Required**. post content |

### create comment
```http
  POST api/comment
```
| Parameter | Type     | Description                   |
| :-------- | :------- | :---------------------------- |
| `content` | `string` | **Required**. comment content |
| `post_id` | `int`    | **Required**. post id         |

### like post
```http
  POST /like
```
| Parameter | Type     | Description                  |
| :-------- | :------- | :--------------------------- |
| `post_id` | `int`    | **Required**. post id        |


### dislike post
```http
  POST /dislike
```
| Parameter | Type     | Description                  |
| :-------- | :------- | :--------------------------- |
| `post_id` | `int`    | **Required**. post id        |



## Contributors


## License

For open source projects, say how it is licensed.

## Support

special thanks to zone01Oujda for the support 