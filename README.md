## **Freep** is an open source upload center that store your file on **Telegram** without expiration.

##### requirements
1. Go installed on your machine
2. a .env file. **You can find the necessary environment variables from .env.example file that is in root directory**
3. this section is Optionals 
  * Build [telegram-bot-api](https://tdlib.github.io/telegram-bot-api/build.html) if you want to upload/download up 2GB or 4GB if you are a premium user.

##### Installation

```
git clone https://github.com/r3a70/freep.git
cd freep && go build . && ./frp
```
Then your server starting on :8000.

#### Task lists
- [ ] Get ip, port and ssl-certs from command line.
- [ ] Dockerization.
- [ ] Replace third-party telegram lib with Self-Write lib.


#### Contribution
Rules are very simple and straightforward.
1. fork repository 
2. change code
3. push and then submit pull request

### Note that work only in develop branch
