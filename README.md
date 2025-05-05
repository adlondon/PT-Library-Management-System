# PT Library Management System

## Approach

This is the first project I've built using Go, and I honestly really enjoyed it. Data structure and methods were fairly intuitive, and I found the docs to be approachable and easy to parse through. I referenced a quick [tutorial](https://dev.to/maxiim3/create-a-fullstack-app-with-vue-and-go-and-tailwindcss-v4-22ai) to help with decisions around build tools and initial setup. I opted to not include a database in the interest of time, and instead initialized the `Books` slice with seed data.

I felt third party libraries would be overkill in the Frontend, but if I were to scale this app I would have incorporated a statemanagement library like Pinia and probably a CSS framework like Tailwind.

## Project Setup

### Backend:

```sh
// make sure you have go installed on your machine
go run main.go
```

### Frontend

```sh
cd ui && npm install
```

```sh
npm run dev
```

Should be running at http://localhost:5173/
