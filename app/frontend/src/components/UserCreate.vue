<template>
   <div class="flex justify-center mt-10">
    <div class="w-full max-w-xs">
      <form class="bg-white shadow-lg rounded px-8 pt-6 pb-8 mb-4" @submit="submit">
        <div class="mb-3">
          <label class="block text-gray-700 text-sm font-bold mb-2" for="login">
            Логин
          </label>
          <input class="shadow appearance-none border border-red-500 rounded w-full py-2 px-3 mb-2 text-gray-700 leading-tight focus:outline-none focus:shadow-outline" v-model="form.login" id="login" type="text" placeholder="Логин">
          <!-- <p class="text-red-500 text-xs italic">Введите логин</p> -->
        </div>
        <div class="mb-3">
          <label class="block text-gray-700 text-sm font-bold mb-2" for="password">
            Пароль
          </label>
          <input class="shadow appearance-none border border-red-500 rounded w-full py-2 px-3 text-gray-700 mb-2 leading-tight focus:outline-none focus:shadow-outline" v-model="form.password" id="password" type="text" placeholder="******************">
          <!-- <p class="text-red-500 text-xs italic">Введите пароль</p> -->
        </div>
        <div class="mb-3">
          <label class="block text-gray-700 text-sm font-bold mb-2" for="surname">
            Фамилия
          </label>
          <input class="shadow appearance-none border border-red-500 rounded w-full py-2 px-3 text-gray-700 mb-2 leading-tight focus:outline-none focus:shadow-outline" v-model="form.surname" id="surname" type="text" placeholder="Иванов">
          <!-- <p class="text-red-500 text-xs italic">Введите фамилию</p> -->
        </div>
        <div class="mb-3">
          <label class="block text-gray-700 text-sm font-bold mb-2" for="name">
            Имя
          </label>
          <input class="shadow appearance-none border border-red-500 rounded w-full py-2 px-3 text-gray-700 mb-2 leading-tight focus:outline-none focus:shadow-outline" v-model="form.name" id="name" type="text" placeholder="Иван">
          <!-- <p class="text-red-500 text-xs italic">Введите имя</p> -->
        </div>
        <div class="flex items-center justify-between">
          <button class="w-full bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline" type="submit">
            Регистрация
          </button>
        </div>
      </form>
    </div>
  </div> 
  
</template>

<script>
export default {
  name: 'UserCreate',
  data() {
    return {
      form: {},
    }
  },
  methods: {
    submit(e) {
      e.preventDefault();
      console.log(this.form);

      this.axios.post('/api/user/create', this.form)
					.then(response => {
            console.log(response);
					})


    }
  },
  mounted: function () {
  }
}
</script>
