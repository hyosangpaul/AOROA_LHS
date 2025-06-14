<template>
  <div class="container">
    <h2>이슈 수정</h2>
    <label>제목:</label>
    <input v-model="issue.title" />

    <label>설명:</label>
    <textarea v-model="issue.description"></textarea>

    <label>담당자:</label>
    <select v-model="issue.userId">
      <option value="">없음</option>
      <option v-for="user in users" :key="user.id" :value="user.id">{{ user.name }}</option>
    </select>

    <label>상태:</label>
    <select v-model="issue.status">
      <option value="PENDING">PENDING</option>
      <option value="IN_PROGRESS">IN_PROGRESS</option>
      <option value="COMPLETED">COMPLETED</option>
      <option value="CANCELLED">CANCELLED</option>
    </select>

    <button @click="saveIssue">저장</button>
    <button @click="$router.push(`/issue/${issue.id}`)">취소</button>
  </div>
</template>

<script>
import { issues, users } from "../data/mockData";

export default {
  data() {
    return {
      issue: {},
      users
    };
  },
  methods: {
    saveIssue() {
      const index = issues.findIndex(i => i.id === this.issue.id);
      if (index !== -1) issues[index] = this.issue;
      alert("이슈가 수정되었습니다.");
      this.$router.push(`/issue/${this.issue.id}`);
    }
  },
  created() {
    const issueId = this.$route.params.id;
    this.issue = issues.find(i => i.id === parseInt(issueId)) || {};
  }
};
</script>