<script setup lang="ts">
const { modelValue, label, placeholder } = defineProps<{
  modelValue: string
  label: string
  placeholder: string
}>()

const emit = defineEmits<{
  'update:modelValue': [value: string]
}>()
</script>

<template>
  <div class="form-input">
    <label>{{ label }}</label>
    <input
      :placeholder="placeholder"
      type="text"
      :value="modelValue"
      @input="emit('update:modelValue', ($event.target as HTMLInputElement).value)"
    />
  </div>
</template>

<style scoped>
.form-input {
  margin-bottom: 10px;
}

label {
  color: #333;
  display: block;
  font-weight: 500;
}

input {
  width: 100%;
  border: none;
  border-bottom: 1px solid black;
  padding: 5px 0;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;

  &:focus {
    outline: none;
  }
}
</style>
