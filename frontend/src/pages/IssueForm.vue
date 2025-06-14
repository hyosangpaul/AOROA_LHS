<template>
  <div class="issue-form">
    <h2>{{ isEditMode ? "이슈 수정" : "새 이슈 생성" }}</h2>
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
    <button @click="$router.push('/')">목록으로 돌아가기</button>
  </div>
</template>

<script>
import { issues, users } from "../data/mockData";

export default {
  data() {
    return {
      issue: {},
      users,
      isEditMode: false
    };
  },
  methods: {
    saveIssue() {
      if (this.isEditMode) {
        // 수정
        const index = issues.findIndex(i => i.id === this.issue.id);
        if (index !== -1) issues[index] = this.issue;
      } else {
        // 생성
        this.issue.id = issues.length + 1;
        this.issue.createdAt = new Date().toISOString();
        issues.push(this.issue);
      }
      this.$router.push("/");
    }
  },
  created() {
    const issueId = this.$route.params.id;
    if (issueId) {
      this.issue = issues.find(i => i.id === parseInt(issueId)) || {};
      this.isEditMode = true;
    } else {
      this.issue = { title: "", description: "", status: "PENDING", userId: null };
    }
  }
};
</script>