<template>
  <div class="container">
    <h2>이슈 목록</h2>
    <label>상태 필터:</label>
    <select v-model="statusFilter">
      <option value="">전체</option>
      <option value="PENDING">PENDING</option>
      <option value="IN_PROGRESS">IN_PROGRESS</option>
      <option value="COMPLETED">COMPLETED</option>
      <option value="CANCELLED">CANCELLED</option>
    </select>
    <button @click="$router.push('/issue/new')">새 이슈 생성</button>

    <ul>
      <li v-for="issue in filteredIssues" :key="issue.id">
        <router-link :to="'/issue/' + issue.id">{{ issue.title }} - {{ issue.status }}</router-link>
      </li>
    </ul>
  </div>
</template>

<script>
import { issues } from "../data/mockData";

export default {
  data() {
    return {
      statusFilter: ""
    };
  },
  computed: {
    filteredIssues() {
      return this.statusFilter
        ? issues.filter(issue => issue.status === this.statusFilter)
        : issues;
    }
  }
};
</script>