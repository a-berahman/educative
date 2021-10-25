# Educative

## Introduction

This repo is a challenge that was implemented with Golang.

We want to develop a web application that maintains information of different students and courses in a relational database.  Design and implement RESTful APIs in Golang to support these capabilities:

- Add a Course [Name, Professor Name, Description]
- Add a Student [Name, Email, Phone] 
- Assign a Student to one or more Courses
- List Students [Name, Email, Phone, Course(s) Enrolled] 
- Change a Student’s contact information
- Delete a given Course 

## Download

clone the repository:

```
git clone git@github.com:a-berahman/educative.git
```

## Run

Start by installing the dev dependencies:

```
make install-dev
```

You can setup the Postgres container to run as a service with the following snippet:

```
make postgres
make createdb
```

Run the database migration

```
Run the service with one of the following:
```

Run the service with one of the following:

```
make run
```
You can also run the service with the docker-compose which is exist in the repo.

