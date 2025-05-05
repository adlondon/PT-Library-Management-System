<script setup lang="ts">
import { computed, onMounted } from 'vue'
import {
  type Book,
  libraryManagementSystemStore,
} from '../composables/libraryManagementSystemStore'

const books = computed(() => libraryManagementSystemStore.value.books)

async function checkoutBook(book: Book) {
  await libraryManagementSystemStore.value.updateBook({ ...book, checkedOut: !book.checkedOut })
}

onMounted(async () => {
  await libraryManagementSystemStore.value.fetchBooks()
})
</script>

<template>
  <div class="checkout-page">
    <h1 class="header">Checkout</h1>
    <div v-for="book in books" :key="book.isbn">
      <table>
        <td>{{ book.title }}</td>
        <td>{{ book.author }}</td>
        <td>{{ book.isbn }}</td>
        <td>{{ book.checkedOut ? 'Checked Out' : 'Available' }}</td>
        <td class="button-container">
          <button v-if="!book.checkedOut" class="checkout-book-button" @click="checkoutBook(book)">
            Checkout
          </button>
          <button v-else class="return-book-button" @click="checkoutBook(book)">Return</button>
        </td>
      </table>
    </div>
  </div>
</template>

<style scoped>
.checkout-page {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;

  .header {
    text-align: center;
    margin-bottom: 30px;
  }

  table {
    border: 1px solid white;
    border-collapse: collapse;
  }

  td {
    border: 1px solid white;
    padding: 10px;
    width: 200px;
    white-space: nowrap;
    overflow: hidden;
    text-overflow: ellipsis;
  }

  .button-container {
    width: 100px;
    display: flex;
    justify-content: center;
    align-items: center;

    button {
      width: 100px;
      height: 30px;
      background-color: transparent;
      color: white;
      cursor: pointer;
      border: none;
      border-radius: 5px;
    }

    .checkout-book-button {
      background-color: #26a96c;
      color: white;
      cursor: pointer;
    }

    .return-book-button {
      background-color: #d1d5db;
      color: #333;
    }
  }
}
</style>
