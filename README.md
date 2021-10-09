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

### Samples:
![image](https://user-images.githubusercontent.com/45752299/136671987-81e9ad09-9af2-48d6-a2b0-4299b724696f.png)
![image](https://user-images.githubusercontent.com/45752299/136672022-0a675851-9a96-4e49-ba6a-e5af4639e65f.png)
