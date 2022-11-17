# go-pkgs

Repository where there is a small development of the various open-source packages that we can find when using golang

## Static Website Generation
This is a relatively new territory for a programming language that is usually used to build backend APIs and microservices, thus a slight novelty feels imminent.

1. Hugo: A great package that lets you build static websites without backend interconnections, all written in Go.

## Dealing with Config Files
Config files are typically written in various formats like JSON and YAML. Go has a very useful package that makes reading and writing all kind of config file formats a piece of cake.

2. Viper: This is a complete configuration solution for Go applications including 12-Factor apps. It is designed to work within an app, and can handle all types of configuration needs and formats.

## Command Line Interfaces
For building CLI apps, Go has a wonderful library that makes it all a breeze.

3. Cobra: A powerful library for creating Go based CLI applications.

## Environment Variables
Go has a number of packages that enable reading .env files that store all kinds of deemed app secrets easier.
	One such great library is:

4. GoDotEnv: It boasts the easiest setup and usage for reading variables from .env files, and is quite lightweight in use too.

## Build Automation
Automation tool that aim to help execute tasks with a simple, concise command.
	One popular tool that you might already know is the Make command, that helps us use the task automation with Makefiles. A good development practise is to keep a list of commands to execute defined in a Makefile, which we can easily reference later on and execute with simple make commands.

5. Task: This library boasts more verbosity and thus, slightly better explainability for executing commands than Make. It also has no dependencies and is quite lightweight in comparison.

## Active Compilation

6. Air: This is a great utility Go package that helps rebuild and execute the project’s main.go on save or virtually any files on save (as we want it) without us typing it out to run it every single time.


## Web Development
Here are the top two web frameworks that are regularly maintained:

7. Gin Web Framework: This is the most popular web development library for Go, and for a number of good reasons.

8. Iris Web Framework: This is also another option for building high-performance web applications and APIs in Go. If you’ve worked with ExpressJS before, this will feel slightly familiar.

## Datetime Management
9. Carbon: This is a great lightweight, easy to use, and semantically intelligent datetime library for Go developers.

## Database ORM
10. Gorm: This is the easiest to use object relational mapping (ORM) Go library I’ve had the pleasure of using with three major kinds of databases SQLite, PostgreSQL, and MySQL.

## Microservices
Microservices are typically used in containerized architectures with Docker and Kubernetes to build robust applications. Here are a couple of the Go microservices packages:

11. Echo: This framework supports RESTful API design and is the most popular Go microservices framework.

12. go-micro: This library is another great option in the same realm with built-in authentication and data storage designs. Quite handy indeed.

## Discord Bots

13. DiscordGo: This is the most useful API wrapper for the Discord API functions and has a great modular structure with all major discord bot actions you might need.
	One small inconvenience is that it still doesn’t have a dedicated documentation with it, but I’ve found over the course of some experimentation with it myself, is that reading the code itself is quite easy with the way it’s neatly separated into modules and packages.

## Web Scraping
14. Colly: This is a wonderful web scraper and crawler framework for Go, especially useful for archiving (which I’ve heavily used it for) and data mining purposes.

## Miscellaneous Mentions
15. Go-redis: This is a great, highly maintained redis database client for Go. It works with both redis 6 and 7 and has a phenomenally easy setup process. Highly recommended.

16. go-elasticsearch — This is the official Elasticsearch client for Go.

17. graphql-go: This is an implementation of GraphQL in Go and supports queries, mutations & subscriptions.

