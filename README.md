# argen

`argen` is an ORM code-generation tool for Go. It provides ActiveRecord-like functionality for your types.

## Installation

```
$ go get -u github.com/monochromegane/argen/...
```

## Quick start

1. Define a table mapping struct in `main.go`.
2. Mark it up with a `+AR` annotation in an adjacent comment like so:

```go
//+AR
type User struct{
	Id int `db:"pk"`
	Name string
	Age int
}
```

And at the command line, simply type:

```
$ argen main.go
```

You should see new files, named `*_gen.go`.

And you can use ActiveRecord-like functions.

```go
db, _ := sql.Open("sqlite3", "foo.db")
Use(db)

u := User{Name: "test", Age: 20}
u.Save()
//// INSERT INTO users (name, age) VALUES (?, ?); [test 20]

User{}.Where("name", "test").And("age", ">", 20).Query()
//// SELECT users.id, users.name, users.age FROM users WHERE name = ? AND age > ?; [test 20]
```

## ActiveRecord-like functions

### Setup

```go
db, _ := sql.Open("sqlite3", "foo.db")

// Set db
Use(db)
```

### Create record

```go
u := User{Name: "test", Age: 20}
u.Save()
```
or

```go
// XxxParams has same fields to original type.
User{}.Create(UserParams{Name: "test", Age: 20})
```

### Query

```go
// Get the first record
User{}.First()
//// SELECT users.id, users.name, users.age FROM users ORDER BY id ASC LIMIT ?; [1]

// Get the last record
User{}.Last()
//// SELECT users.id, users.name, users.age FROM users ORDER BY id DESC LIMIT ?; [1]

// Get the record by Id
User{}.Find(1)
//// SELECT users.id, users.name, users.age FROM users WHERE id = ? LIMIT ?; [1 1]

// Get all record
User{}.All().Query()
//// SELECT users.id, users.name, users.age FROM users;
```

### Query with conditions

```go
// Get the first matched record
User{}.FindBy("name", "test")
//// SELECT users.id, users.name, users.age FROM users WHERE name = ? LIMIT ?; [test 1]

// Get the first matched record
User{}.Where("name", "test").QueryRow()
//// SELECT users.id, users.name, users.age FROM users WHERE name = ?; [test]

// Get the all matched records
User{}.Where("name", "test").Query()
//// SELECT users.id, users.name, users.age FROM users WHERE name = ?; [test]

// And
User{}.Where("name", "test").And("age", ">", 20).Query()
//// SELECT users.id, users.name, users.age FROM users WHERE name = ? AND age > ?; [test 20]
```

```go
// Count
User{}.Count()
//// SELECT COUNT(*) FROM users;

// Exists
User{}.Exists()
//// SELECT 1 FROM users LIMIT ?; [1]

// Select
User{}.Select("id", "name").Query()
//// SELECT users.id, users.name FROM users;

// Order
User{}.Order("name", "ASC").Query()
//// SELECT users.id, users.name, users.age FROM users ORDER BY name ASC;

// Limit and offset
User{}.Limit(1).Offset(2).Query()
//// SELECT users.id, users.name, users.age FROM users LIMIT ? OFFSET ?; [1 2]

// GroupBy and having
User{}.Group("name").Having("count(name)", 2).Query()
//// SELECT users.id, users.name, users.age FROM users GROUP BY name HAVING count(name) = ?; [2]
```

### Update

```go
user, _ := User{}.Find(1)

// Update an existing struct
user.Name = "a"
user.Save()
//// UPDATE users SET id = ?, name = ?, age = ? WHERE id = ?; [1 a 20 1]

// Update attributes with validation
user.Update(UserParams{Name: "b"})
//// UPDATE users SET id = ?, name = ?, age = ? WHERE id = ?; [1 b 20 1]

// Update attributes without validation
user.UpdateColumns(UserParams{Name: "c"})
//// UPDATE users SET id = ?, name = ?, age = ? WHERE id = ?; [1 b 20 1]
```

### Delete

```go
user, _ := User{}.Find(1)

// Delete an existing struct
user.Destroy()
//// DELETE FROM users WHERE id = ?; [1]
```

## Associations

### Has One

Add association function to your type:

```go
func (m User) hasOnePost() *ar.Association {
        return nil
}
```

And type `argen` or `go generate` on your command line.

```go
user, _ := User{}.Create(UserParams{Name: "user1"})
user.BuildPost(PostParams{Name: "post1"}).Save()

// Get the related record
user.Post()
//// SELECT posts.id, posts.user_id, posts.name FROM posts WHERE user_id = ?; [1]

// Join
user.JoinsPost().Query()
//// SELECT users.id, users.name, users.age FROM users INNER JOIN posts ON posts.user_id = users.id;
```

### Has Many

Add association function to your type:

```go
func (m User) hasManyPosts() *ar.Association {
        return nil
}
```

And type `argen` or `go generate` on your command line.

```go
user, _ := User{}.Create(UserParams{Name: "user1"})
user.BuildPost(PostParams{Name: "post1"}).Save()

// Get the related records
user.Posts()
//// SELECT posts.id, posts.user_id, posts.name FROM posts WHERE user_id = ?; [1]

// Join
user.JoinsPosts()
//// SELECT users.id, users.name, users.age FROM users INNER JOIN posts ON posts.user_id = users.id;
```

### Belongs To

Add association function to your type:

```go
func (p Post) belongsToUser() *ar.Association {
        return nil
}
```

And type `argen` or `go generate` on your command line.

```go
u, _ := User{}.Create(UserParams{Name: "user1"})
post, _ := u.BuildPost(PostParams{Name: "post1"})
post.Save()

// Get the related record
post.User()
//// SELECT users.id, users.name, users.age FROM users WHERE id = ?; [1]

// Join
post.JoinsUser().Query()
//// SELECT posts.id, posts.user_id, posts.name FROM posts INNER JOIN users ON users.id = posts.user_id;
```

## Validation

Add validation function to your type:

```go
func (u User) validatesName() ar.Rule {
	// Define rule for "name" column
	return ar.MakeRule().Format().With("[a-z]")
}
```

You specify target column by naming validates`Xxx`.

### Rules

```go
// presence
Presence()

// format
Formart().With("regex")

// numericality
Numericality().OnlyInteger().GreaterThan(10)
```

### Trigger

```go
// OnCreate
OnCreate()

// OnUpdate
OnUpdate()
```

### Method chain

Validation rules has a chainable API. you could use it like this.

```go
ar.MakeRule().
	Presence().
	Format().With("[a-z]").
	OnCreate()
```

### Custom validation

Add custom validation function to your type:

```go
func (p Post) validateCustom() ar.Rule {
        return ar.CustomRule(func(errors *ar.Errors) {
                if p.Name != "name" {
                        errors.Add("name", "must be name")
                }
        }).OnUpdate()
}
```

And type `argen` or `go generate` on your command line.

### Error

```go
u := User{Name: "invalid name"}
_, errs := u.IsValid()
if errs != nil {
	fmt.Errorf("%v\n", errs.Messages["name"])
}
```


## go generate

If you want to generate it by using `go generate`, you mark it up an annotation at top of file.

```go
//go:generate argen
```

and type:

```
$ go generate
```

## TODO

- Transaction
- Callbacks (before/after save)
- Conditions for callbacks and validations.
- `AS` clause

## Author

**monochromegane**

## License

argen is licensed under the MIT
