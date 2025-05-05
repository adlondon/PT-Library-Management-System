<script setup lang="ts">
import FormInput from './FormInput.vue'
import { type Book } from '../composables/libraryManagementSystemStore'

const { modelValue } = defineProps<{
  modelValue: Book
}>()

defineEmits<{
  'update:modelValue': [value: Book]
  update: []
}>()
</script>

<template>
  <FormInput v-model="modelValue.title" label="Title" placeholder="Enter book title..." />
  <FormInput
    v-model="modelValue.description"
    label="Description"
    placeholder="Enter book description..."
  />
  <FormInput v-model="modelValue.author" label="Author" placeholder="Enter book author..." />
  <FormInput v-model="modelValue.isbn" label="ISBN" placeholder="Enter book isbn..." />
</template>
