# SimpleDouyin

<!-- LOGO -->
<div align="center">
  
<a href="https://github.com/zheng-yi-yi/SimpleDouyin/">
  <img src="assets/logo.jpg" alt="Logo">
</a>

<!-- shields -->
![GitHub watchers](https://img.shields.io/github/watchers/zheng-yi-yi/SimpleDouyin?style=flat-square&logo=github)
![GitHub Repo stars](https://img.shields.io/github/stars/zheng-yi-yi/SimpleDouyin?style=flat-square&logo=github&logoColor=orangered)
![GitHub forks](https://img.shields.io/github/forks/zheng-yi-yi/SimpleDouyin?style=flat-square&logo=github&logoColor=lightseagreen)
![GitHub contributors](https://img.shields.io/github/contributors/zheng-yi-yi/SimpleDouyin?style=flat-square&logo=github&logoColor=slateblue)

</br>

[简体中文](README.md) | <strong>English</strong> 


</div>

# Project Introduction

This is a server for a short video sharing platform written in Golang. It covers multiple functional modules, including authentication, video management, comments, likes, follows, etc., to support interactions and content sharing among users.

> ### Module Overview:
>
> - `User` Module: Handles user-related logic, including user registration, user login, and fetching user information.
> - `Video` Module: Handles video-related logic, including video publishing, fetching video feeds, fetching a user's published video list, and fetching video information.
> - `Comment` Module: Handles comment-related logic, such as posting, deleting, and fetching a video's comment list.
> - `Favorite` Module: Handles like-related logic, including liking, unliking, and fetching a user's liked video list.
> - `Relation` Module: Handles user relationship (follow) related logic, including following, unfollowing, fetching followers list, and fetching fans list.

The project uses the `Gin` framework to build an `HTTP` server and define routes. Different route handling functions are implemented to achieve features such as user registration, login, video publishing, liking, following, commenting, etc. The project also uses the `MySQL` database for data storage, and `GORM` as the `ORM` library for database operations.

# Technology Stack


<table>
    <tr>
        <th>Category</th>
        <th>Name</th>
        <th>Description</th>
    </tr>
    <tr>
        <td>Programming Language</td>
        <td><a href="https://go.dev/doc/">Go</a></td>
        <td>A statically typed, compiled, concurrent programming language with garbage collection.</td>
    </tr>
    <tr>
        <td>Framework</td>
        <td><a href="https://gin-gonic.com/docs/">Gin</a></td>
        <td>A lightweight <code>Web</code> framework for building high-performance <code>HTTP</code> services.</td>
    </tr>
    <tr>
        <td>Framework</td>
        <td><a href="https://gorm.io/docs/">Gorm</a></td>
        <td>An Object-Relational Mapping (ORM) library for <code>Go</code> to interact with databases.</td>
    </tr>
    <tr>
        <td>Database</td>
        <td><a href="https://dev.mysql.com/doc/">MySQL</a></td>
        <td>An open-source relational database management system used for storing and managing project data.</td>
    </tr>
</table>


# How to Run

Below are some simple steps to help you quickly start the project and run it locally (make sure you have the `Go` and `MySQL` development environment set up first).

<details>
<summary> Quick Start | Click to View  </summary>

## Clone the Project

```git
git clone https://github.com/zheng-yi-yi/SimpleDouyin.git
```

## Navigate to the Project Root

```bash
cd SimpleDouyin
```

## Install Dependencies

Use the following command to install the required dependencies:

```go
go mod tidy
```

## Configure the Project


pen the `config.yaml` file in the project root, fill in the values for `username` and `password`, and save the file.

Next, modify the `config.go` file in the `config` directory, and assign the constant `Ip_address` with your local `IP` address.

<details>

<summary> How to Find Your Local IP Address | Click to View </summary>

</br>

> 
> 1. Press `Win` + `R` to open the Run dialog.
> 2. Type `cmd` to open the Command Prompt.
> 3. Type `ipconfig` and press Enter to find your local IP address.
> 4. For example: `IPv4 Address . . . . . . . . . . . : 192.168.1.7`
> 


</br>
</details>

## Add the Database

Use the following command to create the douyin database:

```mysql
CREATE DATABASE douyin
```

## Build and Run

Compile the project:

```shell
go build
```

This will generate an executable file. Run it:

```
./SimpleDouyin.exe
```

## APP

Finally, in the SimpleDouyin App, double tap on the bottom right corner's **“Me”** to open advanced settings, enter the server address (your local address), and click **Save and Restart** to see the effect.

> For example, if your local IP address is `192.168.1.7`, you should enter `http://192.168.1.7:8080/` in the advanced settings.

<p><a href="https://bytedance.feishu.cn/docx/NMneddpKCoXZJLxHePUcTzGgnmf">Click here to view detailed usage instructions</a></p>

</details>

---

# Database Design

<div align="center">
    <img src="assets/tables.jpg" alt="tables.jpg">
</div>

# 🎗 Acknowledgments

[Bytedance | Youth Camp](https://youthcamp.bytedance.com/)

<div align="center">
  <a href="https://youthcamp.bytedance.com/">
    <img src="assets/bottom.jpg" alt="bottom.jpg">
  </a>
</div>