<script setup>
import { ref, onMounted } from "vue";

const searchTerm = ref("");
const documents = ref([]);

onMounted(() => {
  const searchInput = document.getElementById("search");
  searchInput.addEventListener("input", handleSearchInput);
  fetchDocuments(searchTerm.value);
});

const handleSearchInput = (event) => {
  searchTerm.value = event.target.value.trim();
  fetchDocuments(searchTerm.value);
};

const fetchDocuments = async (searchTerm) => {
  const requestData = { tag: searchTerm };

  try {
    const response = await fetch("http://127.0.0.1:8080/search/", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(requestData),
    });

    const data = await response.json();
    documents.value = data;
  } catch (error) {
    console.error("Error:", error);
  }
};
</script>

<template>
  <div class="container mx-auto p-6">
    <div class="search-wrapper mb-6">
      <label for="search" class="block text-gray-700 mb-1"
        >Search Document</label
      >
      <input
        v-model="searchTerm"
        id="search"
        data-search
        class="w-full rounded p-2 border border-gray-300 focus:outline-none focus:ring-2 focus:ring-blue-500"
      />
    </div>
    <div
      class="document-cards grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 2xl:grid-cols-5 gap-6"
      data-doc-cards-container
    >
      <div
        v-for="(document, index) in documents"
        :key="index"
        class="card border rounded p-4"
      >
        <div class="header" data-header>{{ document.name }}</div>
        <div class="body" data-body></div>
      </div>
      <div v-if="documents.length == 0">
        <p
          class="text-gray-600 text-center col-span-4 xl:col-span-5 2xl:col-span-5"
        >
          No data found
        </p>
      </div>
    </div>
  </div>
</template>

<style lang="css" scoped></style>
