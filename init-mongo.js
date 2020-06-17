// NOT A REAL PWD used for development purposes
db.createUser(
  {
    user: 'admin',
    pwd: 'ZY+YCzCMlMUB0TdALcNLSQQbGGSCV7hA=',
    roles: [
      {
        role: 'readWrite',
        db: 'test'
      }
    ]
  }
)