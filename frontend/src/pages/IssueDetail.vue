<template>
  <div class="container">
    <h2>이슈 상세</h2>
    <p><strong>제목:</strong> {{ issue.title }}</p>
    <p><strong>설명:</strong> {{ issue.description }}</p>
    <p><strong>상태:</strong> {{ issue.status }}</p>
    <p><strong>담당자:</strong> {{ assignedUser }}</p>

    <button @click="goToEdit">수정</button>
    <button @click="deleteIssue">삭제</button>
    <button @click="$router.push('/')">목록으로 돌아가기</button>
  </div>
</template>

<script>
import { issues, users } from "../data/mockData";

export default {
  data() {
    return {
      issue: {},
    };
  },
  computed: {
    assignedUser() {
      if (!this.issue.userId) return "없음";
      const user = users.find(user => user.id === this.issue.userId);
      return user ? user.name : "알 수 없음";
    }
  },
  methods: {
    goToEdit() {
      this.$router.push(`/issue/edit/${this.issue.id}`);
    },
    deleteIssue() {
      const index = issues.findIndex(i => i.id === this.issue.id);
      if (index !== -1) {
        issues.splice(index, 1);
        alert("이슈가 삭제되었습니다.");
        this.$router.push("/");
      }
    }
  },
  created() {
    const issueId = this.$route.params.id;
    this.issue = issues.find(i => i.id === parseInt(issueId)) || {};
  }
};
</script>