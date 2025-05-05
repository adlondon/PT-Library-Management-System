<script setup lang="ts">
import { defineProps, defineEmits } from 'vue'

type Book = {
  title: string
  author: string
  isbn: string
  description: string
}

const { modelValue } = defineProps<{
  modelValue: Book
}>()

const emit = defineEmits<{
  'update:modelValue': [value: Book]
  update: []
}>()
</script>

<template>
  <input
    placeholder="Enter book title..."
    type="text"
    :value="modelValue.title"
    @input="
      emit('update:modelValue', { ...modelValue, title: ($event.target as HTMLInputElement).value })
    "
  />
  <input
    placeholder="Enter book description..."
    type="text"
    :value="modelValue.description"
    @input="
      emit('update:modelValue', {
        ...modelValue,
        description: ($event.target as HTMLInputElement).value,
      })
    "
  />
  <input
    placeholder="Enter book author..."
    type="text"
    :value="modelValue.author"
    @input="
      emit('update:modelValue', {
        ...modelValue,
        author: ($event.target as HTMLInputElement).value,
      })
    "
  />
  <input
    placeholder="Enter book isbn..."
    type="text"
    :value="modelValue.isbn"
    @input="
      emit('update:modelValue', { ...modelValue, isbn: ($event.target as HTMLInputElement).value })
    "
  />
</template>

<style scoped>
input {
  width: 100%;
  border: none;
  border-bottom: 1px solid black;
  margin-bottom: 10px;

  &:focus {
    outline: none;
  }
}
</style>
