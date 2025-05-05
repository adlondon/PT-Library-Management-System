<script setup lang="ts">
import { ref, onMounted } from 'vue'
import AddBookForm from './AddBookForm.vue'

type Book = {
  title: string
  author: string
  isbn: string
  description: string
}

const data = ref<Book[]>([])
const showAddBook = ref(false)
const newBook = ref<Book>({
  title: '',
  author: '',
  isbn: '',
  description: '',
})

async function fetchData() {
  const prom = await fetch('http://localhost:8888/books')
  const res: Awaited<Book[]> = await prom.json()
  data.value = res
}

async function update() {
  if (
    !newBook.value.title ||
    !newBook.value.author ||
    !newBook.value.isbn ||
    !newBook.value.description
  ) {
    alert('Please fill in all fields')
    return
  }

  const resp = await fetch('http://localhost:8888/book', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
      'Access-Control-Allow-Origin': '*',
    },
    body: JSON.stringify(newBook.value),
  })

  const addedBook: Book = await resp.json()

  data.value.push(addedBook)

  showAddBook.value = false
  newBook.value = {
    title: '',
    author: '',
    isbn: '',
    description: '',
  }
}

onMounted(() => {
  fetchData()
})
</script>

<template>
  <main>
    <h1 class="header">Books in library:</h1>
    <div class="library-container" v-if="data.length > 0">
      <div class="library-item" v-for="book in data" :key="book.isbn">
        <div class="book-info">
          <h3 class="title">{{ book.title }}</h3>
          <div class="description">{{ book.description }}</div>
        </div>
        <div class="book-details">
          <div class="author">{{ book.author }}</div>
          <div class="isbn">{{ book.isbn }}</div>
        </div>
      </div>
      <div class="add-book-container" v-if="!showAddBook">
        <button class="add-book-plus-button" @click="showAddBook = true">+</button>
      </div>
      <div class="add-book-container" v-else>
        <div>
          <AddBookForm v-model="newBook" @update="update" />
        </div>
        <div class="add-book-button-container">
          <button @click="update">Add Book</button>
          <button @click="showAddBook = false">Cancel</button>
        </div>
      </div>
    </div>
    <div v-else>
      <p>No books found. Add a book to the library.</p>
      <div class="add-book-container">
        <AddBookForm v-model="newBook" @update="update" />
      </div>
    </div>
  </main>
</template>

<style scoped>
.header {
  text-align: center;
  margin-bottom: 30px;
}

.library-container {
  display: flex;
  gap: 50px;
  flex-wrap: wrap;
  justify-content: center;

  .library-item,
  .add-book-container {
    display: flex;
    flex-direction: column;
    justify-content: space-between;
    width: 200px;
    height: 300px;
    border: 1px solid white;
    padding: 10px;
    border-radius: 5px;
  }

  .library-item {
    .title {
      text-align: center;
    }

    .author {
      text-align: center;
      font-size: 12px;
    }

    .description {
      text-align: center;
      font-size: 12px;
    }

    .isbn {
      text-align: right;
      font-size: 6px;
    }
  }

  .add-book-container {
    background-color: white;

    .add-book-button-container {
      display: flex;
      gap: 10px;
      justify-content: center;

      button {
        background-color: transparent;
        color: black;
        border: none;
        cursor: pointer;
      }
    }

    .add-book-plus-button {
      background-color: transparent;
      color: black;
      font-size: 120px;
      border: none;
      cursor: pointer;
      width: 100%;
      height: 100%;
      margin-top: -10px;
    }
  }
}
</style>
