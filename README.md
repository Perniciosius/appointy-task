# Appointy

|URL                 |Method |Path Parameter |Query Parameter             |
|--------------------|-------|---------------|----------------------------|
|/users              |POST   |-              |-                           |
|/users/:userId      |GET    |userId         |-                           |
|/posts              |POST   |-              |-                           |
|/posts/:postId      |GET    |postId         |-                           |
|/posts/users/:userId|GET    |userId         |limit=(number)&skip=(number)|


### URL: POST /users
Json Body:
```
{
    "name": "ABC",
    "email": "abc@abd.com",
    "password": "abc123"
}
```


### URL: POST /posts
Json Body:
```
{
    "userId": "348nref89332rew9f8sduf",
    "caption": "Something",
    "imageUrl": "https://images.unsplash.com/photo-1592564630984-7410f94db184?ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8ixlib=rb-1.2.1&auto=format&fit=crop&w=1146&q=80",
}
```