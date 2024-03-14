<template>
  <div class="b-user-detail">
    Страница пользователя
  </div>
  <div>Данные пользователя</div>
  <div v-for="(value, key) in userData" v-bind:key="key">{{ key }}: {{ value }}</div>
</template>

<script>
export default {
  name: 'UserDetail',
  data() {
    return {
      userData: {}
    }
  },
  methods: {
    getData: async function () {
      
				await this.axios.get('/api/user?id=' + this.$route.params.id)
					.then(response => {
            this.userData = response.data;
            console.log(this.userData);
					})
    },
  },
  mounted: async function () {
			await this.getData();
		}
}
</script>
