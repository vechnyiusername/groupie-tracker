# Groupie Tracker 🎵

## Description

**Groupie Tracker** is a web application written in **Go** that consumes a given RESTful API and displays information about music bands and artists in a user-friendly website.

The project focuses on:
- Manipulating and storing API data
- Building a backend server in Go
- Rendering dynamic HTML pages
- Handling client–server interactions (events/actions)
- Proper error handling to ensure the server never crashes

---

## Authors
abotaubay
kmurza
zorymbet

---

## Objectives

The goal of this project is to retrieve data from a provided API and visualize it through a web interface.

The API contains four main endpoints:

- **Artists**  
  Information about bands and artists such as:
  - Name
  - Image
  - Year of creation
  - First album date
  - Members

- **Locations**  
  Last and/or upcoming concert locations

- **Dates**  
  Last and/or upcoming concert dates

- **Relations**  
  Links artists with their corresponding locations and dates

Using this data, the application displays artist information through different visual formats such as cards, lists, and detailed pages.

---

## Features

- Display a list of all artists
- View detailed information about a selected artist
- Show concert locations and dates
- Client-to-server interaction using HTTP requests
- Custom error pages (404, 500)
- Safe error handling (server never crashes)
- Clean separation between backend logic and templates

---

## Technologies Used

- **Go** (standard library only)
- **HTML templates**
- **CSS**
- **Net/HTTP package**
- **JSON parsing**

---

## How It Works

1. The server starts and loads data from the API
2. Data is stored in Go structures
3. Routes handle client requests:
   - `/` → Home page (all artists)
   - `/artist?id=X` → Artist details page
4. HTML templates render the data dynamically
5. Errors are handled gracefully using custom error pages

---

## Error Handling

- Invalid routes return **404 Not Found**
- Internal server issues return **500 Internal Server Error**
- Missing or incorrect data never crashes the server
- Templates are parsed safely without `panic`

---

## Event / Client-Server Interaction

The project implements a client-server event where:
- The client sends a request (e.g. clicking an artist)
- The server processes the request
- The server responds with dynamic data rendered in HTML

This demonstrates request–response communication using HTTP.

---

## Usage

### Run the server

```bash
go run .

