import { ref } from 'vue'

export type Book = {
  title: string
  author: string
  isbn: string
  description: string
  checkedOut?: boolean
}

type LibraryManagementSystemStore = {
  books: Book[]
  fetchBooks: () => Promise<void>
  addBook: (book: Book) => Promise<void>
  deleteBook: (isbn: string) => Promise<void>
  updateBook: (book: Book) => Promise<void>
}

export const libraryManagementSystemStore = ref<LibraryManagementSystemStore>({
  books: [],
  async fetchBooks() {
    const response = await fetch('http://localhost:8888/books')
    const books: Book[] = await response.json()
    this.books = books
  },
  async addBook(book: Book) {
    if (!book.title || !book.author || !book.isbn || !book.description) {
      throw new Error('Please fill in all fields')
    }

    const response = await fetch('http://localhost:8888/book', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
        'Access-Control-Allow-Origin': '*',
      },
      body: JSON.stringify(book),
    })

    const addedBook: Book = await response.json()
    this.books.push(addedBook)
  },
  async deleteBook(isbn: string) {
    await fetch(`http://localhost:8888/book/${isbn}`, {
      method: 'DELETE',
    })
    await this.fetchBooks()
  },
  async updateBook(book: Book) {
    if (!book.title || !book.author || !book.isbn || !book.description) {
      throw new Error('Please fill in all fields')
    }

    const response = await fetch(`http://localhost:8888/book/${book.isbn}`, {
      method: 'PUT',
      headers: {
        'Content-Type': 'application/json',
        'Access-Control-Allow-Origin': '*',
      },
      body: JSON.stringify(book),
    })

    const updatedBook: Book = await response.json()
    const index = this.books.findIndex((b) => b.isbn === book.isbn)
    if (index !== -1) {
      this.books[index] = updatedBook
    }
  },
})
