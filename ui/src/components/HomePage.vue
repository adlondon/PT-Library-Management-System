<script setup lang="ts">
import { ref, onMounted, computed } from 'vue'
import AddEditBookForm from './AddEditBookForm.vue'
import {
  type Book,
  libraryManagementSystemStore,
} from '../composables/libraryManagementSystemStore'

const showAddBook = ref(false)
const showEditBook = ref<string | null>(null)
const editedBook = ref<Book>({
  title: '',
  author: '',
  isbn: '',
  description: '',
})

const newBook = ref<Book>({
  title: '',
  author: '',
  isbn: '',
  description: '',
})

const books = computed(() => libraryManagementSystemStore.value.books)

async function update() {
  try {
    await libraryManagementSystemStore.value.addBook(newBook.value)
    showAddBook.value = false
    newBook.value = {
      title: '',
      author: '',
      isbn: '',
      description: '',
    }
  } catch (error: any) {
    alert(error.message)
  }
}

async function updateBook() {
  try {
    await libraryManagementSystemStore.value.updateBook(editedBook.value)
    showEditBook.value = null
  } catch (error: any) {
    alert(error.message)
  }
}

async function deleteBook(isbn: string) {
  await libraryManagementSystemStore.value.deleteBook(isbn)
}

function startEdit(book: Book) {
  editedBook.value = { ...book }
  showEditBook.value = book.isbn
}

onMounted(async () => {
  await libraryManagementSystemStore.value.fetchBooks()
})
</script>

<template>
  <main>
    <h1 class="header">Books in library:</h1>
    <div class="library-container">
      <div v-if="books.length > 0" v-for="book in books" :key="book.isbn">
        <div v-if="showEditBook === book.isbn" class="edit-book-container library-item">
          <AddEditBookForm v-model="editedBook" @update="updateBook" />
          <div class="edit-book-button-container">
            <button class="confirm-book-button" @click="updateBook">Update</button>
            <button class="cancel-book-button" @click="showEditBook = null">Cancel</button>
          </div>
          <button class="delete-book-button" @click="deleteBook(book.isbn)">Delete Book</button>
        </div>
        <div v-else class="library-item" @click="startEdit(book)">
          <div class="book-info">
            <h3 class="title">{{ book.title }}</h3>
            <div class="description">{{ book.description }}</div>
          </div>
          <div class="book-details">
            <div class="author">{{ book.author }}</div>
            <div class="isbn">{{ book.isbn }}</div>
          </div>
        </div>
      </div>
      <div v-else class="library-item">
        <div class="book-info">
          <h3 class="title">No Books Found</h3>
          <div class="description">Add a book to the library.</div>
        </div>
        <img src="@/assets/404.png" alt="No books found" />
      </div>
      <div class="add-book-container" v-if="!showAddBook">
        <button class="add-book-plus-button" @click="showAddBook = true">+</button>
      </div>
      <div class="add-book-container" v-else>
        <div>
          <AddEditBookForm v-model="newBook" @update="update" />
        </div>
        <div class="add-book-button-container">
          <button class="confirm-book-button" @click="update">Add Book</button>
          <button class="cancel-book-button" @click="showAddBook = false">Cancel</button>
        </div>
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
    height: 325px;
    border: 1px solid white;
    padding: 10px;
    border-radius: 5px;
  }

  .library-item {
    position: relative;
    cursor: pointer;

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

  .add-book-container,
  .edit-book-container {
    background-color: white;

    .add-book-button-container,
    .edit-book-button-container {
      display: flex;
      gap: 10px;
      justify-content: center;

      button {
        cursor: pointer;
        border-radius: 5px;
        padding: 5px;
        width: 100%;
        border: none;
      }
    }

    .delete-book-button {
      background-color: red;
      color: white;
      border: none;
      cursor: pointer;
      margin-top: 10px;
      border-radius: 5px;
      padding: 5px;
    }

    .confirm-book-button {
      background-color: #26a96c;
      color: white;
      cursor: pointer;
    }

    .cancel-book-button {
      background-color: #d1d5db;
      color: #333;
    }

    .add-book-plus-button {
      background-color: transparent;
      color: #333;
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
